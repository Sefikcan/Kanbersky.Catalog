package config

import (
	"fmt"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func ConnectProductDB() *gorm.DB {
	db, err := gorm.Open(mysql.Open(""))

	if err != nil {
		fmt.Println(err.Error())
		panic("Cannot connect to DB")
	}

	return db
}
