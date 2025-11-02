package main

import (
	"middleware-go/middleware"
)

func main() {
	app := middleware.New()
	app.Use(middleware.LoggerWithConfig(middleware.LoggerConfig{Debug: true}))
	app.GET("/ping", func(c *middleware.Context) {
		c.String(200, "pong")
	})

	app.GET("/", func(c *middleware.Context) {
		c.String(200, "Hello, World!")
	})

	app.Run(":8080")
}
