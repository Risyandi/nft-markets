package helpers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type Response struct {
	Code    string      `json:"code"`
	Message string      `json:"message,omitempty"`
	Data    interface{} `json:"data,omitempty"`
}

func Success(ctx *gin.Context, code string, message string, data interface{}) {
	ctx.JSON(http.StatusOK, Response{
		Code:    code,
		Message: message,
		Data:    data,
	})
}

func BadRequest(ctx *gin.Context, code string, message string) {
	ctx.JSON(http.StatusBadRequest, Response{
		Code:    code,
		Message: message,
	})
}

func NotFound(ctx *gin.Context, code string, message string) {
	ctx.JSON(http.StatusNotFound, Response{
		Code:    code,
		Message: message,
	})
}

func ServerError(ctx *gin.Context, code string, message string) {
	ctx.JSON(http.StatusInternalServerError, Response{
		Code:    code,
		Message: message,
	})
}
