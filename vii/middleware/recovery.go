package middleware

import (
	"log"
	"vii/vii"
)

func Recovery() vii.HandlerFunc {
	return func(c *vii.Context) {
		defer func() {
			if err := recover(); err != nil {
				log.Printf("Panic occurred: %s", err)
			}
		}()

		c.Next()
	}
}
