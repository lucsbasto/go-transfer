package usecase

import (
	"context"
	"fmt"
	"time"

	"go-transfer/internal/domain/entities"
	"go-transfer/internal/domain/port"
)

type NotificationUseCase struct {
	notificationRepo    port.NotificationRepository
	notificationService port.NotificationService
}

func NewNotification(notificationRepo port.NotificationRepository, notificationService port.NotificationService) *NotificationUseCase {
	return &NotificationUseCase{
		notificationRepo:    notificationRepo,
		notificationService: notificationService,
	}
}

func (n *NotificationUseCase) Execute(ctx context.Context, receiverID int64, transferID int64, amount float64) error {
	notification := &entities.Notification{
		ReceiverID:    receiverID,
		TransactionID: transferID,
		Amount:        amount,
		Status:        entities.NotificationStatusPending,
		CreatedAt:     time.Now(),
	}

	notificationID, err := n.notificationRepo.Create(ctx, notification)
	if err != nil {
		fmt.Printf("failed to create notification record: %v\n", err)
	}

	err = n.notificationService.Notify(ctx, receiverID, amount)
	if err != nil {
		updateErr := n.notificationRepo.UpdateStatus(ctx, notificationID, entities.NotificationStatusFailed)
		if updateErr != nil {
			fmt.Printf("failed to update notification status to 'failed': %v\n", updateErr)
		}
		fmt.Printf("failed to send notification: %v\n", err)
	}

	if err := n.notificationRepo.UpdateStatus(ctx, notificationID, entities.NotificationStatusSent); err != nil {
		fmt.Printf("failed to update notification status to 'sent': %v\n", err)
	}

	return nil
}

func (n *NotificationUseCase) GetNotificationByID(ctx context.Context, id int64) (*entities.Notification, error) {
	return n.notificationRepo.GetByID(ctx, id)
}

func (n *NotificationUseCase) UpdateNotificationStatus(ctx context.Context, id int64, status entities.NotificationStatus) error {
	return n.notificationRepo.UpdateStatus(ctx, id, status)
}
