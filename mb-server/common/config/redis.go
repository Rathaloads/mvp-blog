package config

import "log"

type Redis struct {
	Host     string `mapstructure:"Host"`
	Password string `mapstructure:"Password"`
	DB       int    `mapstructure:"DB"`
}

var (
	redis *Redis
)

func GetRedis() *Redis {
	if redis == nil {
		if !config.IsSet("Redis") {
			return nil
		}
		redis = &Redis{}
		err := config.UnmarshalKey("Redis", redis)
		if err != nil {
			log.Panic(err)
		}
	}
	return redis
}
