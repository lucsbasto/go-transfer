package setup_usecases

import (
	"fmt"
	"go-transfer/internal/domain/usecase"
	"go-transfer/internal/env"
	"go-transfer/internal/infra/externals"
	"go-transfer/internal/infra/repositories"
)

func SetupTransactionUseCase(
	userRepo *repositories.UserRepository,
	walletRepo *repositories.WalletRepository,
	transactionRepo *repositories.TransactionRepository,
	notificationUseCase *usecase.NotificationUseCase,
) *usecase.Transaction {
	fmt.Println("Configuring Transaction usecases...")
	AppConfig := env.LoadEnv()

	authorizationService := externals.NewAuthorizationService(AppConfig.AuthorizationURL)
	return usecase.NewTransaction(userRepo, walletRepo, transactionRepo, notificationUseCase, authorizationService)
}
