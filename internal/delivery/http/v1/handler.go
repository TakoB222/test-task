package v1

import (
	"github.com/gin-gonic/gin"
	"test-task/internal/service"
)

type Handler struct {
	services *service.Service
}

func NewHandler(service *service.Service) *Handler {
	return &Handler{services: service}
}

func (h *Handler) Init(groupApi *gin.RouterGroup) {
	v1 := groupApi.Group("/v1")
	{
		h.InitCardRoutes(v1)
	}
}
