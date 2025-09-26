package main

import (
	"context"
	"log"
	"net/http"
	"os/signal"
	"sync"
	"syscall"
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

	ctx, stop := signal.NotifyContext(context.Background(), syscall.SIGTERM, syscall.SIGINT)
	defer stop()

	var wg sync.WaitGroup
	svc := s.NewWalletService(dbConn)
	h := h.NewWalletHandler(svc)
	r := r.NewRouter(h)

	wg.Add(1)
	go func() {
		defer wg.Done()

		log.Printf("Server running on :%s", cfg.ServerPort)
		log.Fatal(http.ListenAndServe(":"+cfg.ServerPort, r))
	}()

	<-ctx.Done()
	log.Println("Server will shutdown gracefully...")
	err = svc.Shutdown(context.Background())

	wg.Wait()
}
