package database

import (
	"log"

	"github.com/AstroArbaaz/DevOps/complex_bottle/models"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

type Dbinstance struct {
	Db *gorm.DB
}

var DB Dbinstance

func ConnectDb() {
	dsn := "host=localhost user=postgres password='asdf1234' dbname=devops_test port=5432 sslmode=disable TimeZone=Asia/Kolkata"

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Info),
	})

	if err != nil {
		log.Fatal("Failed to connect to database. \n", err)
	}

	log.Println("database connected")
	db.Logger = logger.Default.LogMode(logger.Info)

	log.Println("Running Maigrations!")
	db.AutoMigrate(&models.User{}, &models.Product{}, &models.Category{}, &models.OrderItem{}, &models.Order{})

	DB = Dbinstance{
		Db: db,
	}
}