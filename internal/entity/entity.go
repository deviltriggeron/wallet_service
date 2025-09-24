package entity

import "github.com/google/uuid"

type WalletRequest struct {
	WalletID      uuid.UUID `json:"walletId"`
	OperationType string    `json:"operationType"`
	Amount        int       `json:"amount"`
}

type Config struct {
	PostgresUser     string
	PostgresPassword string
	PostgresDB       string
	PostgresHost     string
	PostgresPort     string
	ServerPort       string
}
