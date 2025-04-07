package usecase

import (
	"context"
	"errors"
	"go-transfer/internal/domain/entities"
	"go-transfer/internal/domain/port"
)

type Transaction struct {
	userRepo        port.UserRepository
	walletRepo      port.WalletRepository
	TransactionRepo port.TransactionRepository
}

func NewTransaction(
	userRepo port.UserRepository,
	walletRepo port.WalletRepository,
	TransactionRepo port.TransactionRepository,
) *Transaction {
	return &Transaction{
		userRepo:        userRepo,
		walletRepo:      walletRepo,
		TransactionRepo: TransactionRepo,
	}
}

func (t *Transaction) Execute(ctx context.Context, senderID, receiverID int64, amount float64) error {
	if err := t.validateTransaction(ctx, senderID, receiverID, amount); err != nil {
		return err
	}

	Transaction := &entities.Transaction{
		SenderID:   senderID,
		ReceiverID: receiverID,
		Amount:     amount,
		Status:     entities.TransactionStatusPending,
	}
	TransactionId, err := t.TransactionRepo.Create(ctx, Transaction)
	if err != nil {
		return errors.New("failed to create Transaction record: " + err.Error())
	}

	if err := t.updateWallets(ctx, senderID, receiverID, amount); err != nil {
		if err := t.TransactionRepo.UpdateStatus(ctx, TransactionId, entities.TransactionStatusFailed); err != nil {
			return errors.New("failed to create Transaction record: " + err.Error())
		}
		return err
	}
	if err := t.TransactionRepo.UpdateStatus(ctx, TransactionId, entities.TransactionStatusCompleted); err != nil {
		return errors.New("failed to create Transaction record: " + err.Error())
	}

	return nil
}

func (t *Transaction) validateTransaction(ctx context.Context, senderID, receiverID int64, amount float64) error {
	if senderID == receiverID {
		return errors.New("sender and receiver must be different")
	}

	user, err := t.userRepo.GetByID(ctx, senderID)
	if err != nil {
		return errors.New(err.Error())
	}
	if user == nil {
		return errors.New("sender user not found")
	}

	user, err = t.userRepo.GetByID(ctx, receiverID)
	if err != nil {
		return errors.New(err.Error())
	}
	if user == nil {
		return errors.New("receiver user not found")
	}

	senderWallet, err := t.walletRepo.GetByOwnerID(ctx, senderID)
	if err != nil {
		return errors.New(err.Error())
	}

	if senderWallet.Type == entities.MerchantWallet {
		return errors.New("merchant cannot Transaction")
	}

	if senderWallet.Balance < amount {
		return errors.New("insufficient balance")
	}

	_, err = t.walletRepo.GetByOwnerID(ctx, senderID)
	if err != nil {
		return errors.New(err.Error())
	}

	return nil
}

func (t *Transaction) updateWallets(ctx context.Context, senderID, receiverID int64, amount float64) error {
	senderWallet, err := t.walletRepo.GetByOwnerID(ctx, senderID)
	if err != nil {
		return errors.New("sender wallet not found" + err.Error())
	}

	receiverWallet, err := t.walletRepo.GetByOwnerID(ctx, receiverID)
	if err != nil {
		return errors.New("receiver wallet not found" + err.Error())
	}

	if err := t.walletRepo.UpdateBalance(ctx, senderWallet.ID, senderWallet.Balance-amount); err != nil {
		return errors.New("failed to update sender balance" + err.Error())
	}

	if err := t.walletRepo.UpdateBalance(ctx, receiverWallet.ID, receiverWallet.Balance+amount); err != nil {
		return errors.New("failed to update receiver balance" + err.Error())
	}

	return nil
}
