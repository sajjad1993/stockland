package postgres_gorm_db

import (
	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/stdlib"
	_ "github.com/lib/pq"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"log"
	"stockland/config"
)

func NewPostgresConn(cfg config.Config) (*gorm.DB, error) {
	dbConfig, err := pgx.ParseConfig(cfg.DatabaseDsn)
	if err != nil {
		return nil, err
	}
	if err != nil {
		log.Println("DB Connection error : ", err.Error())
		return nil, err
	}
	dbSql := stdlib.OpenDB(*dbConfig)
	gdb, err := gorm.Open(postgres.New(postgres.Config{
		Conn: dbSql,
	}), &gorm.Config{})
	if err != nil {
		return nil, err
	}
	gdb.Debug()
	gdb.Logger = logger.Default.LogMode(logger.Info)
	return gdb, nil

}
