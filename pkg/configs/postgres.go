package configs

import (
	"gorm.io/driver/postgres"
	"os"
	"strconv"
	"time"
	"wynn-member-api/pkg/utils"
)

func PostgresConfig() postgres.Config {
	serverUrl := os.Getenv("PG_SERVER_URL")

	dsn := serverUrl

	return postgres.Config{
		DSN:                  dsn,
		PreferSimpleProtocol: true, // disables implicit prepared statement usage
	}
}

func PostgresPoolConfig() (maxConn, idleConn int, lifeTime time.Duration) {
	var err error

	maxConnStr := os.Getenv("PG_MAX_CONNECTIONS")
	idleConnStr := os.Getenv("PG_MAX_IDLE_CONNECTIONS")
	lifeTimeStr := os.Getenv("PG_MAX_LIFETIME_CONNECTIONS")

	maxConn, err = strconv.Atoi(maxConnStr)
	utils.FailOnErr(err)

	idleConn, err = strconv.Atoi(idleConnStr)
	utils.FailOnErr(err)

	lt, err := strconv.Atoi(lifeTimeStr)
	utils.FailOnErr(err)

	lifeTime = time.Duration(lt) * time.Hour

	return maxConn, idleConn, lifeTime
}
