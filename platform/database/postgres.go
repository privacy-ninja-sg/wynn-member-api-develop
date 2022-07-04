// ./platform/database/postgres.go

package database

import (
	"context"
	"wynn-member-api/pkg/configs"
	"wynn-member-api/pkg/utils"
	_ "github.com/jackc/pgx/v4/stdlib" // load pgx driver for PostgreSQL
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"time"
)

// PostgresConnection func for connection to PostgreSQL database.
func PostgresConnection() *gorm.DB {
	pgConfig := configs.PostgresConfig()

	PgGlobal, err := gorm.Open(postgres.New(pgConfig), &gorm.Config{})
	utils.FailOnErr(err)

	return PgGlobal
}

// PostgresConnectionPool is wraper of Postgres db generic
func PostgresConnectionPool(db *gorm.DB) *gorm.DB {
	sqlDB, err := db.DB()
	utils.FailOnErr(err)

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	sqlDB.PingContext(ctx)

	maxConn, maxIdle, liteTime := configs.PostgresPoolConfig()
	sqlDB.SetMaxOpenConns(maxConn)
	sqlDB.SetConnMaxLifetime(liteTime)
	sqlDB.SetMaxIdleConns(maxIdle)

	return db
}
