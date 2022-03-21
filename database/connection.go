package database

import (
	"fmt"
	"news/models"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

func Connect() {
	dsn := "root:pwd@tcp(0.0.0.0:33060)/dbname?charset=utf8mb4&parseTime=True&loc=Local"
	connection, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})

	if err != nil {
		fmt.Println("Couldn't connect to db")
	}

	DB = connection

	connection.AutoMigrate(&models.News{})
	connection.AutoMigrate(&models.Tags{})
	connection.AutoMigrate(&models.NewsTags{})
}
