package port

import (
	"context"

	"go-transfer/internal/domain/entities"
)

type TransactionRepository interface {
	Create(ctx context.Context, transfer *entities.Transaction) (int64, error)
	UpdateStatus(ctx context.Context, id int64, status entities.TransactionStatus) error
}
