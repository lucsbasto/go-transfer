package setup_usecases

import (
	"fmt"
	"go-transfer/internal/domain/usecase"
	"go-transfer/internal/infra/repositories"
)

func SetupUseCases(
	userRepo *repositories.UserRepository,
	walletRepo *repositories.WalletRepository,
) (*usecase.User, *usecase.Wallet) {
	fmt.Println("Configuring usecases...")
	userUseCase := SetupUserUseCases(userRepo)
	walletUseCase := SetupWalletUseCases(walletRepo)

	return userUseCase, walletUseCase
}
