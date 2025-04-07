package setup_usecases

import (
	"fmt"
	"go-transfer/internal/domain/usecase"
	"go-transfer/internal/infra/repositories"
)

func SetupTransactionUseCase(
	userRepo *repositories.UserRepository,
	walletRepo *repositories.WalletRepository,
	transactionRepo *repositories.TransactionRepository,
	notificationUseCase *usecase.NotificationUseCase,
) *usecase.Transaction {
	fmt.Println("Configuring Transaction usecases...")
	return usecase.NewTransaction(userRepo, walletRepo, transactionRepo, notificationUseCase)
}
