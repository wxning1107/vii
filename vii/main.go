package vii

import (
	"fmt"
	"log"
	"net/http"
	"strings"
)

//func DemoHandler(w http.ResponseWriter, request *http.Request) {
//	_, err := fmt.Fprintf(w, `<h1>Hello Go web</h1>`)
//	if err != nil {
//		panic(err)
//	}
//}

type handlerFunc func(http.ResponseWriter, *http.Request)

type Engine struct {
	router map[string]handlerFunc
}

func NewEngine() *Engine {
	return &Engine{map[string]handlerFunc{}}
}

func (e *Engine) AddRouter(method, pattern string, handler handlerFunc) {
	var b strings.Builder
	b.WriteString(method)
	b.WriteString("-")
	b.WriteString(pattern)

	log.Printf("Route %4s - %s", method, pattern)

	e.router[b.String()] = handler
}

func (e *Engine) Get(pattern string, handler handlerFunc) {
	e.AddRouter("GET", pattern, handler)
}

func (e *Engine) POST(pattern string, handler handlerFunc) {
	e.AddRouter("POST", pattern, handler)
}

func (e Engine) ServeHTTP(w http.ResponseWriter, request *http.Request) {
	key := fmt.Sprintf("%s-%s", request.Method, request.URL.Path)
	if handler, ok := e.router[key]; ok {
		handler(w, request)
	} else {
		_, err := fmt.Fprintf(w, "404 NOT FOUND: %s\n", request.URL)
		if err != nil {
			log.Printf("write to page failed: %v", err)
		}
	}
}

func (e *Engine) Run(addr string) error {
	return http.ListenAndServe(addr, e)
}

//func main() {
//	e := new(Engine)
//	//http.HandleFunc("/vii", e.ServeHTTP)
//	log.Fatal(http.ListenAndServe(":1107", e))
//
//}
