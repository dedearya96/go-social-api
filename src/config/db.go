package config

import (
	"fmt"
	"log"
	"os"
	"github.com/joho/godotenv"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DBAuth *gorm.DB
var DBAccount *gorm.DB

func InitDB() {
	// Load file .env
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	// Koneksi DB Auth
	dsnAuth := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		os.Getenv("DB_USERNAME_AUTH"), os.Getenv("DB_PASSWORD_AUTH"),
		os.Getenv("DB_HOST_AUTH"), os.Getenv("DB_PORT_AUTH"),
		os.Getenv("DB_DATABASE_AUTH"))

	DBAuth, err = gorm.Open(mysql.Open(dsnAuth), &gorm.Config{})
	if err != nil {
		log.Fatal("Gagal koneksi ke DB Auth: ", err)
	}

	// Koneksi DB Account
	dsnAccount := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		os.Getenv("DB_USERNAME_ACCOUNT"), os.Getenv("DB_PASSWORD_ACCOUNT"),
		os.Getenv("DB_HOST_ACCOUNT"), os.Getenv("DB_PORT_ACCOUNT"),
		os.Getenv("DB_DATABASE_ACCOUNT"))

	DBAccount, err = gorm.Open(mysql.Open(dsnAccount), &gorm.Config{})
	if err != nil {
		log.Fatal("Gagal koneksi ke DB Account: ", err)
	}
}