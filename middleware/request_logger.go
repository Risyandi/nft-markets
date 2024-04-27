package middleware

import (
	"fmt"
	"time"

	"github.com/gin-gonic/gin"
)

// RequestLoggerMiddleware logs the incoming requests.
func RequestLoggerMiddleware() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		start := time.Now()

		// Process request
		ctx.Next()

		// After request completes
		end := time.Now()
		latency := end.Sub(start)

		// Log the request
		fmt.Printf(
			"%s - %s %s %s %s\n",
			end.Format("2006/01/02 - 15:04:05"),
			ctx.ClientIP(),
			ctx.Request.Method,
			ctx.Request.URL.Path,
			latency,
		)
	}
}
