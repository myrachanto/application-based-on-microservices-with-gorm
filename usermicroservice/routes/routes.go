package routes

import (
	"log"
	"os"
	"github.com/joho/godotenv"
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
	"github.com/myrachanto/allmicro/gormmicro/usermicroservice/controllers"
)

func CustomerMicroservice() {

	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file in routes")
	}
	PORT := os.Getenv("PORT")
	key := os.Getenv("EncryptionKey")

	e := echo.New()

	// Middleware
	e.Use(middleware.Logger())
	e.Use(middleware.Recover()) 
	e.Use(middleware.CORS())

	JWTgroup := e.Group("/")
	JWTgroup.Use(middleware.JWTWithConfig(middleware.JWTConfig{
		SigningMethod: "HS256",
		SigningKey: []byte(key),
	}))
	//JwtG := e.Group("/users")
	// JwtG.Use(middleware.JWT([]byte(key)))
	// Routes
	e.POST("/register", controllers.UserController.Create)
	e.POST("/login", controllers.UserController.Login)
	JWTgroup.GET("users/logout/:token", controllers.UserController.Logout)
	JWTgroup.GET("users", controllers.UserController.GetAll)
	JWTgroup.GET("users/:id", controllers.UserController.GetOne)
	JWTgroup.PUT("users/:id", controllers.UserController.Update)
	JWTgroup.DELETE("users/:id", controllers.UserController.Delete)
	//e.DELETE("loggoutall/:id", controllers.UserController.DeleteALL) logout all accounts

	// Start server
	e.Logger.Fatal(e.Start(PORT))
}
