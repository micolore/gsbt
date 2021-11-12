package main

import (
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"moppo.com/gsbt/log"
	"net/http"
	"time"
)

// Echo web server
func main() {
	e := echo.New()

	// Middleware
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	// Routes
	e.GET("/", hello)

	log.GetLogger().Info("server start!")

	// Start server
	e.Logger.Fatal(e.Start(":1323"))
}

// Handler
func hello(c echo.Context) error {
	time.Sleep(6 * time.Second)
	return c.String(http.StatusOK, "Hello, World!")
}
