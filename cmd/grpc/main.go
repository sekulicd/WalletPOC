package main

import (
	pbwallet "WalletPOC/apidoc/grpc/gen"
	"WalletPOC/internal/core/application"
	"WalletPOC/internal/infrastructure/persistance/inmemory"
	grpchandler "WalletPOC/internal/interfaces/grpc"
	"WalletPOC/pkg/brontide"
	"github.com/btcsuite/btcutil"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
	"log"
)

const wif = "L4Aak5BJfaMemneuNR11DSQfnRLeVsMzRYx6JiqbyjY74ej3V7kV"

func main() {

	//serverCreds := auth.NewNoiseCredentials()
	//serverOpts := []grpc.ServerOption{grpc.Creds(serverCreds)}
	//walletGrpcServer := grpc.NewServer(serverOpts...)

	walletGrpcServer := grpc.NewServer()

	inMemoryWalletRepository := inmemory.NewWalletRepositoryImpl()
	walletSvc := application.NewWalletService(inMemoryWalletRepository)
	walletHandler := grpchandler.NewWalletHandler(walletSvc)

	pbwallet.RegisterWalletServer(walletGrpcServer, walletHandler)

	wifDecoded, err := btcutil.DecodeWIF(wif)
	if err != nil {
		log.Fatal(err)
	}
	localPriv := wifDecoded.PrivKey
	localKeyECDH := &brontide.PrivKeyECDH{PrivKey: localPriv}

	listener, err := brontide.NewListener(localKeyECDH, ":3333")
	if err != nil {
		log.Fatal(err)
	}

	reflection.Register(walletGrpcServer)
	err = walletGrpcServer.Serve(listener)
	if err != nil {
		log.Fatal(err)
	}
}
