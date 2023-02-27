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
	authHandler.InitializeOAuthGoogle()
	engine := gin.New()
	engine.LoadHTMLGlob("views/*.html")

	// Use logger from Gin
	engine.Use(gin.Logger())

	// Swagger docs
	engine.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerfiles.Handler))

	admin := engine.Group("admin")
	{
		// Authentication
		admin.POST("/login", authHandler.AdminLogin)
		admin.POST("/send/verification", authHandler.SendVerificationMailWorker)
		admin.POST("/verify/account", authHandler.WorkerVerifyAccount)

		admin.Use(middleware.AthoriseJWT)
		{
			admin.GET("/account/verifyJWT", authHandler.AdminHome)

			// Refresh Token
			admin.GET("/refresh-tocken", authHandler.RefreshToken)

			// User Management

			admin.GET("/list-all-users", adminHandler.ListAllUsers)
			admin.GET("/list-new-users", adminHandler.ListNewUsers)
			admin.GET("/list-blockedusers", adminHandler.ListBlockUsers)
			admin.PATCH("/activate-users", adminHandler.ActivateUsers)
			admin.PATCH("/block-users", adminHandler.BlockUsers)

			// Worker Management

			admin.GET("/list-all-workers", adminHandler.ListAllWorkers)
			admin.GET("/list-new-workers", adminHandler.ListNewWorkers)
			admin.GET("/list-blocked-workers", adminHandler.ListBlockWorkers)
			admin.PATCH("/activate-workers", adminHandler.ActivateWorkers)
			admin.PATCH("/block-workers", adminHandler.BlockWorkers)

			// Job Management
			admin.POST("/add-job-category", adminHandler.AddJobCategory)
			admin.GET("/list-job-category", adminHandler.ListJobCategory)
			admin.PATCH("/update-job-category", adminHandler.UpdateJobCategory)
		}

		// Request JWT
		user := engine.Group("user")
		{
			// User Authentication
			user.POST("/signup", authHandler.UserSignUp)
			user.POST("/login", authHandler.UserLogin)
			user.POST("/send/verification", authHandler.SendVerificationMailUser)
			user.GET("/verify/account", authHandler.UserVerifyAccount)

			// Google authentication
			user.GET("/login-gl", authHandler.GoogleAuth)
			user.GET("/callback-gl", authHandler.CallBackFromGoogle)
			// authuser := user.Group("/")

			// Job Payment test
				// Razor-pay
				user.GET("/razor-pay-home", UserHandler.RazorPayHome)
				user.GET("/razor-pay-payment-success", UserHandler.RazorPaySuccess)


			user.Use(middleware.AthoriseJWT)
			{
				user.GET("/account/verifyJWT", authHandler.UserHome)

				// User Profile
				user.POST("/add-profile", UserHandler.UserAddProfile)
				user.PATCH("/edit-profile", UserHandler.UserEditProfile)
				user.PATCH("/change-password", UserHandler.UserChangePassword)

				// List Worker
				user.GET("/list-workers-with-job", UserHandler.ListWorkersWithJob)
				user.GET("/search-workers-with-job", UserHandler.SearchWorkersWithJob)

				// Favorite
				user.POST("/add-to-favorite", UserHandler.UserAddToFavorite)
				user.GET("/list-favorite-list", UserHandler.ListFavorite)

				// User Address
				user.POST("add-address", UserHandler.UserAddAddress)
				user.GET("/list-address", UserHandler.UserListAddress)
				user.DELETE("/delete-address", UserHandler.DeleteAddress)

				// Job Request
				user.POST("/send-job-request", UserHandler.UserSendJobRequest)
				user.DELETE("/cancel-job-request", UserHandler.DeleteJobRequest)
				user.GET("/list-job-request", UserHandler.ListSendRequests)
				user.GET("/view-one-job-request", UserHandler.ViewSendOneRequest)
				user.PATCH("/update-job-complition-status", UserHandler.UpdateJobComplition)

				// Job Payment
				// Razor-pay
				// user.GET("/razor-pay-home",UserHandler.RazorPayHome)
			}
		}

		worker := engine.Group("worker")
		{
			//Worker Authentication
			worker.POST("/signup", authHandler.WorkerSignUp)
			worker.POST("/login", authHandler.WorkerLogin)
			worker.POST("/send/verification", authHandler.SendVerificationMailWorker)
			worker.POST("/verify/account", authHandler.WorkerVerifyAccount)
			// authuser := user.Group("/")
			worker.Use(middleware.AthoriseJWT)
			{
				worker.GET("/account/verifyJWT", authHandler.WorkerHome)

				// Worker Profile
				worker.POST("/add-profile", WorkerHandler.WorkerAddProfile)
				worker.PATCH("/edit-profile", WorkerHandler.WorkerEditProfile)
				worker.PATCH("/change-password", WorkerHandler.WorkerChangePassword)

				// job management
				worker.GET("/list-job-category", WorkerHandler.ListJobCategoryUser)
				worker.POST("/add-job", WorkerHandler.AddJob)
				worker.GET("/view-job", WorkerHandler.ViewJob)
				worker.DELETE("/delete-job", WorkerHandler.DeleteJob)

				// User Job Request
				worker.GET("/list-user-pending-job-request", WorkerHandler.ListPendingJobRequsetFromUser)
				worker.GET("/list-user-accepted-job-request", WorkerHandler.ListAcceptedJobRequsetFromUser)
				worker.PATCH("/accept-job-request", WorkerHandler.AcceptJobRequest)
				worker.PATCH("/reject-job-request", WorkerHandler.RejectJobRequest)
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
