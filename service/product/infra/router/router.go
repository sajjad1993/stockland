package router

import (
	"net/http"
	usecase "stockland/service/product/app"

	"github.com/gin-gonic/gin"
)

type EndPoint struct {
	ProductRequestUsecase *usecase.ProductRequestUsecase
}

func NewRouter(ProductRequestUsecase *usecase.ProductRequestUsecase) http.Handler {
	endpoint := &EndPoint{
		ProductRequestUsecase: ProductRequestUsecase,
	}

	engine := gin.Default()

	engine.Use(gin.Logger())

	engine.SetTrustedProxies(nil)

	engine.Static("/static", "./static")

	engine.GET("/", endpoint.Home)
	engine.GET("/api/search", endpoint.Search)

	return engine
}
