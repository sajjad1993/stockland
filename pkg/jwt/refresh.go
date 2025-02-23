package jwt

import (
	"context"
	"github.com/golang-jwt/jwt"
	"stockland/config"
	"time"
)

type Refresh struct {
	expiredAt    time.Duration
	signingKey   string
	access       IAccess
	cacheTimeOut time.Duration
}

func NewRefresh(cfg config.Config, access IAccess) IRefresh {
	return &Refresh{
		signingKey: cfg.RefreshSignKey,
		expiredAt:  cfg.RefreshExp,
		access:     access,
	}
}
func (j *Refresh) GenerateToken(user *TokenUser) (string, error) {
	return j.access.GenerateToken(user)
}

func (j *Refresh) ParseToken(accessToken string) (*TokenUser, error) {
	return j.access.ParseToken(accessToken)
}

func (j *Refresh) GenerateTokenRefresh(ctx context.Context, user *TokenUser) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS512, customClaims{
		TokenUser: *user,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Add(j.expiredAt).Unix(),
		},
	})
	tokenString, err := token.SignedString([]byte(j.signingKey))
	if err != nil {
		return "", err
	}
	return tokenString, nil
}

func (j *Refresh) ParseTokenRefresh(ctx context.Context, tokenString string) (*TokenUser, error) {
	token, err := jwt.Parse(
		tokenString,
		func(token *jwt.Token) (interface{}, error) {
			if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
				return nil, &TokenInValid{message: "unexpected signing method"}
			}
			return []byte(j.signingKey), nil
		},
	)
	if err != nil {
		return nil, err
	}

	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok {
		return nil, &TokenInValid{message: "cannot get claims from token"}
	}
	user, err := getClaims(claims)
	if err != nil {
		return nil, err
	}
	return user, nil
}
