package brontide

import (
	"context"
	"github.com/btcsuite/btcd/btcec"
	"google.golang.org/grpc/credentials"
	"io"
	"net"
	"time"
)

type NoiseInfo struct {
}

func (n NoiseInfo) AuthType() string {
	return "noise"
}

type noiseCredentials struct {
	clientPrivKeyECDH *PrivKeyECDH
	serverPubKey      *btcec.PublicKey
}

func NewClientCredentials(
	privKey *btcec.PrivateKey,
	serverPubKey *btcec.PublicKey,
) (credentials.TransportCredentials, error) {

	return noiseCredentials{
		clientPrivKeyECDH: &PrivKeyECDH{PrivKey: privKey},
		serverPubKey:      serverPubKey,
	}, nil
}

func NewServerCredentials() credentials.TransportCredentials {
	return noiseCredentials{}
}

func (n noiseCredentials) ClientHandshake(ctx context.Context, s string, conn net.Conn) (net.Conn, credentials.AuthInfo, error) {

	b := &Conn{
		conn:  conn,
		noise: NewBrontideMachine(true, n.clientPrivKeyECDH, n.serverPubKey),
	}

	// Initiate the handshake by sending the first act to the receiver.
	actOne, err := b.noise.GenActOne()
	if err != nil {
		b.conn.Close()
		return nil, nil, err
	}
	if _, err := conn.Write(actOne[:]); err != nil {
		b.conn.Close()
		return nil, nil, err
	}

	// We'll ensure that we get ActTwo from the remote peer in a timely
	// manner. If they don't respond within 1s, then we'll kill the
	// connection.
	err = conn.SetReadDeadline(time.Now().Add(handshakeReadTimeout))
	if err != nil {
		b.conn.Close()
		return nil, nil, err
	}

	// If the first act was successful (we know that address is actually
	// remotePub), then read the second act after which we'll be able to
	// send our static public key to the remote peer with strong forward
	// secrecy.
	var actTwo [ActTwoSize]byte
	if _, err := io.ReadFull(conn, actTwo[:]); err != nil {
		b.conn.Close()
		return nil, nil, err
	}
	if err := b.noise.RecvActTwo(actTwo); err != nil {
		b.conn.Close()
		return nil, nil, err
	}

	// Finally, complete the handshake by sending over our encrypted static
	// key and execute the final ECDH operation.
	actThree, err := b.noise.GenActThree()
	if err != nil {
		b.conn.Close()
		return nil, nil, err
	}
	if _, err := conn.Write(actThree[:]); err != nil {
		b.conn.Close()
		return nil, nil, err
	}

	// We'll reset the deadline as it's no longer critical beyond the
	// initial handshake.
	err = conn.SetReadDeadline(time.Time{})
	if err != nil {
		b.conn.Close()
		return nil, nil, err
	}

	return b, NoiseInfo{}, nil
}

func (n noiseCredentials) ServerHandshake(conn net.Conn) (net.Conn, credentials.AuthInfo, error) {
	panic("implement me")
}

func (n noiseCredentials) Info() credentials.ProtocolInfo {
	return credentials.ProtocolInfo{
		"1.0",
		"noise",
		"1.0",
		"tdex",
	}
}

func (n noiseCredentials) Clone() credentials.TransportCredentials {
	cred, err := NewClientCredentials(
		n.clientPrivKeyECDH.PrivKey,
		n.serverPubKey,
	)
	if err != nil {
		panic(err)
	}
	return cred
}

func (n noiseCredentials) OverrideServerName(s string) error {
	panic("implement me")
}
