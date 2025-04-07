package setup_usecases

import (
	"fmt"
	"go-transfer/internal/domain/usecase"
	"go-transfer/internal/infra/repositories"
)

func SetupTransactionUseCases(
	userRepo *repositories.UserRepository,
	walletRepo *repositories.WalletRepository,
	transactionRepo *repositories.TransactionRepository,
) *usecase.Transaction {
	fmt.Println("Configuring Transaction usecases...")
	return usecase.NewTransaction(userRepo, walletRepo, transactionRepo)
}
