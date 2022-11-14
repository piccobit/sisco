package db

import (
	"context"
	"sisco/internal/auth"
	"sisco/internal/ldapconn"
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

func (c *Client) QuerySecretToken(ctx context.Context, user string, password string) (string, bool, error) {
	var err error

	authToken := auth.GenerateSecureToken(32)
	isAdminToken := false

	t, err := c.dbClient.Token.Query().Where(token.User(user)).Only(ctx)
	if t == nil {
		lc, err := ldapconn.New(&cfg.Config)
		if err != nil {
			return "", false, nil
		}

		authToken, isAdminToken, err = lc.Authenticate(user, password)
		if err != nil {
			return "", false, nil
		}

		_, err = c.dbClient.Token.Create().
			SetUser(user).
			SetToken(authToken).
			SetAdmin(isAdminToken).
			Save(ctx)
		if err != nil {
			return "", false, nil
		}
	} else {
		_, err = c.dbClient.Token.Update().Where(token.User(user)).SetCreated(time.Now()).Save(ctx)
		if err != nil {
			return "", false, nil
		}

		authToken = t.Token
		isAdminToken = t.Admin
	}

	return authToken, isAdminToken, nil
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
