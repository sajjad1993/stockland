package mem

import (
	"fmt"
	"github.com/google/uuid"
	"stockland/service/product/domain"
)

func Seed(memory *Mem) {
	products := []struct {
		Name        string
		Description string
	}{
		{"گوشی هوشمند", "الکترونیک"},
		{"لپ‌تاپ", "کامپیوتر"},
		{"هدفون", "لوازم جانبی"},
		{"ساعت هوشمند", "گجت"},
		{"کوله‌پشتی", "مد و پوشاک"},
	}

	memory.mu.Lock()
	defer memory.mu.Unlock()

	memory.mapper = make(map[string]domain.Product)

	for i, p := range products {
		u := uuid.New().String()
		sku := fmt.Sprintf("SKU-%d", i+1)
		product := domain.Product{
			ID:          uint(i + 1),
			UUID:        u,
			SKU:         sku,
			Name:        p.Name,
			Description: p.Description,
		}
		memory.mapper[u] = product
	}
}
