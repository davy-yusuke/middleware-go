package middleware

import (
	"encoding/json"
	"fmt"
)

func (c *Context) Next() {
	c.index++
	for c.index < len(c.handlers) {
		c.handlers[c.index](c)
		c.index++
	}
}

func (c *Context) Abort() {
	c.index = len(c.handlers)
}

func (c *Context) Param(name string) string {
	if c.Params == nil {
		return ""
	}
	return c.Params[name]
}

func (c *Context) Set(key string, val interface{}) {
	if c.Keys == nil {
		c.Keys = make(map[string]interface{})
	}
	c.Keys[key] = val
}

func (c *Context) Get(key string) (interface{}, bool) {
	if c.Keys == nil {
		return nil, false
	}
	v, ok := c.Keys[key]
	return v, ok
}

func (c *Context) JSON(status int, v interface{}) {
	c.W.Header().Set("Content-Type", "application/json; charset=utf-8")
	c.W.WriteHeader(status)
	enc := json.NewEncoder(c.W)
	_ = enc.Encode(v)
}

func (c *Context) String(status int, format string, a ...interface{}) {
	c.W.Header().Set("Content-Type", "text/plain; charset=utf-8")
	c.W.WriteHeader(status)
	_, _ = fmt.Fprintf(c.W, format, a...)
}
