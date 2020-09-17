package domain

type WalletRepository interface {
	GetSeed() ([]string, error)
}
