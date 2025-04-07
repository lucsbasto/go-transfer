package env

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

type Config struct {
	Port             string
	NotificationURL  string
	AuthorizationURL string
	DatabaseHost     string
	DatabasePort     string
	DatabaseUser     string
	DatabasePassword string
	DatabaseName     string
}

func LoadEnv() *Config {
	err := godotenv.Load(".env")
	if err != nil {
		log.Println("Arquivo .env não encontrado. Usando variáveis de ambiente do sistema.")
	}

	cfg := &Config{
		Port:             os.Getenv("PORT"),
		NotificationURL:  os.Getenv("NOTIFICATION_BASE_URL"),
		AuthorizationURL: os.Getenv("AUTHORIZATION_BASE_URL"),
		DatabaseHost:     os.Getenv("DATABASE_URL"),
		DatabasePort:     os.Getenv("DATABASE_PORT"),
		DatabaseUser:     os.Getenv("DATABASE_USERNAME"),
		DatabasePassword: os.Getenv("DATABASE_PASSWORD"),
		DatabaseName:     os.Getenv("DATABASE_NAME"),
	}

	if cfg.DatabaseHost == "" || cfg.DatabaseUser == "" || cfg.DatabaseName == "" {
		log.Fatal("Variáveis de ambiente essenciais para o banco de dados não foram definidas.")
	}

	return cfg
}
