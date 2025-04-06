package entities

import (
	"time"

	"gorm.io/gorm"
)

type User struct {
	ID                int64          `gorm:"primaryKey"`
	FullName          string         `gorm:"not null"`
	Document          string         `gorm:"unique;not null"`
	Email             string         `gorm:"unique;not null"`
	Password          string         `gorm:"not null"`
	Wallet            Wallet         `gorm:"foreignKey:OwnerID"`
	SentTransfers     []Transaction  `gorm:"foreignKey:SenderID"`
	ReceivedTransfers []Transaction  `gorm:"foreignKey:ReceiverID"`
	CreatedAt         time.Time      `gorm:"autoCreateTime"`
	UpdatedAt         time.Time      `gorm:"autoUpdateTime"`
	DeletedAt         gorm.DeletedAt `gorm:"index"`
}
