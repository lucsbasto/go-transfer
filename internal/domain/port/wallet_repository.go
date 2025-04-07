package port

import (
	"context"

	"go-transfer/internal/domain/entities"
)

type WalletRepository interface {
	GetByID(ctx context.Context, id int64) (*entities.Wallet, error)
	GetByOwnerID(ctx context.Context, ownerID int64) (*entities.Wallet, error)
	UpdateBalance(ctx context.Context, id int64, balance float64) error
	Create(ctx context.Context, wallet *entities.Wallet) error
}
