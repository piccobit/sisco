package client

import (
	"context"
	"google.golang.org/grpc/credentials"
	"log"
	"sisco/internal/cfg"
	"time"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"sisco/pb"
)

func Login(listenAddr string, user string, password string) (string, bool, error) {
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
	defer func(conn *grpc.ClientConn) {
		err := conn.Close()
		if err != nil {

		}
	}(conn)

	c := pb.NewLoginClient(conn)

	// Contact the server and print out its response.
	ctx, cancel := context.WithTimeout(context.Background(), 120*time.Second)
	defer cancel()

	r, err := c.Login(ctx, &pb.LoginRequest{
		User:     user,
		Password: password,
	})
	if err != nil {
		log.Fatalf("login failed: %s", err)
	}

	return r.GetToken(), r.GetIsAdminToken(), err
}
