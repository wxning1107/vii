package vii

import (
	"net/http"
)

type HandlerFunc func(*Context)

type Engine struct {
	router *Router
}

func New() *Engine {
	return &Engine{router: NewRouter()}
}

func (e *Engine) AddRoute(method, pattern string, handler HandlerFunc) {
	e.router.AddRouterHandler(method, pattern, handler)
}

func (e *Engine) Get(pattern string, handler HandlerFunc) {
	e.AddRoute("GET", pattern, handler)
}

func (e *Engine) POST(pattern string, handler HandlerFunc) {
	e.AddRoute("POST", pattern, handler)
}

func (e *Engine) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	c := NewContext(w, req)
	e.router.Handle(c)
}

func (e *Engine) Run(addr string) error {
	return http.ListenAndServe(addr, e)
}
