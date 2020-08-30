package routes

import (
	"log"
	"os"
	"github.com/joho/godotenv"
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
	"github.com/myrachanto/allmicro/gormmicro/categorymicroservice/controllers"
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

	e.Static("/", "assets/images/")
	echoGroupUseJWT := e.Group("/categorys")
	echoGroupUseJWT.Use(middleware.JWT([]byte(key)))
	// Routes
	////////products//////////////////////
	e.POST("/categorys", controllers.CategoryController.Create)
	e.GET("/categorys", controllers.CategoryController.GetAll)
	e.GET("/categorys/:id", controllers.CategoryController.GetOne)
	e.PUT("/categorys/:id", controllers.CategoryController.Update)
	e.DELETE("/categorys/:id", controllers.CategoryController.Delete)
	//e.DELETE("/loggoutall/:id", controllers.UserController.DeleteALL) logout all accounts

	// Start server
	e.Logger.Fatal(e.Start(PORT))
}
