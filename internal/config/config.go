package config

import (
	"fmt"
	"go-transfer/internal/config/handlers"
	"go-transfer/internal/config/setup_repositories"
	"go-transfer/internal/config/setup_routes"
	"go-transfer/internal/config/setup_usecases"
	"go-transfer/internal/env"
	"go-transfer/internal/infra/database"
	"log"
)

func Setup() {
	fmt.Println("Init Setup ...")

	AppConfig := env.LoadEnv()

	db, err := database.SetupDB(AppConfig)
	if err != nil {
		log.Fatalf("Erro ao conectar no banco de dados: %v", err)
	}

	userRepository, walletRepository, transactionRepository, notificationRepository := setup_repositories.SetupRepositories(db)

	userUseCase, walletUseCase, transactionUseCase := setup_usecases.SetupUseCases(
		userRepository,
		walletRepository,
		transactionRepository,
		notificationRepository,
	)

	userHandler, transactionHandler := handlers.SetupHandlers(userUseCase, walletUseCase, transactionUseCase)

	setup_routes.SetupRoutes(userHandler, transactionHandler)
}
