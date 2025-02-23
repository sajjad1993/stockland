package jwt

import (
	"context"
	"stockland/config"
)

type JWT IRefresh

type IAccess interface {
	GenerateToken(user *TokenUser) (string, error)
	ParseToken(accessToken string) (*TokenUser, error)
}

type IRefresh interface {
	IAccess
	GenerateTokenRefresh(ctx context.Context, user *TokenUser) (string, error)
	ParseTokenRefresh(ctx context.Context, tokenString string) (*TokenUser, error)
}

func NewJWT(cfg config.Config) JWT {
	access := NewAccess(cfg)
	refresh := NewRefresh(cfg, access)
	jwt := refresh
	return jwt

}
