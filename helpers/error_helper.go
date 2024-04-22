package helpers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/labstack/gommon/log"
	"github.com/meysamhadeli/problem-details"
)

// GinErrorHandler middleware for handle problem details error on gin
func GinErrorHandler() gin.HandlerFunc {
	return func(c *gin.Context) {

		c.Next()

		for _, err := range c.Errors {

			// problem details handler config
			problem.MapStatus(http.StatusBadRequest, func() problem.ProblemDetailErr {
				return &problem.ProblemDetail{
					Status: http.StatusBadRequest,
					Title:  "Bad Request",
					Detail: err.Error(),
				}
			})

			problem.MapStatus(http.StatusInternalServerError, func() problem.ProblemDetailErr {
				return &problem.ProblemDetail{
					Status: http.StatusInternalServerError,
					Title:  "Internal Server Error",
					Detail: err.Error(),
				}
			})

			if _, err := problem.ResolveProblemDetails(c.Writer, c.Request, err); err != nil {
				log.Error(err)
			}
		}
	}
}
