package vii

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

type V map[string]interface{}

type Context struct {
	Writer     http.ResponseWriter
	Req        *http.Request
	StatusCode int
	Method     string
	Path       string
}

func NewContext(w http.ResponseWriter, req *http.Request) *Context {
	return &Context{
		Writer: w,
		Req:    req,
		Method: req.Method,
		Path:   req.URL.Path,
	}
}

func (c *Context) PostForm(key string) string {
	return c.Req.FormValue(key)
}

func (c *Context) Query(key string) string {
	return c.Req.URL.Query().Get(key)
}

func (c *Context) String(code int, format string, values ...interface{}) {
	log.Printf("Route %4s - %s", c.Method, c.Path)

	c.setHeader("Content-Type", "text/plain")
	c.setStatusCode(code)
	if _, err := c.Writer.Write([]byte(fmt.Sprintf(format, values...))); err != nil {
		http.Error(c.Writer, err.Error(), http.StatusInternalServerError)
	}
}

func (c *Context) setHeader(key, value string) {
	c.Writer.Header().Set(key, value)
}

func (c *Context) setStatusCode(code int) {
	c.StatusCode = code
	c.Writer.WriteHeader(code)
}

func (c *Context) HTML(code int, html string) {
	log.Printf("Route %4s - %s", c.Method, c.Path)

	c.setHeader("Content-Type", "text/html")
	c.setStatusCode(code)
	if _, err := c.Writer.Write([]byte(html)); err != nil {
		panic(err)
	}
}

func (c *Context) JSON(code int, obj interface{}) {
	log.Printf("Route %4s - %s", c.Method, c.Path)

	c.setHeader("Content-Type", "application/json")
	c.setStatusCode(code)

	encoder := json.NewEncoder(c.Writer)
	if err := encoder.Encode(obj); err != nil {
		http.Error(c.Writer, err.Error(), http.StatusInternalServerError)
	}
}
