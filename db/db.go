package db

import (
	"fmt"
	"strail/config"
	"strail/models"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
)

var db *gorm.DB
var err error

func Init() {
	configuration := config.GetConfig()
	connect := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
		configuration.DB_HOST,
		configuration.DB_PORT,
		configuration.DB_USER,
		configuration.DB_PASSWORD,
		configuration.DB_NAME,
	)
	db, err = gorm.Open("postgres", connect)

	if err != nil {
		panic(err.Error())
	}

	db.AutoMigrate(&models.User{}, &models.ShoppingList{}, &models.Schedule{})
	// db.Model(&models.Schedule{}).AddForeignKey("username", "users(username)", "CASCADE", "CASCADE")
}

// func for connect database
func ConnectDB() *gorm.DB {
	return db
}
