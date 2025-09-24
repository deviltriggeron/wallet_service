package main

import (
	"log"
	"net/http"
	cfg "wallet/internal/config"
	"wallet/internal/db"
	h "wallet/internal/handler"
	r "wallet/internal/router"
	s "wallet/internal/service"
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
	r := r.NewRouter(h)

	log.Printf("Server running on :%s", cfg.ServerPort)
	log.Fatal(http.ListenAndServe(":"+cfg.ServerPort, r))
	// add raceful shutdown
}
