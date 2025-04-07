package api

import (
	"encoding/json"
	"go-transfer/internal/domain/usecase"
	"net/http"
)

type TransactionRequest struct {
	Value float64 `json:"value"`
	Payer int64   `json:"payer"`
	Payee int64   `json:"payee"`
}

type TransactionHandler struct {
	TransactionUseCase *usecase.Transaction
}

func NewTransactionHandler(TransactionUseCase *usecase.Transaction) *TransactionHandler {
	return &TransactionHandler{
		TransactionUseCase: TransactionUseCase,
	}
}

func (h *TransactionHandler) Transaction(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Method Not Allowed", http.StatusMethodNotAllowed)
		return
	}

	var req TransactionRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	if err := h.validateTransactionRequest(req); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	if err := h.TransactionUseCase.Execute(r.Context(), req.Payer, req.Payee, req.Value); err != nil {
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}

	h.writeJSON(w, http.StatusOK, map[string]string{"message": "Transaction successful"})
}

func (h *TransactionHandler) validateTransactionRequest(req TransactionRequest) error {
	if req.Value <= 0 {
		return ErrInvalidTransactionValue
	}
	if req.Payer == req.Payee {
		return ErrSamePayerPayee
	}
	return nil
}

func (h *TransactionHandler) writeJSON(w http.ResponseWriter, status int, data interface{}) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	if err := json.NewEncoder(w).Encode(data); err != nil {
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
	}
}

var (
	ErrInvalidTransactionValue = NewError("Transaction value must be greater than zero")
	ErrSamePayerPayee          = NewError("Payer and payee cannot be the same")
)

type Error struct {
	Message string
}

func (e Error) Error() string {
	return e.Message
}

func NewError(msg string) error {
	return Error{Message: msg}
}
