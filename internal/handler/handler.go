package handler

import (
	"encoding/json"
	"net/http"
	"wallet_api_service/internal/model"
	"wallet_api_service/internal/service"

	"github.com/google/uuid"
)

type Handler struct {
	svc *service.WalletService
}

func New(svc *service.WalletService) *Handler {
	return &Handler{svc: svc}
}

type request struct {
	WalletID  uuid.UUID           `json:"walletId"`
	Operation model.OperationType `json:"operationType"`
	Amount    int64               `json:"amount"`
}

func (h *Handler) Process(w http.ResponseWriter, r *http.Request) {
	var req request

	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, err.Error(), 400)
		return
	}

	balance, err := h.svc.Process(r.Context(), req.WalletID, req.Operation, req.Amount)
	if err != nil {
		http.Error(w, err.Error(), 400)
		return
	}

	json.NewEncoder(w).Encode(map[string]int64{"balance": balance})
}

func (h *Handler) GetBalance(w http.ResponseWriter, r *http.Request) {
	idStr := r.URL.Path[len("/api/v1/wallets/"):]
	id, _ := uuid.Parse(idStr)

	balance, err := h.svc.GetBalance(r.Context(), id)
	if err != nil {
		http.Error(w, err.Error(), 404)
		return
	}

	json.NewEncoder(w).Encode(map[string]int64{"balance": balance})
}
