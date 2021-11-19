package config

import (
	"fmt"
	// "os"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func ConnectToDatabase() *gorm.DB {
	DSN := "root:root@tcp(127.0.0.1:3306)/twistcode-test?charset=utf8mb4&parseTime=True&loc=Local"
	DB, err := gorm.Open(mysql.Open(DSN), &gorm.Config{})
	if err != nil {
		fmt.Println(err.Error())
		panic("Failed to Connect to Database")
	}

	return DB

}
