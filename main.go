package main

import (
	"log"
	"net/http"
	"vii/vii"
)

func main() {
	e := vii.New()
	e.Get("/", func(c *vii.Context) {
		c.HTML(http.StatusOK, "<h1>Hello Vii</h1>")
	})

	e.Get("/value", func(c *vii.Context) {
		log.Print('a')
		c.String(http.StatusOK, "Hello %s; your url is %s\n", c.Query("name"), c.Path)
	})

	e.POST("/json", func(c *vii.Context) {
		c.JSON(http.StatusOK, vii.V{
			"username": c.PostForm("username"),
			"email":    c.PostForm("email"),
		})
	})

	_ = e.Run(":1107")
}
