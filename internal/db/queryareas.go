package db

import (
	"context"

	"sisco/ent"
	"sisco/ent/area"
	"sisco/ent/service"
)

func (c *Client) QueryAreas(ctx context.Context) ([]*ent.Area, error) {
	return c.dbClient.Area.Query().WithServices().Order(ent.Asc(area.FieldID)).Order(ent.Asc(service.FieldID)).All(ctx)
}
