package usecase

import (
	"context"
	"stockland/service/product/domain"
)

type ProductRequestUsecase struct {
	repo        domain.ProductRequestRepository
	aiProcessor domain.AIProcessor
}

func NewProductRequestUsecase(repo domain.ProductRequestRepository, ai domain.AIProcessor) *ProductRequestUsecase {
	return &ProductRequestUsecase{repo: repo, aiProcessor: ai}
}

func (uc *ProductRequestUsecase) CreateRequest(ctx context.Context, request domain.ProductRequest) error {

	aiResponse, err := uc.aiProcessor.ProcessRequest(request.Prompt)
	if err == nil {
		request.ProductName, _ = domain.NewProductName(aiResponse)
	}
	request.Status = "pending"
	return uc.repo.Create(request)
}

func (uc *ProductRequestUsecase) AnalyzeImage(ctx, imagePath string) ([]string, error) {
	return uc.aiProcessor.AnalyzeImage(imagePath)
}
