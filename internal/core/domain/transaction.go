package domain

import (
	"time"

	"gorm.io/gorm"
)

type TransactionStatus string

const (
	TransactionStatusPending   TransactionStatus = "PENDING"
	TransactionStatusCompleted TransactionStatus = "COMPLETED"
	TransactionStatusFailed    TransactionStatus = "FAILED"
)

type Transaction struct {
	ID         int64             `gorm:"primaryKey"`
	SenderID   int64             `gorm:"not null;index"`
	ReceiverID int64             `gorm:"not null;index"`
	Amount     float64           `gorm:"not null"`
	Status     TransactionStatus `gorm:"not null default 'PENDING'"`
	Sender     User              `gorm:"foreignKey:SenderID"`
	Receiver   User              `gorm:"foreignKey:ReceiverID"`
	CreatedAt  time.Time         `gorm:"autoCreateTime"`
	UpdatedAt  time.Time         `gorm:"autoUpdateTime"`
	DeletedAt  gorm.DeletedAt    `gorm:"index"`
}
