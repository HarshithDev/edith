package database

import (
	"edith/models"
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

var DB *gorm.DB

func ConnectToDB() {
	err := godotenv.Load(".env.staging")

	if err != nil {
		log.Fatal("Error in loading ENV file \n", err)
	}

	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable TimeZone=Asia/Kolkata", os.Getenv("PSQL_HOST"), os.Getenv("PSQL_USER"), os.Getenv("PSQL_PASS"), os.Getenv("PSQL_DBNAME"), os.Getenv("PSQL_PORT"))

	log.Print("Connecting to Database...")

	DB, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if err != nil {
		log.Fatal("Failed to connect to database. \n", err)
		os.Exit(2)

	}

	log.Println("Connected !")

	// turned on logger on the info mode
	DB.Logger = logger.Default.LogMode(logger.Info)

	log.Print("Running the migration...")
	DB.AutoMigrate(&models.User{}, &models.Claims{})

}
