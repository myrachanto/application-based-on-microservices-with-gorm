package routes

import (
	"log"
	"os"
	"github.com/joho/godotenv"
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
	"github.com/myrachanto/allmicro/gormmicro/customermicroservice/controllers"
)

func CustomerMicroservice() {

	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
	PORT := os.Getenv("PORT")

	e := echo.New()

	// Middleware
	e.Use(middleware.Logger())
	e.Use(middleware.Recover()) 

	// Routes
	e.POST("/customers", controllers.CustomerController.Create)
	e.GET("/customers", controllers.CustomerController.GetAll)
	e.GET("/customers/:id", controllers.CustomerController.GetOne)
	e.PUT("/customers/:id", controllers.CustomerController.Update)
	e.DELETE("/customers/:id", controllers.CustomerController.Delete)

	// Start server
	e.Logger.Fatal(e.Start(PORT))
}
