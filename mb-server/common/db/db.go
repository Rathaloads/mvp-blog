package db

import (
	"mb-server/common/config"
)

func StartMysql(cfg *config.Mysql) {
	// "user:pass@tcp(127.0.0.1:3306)/dbname?charset=utf8mb4&parseTime=True&loc=Local"
	connectMysql(cfg)
}

func StartRedis() {}
