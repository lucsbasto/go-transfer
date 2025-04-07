package api

import (
	"encoding/json"
	"net/http"

	"go-transfer/internal/domain/usecase"
)

type UserHandler struct {
	userUseCase   *usecase.User
	walletUseCase *usecase.Wallet
}

func NewUserHandler(userUseCase *usecase.User, walletUseCase *usecase.Wallet) *UserHandler {
	return &UserHandler{
		userUseCase:   userUseCase,
		walletUseCase: walletUseCase,
	}
}

func (h *UserHandler) CreateUser(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	var input usecase.UserInput
	err := json.NewDecoder(r.Body).Decode(&input)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	user, err := h.userUseCase.CreateUser(r.Context(), input)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	walletInput := usecase.WalletInput{
		OwnerID: user.ID,
		Type:    input.Type,
		Balance: input.Balance,
	}
	err = h.walletUseCase.CreateWallet(r.Context(), walletInput)
	w.Header().Set("Content-Type", "application/json")
	err = json.NewEncoder(w).Encode(user)
	if err != nil {
		return
	}
}
