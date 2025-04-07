package repositories

import (
	"context"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"go-transfer/internal/domain/entities"
)

func TestWalletRepositoryInMemory_Create(t *testing.T) {
	repo := NewWalletRepositoryInMemory()
	ctx := context.Background()

	wallet := &entities.Wallet{
		OwnerID:   1,
		Balance:   100.50,
		CreatedAt: time.Now(),
	}

	err := repo.Create(ctx, wallet)
	assert.NoError(t, err)
	assert.NotZero(t, wallet.ID)
	retrievedWallet, err := repo.GetByOwnerID(ctx, wallet.ID)
	assert.NoError(t, err)
	assert.Equal(t, wallet.OwnerID, retrievedWallet.OwnerID)
	assert.Equal(t, wallet.Balance, retrievedWallet.Balance)
}

func TestWalletRepositoryInMemory_GetByID_Found(t *testing.T) {
	repo := NewWalletRepositoryInMemory()
	ctx := context.Background()

	expectedWallet := &entities.Wallet{
		OwnerID:   2,
		Balance:   50.00,
		CreatedAt: time.Now(),
	}
	err := repo.Create(ctx, expectedWallet)
	assert.NoError(t, err)

	retrievedWallet, err := repo.GetByID(ctx, expectedWallet.ID)
	assert.NoError(t, err)
	assert.Equal(t, expectedWallet, retrievedWallet)
}

func TestWalletRepositoryInMemory_GetByID_NotFound(t *testing.T) {
	repo := NewWalletRepositoryInMemory()
	ctx := context.Background()

	retrievedWallet, err := repo.GetByID(ctx, 999)
	assert.Error(t, err)
	assert.ErrorContains(t, err, "carteira não encontrada")
	assert.Nil(t, retrievedWallet)
}

func TestWalletRepositoryInMemory_GetByOwnerID_Found(t *testing.T) {
	repo := NewWalletRepositoryInMemory()
	ctx := context.Background()

	expectedWallet := &entities.Wallet{
		OwnerID:   3,
		Balance:   120.75,
		CreatedAt: time.Now(),
	}
	err := repo.Create(ctx, expectedWallet)
	assert.NoError(t, err)

	retrievedWallet, err := repo.GetByOwnerID(ctx, expectedWallet.OwnerID)
	assert.NoError(t, err)
	assert.Equal(t, expectedWallet, retrievedWallet)
}

func TestWalletRepositoryInMemory_GetByOwnerID_NotFound(t *testing.T) {
	repo := NewWalletRepositoryInMemory()
	ctx := context.Background()

	retrievedWallet, err := repo.GetByOwnerID(ctx, 999)
	assert.Error(t, err)
	assert.ErrorContains(t, err, "carteira não encontrada para o OwnerID")
	assert.Nil(t, retrievedWallet)
}

func TestWalletRepositoryInMemory_UpdateBalance_Success(t *testing.T) {
	repo := NewWalletRepositoryInMemory()
	ctx := context.Background()

	initialWallet := &entities.Wallet{
		OwnerID:   4,
		Balance:   75.20,
		CreatedAt: time.Now(),
	}
	err := repo.Create(ctx, initialWallet)
	assert.NoError(t, err)

	newBalance := 150.90
	err = repo.UpdateBalance(ctx, initialWallet.ID, newBalance)
	assert.NoError(t, err)

	retrievedWallet, err := repo.GetByOwnerID(ctx, initialWallet.OwnerID)
	assert.NoError(t, err)
	assert.Equal(t, newBalance, retrievedWallet.Balance)
}

func TestWalletRepositoryInMemory_UpdateBalance_NotFound(t *testing.T) {
	repo := NewWalletRepositoryInMemory()
	ctx := context.Background()

	err := repo.UpdateBalance(ctx, 999, 200.00)
	assert.Error(t, err)
	assert.ErrorContains(t, err, "carteira não encontrada")

	retrievedWallet, err := repo.GetByID(ctx, 999)
	assert.Error(t, err)
	assert.ErrorContains(t, err, "carteira não encontrada")
	assert.Nil(t, retrievedWallet)
}
