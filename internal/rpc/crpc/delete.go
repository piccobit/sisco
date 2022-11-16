package crpc

import (
	"context"
	"errors"
	"fmt"
	"sisco/internal/rpc/pb"
	"time"
)

func (c *Client) DeleteArea(bearer string, areaName string) error {
	dac := pb.NewDeleteAreaClient(c.grpcClient)

	// Contact the server and print out its response.
	ctx, cancel := context.WithTimeout(context.Background(), 120*time.Second)
	defer cancel()

	_, err := dac.DeleteArea(ctx, &pb.DeleteAreaRequest{
		Bearer: bearer,
		Area:   areaName,
	})
	if err != nil {
		return errors.New(fmt.Sprintf("delete area failed: %s", err))
	}

	return err
}

func (c *Client) DeleteService(bearer string, serviceName string, areaName string) error {
	dsc := pb.NewDeleteServiceClient(c.grpcClient)

	// Contact the server and print out its response.
	ctx, cancel := context.WithTimeout(context.Background(), 120*time.Second)
	defer cancel()

	_, err := dsc.DeleteService(ctx, &pb.DeleteServiceRequest{
		Bearer:  bearer,
		Service: serviceName,
		Area:    areaName,
	})
	if err != nil {
		return errors.New(fmt.Sprintf("delete area failed: %s", err))
	}

	return err
}
