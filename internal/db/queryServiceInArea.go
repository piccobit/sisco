package db

import (
	"context"

	"sisco/ent"
	"sisco/ent/area"
	"sisco/ent/service"
)

func (c *Client) QueryServiceInArea(ctx context.Context, serviceName string, areaName string) (*ent.Service, error) {
	return c.dbClient.Service.Query().
		Where(service.And(service.Name(serviceName), service.HasAreaWith(area.Name(areaName)))).
		WithTags().
		Order(ent.Asc(service.FieldID)).
		Only(ctx)
}
