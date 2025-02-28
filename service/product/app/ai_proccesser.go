package usecase

import (
	"context"
	"stockland/service/product/domain"
)

type ProductRequestUsecase struct {
	productRepo domain.ProductRepository
	aiProcessor domain.AIProcessor
}

func NewProductRequestUsecase(repo domain.ProductRepository, ai domain.AIProcessor) *ProductRequestUsecase {
	return &ProductRequestUsecase{productRepo: repo, aiProcessor: ai}
}

func (uc *ProductRequestUsecase) CreateRequest(ctx context.Context, request domain.ProductRequest) (domain.Offer, error) {
	products, err := uc.productRepo.GetAll(ctx)
	if err != nil {
		return domain.Offer{}, err
	}

	response, err := uc.aiProcessor.ProcessRequest(ctx, request.Prompt, provideSubPrompt(products))
	if err != nil {
		return response, err
	}
	return response, nil
}

func (uc *ProductRequestUsecase) AnalyzeImage(ctx context.Context, imagePath string) ([]string, error) {
	return uc.aiProcessor.AnalyzeImage(ctx, imagePath)
}

func provideSubPrompt(products []domain.Product) string {
	prompt := ""
	for _, product := range products {
		prompt += product.Name + ","
	}
	if len(prompt) > 0 {
		prompt = prompt[:len(prompt)-1] // remove the extra comma
	}
	return prompt
}
