package repositories

import (
	"context"
	"errors"
	"go-transfer/internal/domain/entities"
	"go-transfer/internal/domain/port"
	"sync"
	"time"
)

type WalletRepositoryInMemory struct {
	wallets map[int64]*entities.Wallet
	mu      sync.RWMutex
	nextID  int64
}

func NewWalletRepositoryInMemory() port.WalletRepository {
	return &WalletRepositoryInMemory{
		wallets: make(map[int64]*entities.Wallet),
		mu:      sync.RWMutex{},
		nextID:  1,
	}
}

func (r *WalletRepositoryInMemory) Create(ctx context.Context, wallet *entities.Wallet) error {
	r.mu.Lock()
	defer r.mu.Unlock()
	wallet.ID = r.nextID
	r.wallets[wallet.ID] = wallet
	r.nextID++
	return nil
}

func (r *WalletRepositoryInMemory) GetByID(ctx context.Context, id int64) (*entities.Wallet, error) {
	r.mu.RLock()
	defer r.mu.RUnlock()
	wallet, ok := r.wallets[id]
	if !ok {
		return nil, errors.New("carteira não encontrada")
	}
	return wallet, nil
}

func (r *WalletRepositoryInMemory) GetByOwnerID(ctx context.Context, ownerID int64) (*entities.Wallet, error) {
	r.mu.RLock()
	defer r.mu.RUnlock()
	for _, wallet := range r.wallets {
		if wallet.OwnerID == ownerID {
			return wallet, nil
		}
	}
	return nil, errors.New("carteira não encontrada para o OwnerID")
}

func (r *WalletRepositoryInMemory) UpdateBalance(ctx context.Context, id int64, balance float64) error {
	r.mu.Lock()
	defer r.mu.Unlock()
	wallet, ok := r.wallets[id]
	if !ok {
		return errors.New("carteira não encontrada")
	}
	wallet.Balance = balance
	wallet.UpdatedAt = time.Now()
	return nil
}
