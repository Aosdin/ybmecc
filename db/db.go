package db

import (
	"ybmecc/config"
	//"ybmecc/model"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

var db *gorm.DB
var err error

func Init() {
	configuration := config.GetConfig()
	connect_string := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8&parseTime=True&loc=Local", configuration.DB_USERNAME, configuration.DB_PASSWORD, configuration.DB_HOST, configuration.DB_PORT, configuration.DB_NAME)

	db, err = gorm.Open("mysql", connect_string)
	db.LogMode(true)
	// defer db.Close()
	if err != nil {
		panic("DB Connection Error")
	}
	//db.AutoMigrate(&model.User{})

}

func DbManager() *gorm.DB {
	return db
}
