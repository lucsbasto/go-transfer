package setup_repositories

import (
	"fmt"
	"go-transfer/internal/infra/repositories"
	"gorm.io/gorm"
)

func NewNotificationRepository(db *gorm.DB) *repositories.NotificationRepository {
	fmt.Println("Configuring notification repository...")
	return repositories.NewNotificationRepository(db)
}
