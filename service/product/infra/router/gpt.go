package router

import (
	"net/http"
	"stockland/service/product/domain"
	"stockland/service/product/infra/router/dto"

	"github.com/gin-gonic/gin"
)

func (endpoint *EndPoint) Search(ctx *gin.Context) {
	var result dto.SearchResponse
	query := ctx.Query("q")

	// اگر جستجو خالی بود، کل لیست رو برگردون
	if query == "" {
		ctx.JSON(http.StatusOK, result)
		return
	}
	productRequest := domain.NewProductRequest(query, "")
	offer, err := endpoint.ProductRequestUsecase.CreateRequest(ctx, productRequest)
	if err != nil {
		Error(ctx, err)
		return
	}
	result.Message = offer.Message
	ctx.JSON(http.StatusOK, offer)
}
