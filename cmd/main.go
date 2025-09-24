package main

import (
	"log"
	"net/http"
	cfg "wallet/internal/config"
	"wallet/internal/db"
	h "wallet/internal/handler"
	s "wallet/internal/service"

	"github.com/gorilla/mux"
)

func main() {
	cfg := cfg.LoadConfig()

	dbConn, err := db.Connect(cfg)
	if err != nil {
		panic(err)
	}
	defer dbConn.Close()

	svc := s.NewWalletService(dbConn)
	h := h.NewWalletHandler(svc)

	r := mux.NewRouter()
	r.HandleFunc("/api/v1/wallet", h.HandleWalletOperation).Methods("POST")
	r.HandleFunc("/api/v1/wallets/{wallet_uuid}", h.GetWalletBalance).Methods("GET")
	r.HandleFunc("/", index).Methods("GET")

	log.Printf("Server running on :%s", cfg.ServerPort)
	log.Fatal(http.ListenAndServe(":"+cfg.ServerPort, r))
	// add raceful shutdown
}

func index(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, "web/index.html")
}
