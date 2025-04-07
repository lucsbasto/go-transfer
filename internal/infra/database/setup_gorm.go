package database

import (
	"fmt"
	"go-transfer/internal/domain/entities"
	"go-transfer/internal/env"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func SetupDB(config *env.Config) (*gorm.DB, error) {
	fmt.Println("Connecting DB...")

	dsn := fmt.Sprintf(
		"user=%s password=%s dbname=%s host=%s port=%s sslmode=disable",
		config.DatabaseUser,
		config.DatabasePassword,
		config.DatabaseName,
		config.DatabaseHost,
		config.DatabasePort,
	)

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
