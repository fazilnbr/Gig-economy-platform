package middleware

import (
	"errors"
	"fmt"
	"net/http"
	"strings"

	"github.com/fazilnbr/project-workey/pkg/common/response"
	service "github.com/fazilnbr/project-workey/pkg/usecase/interface"
	"github.com/fazilnbr/project-workey/pkg/utils"
	"github.com/gin-gonic/gin"
)

type Middleware interface {
	AthoriseJWT(*gin.Context)
}

type middlewar struct {
	jwtUseCase service.JWTUseCase
}

// AthoriseJWT implements Middileware

func (cr *middlewar) AthoriseJWT(c *gin.Context) {
	//getting from header
	autheader := c.Request.Header["Authorization"]
	auth := strings.Join(autheader, " ")
	bearerToken := strings.Split(auth, " ")

	// fmt.Print("\n\n\nff\n\n\ntoken : ", len(bearerToken), ":  :", bearerToken)

	if len(bearerToken) != 2 {
		err := errors.New("request does not contain an access token")
		response := response.ErrorResponse("Failed to create user", err.Error(), nil)
		c.Writer.Header().Set("Content-Type", "application/json")
		c.Writer.WriteHeader(http.StatusUnauthorized)

		utils.ResponseJSON(*c, response)
		c.Abort()
		return
	}

	authtoken := bearerToken[1]
	// fmt.Print(authtoken)
	ok, claims := cr.jwtUseCase.VerifyToken(authtoken)

	fmt.Printf("\n\nok : %v\n\n", ok)

	if !ok {
		err := errors.New("your token is not valid")
		response := response.ErrorResponse("Failed to create user", err.Error(), nil)
		c.Writer.Header().Set("Content-Type", "application/json")
		c.Writer.WriteHeader(http.StatusUnauthorized)

		utils.ResponseJSON(*c, response)
		c.Abort()
		return
	}

	user_email := fmt.Sprint(claims.Username)
	id := fmt.Sprint(claims.User_Id)
	fmt.Printf("\n\nid : %v\n\n", id)
	// r.Header.Set("email", user_email)
	c.Writer.Header().Set("email", user_email)
	// r.Header.Set("id", id)
	c.Writer.Header().Set("id", id)
	// c.Next()
}

func NewUserMiddileware(jwtUserUseCase service.JWTUseCase) Middleware {
	return &middlewar{
		jwtUseCase: jwtUserUseCase,
	}
}
func NewWorkerMiddileware(jwtWorkerUsecase service.JWTUseCase) Middleware {
	return &middlewar{
		jwtUseCase: jwtWorkerUsecase,
	}
}
func NewAdminMiddileware(jwtAdminUseCase service.JWTUseCase) Middleware {
	return &middlewar{
		jwtUseCase: jwtAdminUseCase,
	}
}
