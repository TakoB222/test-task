package v1

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"test-task/internal/service"
)

func (h *Handler) InitCardRoutes(groupApi *gin.RouterGroup) {
	auth := groupApi.Group("/card")
	{
		auth.POST("/validate", h.validate)
	}
}

type (
	validateCardInput struct {
		Number string `json:"number" binding:"required"`
		Month  string `json:"month" binding:"required"`
		Year   string `json:"year" binding:"required"`
	}
)

func (h *Handler) validate(ctx *gin.Context) {
	var input validateCardInput
	if err := ctx.BindJSON(&input); err != nil {
		newResponse(ctx, http.StatusBadRequest, "invalid input body")
		return
	}

	_, err := h.services.CreditCard.Validate(service.CreditCardInput{
		Number: input.Number,
		Month:  input.Month,
		Year:   input.Year,
	})
	if err != nil {
		newResponse(ctx, http.StatusNotAcceptable, err.Error())
		return
	}

	ctx.JSON(http.StatusOK, "card is valid")
}
