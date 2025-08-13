package config

import "log"

// 全局配置
type GlobalConfig struct{}

var (
	globalConfig *GlobalConfig
)

func GetGlobalConfig() *GlobalConfig {
	if globalConfig == nil {
		if !config.IsSet("Global") {
			return nil
		}
		globalConfig = &GlobalConfig{}
		err := config.UnmarshalKey("Global", globalConfig)
		if err != nil {
			log.Panic(err)
		}
	}
	return globalConfig
}
