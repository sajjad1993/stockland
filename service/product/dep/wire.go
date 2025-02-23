//go:build wireinject
// +build wireinject

package dep

import (
	"github.com/google/wire"
	"stockland/config"
	"stockland/pkg/client"
	"stockland/pkg/db/postgres_gorm_db"
	"stockland/pkg/jwt"
	logger "stockland/pkg/log"
	service "stockland/service/product"
)

// initializeContainer is dependency injected form of having *service.Container
func InitializeContainer() (*service.Container, error) {
	wire.Build(
		service.NewContainer,
		config.NewConfigFromViper,
		logger.NewSLogger,
		jwt.NewJWT,
		client.NewClient,
		postgres_gorm_db.NewPostgresConn,
	)
	return new(service.Container), nil
}
