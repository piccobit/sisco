package client

import (
	"context"
	"errors"
	"fmt"
	"time"

	"sisco/pb"
)

func (c *Client) RegisterArea(bearer string, area string, description string) error {
	rac := pb.NewRegisterAreaClient(c.grpcClient)

	// Contact the server and print out its response.
	ctx, cancel := context.WithTimeout(context.Background(), 120*time.Second)
	defer cancel()

	_, err := rac.RegisterArea(ctx, &pb.RegisterAreaRequest{
		Bearer:      bearer,
		Area:        area,
		Description: description,
	})
	if err != nil {
		return errors.New(fmt.Sprintf("adding area failed: %s", err))
	}

	return err
}

func (c *Client) RegisterService(bearer string, serviceName string, areaName string, description string, protocol string, host string, port string, tags ...string) error {
	rsc := pb.NewRegisterServiceClient(c.grpcClient)

	// Contact the server and print out its response.
	ctx, cancel := context.WithTimeout(context.Background(), 120*time.Second)
	defer cancel()

	_, err := rsc.RegisterService(ctx, &pb.RegisterServiceRequest{
		Bearer:      bearer,
		Service:     serviceName,
		Area:        areaName,
		Description: description,
		Protocol:    protocol,
		Host:        host,
		Port:        port,
		Tags:        tags,
	})
	if err != nil {
		return errors.New(fmt.Sprintf("adding area failed: %s", err))
	}

	return err
}
