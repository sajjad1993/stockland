package router

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"stockland/pkg/jwt"
	"stockland/service/product/infra/rest/presenter"
	"strconv"
)

func getID(ctx *gin.Context) (uint, bool) {
	idString := ctx.Param("id")
	id, err := strconv.Atoi(idString)
	if err != nil {
		response := presenter.Response{
			Error: err.Error(),
		}
		ctx.JSON(http.StatusBadRequest, response)

		return 0, false
	}

	return uint(id), true
}

func shouldBindJSON(ctx *gin.Context, object any) bool {
	err := ctx.ShouldBindJSON(object)
	if err != nil {
		response := presenter.Response{
			Error: err.Error(),
		}
		ctx.JSON(http.StatusBadRequest, response)

		return false
	}

	return true
}

func Error(ctx *gin.Context, err error) {
	ctx.JSON(
		http.StatusInternalServerError,
		presenter.Response{
			Error: err.Error(),
		},
	)
}

func OK(ctx *gin.Context, data any) {
	ctx.JSON(
		http.StatusOK,
		presenter.Response{
			Data: data,
		},
	)
}

func getUser(ctx *gin.Context) *jwt.TokenUser {
	// Get user ID
	token, _ := ctx.Get(presenter.UserValueKey)
	user, _ := token.(*jwt.TokenUser)
	return user
}
