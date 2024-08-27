package router

import (
	loancontroller "loanTracker/delivery/controller/loan_controller"
	"loanTracker/delivery/controller/usercontroller" // Add this import statement
	"loanTracker/delivery/middleware"
	"loanTracker/repository"
	"loanTracker/usecase/userusecases" // Add this import statement
	"loanTracker/usecase/loanusecases" // Updated import path
	"time"

	"github.com/gin-contrib/cors"
	"github.com/gin-contrib/secure"
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/mongo"
)
func getLoanController(database *mongo.Database) *loancontroller.LoanController {
	loanRepository := repository.NewLoanRepository(database)
	loanUsecase := loanusecase.NewLoanUsecase(loanRepository)
	loanController := loancontroller.NewLoanController(loanUsecase)

	return loanController
}

func getUserController(database *mongo.Database) *usercontroller.UserController {
	userRepository := repository.NewUserRepository(database)
	authRepository := repository.NewOAuthRepository(database)
	userUsecase := userusecase.NewUserUsecase(userRepository, authRepository)
	userController := usercontroller.NewUserController(userUsecase)


	return userController
}

func publicRouter(router *gin.Engine, userController *usercontroller.UserController) {
	router.POST("/users/register", userController.RegisterUser)
	router.POST("/users/login", userController.LoginUser)
	router.POST("/users/forgot-password", userController.ForgotPassword)

	router.GET("/users/verify", userController.VerifyUser)
	router.GET("/users/reset-password", userController.ResetPassword)

	// router.GET("/oauth2/login/google", userController.GoogleLogin)
	// router.GET("/oauth2/callback/google", userController.GoogleCallback)
}

func protectedRouter(router *gin.Engine, userController *usercontroller.UserController) {
	router.GET(
		"/token/refresh",
		middleware.AuthMiddleware("refresh"),
		userController.RefreshToken,
	)
}

func privateUserRouter(router *gin.RouterGroup, userController *usercontroller.UserController) {
	router.GET("/admin/users", userController.GetUsers)
	router.GET("/users/profile", userController.GetUserProfile)
	router.PATCH("/users", userController.UpdateProfile)

	router.POST("/users/logout", userController.LogoutUser)
	router.DELETE("/admin/users/:username", userController.DeleteUser)
	router.PATCH("/users/password-update", userController.ChangePassword)
}

func privateLoanRouter(router *gin.RouterGroup, loanController *loancontroller.LoanController) {
	router.POST("/loans", loanController.CreateLoan)
	router.GET("/admin/loans", loanController.GetAllLoans)
	router.GET("/loans/:id", loanController.GetLoanById)
	router.PATCH("/admin/loans/:id/status", loanController.UpdateLoanStatus)
	router.DELETE("/admin/loans/:id", loanController.DeleteLoan)
}



func SetupRouter(mongoClient *mongo.Client) *gin.Engine {
	
	router := gin.Default()

	// Secure Headers Configuration
	secureMiddleware := secure.New(secure.Config{
		SSLRedirect:           true,
		STSPreload:            true,
		ContentTypeNosniff:    true,
		BrowserXssFilter:      true,
		ContentSecurityPolicy: "default-src 'self'; script-src 'self'; object-src 'none';",
		ReferrerPolicy:        "no-referrer",
		IsDevelopment:         true,
		BadHostHandler:        func(*gin.Context) {},
	})

	// Apply secure middleware to the router
	router.Use(secureMiddleware)

	// CORS Configuration
	corsConfig := cors.Config{
		AllowOrigins:     []string{"https://trusteddomain.com", "https://anothertrusteddomain.com"}, // Adjust based on your needs
		AllowMethods:     []string{"GET", "POST", "PUT", "DELETE"},
		AllowHeaders:     []string{"Origin", "Content-Type", "Authorization"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
		MaxAge:           12 * time.Hour,
	}

	// Apply CORS middleware with custom configuration
	router.Use(cors.New(corsConfig))

	database := mongoClient.Database("loantracker")
    loanController := getLoanController(database)
	userController := getUserController(database,)

	publicRouter(router, userController)
	protectedRouter(router, userController)

	privateRouter := router.Group("")
	privateRouter.Use(middleware.AuthMiddleware("access"))

	privateUserRouter(privateRouter, userController)
	privateLoanRouter(privateRouter, loanController)
	

	return router
}
