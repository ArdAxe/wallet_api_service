package service

import (
	"context"
	"errors"
	"wallet_api_service/internal/model"
	"wallet_api_service/internal/repository"

	"github.com/google/uuid"
)

type WalletService struct {
	repo *repository.WalletRepository
}

func New(repo *repository.WalletRepository) *WalletService {
	return &WalletService{repo: repo}
}

func (s *WalletService) Process(ctx context.Context, id uuid.UUID, op model.OperationType, amount int64) (int64, error) {
	switch op {
	case model.Deposit:
		return s.repo.Deposit(ctx, id, amount)
	case model.Withdraw:
		return s.repo.Withdraw(ctx, id, amount)
	default:
		return 0, errors.New("invalid operation")
	}
}

func (s *WalletService) GetBalance(ctx context.Context, id uuid.UUID) (int64, error) {
	return s.repo.GetBalance(ctx, id)
}
