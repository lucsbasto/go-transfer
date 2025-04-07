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
) (*usecase.User, *usecase.Wallet, *usecase.Transaction) {
	fmt.Println("Configuring usecases...")
	userUseCase := SetupUserUseCases(userRepo)
	walletUseCase := SetupWalletUseCases(walletRepo)
	transactionRepository := SetupTransactionUseCases(userRepo, walletRepo, transactionRepo)
	return userUseCase, walletUseCase, transactionRepository
}
