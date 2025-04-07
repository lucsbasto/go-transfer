package setup_routes

import (
	"fmt"
	"go-transfer/internal/api"
)

func SetupRoutes(userHandler *api.UserHandler, transactionHandler *api.TransactionHandler) {
	fmt.Println("Configuring routes...")
	SetupUserRoutes(userHandler)
	SetupTransferRoutes(transactionHandler)
}
