package db

import (
	"context"
	"errors"
	"fmt"

	"sisco/ent/area"
	"sisco/ent/service"
)

func (c *Client) DeleteArea(ctx context.Context, areaName string) error {
	numServices, err := c.dbClient.Service.Query().Where(service.HasAreaWith(area.Name(areaName))).Count(ctx)
	if err != nil {
		return err
	}

	if numServices > 0 {
		return errors.New(fmt.Sprintf("area '%s' is not empty", areaName))
	}

	_, err = c.dbClient.Area.Delete().Where(area.Name(areaName)).Exec(ctx)
	if err != nil {
		return err
	}

	return nil
}

func (c *Client) DeleteService(ctx context.Context, serviceName string, areaName string) error {
	_, err := c.dbClient.Service.Delete().
		Where(service.And(service.Name(serviceName), service.HasAreaWith(area.Name(areaName)))).Exec(ctx)
	if err != nil {
		return err
	}

	return nil
}
