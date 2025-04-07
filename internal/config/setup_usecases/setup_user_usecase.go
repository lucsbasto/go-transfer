package setup_usecases

import (
	"fmt"
	"go-transfer/internal/domain/usecase"
	"go-transfer/internal/infra/repositories"
)

func SetupUserUseCase(
	userRepo *repositories.UserRepository,
) *usecase.User {
	fmt.Println("Configuring User usecases...")
	userUseCase := usecase.NewUser(userRepo)

	return userUseCase
}
