package database

import (
	"fmt"
	"log"

	"com.app/pos-app/config"
	"com.app/pos-app/models"
	"github.com/joho/godotenv"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func ConnectionDB() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
		return
	}
	config := config.GetAppConfig()
	dsn := fmt.Sprintf(
		"host=%s user=%s password=%s dbname=%s port=%s sslmode=disable",
		config.DBConfig.Host,
		config.DBConfig.User,
		config.DBConfig.Password,
		config.DBConfig.Name,
		config.DBConfig.Port,
	)
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("❌ Failed to connect to database: ", err)
		return
	}

	DB = db
	log.Println("✅ Database connected")

	Migrate()
}

func Migrate() {
	err := DB.AutoMigrate(
		&models.Product{},
		&models.Category{},
		&models.PaymentMethod{},
		&models.Role{},
		&models.Users{},
		&models.Sales{},
		&models.SaleItem{},
	)
	if err != nil {
		log.Fatal("❌ Failed to migrate database: ", err)
		return
	}
	log.Println("✅ Migration completed")
}
