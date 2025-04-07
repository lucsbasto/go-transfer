package setup_usecases

import (
	"fmt"
	"go-transfer/internal/domain/usecase"
	"go-transfer/internal/env"
	"go-transfer/internal/infra/externals"
	"go-transfer/internal/infra/repositories"
)

func SetupNotificationUseCase(
	notificationRepo *repositories.NotificationRepository,
) *usecase.NotificationUseCase {
	fmt.Println("Configuring Notification usecases...")
	AppConfig := env.LoadEnv()

	notificationService := externals.NewNotificationService(AppConfig.NotificationURL)
	notificationUseCase := usecase.NewNotification(notificationRepo, notificationService)

	return notificationUseCase
}
