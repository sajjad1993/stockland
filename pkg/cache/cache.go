package cache

import (
	"context"
	"time"
)

type Cache interface {
	Get(ctx context.Context, key string) (string, error)
	Set(ctx context.Context, key, val string, exp time.Duration) error
	Delete(ctx context.Context, key string) error
	GetAll(ctx context.Context, keyPattern string) ([]string, error)
	Ping(ctx context.Context) error
	IncrBy(ctx context.Context, key string, increment int) error
}
