package main

import (
	"flag"
	"mb-server/common/config"
	"mb-server/common/db"
	"mb-server/common/logger"
	"mb-server/router"
	"os"
	"os/signal"
	"syscall"
)

func bootStart() {
	var cfg string
	flag.StringVar(&cfg, "c", "config/config.yaml", "配置文件路径,默认为config/config.yaml")
	flag.Parse()
	config.InitConfig(cfg)
	logger.InitLog("./logger")
	if err := db.StartMysql(config.GetMysql()); err != nil {
		logger.Panicf("start mysql fail: %v", err)
	}
	if err := db.StartRedis(config.GetRedis()); err != nil {
		logger.Panic("start redis fail: %v", err)
	}
	logger.Debug("bootstart success.....")
}

func main() {
	bootStart()
	network := router.StartRouter()

	closeCh := make(chan os.Signal, 1)
	signal.Notify(closeCh, syscall.SIGINT, syscall.SIGTERM)
	go network.Run(":8088")
	<-closeCh
	logger.Debugf("server shutdown.....")
}
