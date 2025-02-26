package orm

import (
	"context"
	"gorm.io/gorm"
	"stockland/service/product/domain"
)

type ProductRepository struct {
	db *gorm.DB
}

func (p ProductRepository) Create(ctx context.Context, product *domain.Product) error {
	//TODO implement me
	panic("implement me")
}

func (p ProductRepository) GetByID(ctx context.Context, id uint) (*domain.Product, error) {
	//TODO implement me
	panic("implement me")
}

func (p ProductRepository) GetByUUID(ctx context.Context, uuid string) (*domain.Product, error) {
	//TODO implement me
	panic("implement me")
}

func (p ProductRepository) GetBySKU(ctx context.Context, sku string) (*domain.Product, error) {
	//TODO implement me
	panic("implement me")
}

func (p ProductRepository) GetAll(ctx context.Context) ([]domain.Product, error) {
	var products Products
	err := p.db.WithContext(ctx).Where("1").Find(&products).Error
	if err != nil {
		return nil, err
	}
	return products.toEntity(), nil

}

func NewProductRepository(db *gorm.DB) *ProductRepository {
	return &ProductRepository{
		db: db,
	}
}
