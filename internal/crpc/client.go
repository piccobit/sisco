package crpc

import (
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
	"google.golang.org/grpc/credentials/insecure"
	"log"
)

type Client struct {
	grpcClient  *grpc.ClientConn
	listenAddr  string
	useTLS      bool
	tlsCertFile string
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
			log.Fatalf("could not process the credentials: %v", err)
		}
	} else {
		creds = insecure.NewCredentials()
	}

	conn, err := grpc.Dial(newClient.listenAddr, grpc.WithTransportCredentials(creds))
	if err != nil {
		log.Fatalf("did not connect: %v", err)
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
