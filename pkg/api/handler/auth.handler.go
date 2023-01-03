package handler

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/fazilnbr/project-workey/pkg/common/response"
	"github.com/fazilnbr/project-workey/pkg/domain"
	services "github.com/fazilnbr/project-workey/pkg/usecase/interface"
	"github.com/fazilnbr/project-workey/pkg/utils"
	"github.com/gin-gonic/gin"
)

type AuthHandler struct {
	adminUseCase  services.AdminUseCase
	workerUseCase services.WorkerUseCase
	userUseCase   services.UserUseCase
	jwtUseCase    services.JWTUseCase
	authUseCase   services.AuthUseCase
}

func NewAuthHandler(
	adminUseCase services.AdminUseCase,
	workerUseCase services.WorkerUseCase,
	userusecase services.UserUseCase,
	jwtUseCase services.JWTUseCase,
	authUseCase services.AuthUseCase,

) AuthHandler {
	return AuthHandler{
		adminUseCase:  adminUseCase,
		workerUseCase: workerUseCase,
		userUseCase:   userusecase,
		jwtUseCase:    jwtUseCase,
		authUseCase:   authUseCase,
	}
}

// @Summary Login for admin
// @ID admin login authentication
// @Produce json
// @Param        username   path      string  true  "User Name : "
// @Param        password   path      string  true  "Password : "
// @Success 200 {object} response.Response{Status=bool,Message=string,Errors=string,Data=domain.UserResponse}
// @Failure 422 {object} response.Response{Status=bool,Message=string,Errors=string,Data=string}
// @Router /admin/login [post]
func (cr *AuthHandler) AdminLogin(c *gin.Context) {
	var loginAdmin domain.Login

	fmt.Print("\n\nhi\n\n")
	c.Bind(&loginAdmin)

	//verify User details
	err := cr.authUseCase.VerifyAdmin(loginAdmin.UserName, loginAdmin.Password)

	if err != nil {
		response := response.ErrorResponse("Failed to create user", err.Error(), nil)
		c.Writer.Header().Set("Content-Type", "application/json")
		c.Writer.WriteHeader(http.StatusUnprocessableEntity)

		utils.ResponseJSON(*c, response)
		return
	}

	user, err := cr.adminUseCase.FindAdmin(loginAdmin.UserName)
	// fmt.Printf("\n\n\n%v\n%v\n\n", user, err)

	token := cr.jwtUseCase.GenerateToken(user.ID, user.Username, "user")
	user.Password = ""
	user.Token = token
	response := response.SuccessResponse(true, "SUCCESS", user)
	c.Writer.Header().Set("Content-Type", "application/json")
	c.Writer.WriteHeader(http.StatusOK)
	utils.ResponseJSON(*c, response)
}

// @Summary SignUp for users
// @ID SignUp authentication
// @Produce json
// @Param        username   path      string  true  "User Name : "
// @Param        password   path      string  true  "Password : "
// @Success 200 {object} response.Response{Status=bool,Message=string,Errors=string,Data=domain.Login}
// @Failure 422 {object} response.Response{Status=bool,Message=string,Errors=string,Data=string}
// @Router /user/signup [post]
func (cr *AuthHandler) UserSignUp(c *gin.Context) {
	var newUser domain.Login

	c.Bind(&newUser)

	err := cr.userUseCase.CreateUser(newUser)

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
// @Produce json
// @Param        username   path      string  true  "User Name : "
// @Param        password   path      string  true  "Password : "
// @Success 200 {object} response.Response{Status=bool,Message=string,Errors=string,Data=domain.UserResponse}
// @Failure 422 {object} response.Response{Status=bool,Message=string,Errors=string,Data=string}
// @Router /user/login [post]
func (cr *AuthHandler) UserLogin(c *gin.Context) {
	var loginUser domain.Login

	c.Bind(&loginUser)

	//verify User details
	err := cr.authUseCase.VerifyUser(loginUser.UserName, loginUser.Password)

	if err != nil {
		response := response.ErrorResponse("Failed to create user", err.Error(), nil)
		c.Writer.Header().Set("Content-Type", "application/json")
		c.Writer.WriteHeader(http.StatusUnprocessableEntity)

		utils.ResponseJSON(*c, response)
		return
	}

	user, err := cr.userUseCase.FindUser(loginUser.UserName)
	// fmt.Printf("\n\n\n%v\n%v\n\n", user, err)

	token := cr.jwtUseCase.GenerateToken(user.ID, user.UserName, "user")
	user.Password = ""
	user.Token = token
	response := response.SuccessResponse(true, "SUCCESS", user)
	c.Writer.Header().Set("Content-Type", "application/json")
	c.Writer.WriteHeader(http.StatusOK)
	utils.ResponseJSON(*c, response)
}

// @Summary SignUp for Workers
// @ID Worker SignUp authentication
// @Produce json
// @Param        username   path      string  true  "User Name : "
// @Param        password   path      string  true  "Password : "
// @Success 200 {object} response.Response{Status=bool,Message=string,Errors=string,Data=domain.Login}
// @Failure 422 {object} response.Response{Status=bool,Message=string,Errors=string,Data=string}
// @Router /woker/signup [post]
func (cr *AuthHandler) WorkerSignUp(c *gin.Context) {
	var newUser domain.Login

	c.Bind(&newUser)

	err := cr.workerUseCase.CreateUser(newUser)

	if err != nil {
		response := response.ErrorResponse("Failed to create user", err.Error(), nil)
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
// @Produce json
// @Param        username   path      string  true  "User Name : "
// @Param        password   path      string  true  "Password : "
// @Success 200 {object} response.Response{Status=bool,Message=string,Errors=string,Data=domain.UserResponse}
// @Failure 422 {object} response.Response{Status=bool,Message=string,Errors=string,Data=string}
// @Router /worker/login [post]
func (cr *AuthHandler) WorkerLogin(c *gin.Context) {
	var loginWorker domain.Login

	c.Bind(&loginWorker)

	//verify User details
	err := cr.authUseCase.VerifyWorker(loginWorker.UserName, loginWorker.Password)

	if err != nil {
		response := response.ErrorResponse("Failed to create user", err.Error(), nil)
		c.Writer.Header().Set("Content-Type", "application/json")
		c.Writer.WriteHeader(http.StatusUnprocessableEntity)

		utils.ResponseJSON(*c, response)
		return
	}

	user, err := cr.workerUseCase.FindWorker(loginWorker.UserName)
	// fmt.Printf("\n\n\n%v\n%v\n\n", user, err)

	token := cr.jwtUseCase.GenerateToken(user.ID, user.UserName, "user")
	user.Password = ""
	user.Token = token
	response := response.SuccessResponse(true, "SUCCESS", user)
	c.Writer.Header().Set("Content-Type", "application/json")
	c.Writer.WriteHeader(http.StatusOK)
	utils.ResponseJSON(*c, response)
}

// @Summary Send OTP varification mail to users
// @ID SendVerificationMail authentication
// @Produce json
// @Param        email   path      string  true  "Email : "
// @Success 200 {object} response.Response{Status=bool,Message=string,Errors=string,Data=domain.UserResponse}
// @Failure 422 {object} response.Response{Status=bool,Message=string,Errors=string,Data=string}
// @Router /user/send/verification [post]
func (cr *AuthHandler) SendVerificationMailUser(c *gin.Context) {
	email := c.Query("email")

	user, err := cr.userUseCase.FindUser(email)
	fmt.Printf("\n\n emailvar : %v\n\n", email)

	if err == nil {
		err = cr.authUseCase.SendVerificationEmail(email)
	}

	if err != nil {
		response := response.ErrorResponse("Error while sending verification mail", err.Error(), nil)
		c.Writer.Header().Set("Content-Type", "application/json")
		c.Writer.WriteHeader(http.StatusUnprocessableEntity)

		utils.ResponseJSON(*c, response)
		return
	}

	user, err = cr.userUseCase.FindUser(user.UserName)
	// fmt.Printf("\n\n\n%v\n%v\n\n", user, err)

	token := cr.jwtUseCase.GenerateToken(user.ID, user.UserName, "user")
	user.Password = ""
	user.Token = token
	response := response.SuccessResponse(true, "SUCCESS", user)
	c.Writer.Header().Set("Content-Type", "application/json")
	c.Writer.WriteHeader(http.StatusOK)
	utils.ResponseJSON(*c, response)
}

// @Summary Varify OTP of users
// @ID Varify OTP authentication
// @Produce json
// @Param        email   path      string  true  "Email : "
// @Param        code   path      string  true  "OTP : "
// @Success 200 {object} response.Response{Status=bool,Message=string,Errors=string,Data=string}
// @Failure 422 {object} response.Response{Status=bool,Message=string,Errors=string,Data=string}
// @Router /user/verify/account [post]
func (cr *AuthHandler) UserVerifyAccount(c *gin.Context) {
	email := c.Query("email")
	code, _ := strconv.Atoi(c.Query("code"))

	err := cr.authUseCase.UserVerifyAccount(email, code)

	if err != nil {
		response := response.ErrorResponse("Error while sending verification mail", err.Error(), nil)
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
// @Produce json
// @Param        email   path      string  true  "Email : "
// @Success 200 {object} response.Response{Status=bool,Message=string,Errors=string,Data=domain.UserResponse}
// @Failure 422 {object} response.Response{Status=bool,Message=string,Errors=string,Data=string}
// @Router /worker/send/verification [post]
func (cr *AuthHandler) SendVerificationMailWorker(c *gin.Context) {
	email := c.Query("email")

	user, err := cr.workerUseCase.FindWorker(email)
	fmt.Printf("\n\n emailvar : %v\n\n", email)

	if err == nil {
		err = cr.authUseCase.SendVerificationEmail(email)
	}

	if err != nil {
		response := response.ErrorResponse("Error while sending verification mail", err.Error(), nil)
		c.Writer.Header().Set("Content-Type", "application/json")
		c.Writer.WriteHeader(http.StatusUnprocessableEntity)

		utils.ResponseJSON(*c, response)
		return
	}

	user, err = cr.workerUseCase.FindWorker(user.UserName)
	// fmt.Printf("\n\n\n%v\n%v\n\n", user, err)

	token := cr.jwtUseCase.GenerateToken(user.ID, user.UserName, "user")
	user.Password = ""
	user.Token = token
	response := response.SuccessResponse(true, "SUCCESS", user)
	c.Writer.Header().Set("Content-Type", "application/json")
	c.Writer.WriteHeader(http.StatusOK)
	utils.ResponseJSON(*c, response)
}

// @Summary Varify OTP of users
// @ID Varify worker OTP authentication
// @Produce json
// @Param        email   path      string  true  "Email : "
// @Param        code   path      string  true  "OTP : "
// @Success 200 {object} response.Response{Status=bool,Message=string,Errors=string,Data=string}
// @Failure 422 {object} response.Response{Status=bool,Message=string,Errors=string,Data=string}
// @Router /worker/verify/account [post]
func (cr *AuthHandler) WorkerVerifyAccount(c *gin.Context) {
	email := c.Query("email")
	code, _ := strconv.Atoi(c.Query("code"))

	err := cr.authUseCase.WorkerVerifyAccount(email, code)

	if err != nil {
		response := response.ErrorResponse("Error while sending verification mail", err.Error(), nil)
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
// @Produce json
// @Param        email   path      string  true  "Email : "
// @Success 200 {object} response.Response{Status=bool,Message=string,Errors=string,Data=string}
// @Failure 422 {object} response.Response{Status=bool,Message=string,Errors=string,Data=string}
// @Router /user/account/verifyJWT [get]
func (cr *AuthHandler) UserHome(c *gin.Context) {
	email := c.Query("email")

	response := response.SuccessResponse(true, "welcome home", email)
	c.Writer.Header().Set("Content-Type", "application/json")
	c.Writer.WriteHeader(http.StatusOK)
	utils.ResponseJSON(*c, response)
}

// @Summary Varify JWT of users
// @ID Varify worker JWT authentication
// @Produce json
// @Param        email   path      string  true  "Email : "
// @Success 200 {object} response.Response{Status=bool,Message=string,Errors=string,Data=string}
// @Failure 422 {object} response.Response{Status=bool,Message=string,Errors=string,Data=string}
// @Router /worker/account/verifyJWT [get]
func (cr *AuthHandler) WorkerHome(c *gin.Context) {
	email := c.Query("email")

	response := response.SuccessResponse(true, "welcome home", email)
	c.Writer.Header().Set("Content-Type", "application/json")
	c.Writer.WriteHeader(http.StatusOK)
	utils.ResponseJSON(*c, response)
}

// @Summary Varify JWT of users
// @ID Varify admin JWT authentication
// @Produce json
// @Param        email   path      string  true  "Email : "
// @Success 200 {object} response.Response{Status=bool,Message=string,Errors=string,Data=string}
// @Failure 422 {object} response.Response{Status=bool,Message=string,Errors=string,Data=string}
// @Router /admin/account/verifyJWT [get]
func (cr *AuthHandler) AdminHome(c *gin.Context) {
	email := c.Query("email")

	response := response.SuccessResponse(true, "welcome home", email)
	c.Writer.Header().Set("Content-Type", "application/json")
	c.Writer.WriteHeader(http.StatusOK)
	utils.ResponseJSON(*c, response)
}
