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

type NotificationRepositoryInMemory struct {
	notifications map[int64]*entities.Notification
	mu            sync.RWMutex
	nextID        int64
}

func NewNotificationRepositoryInMemory() port.NotificationRepository {
	return &NotificationRepositoryInMemory{
		notifications: make(map[int64]*entities.Notification),
		mu:            sync.RWMutex{},
		nextID:        1,
	}
}

func (r *NotificationRepositoryInMemory) Create(ctx context.Context, notification *entities.Notification) (int64, error) {
	r.mu.Lock()
	defer r.mu.Unlock()
	notification.ID = r.nextID
	r.notifications[notification.ID] = notification
	r.nextID++
	return notification.ID, nil
}

func (r *NotificationRepositoryInMemory) GetByID(ctx context.Context, id int64) (*entities.Notification, error) {
	r.mu.RLock()
	defer r.mu.RUnlock()
	notification, ok := r.notifications[id]
	if !ok {
		return nil, errors.New("notificação não encontrada")
	}
	return notification, nil
}

func (r *NotificationRepositoryInMemory) UpdateStatus(ctx context.Context, id int64, status entities.NotificationStatus) error {
	r.mu.Lock()
	defer r.mu.Unlock()
	notification, ok := r.notifications[id]
	if !ok {
		return errors.New("notificação não encontrada")
	}
	notification.Status = status
	notification.UpdatedAt = time.Now()
	return nil
}

func TestNotificationRepositoryInMemory_Create(t *testing.T) {
	repo := NewNotificationRepositoryInMemory()
	ctx := context.Background()

	notification := &entities.Notification{
		ReceiverID:    1,
		TransactionID: 100,
		Amount:        50.00,
		Status:        entities.NotificationStatusPending,
		CreatedAt:     time.Now(),
	}

	id, err := repo.Create(ctx, notification)
	assert.NoError(t, err)
	assert.NotZero(t, id)

	retrievedNotification, err := repo.GetByID(ctx, id)
	assert.NoError(t, err)
	assert.Equal(t, notification.ReceiverID, retrievedNotification.ReceiverID)
	assert.Equal(t, notification.TransactionID, retrievedNotification.TransactionID)
	assert.Equal(t, notification.Amount, retrievedNotification.Amount)
	assert.Equal(t, notification.Status, retrievedNotification.Status)
}

func TestNotificationRepositoryInMemory_GetByID_Found(t *testing.T) {
	repo := NewNotificationRepositoryInMemory()
	ctx := context.Background()

	expectedNotification := &entities.Notification{
		ReceiverID:    2,
		TransactionID: 200,
		Amount:        100.00,
		Status:        entities.NotificationStatusSent,
		CreatedAt:     time.Now(),
	}
	id, err := repo.Create(ctx, expectedNotification)
	assert.NoError(t, err)
	expectedNotification.ID = id

	retrievedNotification, err := repo.GetByID(ctx, id)
	assert.NoError(t, err)
	assert.Equal(t, expectedNotification, retrievedNotification)
}

func TestNotificationRepositoryInMemory_GetByID_NotFound(t *testing.T) {
	repo := NewNotificationRepositoryInMemory()
	ctx := context.Background()

	retrievedNotification, err := repo.GetByID(ctx, 999)
	assert.Error(t, err)
	assert.ErrorContains(t, err, "notificação não encontrada")
	assert.Nil(t, retrievedNotification)
}

func TestNotificationRepositoryInMemory_UpdateStatus_Success(t *testing.T) {
	repo := NewNotificationRepositoryInMemory()
	ctx := context.Background()

	initialNotification := &entities.Notification{
		ReceiverID:    3,
		TransactionID: 300,
		Amount:        25.50,
		Status:        entities.NotificationStatusPending,
		CreatedAt:     time.Now(),
	}
	id, err := repo.Create(ctx, initialNotification)
	assert.NoError(t, err)

	newStatus := entities.NotificationStatusFailed
	err = repo.UpdateStatus(ctx, id, newStatus)
	assert.NoError(t, err)

	retrievedNotification, err := repo.GetByID(ctx, id)
	assert.NoError(t, err)
	assert.Equal(t, newStatus, retrievedNotification.Status)
}

func TestNotificationRepositoryInMemory_UpdateStatus_NotFound(t *testing.T) {
	repo := NewNotificationRepositoryInMemory()
	ctx := context.Background()

	err := repo.UpdateStatus(ctx, 999, entities.NotificationStatusSent)
	assert.Error(t, err)
	assert.ErrorContains(t, err, "notificação não encontrada")

	retrievedNotification, err := repo.GetByID(ctx, 999)
	assert.Error(t, err)
	assert.ErrorContains(t, err, "notificação não encontrada")
	assert.Nil(t, retrievedNotification)
}
