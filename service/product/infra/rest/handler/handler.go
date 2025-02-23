package handler

import usecase "stockland/service/product/app"

type Handler struct {
	ProductRequestUsecase usecase.ProductRequestUsecase
}

func New() *Handler {
	return &Handler{}
}
