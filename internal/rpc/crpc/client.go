package crpc

import (
	"fmt"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
	"google.golang.org/grpc/credentials/insecure"
	"sisco/internal/cfg"
	"sisco/internal/exit"
)

type Client struct {
	grpcClient  *grpc.ClientConn
	listenAddr  string
	useTLS      bool
	tlsCertFile string
}

func Default() (*Client, error) {
	listenAddr := fmt.Sprintf(":%d", cfg.Config.GRPCPort)

	rpcClient, err := New(
		ListenAddr(listenAddr),
		UseTLS(cfg.Config.UseTLS),
		TLSCertFile(cfg.Config.TLSCertFile),
	)
	if err != nil {
		return nil, err
	}

	return rpcClient, nil
}

func New(opts ...func(*Client)) (*Client, error) {
	var err error
	var creds credentials.TransportCredentials

	newClient := Client{}

	for _, c := range opts {
		c(&newClient)
	}

	if newClient.useTLS {
		creds, err = credentials.NewClientTLSFromFile(newClient.tlsCertFile, "")
		if err != nil {
			exit.Fatalf(1, "could not process the credentials: %v", err)
		}
	} else {
		creds = insecure.NewCredentials()
	}

	conn, err := grpc.Dial(newClient.listenAddr, grpc.WithTransportCredentials(creds))
	if err != nil {
		exit.Fatalf(1, "did not connect: %v", err)
	}

	newClient.grpcClient = conn

	return &newClient, nil
}

func (c *Client) Close() {
	err := c.grpcClient.Close()
	if err != nil {
		return
	}
}

func ListenAddr(listenAddr string) func(*Client) {
	return func(client *Client) {
		client.listenAddr = listenAddr
	}
}

func UseTLS(yesOrNo bool) func(*Client) {
	return func(client *Client) {
		client.useTLS = yesOrNo
	}
}

func TLSCertFile(tlsCertFile string) func(*Client) {
	return func(client *Client) {
		client.tlsCertFile = tlsCertFile
	}
}
