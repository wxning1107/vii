package vii

import (
	"fmt"
	"log"
	"strings"
)

type Router struct {
	handlers map[string]HandlerFunc
}

func NewRouter() *Router {
	return &Router{handlers: map[string]HandlerFunc{}}
}

func (r *Router) AddRouterHandler(method, pattern string, handler HandlerFunc) {
	var b strings.Builder
	b.WriteString(method)
	b.WriteString("-")
	b.WriteString(pattern)

	r.handlers[b.String()] = handler
}

func (r *Router) Handle(c *Context) {
	key := fmt.Sprintf("%s-%s", c.Method, c.Path)
	if handler, ok := r.handlers[key]; ok {
		handler(c)
	} else {
		if _, err := fmt.Fprintf(c.Writer, "404 NOT FOUND: %s\n", c.Req.URL); err != nil {
			log.Printf("write to page failed: %v", err)
		}
	}
}
