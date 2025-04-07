package port

import (
	"context"

	"go-transfer/internal/domain/entities"
)

type NotificationRepository interface {
	Create(ctx context.Context, notification *entities.Notification) (int64, error)
	UpdateStatus(ctx context.Context, id int64, status entities.NotificationStatus) error
	GetByID(ctx context.Context, id int64) (*entities.Notification, error)
}
