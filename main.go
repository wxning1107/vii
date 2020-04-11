package main

import (
	"fmt"
	"net/http"
	"vii/vii"
)

func main() {
	e := vii.NewEngine()
	e.Get("/", func(w http.ResponseWriter, req *http.Request) {
		_, err := fmt.Fprintf(w, "URL.Path = %q\n", req.URL.Path)
		if err != nil {
			panic(err)
		}
	})

	_ = e.Run(":1107")
}
