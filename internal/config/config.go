package config

import (
	"fmt"
	"go-transfer/internal/config/handlers"
	"go-transfer/internal/config/setup_repositories"
	"go-transfer/internal/config/setup_routes"
	"go-transfer/internal/config/setup_usecases"
	"go-transfer/internal/infra/database"
)

func Setup() {
	fmt.Println("Init Setup ...")

	db, err := database.SetupDB()
	if err != nil {
		fmt.Println("Error connecting database:", err)
		return
	}

	userRepository, walletRepository := setup_repositories.SetupRepositories(db)

	userUseCase, walletUseCase := setup_usecases.SetupUseCases(
		userRepository,
		walletRepository,
	)

	userHandler := handlers.SetupHandlers(userUseCase, walletUseCase)
	setup_routes.SetupUserRoutes(userHandler)
}
