package repository

import (
	"context"
	"database/sql"
	"errors"
	"github.com/google/uuid"
)

type WalletRepository struct {
	Db *sql.DB
}

func New(db *sql.DB) *WalletRepository {
	return &WalletRepository{Db: db}
}

func (r *WalletRepository) Deposit(ctx context.Context, id uuid.UUID, amount int64) (int64, error) {
	var balance int64

	err := r.Db.QueryRowContext(ctx, `
        UPDATE wallets
        SET balance = balance + $1
        WHERE id = $2
        RETURNING balance
    `, amount, id).Scan(&balance)

	return balance, err
}

func (r *WalletRepository) Withdraw(ctx context.Context, id uuid.UUID, amount int64) (int64, error) {
	var balance int64

	err := r.Db.QueryRowContext(ctx, `
        UPDATE wallets
        SET balance = balance - $1
        WHERE id = $2 AND balance >= $1
        RETURNING balance
    `, amount, id).Scan(&balance)

	if errors.Is(err, sql.ErrNoRows) {
		return 0, errors.New("insufficient funds")
	}

	return balance, err
}

func (r *WalletRepository) GetBalance(ctx context.Context, id uuid.UUID) (int64, error) {
	var balance int64
	err := r.Db.QueryRowContext(ctx,
		`SELECT balance FROM wallets WHERE id=$1`, id).Scan(&balance)
	return balance, err
}
