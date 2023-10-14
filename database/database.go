package database

import (
	"fmt"
	"os"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var db *gorm.DB

func GetInstance() *gorm.DB {
	if db == nil {
		db = connectDB()
		runSeeder()
	}

	return db
}

func connectDB() *gorm.DB {
	dsn := fmt.Sprintf("host=%v user=%v password=%v dbname=%v port=%v sslmode=%v TimeZone=%v",
		os.Getenv("DB_HOST"),
		os.Getenv("DB_USER"),
		os.Getenv("DB_PASSWORD"),
		os.Getenv("DB_NAME"),
		os.Getenv("DB_PORT"),
		os.Getenv("DB_SSLMODE"),
		os.Getenv("DB_TIMEZONE"),
	)

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if err != nil {
		panic(err)
	}

	return db
}
