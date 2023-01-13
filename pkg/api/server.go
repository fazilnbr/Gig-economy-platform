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

			admin.GET("/list-all-users", adminHandler.ListAllUsers)
			admin.GET("/list-new-users", adminHandler.ListNewUsers)
			admin.GET("/list-blockedusers", adminHandler.ListBlockUsers)
			admin.PATCH("/activate-users", adminHandler.ActivateUsers)
			admin.PATCH("/block-users", adminHandler.BlockUsers)

			// Worker management

			admin.GET("/list-all-workers", adminHandler.ListAllWorkers)
			admin.GET("/list-new-workers", adminHandler.ListNewWorkers)
			admin.GET("/list-blocked-workers", adminHandler.ListBlockWorkers)
			admin.PATCH("/activate-workers", adminHandler.ActivateWorkers)
			admin.PATCH("/block-workers", adminHandler.BlockWorkers)

			// Job management
			admin.POST("/add-job-category", adminHandler.AddJobCategory)
			admin.GET("/list-job-category", adminHandler.ListJobCategory)
			admin.PATCH("/update-job-category", adminHandler.UpdateJobCategory)
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
				user.POST("/add-profile", UserHandler.UserAddProfile)
				user.PATCH("/edit-profile", UserHandler.UserEditProfile)
				user.PATCH("/change-password", UserHandler.UserChangePassword)
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

				// user profile
				worker.POST("/add-profile", WorkerHandler.WorkerAddProfile)
				worker.PATCH("/edit-profile", WorkerHandler.WorkerEditProfile)
				worker.PATCH("/change-password", WorkerHandler.WorkerChangePassword)

				// job management
				worker.GET("/list-job-category", WorkerHandler.ListJobCategoryUser)
				worker.POST("/add-job", WorkerHandler.AddJob)
				worker.GET("/view-job", WorkerHandler.ViewJob)
				worker.DELETE("/delete-job", WorkerHandler.DeleteJob)
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
