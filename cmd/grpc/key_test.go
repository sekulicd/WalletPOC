package main

import (
	wallet "WalletPOC/apidoc/grpc/gen"
	"WalletPOC/pkg/brontide"
	"context"
	"encoding/hex"
	"fmt"
	"github.com/btcsuite/btcd/btcec"
	"github.com/btcsuite/btcd/chaincfg"
	"github.com/btcsuite/btcutil"
	"github.com/stretchr/testify/assert"
	"google.golang.org/grpc"
	"testing"
)

func TestPrivKey(t *testing.T) {
	localPriv, err := btcec.NewPrivateKey(btcec.S256())
	if err != nil {
		t.Fatal(err)
	}

	wif1, err := btcutil.NewWIF(localPriv, &chaincfg.MainNetParams, true)
	if err != nil {
		t.Fatal(err)
	}

	fmt.Println(wif1.String())

	wifDecoded, err := btcutil.DecodeWIF(wif1.String())
	if err != nil {
		t.Fatal(err)
	}

	assert.Equal(
		t,
		hex.EncodeToString(wifDecoded.PrivKey.PubKey().SerializeCompressed()),
		hex.EncodeToString(localPriv.PubKey().SerializeCompressed()))
}

func TestConnection(t *testing.T) {
	wifDecoded, err := btcutil.DecodeWIF(wif)
	if err != nil {
		t.Fatal(err)
	}
	localPriv := wifDecoded.PrivKey

	remotePriv, err := btcec.NewPrivateKey(btcec.S256())
	if err != nil {
		t.Fatal(err)
	}

	credentials, err := brontide.NewClientCredentials(
		remotePriv,
		localPriv.PubKey(),
	)
	if err != nil {
		t.Fatal(err)
	}

	conn, err := grpc.Dial(":3333", grpc.WithTransportCredentials(credentials))
	if err != nil {
		t.Fatal(err)
	}
	client := wallet.NewWalletClient(conn)
	seed, err := client.GenSeed(context.Background(), &wallet.GenSeedRequest{})
	if err != nil {
		t.Fatal(err)
	}
	fmt.Println(seed.SeedMnemonic)

	defer conn.Close()
}
