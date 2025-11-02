# middleware-go

A simple and extensible HTTP middleware engine for Go, inspired by the elegant Gin style syntax.

## Features

- Minimal core, easy to read and extend
- Gin-like handler registration and middleware style
- Detailed logging and debug mode support
- Custom handler and middleware chaining
- No third-party dependencies

## Getting Started

### Installation

Clone this repository and use it in your own Go project:
```bash
git clone https://github.com/yourname/middleware-go.git
```
Or simply copy the `middleware/` folder into your project.

### Quick Example

```go
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
    app.Run(":8080")
}
```

You can now access [http://localhost:8080/ping](http://localhost:8080/ping) and see detailed logs.

## Custom Middleware

You can add your own middlewares by implementing the following signature:
```go
func MyMiddleware() middleware.HandlerFunc {
    return func(c *middleware.Context) {
        // your code
    }
}
```
Register it with:
```go
app.Use(MyMiddleware())
```

## Debug Logging

Use `LoggerWithConfig` for detailed request/response logging and header examination:
```go
app.Use(middleware.LoggerWithConfig(middleware.LoggerConfig{Debug: true}))
```

## License

MIT License. See [LICENSE](LICENSE) for details.
