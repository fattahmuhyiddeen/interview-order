package routes

import (
	controller "github.com/fattahmuhyiddeen/interview-order/controller"
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
)

// APIRoutes :
func APIRoutes() *echo.Echo {
	api := echo.New()
	api.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins: []string{"*"},
		AllowHeaders: []string{"*"},
	}))

	api.Use(middleware.Recover())

	// Routes
	api.GET("/", controller.HomePage)
	api.GET("/timestamp", controller.Timestamp)
	api.GET("/check_db", controller.CheckDB)

	api.POST("/create_order", controller.CreateOrder)
	api.DELETE("/cancel_order", controller.CancelOrder)
	api.GET("/check_order", controller.CheckOrder)
	api.GET("/list_orders", controller.CheckDB)

	return api
}
