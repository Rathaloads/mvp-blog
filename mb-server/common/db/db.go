package db

import (
	"fmt"
	"log"
	"mb-server/common/config"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var (
	Mysql *gorm.DB
)

func StartDb(cfg *config.Mysql) {
	// "user:pass@tcp(127.0.0.1:3306)/dbname?charset=utf8mb4&parseTime=True&loc=Local"
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local", cfg.Username, cfg.Password, cfg.Host, cfg.Port, cfg.DatabaseName)
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Panicf("cant not open database with: %v", err)
	}
	Mysql = db
}

func StartRedis() {}
