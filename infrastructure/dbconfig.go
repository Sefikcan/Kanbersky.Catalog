package dbConfig

import (
	"fmt"
	"kanbersky/common/constants"
	"kanbersky/infrastructure/entities"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func ConnectProductDB() *gorm.DB {
	db, err := gorm.Open(mysql.Open(constants.PRODUCTCONNSTRING))

	if err != nil {
		fmt.Println(err.Error())
		panic("Cannot connect to DB")
	}

	db.AutoMigrate(&entities.Product{})
	return db
}
