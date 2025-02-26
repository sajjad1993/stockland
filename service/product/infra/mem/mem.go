package mem

import (
	"context"
	"stockland/service/product/domain"
	"sync"
)

type Mem struct {
	mapper map[string]domain.Product
	mu     sync.RWMutex
}

func (m *Mem) Create(ctx context.Context, product *domain.Product) error {
	m.mu.Lock()
	defer m.mu.Unlock()
	m.mapper[domain.GenerateUUID()] = *product
	return nil
}

func (m *Mem) GetByID(ctx context.Context, id uint) (*domain.Product, error) {
	//TODO implement me
	panic("implement me")
}

func (m *Mem) GetByUUID(ctx context.Context, uuid string) (*domain.Product, error) {
	//TODO implement me
	panic("implement me")
}

func (m *Mem) GetBySKU(ctx context.Context, sku string) (*domain.Product, error) {
	//TODO implement me
	panic("implement me")
}

func (m *Mem) GetAll(ctx context.Context) ([]domain.Product, error) {
	var products []domain.Product
	m.mu.RLock()
	defer m.mu.RUnlock()
	for _, product := range m.mapper {
		products = append(products, product)
	}
	return products, nil
}
