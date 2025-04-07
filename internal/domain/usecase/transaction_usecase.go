package usecase

import (
	"context"
	"errors"
	"fmt"
	"go-transfer/internal/domain/entities"
	"go-transfer/internal/domain/port"
	"sync"
)

type Transaction struct {
	userRepo             port.UserRepository
	walletRepo           port.WalletRepository
	transactionRepo      port.TransactionRepository
	notificationUseCase  NotificationUseCase
	authorizationService port.AuthorizationService
	walletLocker         *sync.Map // sync.Map[int64]*sync.Mutex
}

func NewTransaction(
	userRepo port.UserRepository,
	walletRepo port.WalletRepository,
	transactionRepo port.TransactionRepository,
	notificationUseCase *NotificationUseCase,
	authorizationService port.AuthorizationService,
) *Transaction {
	return &Transaction{
		userRepo:             userRepo,
		walletRepo:           walletRepo,
		transactionRepo:      transactionRepo,
		notificationUseCase:  *notificationUseCase,
		authorizationService: authorizationService,
		walletLocker:         &sync.Map{},
	}
}

func (t *Transaction) Execute(ctx context.Context, senderID, receiverID int64, amount float64) error {
	if err := t.checkAuthorization(ctx); err != nil {
		return err
	}

	if err := t.validateTransaction(ctx, senderID, receiverID, amount); err != nil {
		return err
	}

	unlock := t.lockWallets(senderID, receiverID)
	defer unlock()

	transactionID, err := t.createTransaction(ctx, senderID, receiverID, amount)
	if err != nil {
		return err
	}

	if err := t.updateWallets(ctx, senderID, receiverID, amount); err != nil {
		_ = t.transactionRepo.UpdateStatus(ctx, transactionID, entities.TransactionStatusFailed)
		return err
	}

	if err := t.transactionRepo.UpdateStatus(ctx, transactionID, entities.TransactionStatusCompleted); err != nil {
		return err
	}

	if err := t.sendNotification(ctx, senderID, receiverID, amount); err != nil {
		fmt.Print("failed to send notification: " + err.Error())
	}

	return nil
}

func (t *Transaction) checkAuthorization(ctx context.Context) error {
	isAuthorized, err := t.authorizationService.Authorize(ctx)
	if err != nil {
		return err
	}
	if !isAuthorized {
		return errors.New("unauthorized")
	}
	return nil
}

func (t *Transaction) validateTransaction(ctx context.Context, senderID, receiverID int64, amount float64) error {
	if senderID == receiverID {
		return errors.New("sender and receiver must be different")
	}

	if err := t.checkUserExists(ctx, senderID, receiverID); err != nil {
		return err
	}

	senderWallet, err := t.walletRepo.GetByOwnerID(ctx, senderID)
	if err != nil {
		return err
	}
	if senderWallet.Type == entities.MerchantWallet {
		return errors.New("merchant cannot transfer")
	}
	if senderWallet.Balance < amount {
		return errors.New("insufficient balance")
	}

	return nil
}

func (t *Transaction) checkUserExists(ctx context.Context, senderID, receiverID int64) error {
	user, err := t.userRepo.GetByID(ctx, senderID)
	if err != nil || user == nil {
		return errors.New("sender not found")
	}

	user, err = t.userRepo.GetByID(ctx, receiverID)
	if err != nil || user == nil {
		return errors.New("receiver not found")
	}

	return nil
}

func (t *Transaction) lockWallets(senderID, receiverID int64) func() {
	if senderID > receiverID {
		senderID, receiverID = receiverID, senderID
	}
	senderLock := t.getWalletLock(senderID)
	receiverLock := t.getWalletLock(receiverID)

	senderLock.Lock()
	receiverLock.Lock()

	return func() {
		receiverLock.Unlock()
		senderLock.Unlock()
	}
}

func (t *Transaction) getWalletLock(walletID int64) *sync.Mutex {
	lock, ok := t.walletLocker.Load(walletID)
	if !ok {
		lock = &sync.Mutex{}
		t.walletLocker.Store(walletID, lock)
	}
	return lock.(*sync.Mutex)
}

func (t *Transaction) createTransaction(ctx context.Context, senderID, receiverID int64, amount float64) (int64, error) {
	transaction := &entities.Transaction{
		SenderID:   senderID,
		ReceiverID: receiverID,
		Amount:     amount,
		Status:     entities.TransactionStatusPending,
	}
	transactionID, err := t.transactionRepo.Create(ctx, transaction)
	if err != nil {
		return 0, errors.New("failed to create transaction record: " + err.Error())
	}
	return transactionID, nil
}

func (t *Transaction) updateWallets(ctx context.Context, senderID, receiverID int64, amount float64) error {
	senderWallet, err := t.walletRepo.GetByOwnerID(ctx, senderID)
	if err != nil {
		return err
	}

	receiverWallet, err := t.walletRepo.GetByOwnerID(ctx, receiverID)
	if err != nil {
		return err
	}

	if err := t.walletRepo.UpdateBalance(ctx, senderWallet.ID, senderWallet.Balance-amount); err != nil {
		return err
	}

	if err := t.walletRepo.UpdateBalance(ctx, receiverWallet.ID, receiverWallet.Balance+amount); err != nil {
		return err
	}

	return nil
}

func (t *Transaction) sendNotification(ctx context.Context, senderID, receiverID int64, amount float64) error {
	return t.notificationUseCase.Execute(ctx, senderID, receiverID, amount)
}
