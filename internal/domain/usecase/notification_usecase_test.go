package usecase

import (
	"context"
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"go-transfer/internal/domain/entities"
)

type MockNotificationRepository struct {
	mock.Mock
}

func (m *MockNotificationRepository) Create(ctx context.Context, notification *entities.Notification) (int64, error) {
	args := m.Called(ctx, notification)
	return args.Get(0).(int64), args.Error(1)
}

func (m *MockNotificationRepository) UpdateStatus(ctx context.Context, id int64, status entities.NotificationStatus) error {
	args := m.Called(ctx, id, status)
	return args.Error(0)
}

func (m *MockNotificationRepository) GetByID(ctx context.Context, id int64) (*entities.Notification, error) {
	args := m.Called(ctx, id)
	return args.Get(0).(*entities.Notification), args.Error(1)
}

type MockNotificationService struct {
	mock.Mock
}

func (m *MockNotificationService) Notify(ctx context.Context, receiverID int64, amount float64) error {
	args := m.Called(ctx, receiverID, amount)
	return args.Error(0)
}

func TestNotificationUseCase_Execute_Success(t *testing.T) {
	ctx := context.Background()
	mockRepo := new(MockNotificationRepository)
	mockService := new(MockNotificationService)

	uc := NewNotification(mockRepo, mockService)

	receiverID := int64(1)
	transferID := int64(101)
	amount := 250.0
	notificationID := int64(999)

	mockRepo.
		On("Create", mock.Anything, mock.AnythingOfType("*entities.Notification")).
		Return(notificationID, nil)

	mockService.
		On("Notify", mock.Anything, receiverID, amount).
		Return(nil)

	mockRepo.
		On("UpdateStatus", mock.Anything, notificationID, entities.NotificationStatusSent).
		Return(nil)

	err := uc.Execute(ctx, receiverID, transferID, amount)

	assert.NoError(t, err)
	mockRepo.AssertExpectations(t)
	mockService.AssertExpectations(t)
}

func TestNotificationUseCase_Execute_FailedNotify(t *testing.T) {
	ctx := context.Background()
	mockRepo := new(MockNotificationRepository)
	mockService := new(MockNotificationService)

	uc := NewNotification(mockRepo, mockService)

	receiverID := int64(1)
	transferID := int64(102)
	amount := 500.0
	notificationID := int64(1000)

	mockRepo.
		On("Create", mock.Anything, mock.AnythingOfType("*entities.Notification")).
		Return(notificationID, nil)

	mockService.
		On("Notify", mock.Anything, receiverID, amount).
		Return(errors.New("notify error"))

	mockRepo.
		On("UpdateStatus", mock.Anything, notificationID, entities.NotificationStatusFailed).
		Return(nil)

	mockRepo.
		On("UpdateStatus", mock.Anything, notificationID, entities.NotificationStatusSent).
		Return(nil)

	err := uc.Execute(ctx, receiverID, transferID, amount)

	assert.NoError(t, err)
	mockRepo.AssertExpectations(t)
	mockService.AssertExpectations(t)
}
