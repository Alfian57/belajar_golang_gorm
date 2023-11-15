package util

import (
	"github.com/Alfian57/belajar_golang_gorm/helper"
	"github.com/spf13/viper"
)

func GetConfig() *viper.Viper {
	config := viper.New()
	config.SetConfigFile(".env")
	config.AddConfigPath(".")

	err := config.ReadInConfig()
	helper.PanicIfError(err)

	return config
}
