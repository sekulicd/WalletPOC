package inmemory

import "WalletPOC/internal/core/domain"

type walletRepositoryImpl struct {
}

func NewWalletRepositoryImpl() domain.WalletRepository {
	return &walletRepositoryImpl{}
}

func (w walletRepositoryImpl) GetSeed() ([]string, error) {
	return []string{"color", "house", "work", "ball"}, nil
}
