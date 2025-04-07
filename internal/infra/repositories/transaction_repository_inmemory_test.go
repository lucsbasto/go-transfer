package repositories_test

import (
	"context"
	"errors"
	"sync"
	"testing"
	"time"

	"go-transfer/internal/domain/entities"
	"go-transfer/internal/domain/port"

	"github.com/stretchr/testify/assert"
)

type TransactionRepositoryInMemory struct {
	transactions map[int64]*entities.Transaction
	mu           sync.RWMutex
	nextID       int64
}

func NewTransactionRepositoryInMemory() port.TransactionRepository {
	return &TransactionRepositoryInMemory{
		transactions: make(map[int64]*entities.Transaction),
		mu:           sync.RWMutex{},
		nextID:       1,
	}
}

func (r *TransactionRepositoryInMemory) Create(ctx context.Context, transfer *entities.Transaction) (int64, error) {
	r.mu.Lock()
	defer r.mu.Unlock()
	transfer.ID = r.nextID
	r.transactions[transfer.ID] = transfer
	r.nextID++
	return transfer.ID, nil
}

func (r *TransactionRepositoryInMemory) UpdateStatus(ctx context.Context, id int64, status entities.TransactionStatus) error {
	r.mu.Lock()
	defer r.mu.Unlock()
	transaction, ok := r.transactions[id]
	if !ok {
		return errors.New("transação não encontrada")
	}
	transaction.Status = status
	transaction.UpdatedAt = time.Now()
	return nil
}

func (r *TransactionRepositoryInMemory) GetByID(ctx context.Context, id int64) (*entities.TransactionStatus, error) {
	r.mu.RLock()
	defer r.mu.RUnlock()
	transaction, ok := r.transactions[id]
	if !ok {
		return nil, errors.New("transação não encontrada")
	}
	return &transaction.Status, nil
}

func TestTransactionRepositoryInMemory_Create(t *testing.T) {
	repo := NewTransactionRepositoryInMemory()
	ctx := context.Background()

	transfer := &entities.Transaction{
		SenderID:   1,
		ReceiverID: 2,
		Amount:     50.00,
		Status:     entities.TransactionStatusPending,
		CreatedAt:  time.Now(),
	}

	id, err := repo.Create(ctx, transfer)
	assert.NoError(t, err)
	assert.NotZero(t, id)

	status, err := repo.GetByID(ctx, id)
	assert.NoError(t, err)
	assert.Equal(t, transfer.Status, *status)
}

func TestTransactionRepositoryInMemory_UpdateStatus_Success(t *testing.T) {
	repo := NewTransactionRepositoryInMemory()
	ctx := context.Background()

	initialTransfer := &entities.Transaction{
		SenderID:   1,
		ReceiverID: 2,
		Amount:     50.00,
		Status:     entities.TransactionStatusPending,
		CreatedAt:  time.Now(),
	}
	id, err := repo.Create(ctx, initialTransfer)
	assert.NoError(t, err)

	newStatus := entities.TransactionStatusCompleted
	err = repo.UpdateStatus(ctx, id, newStatus)
	assert.NoError(t, err)

	retrievedStatus, err := repo.GetByID(ctx, id)
	assert.NoError(t, err)
	assert.Equal(t, newStatus, *retrievedStatus)
}

func TestTransactionRepositoryInMemory_UpdateStatus_NotFound(t *testing.T) {
	repo := NewTransactionRepositoryInMemory()
	ctx := context.Background()

	nonExistentID := int64(999)
	newStatus := entities.TransactionStatusFailed
	err := repo.UpdateStatus(ctx, nonExistentID, newStatus)
	assert.Error(t, err)
	assert.ErrorContains(t, err, "transação não encontrada")

	status, err := repo.GetByID(ctx, nonExistentID)
	assert.Error(t, err)
	assert.ErrorContains(t, err, "transação não encontrada")
	assert.Nil(t, status)
}

func TestTransactionRepositoryInMemory_GetByID_Found(t *testing.T) {
	repo := NewTransactionRepositoryInMemory()
	ctx := context.Background()

	transfer := &entities.Transaction{
		SenderID:   3,
		ReceiverID: 4,
		Amount:     100.00,
		Status:     entities.TransactionStatusPending,
		CreatedAt:  time.Now(),
	}
	id, err := repo.Create(ctx, transfer)
	assert.NoError(t, err)

	status, err := repo.GetByID(ctx, id)
	assert.NoError(t, err)
	assert.Equal(t, transfer.Status, *status)
}

func TestTransactionRepositoryInMemory_GetByID_NotFound(t *testing.T) {
	repo := NewTransactionRepositoryInMemory()
	ctx := context.Background()

	status, err := repo.GetByID(ctx, 999)
	assert.Error(t, err)
	assert.ErrorContains(t, err, "transação não encontrada")
	assert.Nil(t, status)
}
