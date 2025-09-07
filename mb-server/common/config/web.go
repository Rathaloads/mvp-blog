package config

import "log"

type Web struct {
	SecretKey string `mapstructure:"SecretKey"`
	Mod       string `mapstructure:"Mod"` // "debug", "release"
}

var (
	web *Web
)

func GetWebConfig() *Web {
	if web == nil {
		if !config.IsSet("Web") {
			return nil
		}
		web = &Web{}
		err := config.UnmarshalKey("Web", web)
		if err != nil {
			log.Panic(err)
		}
	}
	return web
}
