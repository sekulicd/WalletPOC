package grpchandler

import (
	pbwallet "WalletPOC/apidoc/grpc/gen"
	"WalletPOC/internal/core/application"
	"context"
)

type walletHandler struct {
	walletSvc application.WalletService
}

func NewWalletHandler(walletSvc application.WalletService) pbwallet.WalletServer {
	return walletHandler{
		walletSvc: walletSvc,
	}
}

func (w walletHandler) GenSeed(ctx context.Context, request *pbwallet.GenSeedRequest) (
	reply *pbwallet.GenSeedReply,
	errResp error,
) {
	seed, err := w.walletSvc.GenSeed(ctx)
	if err != nil {
		errResp = err
		return
	}
	reply = &pbwallet.GenSeedReply{
		SeedMnemonic: seed,
	}

	return
}

func (w walletHandler) InitWallet(ctx context.Context, request *pbwallet.InitWalletRequest) (*pbwallet.InitWalletReply, error) {
	panic("implement me")
}

func (w walletHandler) UnlockWallet(ctx context.Context, request *pbwallet.UnlockWalletRequest) (*pbwallet.UnlockWalletReply, error) {
	panic("implement me")
}

func (w walletHandler) ChangePassword(ctx context.Context, request *pbwallet.ChangePasswordRequest) (*pbwallet.ChangePasswordReply, error) {
	panic("implement me")
}

func (w walletHandler) WalletAddress(ctx context.Context, request *pbwallet.WalletAddressRequest) (*pbwallet.WalletAddressReply, error) {
	panic("implement me")
}

func (w walletHandler) WalletBalance(ctx context.Context, request *pbwallet.WalletBalanceRequest) (*pbwallet.WalletBalanceReply, error) {
	panic("implement me")
}

func (w walletHandler) SendToMany(ctx context.Context, request *pbwallet.SendToManyRequest) (*pbwallet.SendToManyReply, error) {
	panic("implement me")
}
