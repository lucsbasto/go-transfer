package setup_routes

import (
	"fmt"
	"go-transfer/internal/api"
	"net/http"
)

func SetupTransferRoutes(transactionHandler *api.TransactionHandler) {
	fmt.Println("Configuring routes...")
	http.HandleFunc("/transfers", transactionHandler.Transaction)
}
