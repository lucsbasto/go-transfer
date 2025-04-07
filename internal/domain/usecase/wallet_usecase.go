package usecase

import (
	"context"
	"go-transfer/internal/domain/entities"
	"go-transfer/internal/domain/port"
)

type WalletInput struct {
	OwnerID int64               `json:"owner_id"`
	Type    entities.WalletType `json:"type"`
	Balance float64             `json:"balance"`
}

type Wallet struct {
	walletRepo port.WalletRepository
}

func NewWallet(walletRepo port.WalletRepository) *Wallet {
	return &Wallet{
		walletRepo: walletRepo,
	}
}

func (w *Wallet) CreateWallet(ctx context.Context, input WalletInput) error {
	wallet := &entities.Wallet{
		OwnerID: input.OwnerID,
		Type:    input.Type,
		Balance: input.Balance,
	}

	return w.walletRepo.Create(ctx, wallet)
}

func (w *Wallet) GetWalletByID(ctx context.Context, id int64) (*entities.Wallet, error) {
	return w.walletRepo.GetByID(ctx, id)
}

func (w *Wallet) GetWalletByOwnerID(ctx context.Context, ownerID int64) (*entities.Wallet, error) {
	return w.walletRepo.GetByOwnerID(ctx, ownerID)
}

func (w *Wallet) UpdateWalletBalance(ctx context.Context, id int64, balance float64) error {
	return w.walletRepo.UpdateBalance(ctx, id, balance)
}
