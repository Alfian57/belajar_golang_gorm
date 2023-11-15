package util

import (
	"fmt"

	"github.com/Alfian57/belajar_golang_gorm/helper"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func OpenDbConnection() *gorm.DB {
	config := GetConfig()

	databaseUser := config.GetString("DATABASE_USER")
	databasePassword := config.GetString("DATABASE_PASSWORD")
	databaseHost := config.GetString("DATABASE_HOST")
	databasePort := config.GetString("DATABASE_PORT")
	databaseName := config.GetString("DATABASE_NAME")

	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?parseTime=true", databaseUser, databasePassword, databaseHost, databasePort, databaseName)
	dialect := mysql.Open(dsn)
	db, err := gorm.Open(dialect, &gorm.Config{})
	helper.PanicIfError(err)

	return db
}
