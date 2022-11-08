package db

import (
	"context"

	"sisco/ent"
	"sisco/ent/area"
	"sisco/ent/service"
)

func (c *Client) QueryServicesInArea(ctx context.Context, areaName string) ([]*ent.Service, error) {
	return c.dbClient.Service.Query().
		Where(service.HasAreaWith(area.Name(areaName))).
		WithTags().
		Order(ent.Asc(service.FieldID)).
		All(ctx)
}
