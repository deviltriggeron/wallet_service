package handler

import (
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

func (h *WalletHandler) HandleWalletOperation(w http.ResponseWriter, r *http.Request) {
	var wr e.WalletRequest

	switch wr.OperationType {
	case "deposit":
		h.svc.Deposit(r.Context(), wr.WalletID, wr.Amount)
	case "withdraw":
		h.svc.WithDraw(r.Context(), wr.WalletID, wr.Amount)
	default:
		http.Error(w, "invalid operation type", http.StatusBadRequest)
		return
	}
}

func (h *WalletHandler) GetWalletBalance(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	walletID, err := uuid.Parse(vars["wallet_uuid"])
	if err != nil {
		http.Error(w, "invalid wallet id", http.StatusBadRequest)
		return
	}
	h.svc.GetBalance(r.Context(), walletID)
}
