package repositories

import (
	"context"

	"go-transfer/internal/domain/entities"

	"gorm.io/gorm"
)

type WalletRepository struct {
	db *gorm.DB
}

func NewWalletRepository(db *gorm.DB) *WalletRepository {
	return &WalletRepository{
		db: db,
	}
}

func (r *WalletRepository) Create(ctx context.Context, wallet *entities.Wallet) error {
	return r.db.WithContext(ctx).Create(wallet).Error
}

func (r *WalletRepository) GetByID(ctx context.Context, id int64) (*entities.Wallet, error) {
	wallet := &entities.Wallet{}
	err := r.db.WithContext(ctx).First(wallet, id).Error
	if err != nil {
		return nil, err
	}
	return wallet, nil
}

func (r *WalletRepository) GetByOwnerID(ctx context.Context, ownerID int64) (*entities.Wallet, error) {
	wallet := &entities.Wallet{}
	err := r.db.WithContext(ctx).Where("owner_id = ?", ownerID).First(wallet).Error
	if err != nil {
		return nil, err
	}
	return wallet, nil
}

func (r *WalletRepository) UpdateBalance(ctx context.Context, id int64, balance float64) error {
	return r.db.WithContext(ctx).Model(&entities.Wallet{}).Where("id = ?", id).Update("balance", balance).Error
}
