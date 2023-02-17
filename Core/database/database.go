package infras

import (
	"fmt"
	"log"
	"os"

	"github.com/Jingyu-shanshan/Xiaoxuan-Thesis/Core/models"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

type Dbinstance struct {
	Db *gorm.DB
}

var DB Dbinstance

func ConnectDb() {
	var host, user, password, dbname string = "db", os.Getenv("DB_USER"), os.Getenv("DB_PASSWORD"), os.Getenv("DB_NAME")

	dsn := fmt.Sprintf(
		"host=%s user=%s password=%s dbname=%s port=5432 sslmode=disable",
		host, user, password, dbname)

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Info),
	})

	if err != nil {
		log.Fatal("Failed to connect to database. \n", err)
		os.Exit(1)
	}

	log.Println("connected")
	db.Logger = logger.Default.LogMode(logger.Info)

	log.Println("running migrations")
	db.AutoMigrate(&models.User{})

	DB = Dbinstance{
		Db: db,
	}
}
