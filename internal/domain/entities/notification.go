package entities

import (
	"time"

	"gorm.io/gorm"
)

type NotificationStatus string

const (
	NotificationStatusPending NotificationStatus = "PENDING"
	NotificationStatusSent    NotificationStatus = "SENT"
	NotificationStatusFailed  NotificationStatus = "FAILED"
)

type Notification struct {
	ID            int64              `gorm:"primaryKey"`
	ReceiverID    int64              `gorm:"not null;index"`
	TransactionID int64              `gorm:"not null;index"`
	Amount        float64            `gorm:"not null"`
	Status        NotificationStatus `gorm:"not null default 'PENDING'"`
	CreatedAt     time.Time          `gorm:"autoCreateTime"`
	UpdatedAt     time.Time          `gorm:"autoUpdateTime"`
	DeletedAt     gorm.DeletedAt     `gorm:"index"`
	Receiver      User               `gorm:"foreignKey:ReceiverID"`
	Transaction   Transaction        `gorm:"foreignKey:TransactionID"`
}
