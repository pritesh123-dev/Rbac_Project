package config

import (
    "gorm.io/driver/postgres"
    "gorm.io/gorm"
	"github.com/pritesh/gobackend/models"
    "log"
)

var DB *gorm.DB

func Connect() {
    dsn := "host=localhost user=postgres password=BlueLock@321 dbname=document_upload port=5432 sslmode=disable"
    var err error
    DB, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
    if err != nil {
        log.Fatal("Failed to connect to database:", err)
    }
	DB.AutoMigrate(&models.Document{})
	DB.AutoMigrate(&models.User{})
    log.Println("Connected to database.")
}
