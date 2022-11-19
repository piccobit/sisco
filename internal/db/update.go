package db

import (
	"context"
	"sisco/ent/service"
	"time"
)

func (c *Client) UpdateServiceAvailable(ctx context.Context, serviceName string, available bool) error {
	return c.dbClient.Service.
		Update().
		Where(service.NameEqualFold(serviceName)).
		SetAvailable(available).
		Exec(ctx)
}

func (c *Client) UpdateServiceHeartbeat(ctx context.Context, serviceName string, heartbeat time.Time) error {
	return c.dbClient.Service.
		Update().
		Where(service.NameEqualFold(serviceName)).
		SetHeartbeat(heartbeat).
		Exec(ctx)
}
