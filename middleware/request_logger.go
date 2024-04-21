package middleware

import (
	"fmt"
	"time"

	"github.com/gin-gonic/gin"
)

// RequestLoggerMiddleware logs the incoming requests.
func RequestLoggerMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		start := time.Now()

		// Process request
		c.Next()

		// After request completes
		end := time.Now()
		latency := end.Sub(start)

		// Log the request
		fmt.Printf(
			"%s - %s %s %s %s\n",
			end.Format("2006/01/02 - 15:04:05"),
			c.ClientIP(),
			c.Request.Method,
			c.Request.URL.Path,
			latency,
		)
	}
}
