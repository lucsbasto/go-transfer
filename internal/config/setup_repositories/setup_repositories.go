package setup_repositories

import (
	"fmt"
	"go-transfer/internal/infra/repositories"
	"gorm.io/gorm"
)

func SetupRepositories(db *gorm.DB) (
	*repositories.UserRepository,
	*repositories.WalletRepository,
	*repositories.TransactionRepository,
	*repositories.NotificationRepository,
) {
	fmt.Println("Configuring repositories...")
	userRepository := NewUserRepository(db)
	walletRepository := NewWalletRepository(db)
	transactionRepository := NewTransactionRepository(db)
	notificationRepository := NewNotificationRepository(db)
	return userRepository, walletRepository, transactionRepository, notificationRepository
}
