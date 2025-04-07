package setup_routes

import (
	"fmt"
	"go-transfer/internal/api"
	"net/http"
)

func SetupUserRoutes(userHandler *api.UserHandler) {
	fmt.Println("Configuring user routes...")
	http.HandleFunc("/users", userHandler.CreateUser)
}
