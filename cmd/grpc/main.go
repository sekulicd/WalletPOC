package main

import (
	pbwallet "WalletPOC/apidoc/grpc/gen"
	"WalletPOC/internal/core/application"
	"WalletPOC/internal/infrastructure/persistance/inmemory"
	grpchandler "WalletPOC/internal/interfaces/grpc"
	"fmt"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
	"log"
	"net"
)

func main() {

	walletGrpcServer := grpc.NewServer()

	inMemoryWalletRepository := inmemory.NewWalletRepositoryImpl()
	walletSvc := application.NewWalletService(inMemoryWalletRepository)
	walletHandler := grpchandler.NewWalletHandler(walletSvc)

	pbwallet.RegisterWalletServer(walletGrpcServer, walletHandler)

	listen, err := net.Listen("tcp", ":3333")
	if err != nil {
		fmt.Printf("failed to listen: %v\n", err)
		return
	}

	reflection.Register(walletGrpcServer)
	err = walletGrpcServer.Serve(listen)
	if err != nil {
		log.Fatal(err)
	}
}
