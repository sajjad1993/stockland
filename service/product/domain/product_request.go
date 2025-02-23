package domain

import "stockland/pkg/errs"

type ProductRequest struct {
	ID          uint
	UUID        string
	Prompt      string
	Image       string
	UserID      uint
	UserUUID    string
	ProductID   *uint
	ProductUUID *string
	ProductName string
	Status      string
}

type ProductStatus string

const (
	ProductStatusActive  ProductStatus = "active"
	ProductStatusPending ProductStatus = "pending"
)

func NewProductRequest(Prompt string, Image string) ProductRequest {
	return ProductRequest{
		ID:     0,
		UUID:   GenerateUUID(),
		Prompt: Prompt,
		Image:  Image,
	}
}

func (p ProductRequest) IsValid() error {
	if p.Prompt == "" {
		return errs.NewValidationError("empty prompt")
	}
	return nil
}
