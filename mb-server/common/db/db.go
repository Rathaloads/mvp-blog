package db

import (
	"mb-server/common/config"
)

func StartMysql(cfg *config.Mysql) error {
	// "user:pass@tcp(127.0.0.1:3306)/dbname?charset=utf8mb4&parseTime=True&loc=Local"
	return connectMysql(cfg)
}

func StartRedis(cfg *config.Redis) error {
	return connectRedis(cfg)
}
