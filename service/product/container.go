package service

import (
	"gorm.io/gorm"
	"stockland/config"
	"stockland/pkg/client"
	"stockland/pkg/jwt"
	logger "stockland/pkg/log"
)

type Container struct {
	Config config.Config
	Logger logger.Logger
	Client client.Client
	Jwt    jwt.JWT
	DB     *gorm.DB
}

func NewContainer(Client client.Client,
	Config config.Config,
	Jwt jwt.JWT,
	Logger logger.Logger,
	DB *gorm.DB,
) *Container {
	return &Container{
		Config: Config,
		Logger: Logger,
		Client: Client,
		Jwt:    Jwt,
		DB:     DB,
	}

}
