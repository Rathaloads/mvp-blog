package config

import "log"

type Mysql struct {
	Username     string `mapstructure:"Username"`
	Password     string `mapstructure:"Password"`
	Host         string `mapstructure:"Host"`
	Port         string `mapstructure:"Port"`
	DatabaseName string `mapstructure:"Database"`
}

var (
	mysql *Mysql
)

func GetMysql() *Mysql {
	if mysql == nil {
		if !config.IsSet("Mysql") {
			return nil
		}
		mysql = &Mysql{}
		err := config.UnmarshalKey("Mysql", mysql)
		if err != nil {
			log.Panic(err)
		}
	}
	return mysql
}
