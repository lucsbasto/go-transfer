package setup_usecases

import (
	"fmt"
	"go-transfer/internal/domain/usecase"
	"go-transfer/internal/infra/repositories"
)

func SetupUseCases(
	userRepo *repositories.UserRepository,
	walletRepo *repositories.WalletRepository,
	transactionRepo *repositories.TransactionRepository,
	notificationRepo *repositories.NotificationRepository,
) (*usecase.User, *usecase.Wallet, *usecase.Transaction) {
	fmt.Println("Configuring usecases...")
	userUseCase := SetupUserUseCase(userRepo)
	walletUseCase := SetupWalletUseCase(walletRepo)
	notificationUseCase := SetupNotificationUseCase(notificationRepo)
	transactionRepository := SetupTransactionUseCase(userRepo, walletRepo, transactionRepo, notificationUseCase)
	return userUseCase, walletUseCase, transactionRepository
}
