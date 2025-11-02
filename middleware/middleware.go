package middleware

import (
	"net/http"
)

type HandlerFunc func(*Context)

type Engine struct {
	handlers []HandlerFunc
	routes   map[string]HandlerFunc
}

func New() *Engine {
	return &Engine{
		routes: make(map[string]HandlerFunc),
	}
}

func (e *Engine) Use(mw HandlerFunc) {
	e.handlers = append(e.handlers, mw)
}

func (e *Engine) GET(path string, handler HandlerFunc) {
	e.routes[path] = handler
}

func (e *Engine) Run(addr string) error {
	return http.ListenAndServe(addr, e)
}

func (e *Engine) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	c := &Context{Writer: w, Request: r}
	var allHandlers []HandlerFunc
	if handler, ok := e.routes[r.URL.Path]; ok {
		allHandlers = append(allHandlers, e.handlers...)
		allHandlers = append(allHandlers, handler)
		for _, h := range allHandlers {
			h(c)
		}
		return
	}
	w.WriteHeader(http.StatusNotFound)
}
