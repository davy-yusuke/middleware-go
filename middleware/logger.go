package middleware

import (
	"fmt"
	"time"
)

type LoggerConfig struct {
	Debug bool
}

func LoggerWithConfig(cfg LoggerConfig) func(*Context) {
	return func(c *Context) {
		start := time.Now()
		fmt.Printf("[LOG] %s %s from %s\n", c.Request.Method, c.Request.URL.Path, c.Request.RemoteAddr)

		if cfg.Debug {
			fmt.Printf("[DEBUG] Request Headers: %v\n", c.Request.Header)
		}

		defer func() {
			duration := time.Since(start)
			fmt.Printf("[LOG] completed in %v\n", duration)
			if cfg.Debug {
				fmt.Printf("[DEBUG] Response Headers: %v\n", c.Writer.Header())
			}
		}()
	}
}

func Logger() func(*Context) {
	return LoggerWithConfig(LoggerConfig{Debug: false})
}
