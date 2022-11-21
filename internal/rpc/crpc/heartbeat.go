package crpc

import (
	"context"
	"errors"
	"fmt"
	"sisco/internal/rpc/pb"
	"time"
)

func (c *Client) Heartbeat(bearer string, serviceName string, areaName string) error {
	hb := pb.NewHeartbeatClient(c.grpcClient)

	// Contact the server and print out its response.
	ctx, cancel := context.WithTimeout(context.Background(), 120*time.Second)
	defer cancel()

	_, err := hb.Heartbeat(ctx, &pb.HeartbeatRequest{
		Bearer:  bearer,
		Service: serviceName,
		Area:    areaName,
	})
	if err != nil {
		return errors.New(fmt.Sprintf("heartbeat failed: %s", err))
	}

	return nil
}
