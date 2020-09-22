package main

import (
	"WalletPOC/pkg/brontide"
	"encoding/hex"
	"fmt"
	"github.com/btcsuite/btcd/btcec"
	"github.com/btcsuite/btcd/chaincfg"
	"github.com/btcsuite/btcutil"
	"github.com/stretchr/testify/assert"
	"log"
	"net"
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
		log.Fatal(err)
	}
	localPriv := wifDecoded.PrivKey

	//localPriv, err := btcec.NewPrivateKey(btcec.S256())
	//if err != nil {
	//	t.Fatal(err)
	//}

	remotePriv, err := btcec.NewPrivateKey(btcec.S256())
	if err != nil {
		t.Fatal(err)
	}
	remoteKeyECDH := &brontide.PrivKeyECDH{PrivKey: remotePriv}

	addr, err := net.ResolveTCPAddr("tcp", ":3333")
	if err != nil {
		t.Fatal(err)
	}

	netAddr := &brontide.NetAddress{
		IdentityKey: localPriv.PubKey(),
		Address:     addr,
	}

	remoteConn, err := brontide.Dial(remoteKeyECDH, netAddr, net.Dial)
	if err != nil {
		t.Fatal(err)
	}

	//conn, err := grpc.Dial(":3333", grpc.WithInsecure())
	//if err != nil {
	//	fmt.Println(err)
	//}
	//defer conn.Close()
	//wallet.NewWalletClient(conn)

	defer remoteConn.Close()
}
