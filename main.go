package main

import (
	"net/http"
	"time"

	config "github.com/fattahmuhyiddeen/interview-order/config"
	"github.com/fattahmuhyiddeen/interview-order/routes"
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
)

func main() {
	e := echo.New()

	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins: []string{"*"},
		AllowHeaders: []string{"*"},
		AllowMethods: []string{"*"},
	}))

	e.Use(middleware.Logger())

	if config.Env == "heroku" {
		e.Pre(middleware.HTTPSRedirect())
	}
	e.Any("/*", func(c echo.Context) (err error) {
		routes.APIRoutes().ServeHTTP(c.Response(), c.Request())
		return
	})
	serverConfig := &http.Server{
		Addr:         ":" + config.Port,
		ReadTimeout:  20 * time.Minute,
		WriteTimeout: 20 * time.Minute,
	}
	e.Logger.Fatal(e.StartServer(serverConfig))
}
