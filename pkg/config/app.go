package config

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

const (
	dbUserName    = "root"
	dbPassword    = "secret"
	dbHost        = "localhost"
	confCharset   = "utf8"
	confParseTime = true
	confLocation  = "Local"
)

var (
	db *gorm.DB
)

func Connect() {
	//In order to handle time.Time correctly, you need to include parseTime as a parameter.
	//In order to fully support UTF-8 encoding, you need to change charset=utf8 to charset=utf8mb4
	//If you want to specify the host, you need to use ()
	d, err := gorm.Open("mysql", "user:password@(localhost)/dbname?charset=utf8&parseTime=True&loc=Local\n")
	if err != nil {
		panic(err)
	}
	db = d
}
func GetDB() *gorm.DB {
	return db
}
