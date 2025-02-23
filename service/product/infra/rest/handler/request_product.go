package handler

import (
	"github.com/gin-gonic/gin"
	"stockland/service/product/domain"
	"stockland/service/product/infra/rest/presenter"
)

// swagger:route POST /api/v1/telegram/product/request telegram productRequest
//
// # requests a product by giving a prompt
//
// Produces:
// application/json
//
// Responses:
// 200: Response
func (handler *Handler) ProductRequest(ctx *gin.Context) {

	var request presenter.RequestProduct

	if !shouldBindJSON(ctx, &request) {
		return
	}

	productRequestEntity := domain.NewProductRequest(request.Prompt, request.Image)
	err := handler.ProductRequestUsecase.CreateRequest(ctx, productRequestEntity)
	if err != nil {
		Error(ctx, err)
		return
	}

	OK(
		ctx, "successfully cleared",
	)
}
