package databse

import (
	"rentoutlkApi/models"
	"gorm.io/gorm"
	"gorm.io/driver/mysql"
)

var DB *gorm.DB

func Connect(){
	connection, error := gorm.Open(mysql.Open("root:root@/rentout"), &gorm.Config{})

	if error != nil {
		panic("couldn't connect to the database")
	}

	DB = connection

	connection.AutoMigrate(&models.User{}, &models.Product{},&models.ProductCategory{}, &models.PropertyDetail{})
}

