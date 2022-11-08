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

func DeleteArea(listenAddr string, bearer string, area string) error {
	conn, err := grpc.Dial(listenAddr, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		return errors.New(fmt.Sprintf("did not connect: %v", err))
	}
	defer func(conn *grpc.ClientConn) {
		err := conn.Close()
		if err != nil {

		}
	}(conn)

	c := pb.NewDeleteAreaClient(conn)

	// Contact the server and print out its response.
	ctx, cancel := context.WithTimeout(context.Background(), 120*time.Second)
	defer cancel()

	_, err = c.DeleteArea(ctx, &pb.DeleteAreaRequest{
		Bearer: bearer,
		Area:   area,
	})
	if err != nil {
		return errors.New(fmt.Sprintf("delete area failed: %s", err))
	}

	return err
}
