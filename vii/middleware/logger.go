package middleware

import (
	"log"
	"time"
	"vii/vii"
)

func Logger() vii.HandlerFunc {
	return func(c *vii.Context) {
		// Start time
		t := time.Now()
		// Process request
		c.Next()
		// Calculate resolution time
		log.Printf("[%d] %s in %v", c.StatusCode, c.Req.RequestURI, time.Since(t))
	}
}
