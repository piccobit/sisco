package db

import (
	"context"
	"time"

	"sisco/ent"
	"sisco/ent/area"
	"sisco/ent/service"
	"sisco/ent/token"
	"sisco/internal/cfg"
)

func (c *Client) CheckToken(ctx context.Context, bearer string, isAdminToken bool) (bool, error) {
	t, err := c.dbClient.Token.Query().Where(token.Token(bearer)).Only(ctx)
	if err != nil {
		return false, err
	}

	if int(time.Now().Sub(t.Created).Seconds()) > cfg.Config.TokenValidInSeconds {
		return false, err
	}

	if isAdminToken {
		if !t.Admin {
			return false, err
		}
	}

	return true, nil
}

func (c *Client) QueryAreas(ctx context.Context) ([]*ent.Area, error) {
	return c.dbClient.Area.Query().WithServices().Order(ent.Asc(area.FieldID)).Order(ent.Asc(service.FieldID)).All(ctx)
}

func (c *Client) QueryServiceInArea(ctx context.Context, serviceName string, areaName string) (*ent.Service, error) {
	return c.dbClient.Service.Query().
		Where(service.And(service.Name(serviceName), service.HasAreaWith(area.Name(areaName)))).
		WithTags().
		Order(ent.Asc(service.FieldID)).
		Only(ctx)
}

func (c *Client) QueryServicesInArea(ctx context.Context, areaName string) ([]*ent.Service, error) {
	return c.dbClient.Service.Query().
		Where(service.HasAreaWith(area.Name(areaName))).
		WithTags().
		Order(ent.Asc(service.FieldID)).
		All(ctx)
}
