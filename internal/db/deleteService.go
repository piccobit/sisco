package db

import (
	"context"

	"sisco/ent/area"
	"sisco/ent/service"
)

func (c *Client) DeleteService(ctx context.Context, serviceName string, areaName string) error {
	_, err := c.dbClient.Service.Delete().
		Where(service.And(service.Name(serviceName), service.HasAreaWith(area.Name(areaName)))).Exec(ctx)
	if err != nil {
		return err
	}

	return nil
}
