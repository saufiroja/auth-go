package config

import (
	"echo/auth/entity"
	"log"
	"os"

	"github.com/joho/godotenv"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type DBConfig struct {
	DB_HOST string
	DB_PORT string
	DB_USER string
	DB_PASS string
	DB_NAME string
}

func InitDB(conf DBConfig) *gorm.DB {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	username := os.Getenv("DB_USER")
	password := os.Getenv("DB_PASS")
	host := os.Getenv("DB_HOST")
	port := os.Getenv("DB_PORT")
	name := os.Getenv("DB_NAME")

	dns := "host=" + host + " port=" + port + " user=" + username + " dbname=" + name + " password=" + password + " sslmode=disable"
	db, err := gorm.Open(postgres.Open(dns), &gorm.Config{})

	db.AutoMigrate(&entity.User{})

	if err != nil {
		log.Fatal(err)
	}

	return db
}
