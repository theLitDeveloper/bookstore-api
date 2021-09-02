package bookstore

import (
	"net/http"
	"os"

	"github.com/labstack/echo-contrib/prometheus"
	"github.com/labstack/echo/v4"
)

func Run() {

	e := echo.New()

	// Init in-memory datastore
	initDatastore()

	// Init bookstore API routes
	initRoutes(e)

	// Health check
	e.GET("/health", func(ctx echo.Context) error {
		return ctx.String(http.StatusOK, "OK")
	})

	// Version info
	e.GET("/version", func(ctx echo.Context) error {
		resp := make(map[string]interface{})
		resp["version"] = os.Getenv("TAG")
		return ctx.JSON(http.StatusOK, resp)
	})

	// Enable metrics middleware
	p := prometheus.NewPrometheus("echo", nil)
	p.Use(e)

	// Spin up the HTTP server
	e.Logger.Fatal(e.Start(":8080"))
}
