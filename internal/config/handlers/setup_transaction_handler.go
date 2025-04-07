package handlers

import (
	"fmt"
	"go-transfer/internal/api"
	"go-transfer/internal/domain/usecase"
)

func SetupTransactionHandlers(
	transactionUseCase *usecase.Transaction,
) *api.TransactionHandler {
	fmt.Println("Configuring Transaction handler...")
	return api.NewTransactionHandler(transactionUseCase)
}
