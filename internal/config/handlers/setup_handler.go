package handlers

import (
	"fmt"
	"go-transfer/internal/api"
	"go-transfer/internal/domain/usecase"
)

func SetupHandlers(
	userUseCase *usecase.User,
	walletUseCase *usecase.Wallet,
) *api.UserHandler {
	fmt.Println("Configuring handlers...")
	userHandler := SetupUserHandlers(userUseCase, walletUseCase)
	return userHandler
}
