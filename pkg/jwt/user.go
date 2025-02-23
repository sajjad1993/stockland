package jwt

import "github.com/golang-jwt/jwt"

type TokenUser struct {
	ID       uint   `json:"id"`
	Username string `json:"username"`
	Role     string `json:"role"`
	IsAdmin  bool   `json:"is_admin"`
}

type customClaims struct {
	TokenUser
	jwt.StandardClaims
}
