package setup_usecases

import (
	"fmt"
	"go-transfer/internal/domain/usecase"
	"go-transfer/internal/infra/repositories"
)

func SetupWalletUseCase(
	walletRepo *repositories.WalletRepository,
) *usecase.Wallet {
	fmt.Println("Configuring Wallet usecases...")
	walletUseCase := usecase.NewWallet(walletRepo)

	return walletUseCase
}
