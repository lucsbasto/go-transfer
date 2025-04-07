package usecase

import (
	"context"
	"errors"
	"testing"
	"time"

	"go-transfer/internal/domain/entities"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

// MockWalletRepository é um mock da interface port.WalletRepository para testes.
type MockWalletRepository struct {
	mock.Mock
}

func (m *MockWalletRepository) Create(ctx context.Context, wallet *entities.Wallet) error {
	args := m.Called(ctx, wallet)
	return args.Error(0)
}

func (m *MockWalletRepository) GetByID(ctx context.Context, id int64) (*entities.Wallet, error) {
	args := m.Called(ctx, id)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).(*entities.Wallet), args.Error(1)
}

func (m *MockWalletRepository) GetByOwnerID(ctx context.Context, ownerID int64) (*entities.Wallet, error) {
	args := m.Called(ctx, ownerID)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).(*entities.Wallet), args.Error(1)
}

func (m *MockWalletRepository) UpdateBalance(ctx context.Context, id int64, balance float64) error {
	args := m.Called(ctx, id, balance)
	return args.Error(0)
}

func TestWalletUseCase_CreateWallet_Success(t *testing.T) {
	mockRepo := new(MockWalletRepository)
	walletUseCase := NewWallet(mockRepo)
	ctx := context.Background()

	input := WalletInput{
		OwnerID: 1,
		Type:    entities.CommonWallet,
		Balance: 100.0,
	}

	wallet := &entities.Wallet{
		OwnerID: input.OwnerID,
		Type:    input.Type,
		Balance: input.Balance,
	}

	mockRepo.On("Create", ctx, wallet).Return(nil)

	err := walletUseCase.CreateWallet(ctx, input)
	assert.NoError(t, err)
	mockRepo.AssertExpectations(t)
}

func TestWalletUseCase_CreateWallet_Error(t *testing.T) {
	mockRepo := new(MockWalletRepository)
	walletUseCase := NewWallet(mockRepo)
	ctx := context.Background()

	input := WalletInput{
		OwnerID: 1,
		Type:    entities.MerchantWallet,
		Balance: 100.0,
	}

	wallet := &entities.Wallet{
		OwnerID: input.OwnerID,
		Type:    input.Type,
		Balance: input.Balance,
	}

	mockRepo.On("Create", ctx, wallet).Return(errors.New("database error"))

	err := walletUseCase.CreateWallet(ctx, input)
	assert.Error(t, err)
	assert.Equal(t, "database error", err.Error())
	mockRepo.AssertExpectations(t)
}

func TestWalletUseCase_GetWalletByID_Success(t *testing.T) {
	mockRepo := new(MockWalletRepository)
	walletUseCase := NewWallet(mockRepo)
	ctx := context.Background()
	walletID := int64(1)

	expectedWallet := &entities.Wallet{
		ID:        walletID,
		OwnerID:   1,
		Type:      entities.MerchantWallet,
		Balance:   100.0,
		CreatedAt: time.Now(),
	}

	mockRepo.On("GetByID", ctx, walletID).Return(expectedWallet, nil)

	retrievedWallet, err := walletUseCase.GetWalletByID(ctx, walletID)
	assert.NoError(t, err)
	assert.Equal(t, expectedWallet, retrievedWallet)
	mockRepo.AssertExpectations(t)
}

func TestWalletUseCase_GetWalletByID_NotFound(t *testing.T) {
	mockRepo := new(MockWalletRepository)
	walletUseCase := NewWallet(mockRepo)
	ctx := context.Background()
	walletID := int64(1)

	mockRepo.On("GetByID", ctx, walletID).Return(nil, errors.New("wallet not found"))

	retrievedWallet, err := walletUseCase.GetWalletByID(ctx, walletID)
	assert.Error(t, err)
	assert.Equal(t, "wallet not found", err.Error())
	assert.Nil(t, retrievedWallet)
	mockRepo.AssertExpectations(t)
}

func TestWalletUseCase_GetWalletByOwnerID_Success(t *testing.T) {
	mockRepo := new(MockWalletRepository)
	walletUseCase := NewWallet(mockRepo)
	ctx := context.Background()
	ownerID := int64(1)

	expectedWallet := &entities.Wallet{
		ID:        1,
		OwnerID:   ownerID,
		Type:      entities.MerchantWallet,
		Balance:   100.0,
		CreatedAt: time.Now(),
	}

	mockRepo.On("GetByOwnerID", ctx, ownerID).Return(expectedWallet, nil)

	retrievedWallet, err := walletUseCase.GetWalletByOwnerID(ctx, ownerID)
	assert.NoError(t, err)
	assert.Equal(t, expectedWallet, retrievedWallet)
	mockRepo.AssertExpectations(t)
}

func TestWalletUseCase_GetWalletByOwnerID_NotFound(t *testing.T) {
	mockRepo := new(MockWalletRepository)
	walletUseCase := NewWallet(mockRepo)
	ctx := context.Background()
	ownerID := int64(1)

	mockRepo.On("GetByOwnerID", ctx, ownerID).Return(nil, errors.New("wallet not found for owner"))

	retrievedWallet, err := walletUseCase.GetWalletByOwnerID(ctx, ownerID)
	assert.Error(t, err)
	assert.Equal(t, "wallet not found for owner", err.Error())
	assert.Nil(t, retrievedWallet)
	mockRepo.AssertExpectations(t)
}

func TestWalletUseCase_UpdateWalletBalance_Success(t *testing.T) {
	mockRepo := new(MockWalletRepository)
	walletUseCase := NewWallet(mockRepo)
	ctx := context.Background()
	walletID := int64(1)
	newBalance := 150.0

	mockRepo.On("UpdateBalance", ctx, walletID, newBalance).Return(nil)

	err := walletUseCase.UpdateWalletBalance(ctx, walletID, newBalance)
	assert.NoError(t, err)
	mockRepo.AssertExpectations(t)
}

func TestWalletUseCase_UpdateWalletBalance_Error(t *testing.T) {
	mockRepo := new(MockWalletRepository)
	walletUseCase := NewWallet(mockRepo)
	ctx := context.Background()
	walletID := int64(1)
	newBalance := 150.0

	mockRepo.On("UpdateBalance", ctx, walletID, newBalance).Return(errors.New("failed to update balance"))

	err := walletUseCase.UpdateWalletBalance(ctx, walletID, newBalance)
	assert.Error(t, err)
	assert.Equal(t, "failed to update balance", err.Error())
	mockRepo.AssertExpectations(t)
}
