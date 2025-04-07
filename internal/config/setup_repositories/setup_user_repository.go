package setup_repositories

import (
	"fmt"
	"go-transfer/internal/infra/repositories"
	"gorm.io/gorm"
)

func NewUserRepository(db *gorm.DB) *repositories.UserRepository {
	fmt.Println("Configuring user repository...")
	return repositories.NewUserRepository(db)
}
