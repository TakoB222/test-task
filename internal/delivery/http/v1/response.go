package v1

import (
	"github.com/gin-gonic/gin"
	"test-task/pkg/logger"
)

type response struct {
	Message string `json:"message"`
}

func newResponse(ctx *gin.Context, statusCode int, message string) {
	logger.Error(message)
	ctx.AbortWithStatusJSON(statusCode, response{message})
}
