package database

import (
	"fmt"
	"log"
	"os"

	"github.com/hylim-tech-lover/restful-services-go/models"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

type DbInstance struct {
	Db *gorm.DB
}

var DB DbInstance

func ConnectDb() {
	dsn := fmt.Sprintf("host=db user=%s password=%s dbname =%s port=5432 sslmode=disable TimeZone=Asia/Singapore",
		os.Getenv("DB_USER"),
		os.Getenv("DB_PASSWORD"),
		os.Getenv("DB_NAME"),
	)
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Info),
	})

	if err != nil {
		log.Fatalf("Failed to connect to database \n %v", err)
	}

	log.Println("Database connected")
	db.Logger = logger.Default.LogMode(logger.Info)

	log.Println("Running database migration")
	db.AutoMigrate(&models.Fact{})

	DB = DbInstance{
		Db: db,
	}
}
