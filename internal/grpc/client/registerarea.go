package client

import (
	"context"
	"errors"
	"fmt"
	"time"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"sisco/pb"
)

func RegisterArea(listenAddr string, bearer string, area string, description string) error {
	conn, err := grpc.Dial(listenAddr, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		return errors.New(fmt.Sprintf("did not connect: %v", err))
	}
	defer func(conn *grpc.ClientConn) {
		err := conn.Close()
		if err != nil {

		}
	}(conn)

	c := pb.NewRegisterAreaClient(conn)

	// Contact the server and print out its response.
	ctx, cancel := context.WithTimeout(context.Background(), 120*time.Second)
	defer cancel()

	_, err = c.RegisterArea(ctx, &pb.RegisterAreaRequest{
		Bearer:      bearer,
		Area:        area,
		Description: description,
	})
	if err != nil {
		return errors.New(fmt.Sprintf("adding area failed: %s", err))
	}

	return err
}
