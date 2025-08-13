package config

import (
	"log"
	"path/filepath"
	"strings"

	"github.com/spf13/viper"
)

var (
	config *viper.Viper
)

func InitConfig(cfg string) {
	v := viper.New()
	v.SetConfigType(strings.ReplaceAll(filepath.Ext(cfg), ".", ""))
	v.SetConfigName(strings.ReplaceAll(filepath.Base(cfg), filepath.Ext(cfg), ""))
	v.AddConfigPath(filepath.Dir(cfg))
	err := v.ReadInConfig()
	if err != nil {
		log.Panic(err)
	}
	v.WatchConfig()
	config = v
}

func Get() *viper.Viper {
	return config
}
