package database

import (
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"log"
	"os"
	"todo-gorilla/models"
)

func Connect() {
	var err error

	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		os.Getenv("DB_USER"),
		os.Getenv("DB_PASS"),
		os.Getenv("DB_HOST"),
		os.Getenv("DB_PORT"),
		os.Getenv("DB_NAME"))

	DB, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("Error While Connecting Database")
	}
	fmt.Println("DB Connection has been established successfully.")

	if err = DB.AutoMigrate(&models.User{}, &models.Todo{}); err != nil {
		log.Fatal("Error While Migrating.")
	}
	fmt.Println("Migrate is done.")
}
