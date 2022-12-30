package middleware

import (
	"fmt"
	"net/http"

	"github.com/fazilnbr/project-workey/pkg/common/response"
	"github.com/fazilnbr/project-workey/pkg/domain"
	services "github.com/fazilnbr/project-workey/pkg/usecase/interface"
	"github.com/fazilnbr/project-workey/pkg/utils"
	"github.com/gin-gonic/gin"
)

type UserHandler struct {
	userUseCase services.UserUseCase
	jwtUseCase  services.JWTUseCase
}

func NewUserHandler(
	usecase services.UserUseCase,
	jwtUseCase services.JWTUseCase,

) *UserHandler {
	return &UserHandler{
		userUseCase: usecase,
		jwtUseCase:  jwtUseCase,
	}
}

// @Summary SignUp for users
// @ID SignUp authentication
// @Produce json
// @Param        username   path      string  true  "User Name : "
// @Param        password   path      string  true  "Password : "
// @Success 200 {object} response.Response{Status=bool,Message=string,Errors=string,Data=domain.Login}
// @Failure 422 {object} response.Response{Status=bool,Message=string,Errors=string,Data=string}
// @Router /user/signup [post]
func (cr *UserHandler) UserSignUp(c *gin.Context) {
	var newUser domain.Login

	c.Bind(&newUser)
	fmt.Printf("\n\n user : %v\n\n", newUser)
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
// @Success 200 {object} response.Response{Status=bool,Message=string,Errors=string,Data=domain.Login}
// @Failure 422 {object} response.Response{Status=bool,Message=string,Errors=string,Data=string}
// @Router /user/login [post]
func (cr *UserHandler) UserLogin(c *gin.Context) {
	var loginUser domain.Login

	c.Bind(&loginUser)

	//verify User details
	err := cr.userUseCase.VerifyUser(loginUser.UserName, loginUser.Password)

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
