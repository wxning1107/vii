package vii

import (
	"net/http"
)

type HandlerFunc func(*Context)

type RouterGroup struct {
	prefix      string
	middlewares []HandlerFunc
	parent      *RouterGroup
	engine      *Engine
}

type Engine struct {
	*RouterGroup
	router *Router
	groups []*RouterGroup
}

func New() *Engine {
	engine := &Engine{router: NewRouter()}
	engine.RouterGroup = &RouterGroup{engine: engine}
	engine.groups = []*RouterGroup{engine.RouterGroup}

	return engine
}

func (g *RouterGroup) Group(prefix string) *RouterGroup {
	newGroup := &RouterGroup{
		prefix: g.prefix + prefix,
		parent: g,
		engine: g.engine,
	}
	g.engine.groups = append(g.engine.groups, newGroup)

	return newGroup
}

func (rg *RouterGroup) AddRoute(method, comp string, handler HandlerFunc) {
	pattern := rg.prefix + comp
	rg.engine.router.AddRouterHandler(method, pattern, handler)
}

func (rg *RouterGroup) GET(pattern string, handler HandlerFunc) {
	rg.AddRoute("GET", pattern, handler)
}

func (rg *RouterGroup) POST(pattern string, handler HandlerFunc) {
	rg.AddRoute("POST", pattern, handler)
}

func (e *Engine) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	c := NewContext(w, req)
	e.router.Handle(c)
}

func (e *Engine) Run(addr string) error {
	return http.ListenAndServe(addr, e)
}
