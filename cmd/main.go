package main

import (
	"go-transfer/internal/config"
	"net/http"
)

func main() {
	config.Setup()
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		return
	}
}
