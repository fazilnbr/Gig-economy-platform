package api

import (
	"fmt"
	"log"

	"github.com/fazilnbr/project-workey/pkg/api/handler"
	"github.com/fazilnbr/project-workey/pkg/api/middleware"
	"github.com/gin-gonic/gin"
	swaggerfiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

type ServerHTTP struct {
	engine *gin.Engine
}

func NewServerHTTP(authHandler handler.AuthHandler, adminHandler handler.AdminHandler, UserHandler handler.UserHandler, WorkerHandler handler.WorkerHandler, middleware middleware.Middleware) *ServerHTTP {
	engine := gin.New()

	// Use logger from Gin
	engine.Use(gin.Logger())

	// Swagger docs
	engine.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerfiles.Handler))

	admin := engine.Group("admin")
	{
		admin.POST("/login", authHandler.AdminLogin)
		admin.POST("/send/verification", authHandler.SendVerificationMailWorker)
		admin.POST("/verify/account", authHandler.WorkerVerifyAccount)

		admin.Use(middleware.AthoriseJWT)
		{
			admin.GET("/account/verifyJWT", authHandler.AdminHome)

			admin.GET("/refresh-tocken", authHandler.RefreshToken)

			// User management

			admin.GET("/listallusers", adminHandler.ListAllUsers)
			admin.GET("/listnewusers", adminHandler.ListNewUsers)
			admin.GET("/listblockedusers", adminHandler.ListBlockUsers)
			admin.PUT("/activateuser", adminHandler.ActivateUsers)
			admin.PUT("/blockusers", adminHandler.BlockUsers)

			// Worker management

			admin.GET("/listallworkers", adminHandler.ListAllWorkers)
			admin.GET("/listnewworkers", adminHandler.ListNewWorkers)
			admin.GET("/listblockedworkers", adminHandler.ListBlockWorkers)
			admin.PATCH("/activateworkers", adminHandler.ActivateWorkers)
			admin.PATCH("/blockworkers", adminHandler.BlockWorkers)

			// Job management
			admin.POST("/addjobcategory", adminHandler.AddJobCategory)
			admin.GET("/listjobcategory", adminHandler.ListJobCategory)
			admin.PATCH("/updatejobcategory", adminHandler.UpdateJobCategory)
		}

		// Request JWT
		user := engine.Group("user")
		{
			user.POST("/signup", authHandler.UserSignUp)
			user.POST("/login", authHandler.UserLogin)
			user.POST("/send/verification", authHandler.SendVerificationMailUser)
			user.POST("/verify/account", authHandler.UserVerifyAccount)
			// authuser := user.Group("/")
			user.Use(middleware.AthoriseJWT)
			{
				user.GET("/account/verifyJWT", authHandler.UserHome)
				user.POST("/addprofile", UserHandler.UserAddProfile)
				user.PATCH("/editprofile", UserHandler.UserEditProfile)
				user.PATCH("/changepassword", UserHandler.UserChangePassword)
			}
		}

		worker := engine.Group("worker")
		{
			worker.POST("/signup", authHandler.WorkerSignUp)
			worker.POST("/login", authHandler.WorkerLogin)
			worker.POST("/send/verification", authHandler.SendVerificationMailWorker)
			worker.POST("/verify/account", authHandler.WorkerVerifyAccount)
			// authuser := user.Group("/")
			worker.Use(middleware.AthoriseJWT)
			{
				worker.GET("/account/verifyJWT", authHandler.WorkerHome)
				worker.POST("/addprofile", WorkerHandler.WorkerAddProfile)
				worker.PATCH("/editprofile", WorkerHandler.WorkerEditProfile)
				worker.PATCH("/changepassword", WorkerHandler.WorkerChangePassword)
			}
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
