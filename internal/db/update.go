package db

import (
	"context"
	"sisco/ent/area"
	"sisco/ent/service"
	"time"
)

func (c *Client) UpdateServiceAvailable(ctx context.Context, serviceName string, areaName string, available bool) error {
	return c.dbClient.Service.
		Update().
		Where(
			service.And(
				service.NameEqualFold(serviceName),
				service.HasAreaWith(area.NameEqualFold(areaName)),
			),
		).
		SetAvailable(available).
		Exec(ctx)
}

func (c *Client) UpdateServiceHeartbeat(ctx context.Context, serviceName string, areaName string, heartbeat time.Time) error {
	return c.dbClient.Service.
		Update().
		Where(
			service.And(
				service.NameEqualFold(serviceName),
				service.HasAreaWith(area.NameEqualFold(areaName)),
			),
		).
		SetHeartbeat(heartbeat).
		Exec(ctx)
}

func (c *Client) UpdateServiceAvailableHeartbeat(ctx context.Context, serviceName string, areaName string, available bool, heartbeat time.Time) error {
	return c.dbClient.Service.
		Update().
		Where(
			service.And(
				service.NameEqualFold(serviceName),
				service.HasAreaWith(area.NameEqualFold(areaName)),
			),
		).
		SetAvailable(available).
		SetHeartbeat(heartbeat).
		Exec(ctx)
}
