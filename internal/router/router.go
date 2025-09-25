package router

import (
	"net/http"
	h "wallet/internal/handler"

	"github.com/gorilla/mux"
)

func NewRouter(h *h.WalletHandler) *mux.Router {
	r := mux.NewRouter()
	r.HandleFunc("/api/v1/wallet", h.HandleWalletOperation).Methods("POST")
	r.HandleFunc("/api/v1/wallets/{wallet_uuid}", h.GetWalletBalance).Methods("GET")
	r.HandleFunc("/api/v1/createWallets", h.CreateWallet).Methods("POST")
	r.HandleFunc("/", index).Methods("GET")

	return r
}

func index(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, "web/index.html")
}
