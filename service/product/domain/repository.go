package domain

import "context"

type UserRepository interface {
	Create(user *User) error
	GetByID(id uint) (*User, error)
	GetByUUID(uuid string) (*User, error)
}

type ProductRepository interface {
	Create(ctx context.Context, product *Product) error
	GetByID(ctx context.Context, id uint) (*Product, error)
	GetByUUID(ctx context.Context, uuid string) (*Product, error)
	GetBySKU(ctx context.Context, sku string) (*Product, error)
	GetAll(ctx context.Context) ([]Product, error)
}

type ProductRequestRepository interface {
	Create(ctx context.Context, request *ProductRequest) error
	GetByID(ctx context.Context, id uint) (*ProductRequest, error)
	GetByUUID(ctx context.Context, uuid string) (*ProductRequest, error)
	GetByUserID(ctx context.Context, userID uint) ([]ProductRequest, error)
	GetPendingRequests(ctx context.Context) ([]ProductRequest, error)
	UpdateStatus(ctx context.Context, id uint, status string) error
}
