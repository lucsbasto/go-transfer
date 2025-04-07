package database

import (
	"fmt"
	"go-transfer/internal/domain/entities"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func SetupDB() (*gorm.DB, error) {
	fmt.Println("Connecting DB...")

	dsn := "user=postgres password=postgres dbname=go-transfer sslmode=disable"

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		return nil, err
	}
	err = AutoMigrate(db)
	if err != nil {
		return nil, err
	}
	return db, nil
}

func AutoMigrate(db *gorm.DB) error {
	return db.AutoMigrate(&entities.User{}, &entities.Wallet{})
}
