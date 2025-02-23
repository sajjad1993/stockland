package rest

import (
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"net/http"
	"stockland/config"
	service "stockland/service/product"
	"stockland/service/product/infra/rest/handler"
	"stockland/service/product/infra/rest/middleware"
)

type RouterGroup struct {
	users    *gin.RouterGroup
	telegram *gin.RouterGroup
}

type ManagerHandler struct {
	group   *RouterGroup
	handler *handler.Handler
	config  *config.Config
}

func New(container *service.Container) http.Handler {
	manager := &ManagerHandler{
		config:  &container.Config,
		group:   &RouterGroup{},
		handler: handler.New(),
	}

	router := gin.New()
	router.Use(gin.Recovery())
	router.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"https://panel.bestv2.co"},
		AllowMethods:     []string{"PUT", "PATCH", "GET", "POST", "HEAD", "DELETE", "OPTIONS"},    // Added DELETE and OPTIONS
		AllowHeaders:     []string{"Origin", "Authorization", "X-Requested-With", "Content-Type"}, // Added Authorization and X-Requested-With
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
		AllowOriginFunc: func(origin string) bool {
			return origin == "https://panel.bestv2.co"
		},
	}))

	//router.POST("/login", manager.handler.Login)
	//router.POST("/logout", manager.handler.Logout)
	//router.POST("/access", manager.handler.Access)

	v1 := router.Group("/api/v1")
	v1.Use(middleware.CheckToken(container.Jwt)) //todo telegram check
	{
		manager.group.users = v1.Group("/users")

		registerUserRoutes(manager)
		registerAdminRoutes(manager)
	}

	return router
}
