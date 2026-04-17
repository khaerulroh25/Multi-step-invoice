package config

import (
	"fmt"
	"log"
	"os"
	"time"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

func ConnectDB() {
	var db *gorm.DB
	var err error

	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		os.Getenv("DB_USER"),
		os.Getenv("DB_PASS"),
		os.Getenv("DB_HOST"),
		os.Getenv("DB_PORT"),
		os.Getenv("DB_NAME"),
	)

	for i := 0; i < 10; i++ {
	db, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err == nil {
			break
		}

		fmt.Println("Waiting for database...")
		time.Sleep(2 * time.Second)
	}

	if err != nil {
		log.Fatal("DB Connection Failed:", err)
	}

	DB = db
	fmt.Println("Database connected!")
}