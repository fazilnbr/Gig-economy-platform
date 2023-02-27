package handler

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
	"strconv"
	"strings"

	"github.com/fazilnbr/project-workey/pkg/common/response"
	"github.com/fazilnbr/project-workey/pkg/config"
	"github.com/fazilnbr/project-workey/pkg/domain"
	services "github.com/fazilnbr/project-workey/pkg/usecase/interface"
	"github.com/fazilnbr/project-workey/pkg/utils"
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v4"
	"golang.org/x/oauth2"
	"golang.org/x/oauth2/google"
)

type AuthHandler struct {
	adminUseCase  services.AdminUseCase
	workerUseCase services.WorkerUseCase
	userUseCase   services.UserUseCase
	jwtUseCase    services.JWTUseCase
	authUseCase   services.AuthUseCase
	cfg           config.Config
}

func NewAuthHandler(
	adminUseCase services.AdminUseCase,
	workerUseCase services.WorkerUseCase,
	userusecase services.UserUseCase,
	jwtUseCase services.JWTUseCase,
	authUseCase services.AuthUseCase,
	cfg config.Config,

) AuthHandler {
	return AuthHandler{
		adminUseCase:  adminUseCase,
		workerUseCase: workerUseCase,
		userUseCase:   userusecase,
		jwtUseCase:    jwtUseCase,
		authUseCase:   authUseCase,
		cfg:           cfg,
	}
}

var (
	oauthConfGl = &oauth2.Config{
		ClientID:     "",
		ClientSecret: "",
		RedirectURL:  "http://localhost:8080/user/callback-gl",
		Scopes:       []string{"https://www.googleapis.com/auth/userinfo.email"},
		Endpoint:     google.Endpoint,
	}
	oauthStateStringGl = ""
)

func (cr *AuthHandler) InitializeOAuthGoogle() {
	oauthConfGl.ClientID = cr.cfg.ClientID
	oauthConfGl.ClientSecret = cr.cfg.ClientSecret
	oauthStateStringGl = cr.cfg.OauthStateString
	fmt.Printf("\n\n%v\n\n", oauthConfGl)
}

// @Summary Authenticate With Google
// @ID Authenticate With Google
// @Tags User Authentication
// @Security BearerAuth
// @Produce json
// @Success 200 {object} response.Response{}
// @Failure 422 {object} response.Response{}
// @Router /user/login-gl [get]
func (cr *AuthHandler) GoogleAuth(c *gin.Context) {
	HandileLogin(c, oauthConfGl, oauthStateStringGl)
}

func HandileLogin(c *gin.Context, oauthConf *oauth2.Config, oauthStateString string) error {
	URL, err := url.Parse(oauthConf.Endpoint.AuthURL)
	if err != nil {
		fmt.Printf("\n\n\nerror in handile login :%v\n\n", err)
		return err
	}
	parameters := url.Values{}
	parameters.Add("client_id", oauthConf.ClientID)
	parameters.Add("scope", strings.Join(oauthConf.Scopes, " "))
	parameters.Add("redirect_uri", oauthConf.RedirectURL)
	parameters.Add("response_type", "code")
	parameters.Add("state", oauthStateString)
	URL.RawQuery = parameters.Encode()
	url := URL.String()
	fmt.Printf("\n\nurl : %v\n\n", oauthConf.RedirectURL)
	c.Redirect(http.StatusTemporaryRedirect, url)
	return nil

}

func (cr *AuthHandler) CallBackFromGoogle(c *gin.Context) {
	fmt.Print("\n\nfuck\n\n")
	c.Request.ParseForm()
	state := c.Request.FormValue("state")

	if state != oauthStateStringGl {
		c.Redirect(http.StatusTemporaryRedirect, "/")
		return
	}

	code := c.Request.FormValue("code")

	if code == "" {
		c.JSON(http.StatusBadRequest, "Code Not Found to provide AccessToken..\n")

		reason := c.Request.FormValue("error_reason")
		if reason == "user_denied" {
			c.JSON(http.StatusBadRequest, "User has denied Permission..")
		}
	} else {
		token, err := oauthConfGl.Exchange(oauth2.NoContext, code)
		if err != nil {
			return
		}
		resp, err := http.Get("https://www.googleapis.com/oauth2/v2/userinfo?access_token=" + url.QueryEscape(token.AccessToken))
		if err != nil {
			c.Redirect(http.StatusTemporaryRedirect, "/")
			return
		}
		defer resp.Body.Close()

		respons, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			c.Redirect(http.StatusTemporaryRedirect, "/")
			return
		}
		type data struct {
			Id             string
			Email          string
			Verified_email bool
			Picture        string
			// data           string
		}
		var gdata data
		json.Unmarshal(respons, &gdata)
		fmt.Printf("\n\ndata :%v\n\n", string(respons))
		fmt.Printf("\n\ndata :%v\n\n", gdata)

		if !gdata.Verified_email {
			response := response.ErrorResponse("Failed to Login ", "Your email is not varified by google ", nil)
			c.Writer.Header().Add("Content-Type", "application/json")
			c.Writer.WriteHeader(http.StatusUnauthorized)
			utils.ResponseJSON(*c, response)
			return
		}
		var newUser domain.User

		newUser.UserName, newUser.Verification = gdata.Email, gdata.Verified_email

		err = cr.userUseCase.CreateUser(newUser)
		fmt.Printf("\n\nerrrrorr  :  %v\n\n", err)

		if err == nil || err.Error() == "Username already exists" {
			user, err := cr.userUseCase.FindUser(newUser.UserName)
			fmt.Printf("\n\n\n%v\n%v\n\n", user.ID, err)

			token, err := cr.jwtUseCase.GenerateAccessToken(user.ID, user.UserName, "admin")
			if err != nil {
				response := response.ErrorResponse("Failed to generate access token", err.Error(), nil)
				c.Writer.Header().Add("Content-Type", "application/json")
				c.Writer.WriteHeader(http.StatusUnauthorized)
				utils.ResponseJSON(*c, response)
				return
			}
			user.AccessToken = token

			token, err = cr.jwtUseCase.GenerateRefreshToken(user.ID, user.UserName, "admin")

			if err != nil {
				response := response.ErrorResponse("Failed to generate refresh token please login again", err.Error(), nil)
				c.Writer.Header().Add("Content-Type", "application/json")
				c.Writer.WriteHeader(http.StatusUnauthorized)
				utils.ResponseJSON(*c, response)
				return
			}
			user.RefreshToken = token

			user.Password = ""
			response := response.SuccessResponse(true, "SUCCESS", user)
			c.Writer.Header().Set("Content-Type", "application/json")
			c.Writer.WriteHeader(http.StatusOK)
			utils.ResponseJSON(*c, response)

		} else {
			response := response.ErrorResponse("Failed to Login please login again", err.Error(), nil)
			c.Writer.Header().Add("Content-Type", "application/json")
			c.Writer.WriteHeader(http.StatusUnauthorized)
			utils.ResponseJSON(*c, response)
			return
		}

		c.JSON(http.StatusOK, "Hello, I'm protected\n")
		c.JSON(http.StatusOK, string(respons))
		return
	}
}

// @Summary Refresh The Access Token
// @ID Refresh access token
// @Tags Refresh Token
// @Security BearerAuth
// @Produce json
// @Success 200 {object} response.Response{}
// @Failure 422 {object} response.Response{}
// @Router /admin/refresh-tocken [get]
func (cr *AuthHandler) RefreshToken(c *gin.Context) {

	autheader := c.Request.Header["Authorization"]
	auth := strings.Join(autheader, " ")
	bearerToken := strings.Split(auth, " ")
	fmt.Printf("\n\ntocen : %v\n\n", autheader)
	token := bearerToken[1]
	ok, claims := cr.jwtUseCase.VerifyToken(token)
	if !ok {
		log.Fatal("referesh token not valid")
	}

	fmt.Println("//////////////////////////////////", claims.UserName)
	accesstoken, err := cr.jwtUseCase.GenerateAccessToken(claims.UserId, claims.UserName, claims.Role)

	if err != nil {
		response := response.ErrorResponse("Failed to generating refresh token please login again", err.Error(), nil)
		c.Writer.Header().Add("Content-Type", "application/json")
		c.Writer.WriteHeader(http.StatusUnprocessableEntity)
		utils.ResponseJSON(*c, response)
		return
	}

	response := response.SuccessResponse(true, "SUCCESS", accesstoken)
	c.Writer.Header().Add("Content-Type", "application/json")
	c.Writer.WriteHeader(http.StatusOK)
	utils.ResponseJSON(*c, response)

}

// @Summary Login for admin
// @ID admin login authentication
// @Tags Admin Authentication
// @accept json
// @Produce json
// @Param AdminLogin body domain.User{username=string,password=string} true "admin login"
// @Success 200 {object} response.Response{}
// @Failure 422 {object} response.Response{}
// @Router /admin/login [post]
func (cr *AuthHandler) AdminLogin(c *gin.Context) {
	var loginAdmin domain.User

	fmt.Print("\n\nhi\n\n")
	c.Bind(&loginAdmin)

	//verify User details
	err := cr.authUseCase.VerifyAdmin(loginAdmin.UserName, loginAdmin.Password)

	if err != nil {
		response := response.ErrorResponse("Failed to varifing Admin", err.Error(), nil)
		c.Writer.Header().Set("Content-Type", "application/json")
		c.Writer.WriteHeader(http.StatusUnprocessableEntity)

		utils.ResponseJSON(*c, response)
		return
	}

	user, err := cr.adminUseCase.FindAdmin(loginAdmin.UserName)
	// fmt.Printf("\n\n\n%v\n%v\n\n", user, err)

	token, err := cr.jwtUseCase.GenerateAccessToken(user.ID, user.Username, "admin")
	if err != nil {
		response := response.ErrorResponse("Failed to generate access token please login again", err.Error(), nil)
		c.Writer.Header().Add("Content-Type", "application/json")
		c.Writer.WriteHeader(http.StatusUnauthorized)
		utils.ResponseJSON(*c, response)
		return
	}
	user.AccessToken = token

	token, err = cr.jwtUseCase.GenerateRefreshToken(user.ID, user.Username, "admin")

	if err != nil {
		response := response.ErrorResponse("Failed to generate refresh token please login", err.Error(), nil)
		c.Writer.Header().Add("Content-Type", "application/json")
		c.Writer.WriteHeader(http.StatusUnauthorized)
		utils.ResponseJSON(*c, response)
		return
	}
	user.RefreshToken = token

	user.Password = ""
	response := response.SuccessResponse(true, "SUCCESS", user)
	c.Writer.Header().Set("Content-Type", "application/json")
	c.Writer.WriteHeader(http.StatusOK)
	utils.ResponseJSON(*c, response)
}

// @Summary SignUp for users
// @ID SignUp authentication
// @Tags User Authentication
// @Produce json
// @Tags User Authentication
// @Param WorkerLogin body domain.User{username=string,password=string} true "Worker Login"
// @Success 200 {object} response.Response{}
// @Failure 422 {object} response.Response{}
// @Router /user/signup [post]
func (cr *AuthHandler) UserSignUp(c *gin.Context) {
	var newUser domain.User
	fmt.Printf("\n\nerrrrrrr : %v\n\n", c.Bind(&newUser))

	err := c.Bind(&newUser)
	if err != nil {
		fmt.Printf("\n\nerr : %v\n\n", err)
	}
	fmt.Printf("\n\nname ; %v  %v\n\n", newUser, err)
	err = cr.userUseCase.CreateUser(newUser)

	if err != nil {
		response := response.ErrorResponse("Failed to create user", err.Error(), nil)
		c.Writer.Header().Set("Content-Type", "application/json")
		c.Writer.WriteHeader(http.StatusUnprocessableEntity)

		utils.ResponseJSON(*c, response)
		return
	}

	user, err := cr.userUseCase.FindUser(newUser.UserName)
	fmt.Printf("\n\n\n%v\n%v\n\n", user, err)
	fmt.Printf("\n\n user : %v\n\n", user)

	user.Password = ""
	response := response.SuccessResponse(true, "SUCCESS", user)
	c.Writer.Header().Set("Content-Type", "application/json")
	c.Writer.WriteHeader(http.StatusOK)
	utils.ResponseJSON(*c, response)
}

// @Summary Login for users
// @ID login authentication
// @Tags User Authentication
// @Produce json
// @Param WorkerLogin body domain.User{username=string,password=string} true "Worker Login"
// @Success 200 {object} response.Response{}
// @Failure 422 {object} response.Response{}
// @Router /user/login [post]
func (cr *AuthHandler) UserLogin(c *gin.Context) {
	var loginUser domain.User

	c.Bind(&loginUser)

	//verify User details
	err := cr.authUseCase.VerifyUser(loginUser.UserName, loginUser.Password)

	if err != nil {
		response := response.ErrorResponse("Failed to login", err.Error(), nil)
		c.Writer.Header().Set("Content-Type", "application/json")
		c.Writer.WriteHeader(http.StatusUnprocessableEntity)

		utils.ResponseJSON(*c, response)
		return
	}

	user, err := cr.userUseCase.FindUser(loginUser.UserName)
	fmt.Printf("\n\n\n%v\n%v\n\n", user.ID, err)

	token, err := cr.jwtUseCase.GenerateAccessToken(user.ID, user.UserName, "admin")
	if err != nil {
		response := response.ErrorResponse("Failed to generate access token", err.Error(), nil)
		c.Writer.Header().Add("Content-Type", "application/json")
		c.Writer.WriteHeader(http.StatusUnauthorized)
		utils.ResponseJSON(*c, response)
		return
	}
	user.AccessToken = token

	token, err = cr.jwtUseCase.GenerateRefreshToken(user.ID, user.UserName, "admin")

	if err != nil {
		response := response.ErrorResponse("Failed to generate refresh token please login again", err.Error(), nil)
		c.Writer.Header().Add("Content-Type", "application/json")
		c.Writer.WriteHeader(http.StatusUnauthorized)
		utils.ResponseJSON(*c, response)
		return
	}
	user.RefreshToken = token

	user.Password = ""
	response := response.SuccessResponse(true, "SUCCESS", user)
	c.Writer.Header().Set("Content-Type", "application/json")
	c.Writer.WriteHeader(http.StatusOK)
	utils.ResponseJSON(*c, response)
}

// @Summary SignUp for Workers
// @ID Worker SignUp authentication
// @Tags Worker Authentication
// @Produce json
// @Param WorkerSignup body domain.User{username=string,password=string} true "Worker Signup"
// @Success 200 {object} response.Response{}
// @Failure 422 {object} response.Response{}
// @Router /worker/signup [post]
func (cr *AuthHandler) WorkerSignUp(c *gin.Context) {
	var newUser domain.User

	c.Bind(&newUser)

	err := cr.workerUseCase.CreateUser(newUser)

	if err != nil {
		response := response.ErrorResponse("Failed to create worker", err.Error(), nil)
		c.Writer.Header().Set("Content-Type", "application/json")
		c.Writer.WriteHeader(http.StatusUnprocessableEntity)

		utils.ResponseJSON(*c, response)
		return
	}

	user, err := cr.workerUseCase.FindWorker(newUser.UserName)
	fmt.Printf("\n\n\n%v\n%v\n\n", user, err)
	fmt.Printf("\n\n user : %v\n\n", user)

	user.Password = ""
	response := response.SuccessResponse(true, "SUCCESS", user)
	c.Writer.Header().Set("Content-Type", "application/json")
	c.Writer.WriteHeader(http.StatusOK)
	utils.ResponseJSON(*c, response)
}

// @Summary Login for worker
// @ID worker login authentication
// @Tags Worker Authentication
// @Produce json
// @Param WorkerLogin body domain.User{username=string,password=string} true "Worker Login"
// @Success 200 {object} response.Response{}
// @Failure 422 {object} response.Response{}
// @Router /worker/login [post]
func (cr *AuthHandler) WorkerLogin(c *gin.Context) {
	var loginWorker domain.User

	c.Bind(&loginWorker)

	//verify User details
	err := cr.authUseCase.VerifyWorker(loginWorker.UserName, loginWorker.Password)

	if err != nil {
		response := response.ErrorResponse("Failed to login", err.Error(), nil)
		c.Writer.Header().Set("Content-Type", "application/json")
		c.Writer.WriteHeader(http.StatusUnprocessableEntity)

		utils.ResponseJSON(*c, response)
		return
	}

	user, err := cr.workerUseCase.FindWorker(loginWorker.UserName)
	// fmt.Printf("\n\n\n%v\n%v\n\n", user, err)

	token, err := cr.jwtUseCase.GenerateAccessToken(user.ID, user.UserName, "admin")
	if err != nil {
		response := response.ErrorResponse("Failed to generate access token please login again", err.Error(), nil)
		c.Writer.Header().Add("Content-Type", "application/json")
		c.Writer.WriteHeader(http.StatusUnauthorized)
		utils.ResponseJSON(*c, response)
		return
	}
	user.AccessToken = token

	token, err = cr.jwtUseCase.GenerateRefreshToken(user.ID, user.UserName, "admin")

	if err != nil {
		response := response.ErrorResponse("Failed to generate refresh token please login again", err.Error(), nil)
		c.Writer.Header().Add("Content-Type", "application/json")
		c.Writer.WriteHeader(http.StatusUnauthorized)
		utils.ResponseJSON(*c, response)
		return
	}
	user.RefreshToken = token

	user.Password = ""
	response := response.SuccessResponse(true, "SUCCESS", user)
	c.Writer.Header().Set("Content-Type", "application/json")
	c.Writer.WriteHeader(http.StatusOK)
	utils.ResponseJSON(*c, response)
}

// @Summary Send OTP varification mail to users
// @ID SendVerificationMail authentication
// @Tags User Authentication
// @Produce json
// @Param        email   query      string  true  "Email : "
// @Success 200 {object} response.Response{}
// @Failure 422 {object} response.Response{}
// @Router /user/send/verification [post]
func (cr *AuthHandler) SendVerificationMailUser(c *gin.Context) {
	email := c.Query("email")

	user, err := cr.userUseCase.FindUser(email)
	fmt.Printf("\n\n emailvar : %v\n%v\n", email, err)

	if err == nil {
		err = cr.authUseCase.SendVerificationEmail(email)
	}

	if err != nil {
		response := response.ErrorResponse("Error while sending verification mail to user", err.Error(), nil)
		c.Writer.Header().Set("Content-Type", "application/json")
		c.Writer.WriteHeader(http.StatusUnprocessableEntity)

		utils.ResponseJSON(*c, response)
		return
	}

	user, err = cr.userUseCase.FindUser(user.UserName)
	// fmt.Printf("\n\n\n%v\n%v\n\n", user, err)

	// token, err := cr.jwtUseCase.GenerateAccessToken(user.ID, user.UserName, "admin")
	// if err != nil {
	// 	response := response.ErrorResponse("Failed to generate access token", err.Error(), nil)
	// 	c.Writer.Header().Add("Content-Type", "application/json")
	// 	c.Writer.WriteHeader(http.StatusUnauthorized)
	// 	utils.ResponseJSON(*c, response)
	// 	return
	// }
	// user.AccessToken = token

	// token, err = cr.jwtUseCase.GenerateRefreshToken(user.ID, user.UserName, "admin")

	// if err != nil {
	// 	response := response.ErrorResponse("Failed to generate refresh token", err.Error(), nil)
	// 	c.Writer.Header().Add("Content-Type", "application/json")
	// 	c.Writer.WriteHeader(http.StatusUnauthorized)
	// 	utils.ResponseJSON(*c, response)
	// 	return
	// }
	// user.RefreshToken = token

	user.Password = ""
	response := response.SuccessResponse(true, "SUCCESS", user)
	c.Writer.Header().Set("Content-Type", "application/json")
	c.Writer.WriteHeader(http.StatusOK)
	utils.ResponseJSON(*c, response)
}

// @Summary Varify OTP of users
// @ID Varify OTP authentication
// @Tags User Authentication
// @Produce json
// @Param        email   query      string  true  "Email : "
// @Param        code   query      string  true  "OTP : "
// @Success 200 {object} response.Response{}
// @Failure 422 {object} response.Response{}
// @Router /user/verify/account [post]
func (cr *AuthHandler) UserVerifyAccount(c *gin.Context) {

	fmt.Println("ggggggghggggggggggggggggggggggggggg")
	// email := c.Query("email")
	tokenString := c.Query("token")
	fmt.Printf("\n\ntoken   :   %v\n\n", tokenString)
	fmt.Println("varify account from authhandler called , ", tokenString)
	var email string
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		return []byte("secret"), nil
	})
	if err != nil {
		c.String(http.StatusBadRequest, "Invalid verification token")
		return
	}
	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		// get the username from the claims
		email = claims["username"].(string)

	} else {
		c.String(http.StatusBadRequest, "Invalid verification token")
		return
	}

	err = cr.authUseCase.UserVerifyAccount(email, tokenString)

	if err != nil {
		response := response.ErrorResponse("Error while verifing user mail", err.Error(), nil)
		c.Writer.Header().Set("Content-Type", "application/json")
		c.Writer.WriteHeader(http.StatusUnprocessableEntity)

		utils.ResponseJSON(*c, response)
		return
	}

	response := response.SuccessResponse(true, "SUCCESS", email)
	c.Writer.Header().Set("Content-Type", "application/json")
	c.Writer.WriteHeader(http.StatusOK)
	utils.ResponseJSON(*c, response)
}

// @Summary Send OTP varification mail to worker
// @ID Worker SendVerificationMail authentication
// @Tags Worker Authentication
// @Produce json
// @Param        email   query      string  true  "Email : "
// @Success 200 {object} response.Response{}
// @Failure 422 {object} response.Response{}
// @Router /worker/send/verification [post]
func (cr *AuthHandler) SendVerificationMailWorker(c *gin.Context) {
	email := c.Query("email")

	user, err := cr.workerUseCase.FindWorker(email)
	fmt.Printf("\n\n emailvar : %v\n\n", email)

	if err == nil {
		err = cr.authUseCase.SendVerificationEmail(email)
	}

	if err != nil {
		response := response.ErrorResponse("Error while sending verification mail to worker", err.Error(), nil)
		c.Writer.Header().Set("Content-Type", "application/json")
		c.Writer.WriteHeader(http.StatusUnprocessableEntity)

		utils.ResponseJSON(*c, response)
		return
	}

	user, err = cr.workerUseCase.FindWorker(user.UserName)
	// fmt.Printf("\n\n\n%v\n%v\n\n", user, err)

	// token, err := cr.jwtUseCase.GenerateAccessToken(user.ID, user.UserName, "admin")
	// if err != nil {
	// 	response := response.ErrorResponse("Failed to generate access token", err.Error(), nil)
	// 	c.Writer.Header().Add("Content-Type", "application/json")
	// 	c.Writer.WriteHeader(http.StatusUnauthorized)
	// 	utils.ResponseJSON(*c, response)
	// 	return
	// }
	// user.AccessToken = token

	// token, err = cr.jwtUseCase.GenerateRefreshToken(user.ID, user.UserName, "admin")

	// if err != nil {
	// 	response := response.ErrorResponse("Failed to generate refresh token", err.Error(), nil)
	// 	c.Writer.Header().Add("Content-Type", "application/json")
	// 	c.Writer.WriteHeader(http.StatusUnauthorized)
	// 	utils.ResponseJSON(*c, response)
	// 	return
	// }
	// user.RefreshToken = token

	user.Password = ""
	response := response.SuccessResponse(true, "SUCCESS", user)
	c.Writer.Header().Set("Content-Type", "application/json")
	c.Writer.WriteHeader(http.StatusOK)
	utils.ResponseJSON(*c, response)
}

// @Summary Varify OTP of users
// @ID Varify worker OTP authentication
// @Tags Worker Authentication
// @Produce json
// @Param        email   query      string  true  "Email : "
// @Param        code   query      string  true  "OTP : "
// @Success 200 {object} response.Response{}
// @Failure 422 {object} response.Response{}
// @Router /worker/verify/account [post]
func (cr *AuthHandler) WorkerVerifyAccount(c *gin.Context) {
	email := c.Query("email")
	code, _ := strconv.Atoi(c.Query("code"))

	err := cr.authUseCase.WorkerVerifyAccount(email, code)

	if err != nil {
		response := response.ErrorResponse("Error while verifing worker mail", err.Error(), nil)
		c.Writer.Header().Set("Content-Type", "application/json")
		c.Writer.WriteHeader(http.StatusUnprocessableEntity)

		utils.ResponseJSON(*c, response)
		return
	}

	response := response.SuccessResponse(true, "SUCCESS", email)
	c.Writer.Header().Set("Content-Type", "application/json")
	c.Writer.WriteHeader(http.StatusOK)
	utils.ResponseJSON(*c, response)
}

// @Summary Varify JWT of users
// @ID Varify JWT authentication
// @Tags User
// @Security BearerAuth
// @param Authorization header string true "Authorization"
// @Produce json
// @Param        email   query      string  true  "Email : "
// @Success 200 {object} response.Response{}
// @Failure 422 {object} response.Response{}
// @Router /user/account/verifyJWT [get]
func (cr *AuthHandler) UserHome(c *gin.Context) {
	email := c.Query("email")

	response := response.SuccessResponse(true, "Welcome to user home", email)
	c.Writer.Header().Set("Content-Type", "application/json")
	c.Writer.WriteHeader(http.StatusOK)
	utils.ResponseJSON(*c, response)
}

// @Summary Varify JWT of users
// @ID Varify worker JWT authentication
// @Tags Worker
// @Security BearerAuth
// @Produce json
// @Param        email   query      string  true  "Email : "
// @Success 200 {object} response.Response{}
// @Failure 422 {object} response.Response{}
// @Router /worker/account/verifyJWT [get]
func (cr *AuthHandler) WorkerHome(c *gin.Context) {
	email := c.Query("email")

	response := response.SuccessResponse(true, "Welcome eorker home", email)
	c.Writer.Header().Set("Content-Type", "application/json")
	c.Writer.WriteHeader(http.StatusOK)
	utils.ResponseJSON(*c, response)
}

// @Summary Varify JWT of Admin
// @ID Varify admin JWT authentication
// @Tags Admin
// @Produce json
// @Param        email   query      string  true  "Email : "
// @Success 200 {object} response.Response{}
// @Failure 422 {object} response.Response{}
// @Router /admin/account/verifyJWT [get]
func (cr *AuthHandler) AdminHome(c *gin.Context) {
	email := c.Query("email")

	response := response.SuccessResponse(true, "Welcome admin home", email)
	c.Writer.Header().Set("Content-Type", "application/json")
	c.Writer.WriteHeader(http.StatusOK)
	utils.ResponseJSON(*c, response)
}
