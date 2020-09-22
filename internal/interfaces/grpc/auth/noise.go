package auth

import (
	"context"
	"google.golang.org/grpc/credentials"
	"net"
)

type noiseCredenstials struct {
}

func NewNoiseCredentials() credentials.TransportCredentials {

	return noiseCredenstials{}
}

func (n noiseCredenstials) ClientHandshake(ctx context.Context, s string, conn net.Conn) (net.Conn, credentials.AuthInfo, error) {
	panic("implement me")
}

func (n noiseCredenstials) ServerHandshake(conn net.Conn) (net.Conn, credentials.AuthInfo, error) {
	panic("implement me")
}

func (n noiseCredenstials) Info() credentials.ProtocolInfo {
	panic("implement me")
}

func (n noiseCredenstials) Clone() credentials.TransportCredentials {
	panic("implement me")
}

func (n noiseCredenstials) OverrideServerName(s string) error {
	panic("implement me")
}
