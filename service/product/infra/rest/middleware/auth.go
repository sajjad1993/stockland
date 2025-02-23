package middleware

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"stockland/pkg/jwt"
	"stockland/service/product/infra/rest/presenter"
	"strings"
)

const bearerPrefix = "Bearer "

// CheckToken just check token => save user in header if token exist
func CheckToken(jwtService jwt.JWT) func(ctx *gin.Context) {
	return func(ctx *gin.Context) {
		authHeader := ctx.GetHeader("Authorization")
		if !strings.HasPrefix(authHeader, bearerPrefix) {
			ctx.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"result": "unauthorized"})
			return
		}
		tokenUser, err := jwtService.ParseToken(strings.TrimPrefix(authHeader, bearerPrefix))
		if err != nil {
			ctx.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"result": "unauthorized"})
			return
		}
		if tokenUser != nil {
			ctx.Set(presenter.UserValueKey, tokenUser)
		}
		ctx.Next()
	}
}

// AdminRequired , middleware for admin permission
func AdminRequired() func(ctx *gin.Context) {
	return func(ctx *gin.Context) {
		{
			userToken, _ := ctx.Get(presenter.UserValueKey)
			user, ok := userToken.(*jwt.TokenUser)
			if !ok {
				ctx.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"result": "unauthorized"})
				return
			}
			if !user.IsAdmin {
				errMessage := presenter.ResponseMessages[presenter.LoginForbidden]
				ctx.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"result": errMessage})
				return
			}
			ctx.Next()
		}
	}
}
