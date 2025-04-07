package handlers

import (
	"fmt"
	"go-transfer/internal/api"
	"go-transfer/internal/domain/usecase"
)

func SetupUserHandlers(
	userUseCase *usecase.User,
	walletUseCase *usecase.Wallet,
) *api.UserHandler {
	fmt.Println("Configuring User handler...")
	return api.NewUserHandler(userUseCase, walletUseCase)
}
