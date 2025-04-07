package setup_repositories

import (
	"fmt"
	"go-transfer/internal/infra/repositories"
	"gorm.io/gorm"
)

func NewWalletRepository(db *gorm.DB) *repositories.WalletRepository {
	fmt.Println("Configuring wallet repository...")
	return repositories.NewWalletRepository(db)
}
