package service

import (
	"context"
	"database/sql"

	"github.com/google/uuid"
)

type WalletService struct {
	db *sql.DB
}

func NewWalletService(db *sql.DB) *WalletService {
	return &WalletService{db: db}
}

func (w *WalletService) Deposit(ctx context.Context, walletID uuid.UUID, amount int) {
	w.updateBalance(ctx, walletID, amount)
}

func (w *WalletService) WithDraw(ctx context.Context, walletID uuid.UUID, amount int) {
	w.updateBalance(ctx, walletID, -amount)
}

func (w *WalletService) GetBalance(ctx context.Context, walletID uuid.UUID) error {
	var balance int
	err := w.db.QueryRowContext(ctx, "SELECT balance FROM wallets WHERE id=$1", walletID).Scan(&balance)
	return err
}

func (w *WalletService) updateBalance(ctx context.Context, walletID uuid.UUID, amount int) {

}
