package service

import (
	"context"
	"database/sql"
	"errors"

	"github.com/google/uuid"
)

type WalletService struct {
	db *sql.DB
}

func NewWalletService(db *sql.DB) *WalletService {
	return &WalletService{db: db}
}

func (s *WalletService) CreateWallet(ctx context.Context, walletID uuid.UUID, amount int) error {
	_, err := s.db.ExecContext(ctx,
		"INSERT INTO wallets (id, balance) VALUES ($1, $2)",
		walletID, amount,
	)
	return err
}

func (s *WalletService) Deposit(ctx context.Context, walletID uuid.UUID, amount int) error {
	return s.updateBalance(ctx, walletID, amount)
}

func (s *WalletService) WithDraw(ctx context.Context, walletID uuid.UUID, amount int) error {
	return s.updateBalance(ctx, walletID, -amount)
}

func (s *WalletService) GetBalance(ctx context.Context, walletID uuid.UUID) (int, error) {
	var balance int
	err := s.db.QueryRowContext(ctx, "SELECT balance FROM wallets WHERE id=$1", walletID).Scan(&balance)
	return balance, err
}

func (s *WalletService) updateBalance(ctx context.Context, walletID uuid.UUID, amount int) error {
	tx, err := s.db.BeginTx(ctx, &sql.TxOptions{Isolation: sql.LevelSerializable})
	if err != nil {
		return err
	}
	defer tx.Rollback()

	var balance int
	err = tx.QueryRowContext(ctx, "SELECT balance FROM wallets WHERE id=$1 FOR UPDATE", walletID).Scan(&balance)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return errors.New("wallet not found")
		}
		return err
	}

	newBalance := balance + amount
	if newBalance < 0 {
		return errors.New("insufficient funds")
	}

	_, err = tx.ExecContext(ctx, "UPDATE wallets SET balance=$1 WHERE id=$2", newBalance, walletID)
	if err != nil {
		return err
	}

	return tx.Commit()
}
