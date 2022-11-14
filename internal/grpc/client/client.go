package client

import (
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
	"google.golang.org/grpc/credentials/insecure"
	"log"
	"sisco/internal/cfg"
)

type Client struct {
	grpcClient *grpc.ClientConn
}

func New(listenAddr string) (*Client, error) {
	var err error
	var creds credentials.TransportCredentials

	if cfg.Config.UseTLS {
		creds, err = credentials.NewClientTLSFromFile(cfg.Config.TLSCertFile, "")
		if err != nil {
			log.Fatalf("could not process the credentials: %v", err)
		}
	} else {
		creds = insecure.NewCredentials()
	}

	conn, err := grpc.Dial(listenAddr, grpc.WithTransportCredentials(creds))
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}

	client := Client{grpcClient: conn}

	return &client, nil
}

func (c *Client) Close() {
	err := c.grpcClient.Close()
	if err != nil {
		return
	}
}
