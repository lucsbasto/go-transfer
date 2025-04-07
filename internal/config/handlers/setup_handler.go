package handlers

import (
	"fmt"
	"go-transfer/internal/api"
	"go-transfer/internal/domain/usecase"
)

func SetupHandlers(
	userUseCase *usecase.User,
	walletUseCase *usecase.Wallet,
	transactionUseCase *usecase.Transaction,
) (*api.UserHandler, *api.TransactionHandler) {
	fmt.Println("Configuring handlers...")
	userHandler := SetupUserHandlers(userUseCase, walletUseCase)
	transactionHandler := SetupTransactionHandlers(transactionUseCase)
	return userHandler, transactionHandler
}
