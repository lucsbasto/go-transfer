package usecase

import (
	"context"
	"testing"

	"go-transfer/internal/domain/entities"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

type mockUserRepo struct{ mock.Mock }

func (m *mockUserRepo) GetByID(ctx context.Context, id int64) (*entities.User, error) {
	args := m.Called(ctx, id)
	return args.Get(0).(*entities.User), args.Error(1)
}

func (m *mockUserRepo) Create(ctx context.Context, user *entities.User) error {
	args := m.Called(ctx, user)
	return args.Error(0)
}

func (m *mockUserRepo) GetByDocument(ctx context.Context, document string) (*entities.User, error) {
	args := m.Called(ctx, document)
	return args.Get(0).(*entities.User), args.Error(1)
}

func (m *mockUserRepo) GetByEmail(ctx context.Context, email string) (*entities.User, error) {
	args := m.Called(ctx, email)
	return args.Get(0).(*entities.User), args.Error(1)
}

type mockWalletRepo struct{ mock.Mock }

func (m *mockWalletRepo) GetByID(ctx context.Context, id int64) (*entities.Wallet, error) {
	args := m.Called(ctx, id)
	return args.Get(0).(*entities.Wallet), args.Error(1)
}

func (m *mockWalletRepo) GetByOwnerID(ctx context.Context, ownerID int64) (*entities.Wallet, error) {
	args := m.Called(ctx, ownerID)
	return args.Get(0).(*entities.Wallet), args.Error(1)
}

func (m *mockWalletRepo) Create(ctx context.Context, wallet *entities.Wallet) error {
	args := m.Called(ctx, wallet)
	return args.Error(0)
}

func (m *mockWalletRepo) UpdateBalance(ctx context.Context, walletID int64, newBalance float64) error {
	args := m.Called(ctx, walletID, newBalance)
	return args.Error(0)
}

type mockTransactionRepo struct{ mock.Mock }

func (m *mockTransactionRepo) Create(ctx context.Context, transaction *entities.Transaction) (int64, error) {
	args := m.Called(ctx, transaction)
	return args.Get(0).(int64), args.Error(1)
}

func (m *mockTransactionRepo) UpdateStatus(ctx context.Context, id int64, status entities.TransactionStatus) error {
	args := m.Called(ctx, id, status)
	return args.Error(0)
}

func (m *mockTransactionRepo) GetByID(ctx context.Context, id int64) (*entities.TransactionStatus, error) {
	args := m.Called(ctx, id)
	return args.Get(0).(*entities.TransactionStatus), args.Error(1)
}

type mockAuthService struct{ mock.Mock }

func (m *mockAuthService) Authorize(ctx context.Context) (bool, error) {
	args := m.Called(ctx)
	return args.Bool(0), args.Error(1)
}

type mockNotificationUseCase struct{ mock.Mock }

func (m *mockNotificationUseCase) Execute(ctx context.Context, senderID, receiverID int64, amount float64) error {
	args := m.Called(ctx, senderID, receiverID, amount)
	return args.Error(0)
}

func TestTransaction_Execute_Success(t *testing.T) {
	ctx := context.Background()
	senderID := int64(1)
	receiverID := int64(2)
	amount := 50.0

	userRepo := new(mockUserRepo)
	walletRepo := new(mockWalletRepo)
	transactionRepo := new(mockTransactionRepo)
	authService := new(mockAuthService)
	notificationUseCase := new(mockNotificationUseCase)

	userRepo.On("GetByID", ctx, senderID).Return(&entities.User{ID: senderID}, nil)
	userRepo.On("GetByID", ctx, receiverID).Return(&entities.User{ID: receiverID}, nil)

	senderWallet := &entities.Wallet{ID: 10, OwnerID: senderID, Type: entities.CommonWallet, Balance: 100}
	receiverWallet := &entities.Wallet{ID: 20, OwnerID: receiverID, Type: entities.MerchantWallet, Balance: 25}

	walletRepo.On("GetByOwnerID", ctx, senderID).Return(senderWallet, nil)
	walletRepo.On("GetByOwnerID", ctx, receiverID).Return(receiverWallet, nil)
	walletRepo.On("UpdateBalance", ctx, senderWallet.ID, senderWallet.Balance-amount).Return(nil)
	walletRepo.On("UpdateBalance", ctx, receiverWallet.ID, receiverWallet.Balance+amount).Return(nil)

	transactionRepo.On("Create", ctx, mock.Anything).Return(int64(99), nil)
	transactionRepo.On("UpdateStatus", ctx, int64(99), entities.TransactionStatusCompleted).Return(nil)

	authService.On("Authorize", ctx).Return(true, nil)
	notificationUseCase.On("Execute", ctx, senderID, receiverID, amount).Return(nil)

	tx := NewTransaction(userRepo, walletRepo, transactionRepo, &NotificationUseCase{notificationRepo: nil, notificationService: nil}, authService)
	txWithNotif := *tx
	txWithNotif.notificationUseCase = notificationUseCase

	err := txWithNotif.Execute(ctx, senderID, receiverID, amount)
	assert.NoError(t, err)

	userRepo.AssertExpectations(t)
	walletRepo.AssertExpectations(t)
	transactionRepo.AssertExpectations(t)
	authService.AssertExpectations(t)
	notificationUseCase.AssertExpectations(t)
}
