package domain

import (
	"time"

	"gorm.io/gorm"
)

type WalletType string

const (
	CommonWallet   WalletType = "COMMON"
	MerchantWallet WalletType = "MERCHANT"
)

type Wallet struct {
	ID        int64          `gorm:"primaryKey"`
	OwnerID   int64          `gorm:"not null;index"`
	Balance   float64        `gorm:"default:0.00"`
	Type      WalletType     `gorm:"type:text;default:'COMMON'"`
	CreatedAt time.Time      `gorm:"autoCreateTime"`
	UpdatedAt time.Time      `gorm:"autoUpdateTime"`
	DeletedAt gorm.DeletedAt `gorm:"index"`
}
