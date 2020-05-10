package main

import (
	"log"
	"net/http"
	"time"
	"vii/vii"
	"vii/vii/middleware"
)

func myMiddleware() vii.HandlerFunc {
	return func(c *vii.Context) {
		// Start timer
		t := time.Now()
		// if a server error occurred
		c.Fail(http.StatusInternalServerError, "Internal Server Error")
		// Calculate resolution time
		log.Printf("[%d] %s in %v for test middleware", c.StatusCode, c.Req.RequestURI, time.Since(t))
	}
}

func main() {
	e := vii.New()

	e.Use(middleware.Logger())

	e.GET("/", func(c *vii.Context) {
		c.HTML(http.StatusOK, "<h1>Hello Vii</h1>")
	})

	e.GET("/value", func(c *vii.Context) {
		c.String(http.StatusOK, "Hello %s; your url is %s\n", c.Query("name"), c.Path)
	})

	e.POST("/json", func(c *vii.Context) {
		c.JSON(http.StatusOK, vii.V{
			"username": c.PostForm("username"),
			"email":    c.PostForm("email"),
		})
	})

	e.GET("/hello/:name", func(c *vii.Context) {
		c.String(http.StatusOK, "hey %s, you're at %s\n", c.Param("name"), c.Path)
	})

	e.GET("/assets/*filepath", func(c *vii.Context) {
		c.JSON(http.StatusOK, vii.V{"filepath": c.Param("filepath")})
	})

	v1 := e.Group("/v1")

	v1.Use(middleware.Recovery())

	{

		v1.GET("/", func(c *vii.Context) {
			c.HTML(http.StatusOK, "<h1>Hello Gee</h1>")
		})

		v1.GET("/hello", func(c *vii.Context) {
			panic("My panic occurred")
			c.String(http.StatusOK, "hello %s, you're at %s\n", c.Query("name"), c.Path)
		})
	}

	v2 := e.Group("/v2")

	v2.Use(myMiddleware())

	{
		v2.GET("/hello/:name", func(c *vii.Context) {
			c.String(http.StatusOK, "hello %s, you're at %s\n", c.Param("name"), c.Path)
		})

		v2.POST("/login", func(c *vii.Context) {
			c.JSON(http.StatusOK, vii.V{
				"username": c.PostForm("username"),
				"password": c.PostForm("password"),
			})
		})

	}

	_ = e.Run(":1107")
}
