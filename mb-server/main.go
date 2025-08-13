package main

import (
	"flag"
	"mb-server/common/config"
	"mb-server/common/db"
	"mb-server/common/logger"
	"mb-server/router"
)

func bootStart() {
	var cfg string
	flag.StringVar(&cfg, "c", "config/config.yaml", "配置文件路径,默认为config/config.yaml")
	flag.Parse()
	config.InitConfig(cfg)
	logger.InitLog("./logger")
	logger.Debugf("init base success!!!")
	db.StartDb(config.GetMysql())
	logger.Debugf("init database success....")
}

func main() {
	bootStart()
	network := router.StartRouter()
	network.Run(":8088")
}
