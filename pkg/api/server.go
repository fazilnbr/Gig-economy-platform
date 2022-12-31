package api

import (
	"fmt"
	"log"

	"github.com/fazilnbr/project-workey/pkg/api/middleware"
	"github.com/gin-gonic/gin"
	swaggerfiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

type ServerHTTP struct {
	engine *gin.Engine
}

func NewServerHTTP(authHandler *middleware.AuthHandler, middleware middleware.Middleware) *ServerHTTP {
	engine := gin.New()

	// Use logger from Gin
	engine.Use(gin.Logger())

	// Swagger docs
	engine.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerfiles.Handler))

	// Request JWT
	user := engine.Group("user")
	{
		user.POST("/signup", authHandler.UserSignUp)
		user.POST("/login", authHandler.UserLogin)
		user.POST("/send/verification", authHandler.SendVerificationMail)
		user.POST("/verify/account", authHandler.VerifyAccount)
		// authuser := user.Group("/")
		user.Use(middleware.AthoriseJWT)
		{
			user.GET("/account/verifyJWT", authHandler.UserHome)
		}
	}

	return &ServerHTTP{engine: engine}
}

func (sh *ServerHTTP) Start() {
	fmt.Print("\n\nddddddddd\n\n")
	err := sh.engine.Run(":8080")
	if err != nil {
		log.Fatalln(err)
	}
}
