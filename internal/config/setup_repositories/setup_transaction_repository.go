package setup_repositories

import (
	"fmt"
	"go-transfer/internal/infra/repositories"
	"gorm.io/gorm"
)

func NewTransactionRepository(db *gorm.DB) *repositories.TransactionRepository {
	fmt.Println("Configuring transaction repository...")
	return repositories.NewTransactionRepository(db)
}
