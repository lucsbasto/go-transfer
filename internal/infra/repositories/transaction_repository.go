package repositories

import (
	"context"

	"go-transfer/internal/domain/entities"

	"gorm.io/gorm"
)

type TransactionRepository struct {
	db *gorm.DB
}

func NewTransactionRepository(db *gorm.DB) *TransactionRepository {
	return &TransactionRepository{
		db: db,
	}
}

func (r *TransactionRepository) Create(ctx context.Context, transfer *entities.Transaction) (int64, error) {
	result := r.db.WithContext(ctx).Create(transfer)
	if result.Error != nil {
		return 0, result.Error
	}
	return transfer.ID, nil
}

func (r *TransactionRepository) UpdateStatus(ctx context.Context, id int64, status entities.TransactionStatus) error {
	return r.db.WithContext(ctx).Model(&entities.Transaction{}).Where("id = ?", id).Update("status", status).Error
}

func (r *TransactionRepository) GetByID(ctx context.Context, id int64) (*entities.TransactionStatus, error) {
	transaction := &entities.Transaction{}
	err := r.db.WithContext(ctx).First(transaction, id).Error
	if err != nil {
		return nil, err
	}
	return &transaction.Status, nil
}
