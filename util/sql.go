package util

import (
	"fmt"
	"github.com/spf13/viper"
)

func MysqlDSN() string {
	return fmt.Sprintf(
		"%s:%s@tcp(%s:%s)/%s?charset=utf8&parseTime=True&loc=Local",
		viper.Get("database.user"),
		viper.Get("database.password"),
		viper.Get("database.host"),
		viper.Get("database.port"),
		viper.Get("database.name"),
	)
}
