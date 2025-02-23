package cache

import (
	"context"
	"fmt"
	"stockland/pkg/errs"
	"time"

	"github.com/go-redis/redis/v8"
)

type redisCache struct {
	rdb *redis.Client
}

func (r *redisCache) Get(ctx context.Context, key string) (string, error) {
	val, err := r.rdb.Get(ctx, key).Result()
	if err != nil {
		if err == redis.Nil {
			return "", errs.ErrCacheMis
		}
		return "", err
	}

	return val, nil
}

func (r *redisCache) GetAll(ctx context.Context, keyPattern string) ([]string, error) {
	keys, err := r.rdb.Keys(ctx, keyPattern).Result()
	if err != nil {
		return nil, err
	}
	if len(keys) == 0 {
		return nil, nil
	}
	vals, err := r.rdb.MGet(ctx, keys...).Result()
	if err != nil {
		return nil, err
	}
	resp := []string{}
	for _, v := range vals {
		resp = append(resp, fmt.Sprintf("%v", v))
	}

	return resp, nil
}

func (r *redisCache) Set(ctx context.Context, key, val string, exp time.Duration) error {
	err := r.rdb.Set(ctx, key, val, exp).Err()
	if err != nil {
		return err
	}
	return nil
}

func (r *redisCache) Delete(ctx context.Context, key string) error {
	err := r.rdb.Del(ctx, key).Err()
	if err != nil {
		return err
	}
	return nil
}

func (r *redisCache) Ping(ctx context.Context) error {
	_, err := r.rdb.Ping(ctx).Result()
	return err
}

func (r *redisCache) IncrBy(ctx context.Context, key string, increment int) error {
	err := r.rdb.IncrBy(ctx, key, int64(increment)).Err()
	if err != nil {
		return err
	}
	return nil
}
