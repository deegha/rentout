package databse

import (
	"rentoutlkApi/models"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

func Connect() {
	connection, error := gorm.Open(mysql.Open("root:root@/rentout"), &gorm.Config{})

	if error != nil {
		panic("couldn't connect to the database")
	}

	DB = connection

	connection.AutoMigrate(&models.User{}, &models.Product{}, &models.ProductCategory{}, &models.PropertyDetail{}, &models.ImageDetail{}, &models.Location{})
}
