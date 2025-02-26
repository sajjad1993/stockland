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

func (uc *ProductRequestUsecase) CreateRequest(ctx context.Context, request domain.ProductRequest) error {
	products, err := uc.productRepo.GetAll(ctx)
	if err != nil {
		return err
	}

	_, err = uc.aiProcessor.ProcessRequest(ctx, request.Prompt, provideSubPrompt(products))
	if err != nil {
		return err
	}
	//todo  get the actual products from the database
	request.Status = "pending"
	return nil
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
