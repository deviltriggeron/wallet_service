package handler

import (
	"encoding/json"
	"net/http"
	e "wallet/internal/entity"
	s "wallet/internal/service"

	"github.com/google/uuid"
	"github.com/gorilla/mux"
)

type WalletHandler struct {
	svc *s.WalletService
}

func NewWalletHandler(svc *s.WalletService) *WalletHandler {
	return &WalletHandler{svc: svc}
}

func (h *WalletHandler) CreateWallet(w http.ResponseWriter, r *http.Request) {
	var wr e.CreateWalletRequest
	if err := json.NewDecoder(r.Body).Decode(&wr); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	id := uuid.New()
	if err := h.svc.CreateWallet(r.Context(), id, wr.Amount); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string]interface{}{
		"walletId": id,
		"balance":  wr.Amount,
	})
}

func (h *WalletHandler) HandleWalletOperation(w http.ResponseWriter, r *http.Request) {
	var wr e.WalletRequest
	if err := json.NewDecoder(r.Body).Decode(&wr); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	var err error
	var operationType string
	switch wr.OperationType {
	case "DEPOSIT":
		err = h.svc.Deposit(r.Context(), wr.WalletID, wr.Amount)
		operationType = "DEPOSIT"
	case "WITHDRAW":
		err = h.svc.WithDraw(r.Context(), wr.WalletID, wr.Amount)
		operationType = "WITHDRAW"
	default:
		http.Error(w, "invalid operation type", http.StatusBadRequest)
		return
	}

	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	json.NewEncoder(w).Encode(map[string]interface{}{
		"walletId":    wr.WalletID,
		operationType: wr.Amount,
	})
}

func (h *WalletHandler) GetWalletBalance(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	walletID, err := uuid.Parse(vars["wallet_uuid"])
	if err != nil {
		http.Error(w, "invalid wallet id", http.StatusBadRequest)
		return
	}

	balance, err := h.svc.GetBalance(r.Context(), walletID)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(map[string]int{"balance": balance})
}
