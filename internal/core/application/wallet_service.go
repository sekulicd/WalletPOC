package application

import (
	"WalletPOC/internal/core/domain"
	"context"
)

type WalletService interface {
	GenSeed(ctx context.Context) ([]string, error)
}

type walletService struct {
	walletRepository domain.WalletRepository
}

func NewWalletService(walletRepository domain.WalletRepository) WalletService {
	return &walletService{
		walletRepository: walletRepository,
	}
}

func (w walletService) GenSeed(ctx context.Context) ([]string, error) {
	return w.walletRepository.GetSeed()
}
