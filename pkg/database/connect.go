package database

import (
	"fmt"
	"quizcms/pkg/config"
	"quizcms/pkg/model"
	"strconv"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
)

var DB *gorm.DB

func ConnectDB() {
	var err error
	p := config.Config("DB_PORT")
	port, err := strconv.ParseUint(p, 10, 32)
	if err != nil {
		panic("Can't parse the port: " + p)
	}
	DB, err = gorm.Open("postgres", fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable", config.Config("DB_HOST"), port, config.Config("DB_USER"), config.Config("DB_PASSWORD"), config.Config("DB_NAME")))

	if err != nil {
		panic("Failed to connect database")
	}

	fmt.Println("Connection opened to databae")
	DB.AutoMigrate(&model.Category{})
	DB.AutoMigrate(&model.Language{})
	DB.AutoMigrate(&model.Platform{})
	DB.AutoMigrate(&model.Question{})
	DB.AutoMigrate(&model.Quiz{})
	fmt.Println("Database Migrated")
}
