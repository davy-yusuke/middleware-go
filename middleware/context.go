package middleware

import (
	"fmt"
	"net/http"
)

type Context struct {
	Writer     http.ResponseWriter
	Request    *http.Request
	StatusCode int
}

func (c *Context) String(code int, msg string) {
	c.StatusCode = code
	c.Writer.WriteHeader(code)
	fmt.Fprint(c.Writer, msg)
}
