package repositories

import (
	"context"

	"go-transfer/internal/domain/entities"

	"gorm.io/gorm"
)

type NotificationRepository struct {
	db *gorm.DB
}

func NewNotificationRepository(db *gorm.DB) *NotificationRepository {
	return &NotificationRepository{
		db: db,
	}
}
func (r *NotificationRepository) Create(ctx context.Context, notification *entities.Notification) (int64, error) {
	result := r.db.WithContext(ctx).Create(notification)
	if result.Error != nil {
		return 0, result.Error
	}
	return notification.ID, nil
}

func (r *NotificationRepository) GetByID(ctx context.Context, id int64) (*entities.Notification, error) {
	notification := &entities.Notification{}
	err := r.db.WithContext(ctx).First(notification, id).Error
	if err != nil {
		return nil, err
	}
	return notification, nil
}

func (r *NotificationRepository) UpdateStatus(ctx context.Context, id int64, status entities.NotificationStatus) error {
	return r.db.WithContext(ctx).Model(&entities.Notification{}).Where("id = ?", id).Update("status", status).Error
}
