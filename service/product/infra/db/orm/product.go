package orm

import (
	"gorm.io/gorm"
	"stockland/service/product/domain"
)

type Product struct {
	gorm.Model
	Name        string `json:"name" gorm:"Index;not null;size:32"`
	Description string `json:"description" gorm:"not null"`
}
type Products []Product

func (t Product) toEntity() domain.Product {
	return domain.Product{
		ID:          t.ID,
		Name:        t.Name,
		Description: t.Description,
	}
}

func FromEntity(e domain.Product) Product {

	return Product{
		Name:        e.Name,
		Description: e.Description,
	}
}

func (p Products) toEntity() []domain.Product {
	var products []domain.Product
	for _, product := range p {
		products = append(products, product.toEntity())
	}
	return products
}
