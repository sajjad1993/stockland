package jwt

import (
	"github.com/golang-jwt/jwt"
	"stockland/config"
	"time"
)

type Error string

func (e Error) Error() string {
	return string(e)
}

var (
	ErrInvalidToken error = Error("invalid or malformed token")
)

type Access struct {
	expiredAt  time.Duration
	signingKey string
}

func (j Access) GenerateToken(user *TokenUser) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, customClaims{
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

func (j Access) ParseToken(accessToken string) (*TokenUser, error) {
	token, err := jwt.Parse(
		accessToken,
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

type AccessOld struct {
	secret      string
	secretBytes []byte
	exp         time.Duration
}

func NewAccess(cfg config.Config) IAccess {
	return &Access{
		signingKey: cfg.AccessSignKey,
		expiredAt:  cfg.AccessJWTExp,
	}
}

func getClaims(claims jwt.MapClaims) (*TokenUser, error) {

	isAdmin, _ := claims["is_admin"].(bool) // in not admin user , this field is nil

	userId, ok := claims["id"]
	if !ok {
		return nil, ErrInvalidToken
	}
	username, ok := claims["username"]
	if !ok {
		return nil, ErrInvalidToken
	}
	role, ok := claims["role"]
	if !ok {
		return nil, ErrInvalidToken
	}
	user := &TokenUser{
		ID:       uint(userId.(float64)),
		Username: username.(string),
		Role:     role.(string),
		IsAdmin:  isAdmin,
	}

	return user, nil
}
