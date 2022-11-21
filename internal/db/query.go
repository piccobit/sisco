package db

import (
	"context"
	"errors"
	"sisco/ent/tag"
	"sisco/internal/auth"
	"sisco/internal/ldapconn"
	"time"

	"sisco/ent"
	"sisco/ent/area"
	"sisco/ent/service"
	"sisco/ent/token"
	"sisco/internal/cfg"
)

func (c *Client) CheckToken(ctx context.Context, bearer string, permissions uint64) (bool, error) {
	t, err := c.dbClient.Token.Query().Where(token.Token(bearer)).Only(ctx)
	if err != nil {
		return false, err
	}

	if int(time.Now().Sub(t.Created).Seconds()) > cfg.Config.TokenValidInSeconds {
		err = errors.New("token is not valid anymore")
		return false, err
	}

	if (t.Permissions & permissions) == 0 {
		err = errors.New("token is not an admin token")
		return false, err
	}

	return true, nil
}

func (c *Client) QuerySecretToken(ctx context.Context, user string, password string) (string, auth.Permissions, error) {
	var err error

	authToken := auth.GenerateSecureToken(32)

	lc, err := ldapconn.New(&cfg.Config)
	if err != nil {
		return "", auth.Unknown, nil
	}

	permissions, err := lc.Authenticate(user, password)
	if err != nil {
		return "", auth.Unknown, nil
	}

	t, err := c.dbClient.Token.Query().Where(token.User(user)).Only(ctx)
	if t == nil {
		_, err = c.dbClient.Token.Create().
			SetUser(user).
			SetToken(authToken).
			SetPermissions(uint64(permissions)).
			Save(ctx)
		if err != nil {
			return "", auth.Unknown, nil
		}
	} else {
		_, err = c.dbClient.Token.Update().
			Where(token.User(user)).
			SetToken(authToken).
			SetPermissions(uint64(permissions)).
			SetCreated(time.Now()).
			Save(ctx)
		if err != nil {
			return "", auth.Unknown, nil
		}
	}

	return authToken, permissions, nil
}
func (c *Client) QueryAreas(ctx context.Context) ([]*ent.Area, error) {
	return c.dbClient.Area.Query().WithServices().Order(ent.Asc(area.FieldID)).Order(ent.Asc(service.FieldID)).All(ctx)
}

func (c *Client) QueryService(ctx context.Context, serviceName string, areaName string) (*ent.Service, error) {
	return c.dbClient.Service.Query().
		Where(service.And(service.Name(serviceName), service.HasAreaWith(area.Name(areaName)))).
		WithTags().
		Order(ent.Asc(service.FieldID)).
		Only(ctx)
}

func (c *Client) QueryServices(ctx context.Context, areaName string, tagName string) ([]*ent.Service, error) {
	if len(areaName) == 0 && len(tagName) == 0 {
		// Query all services unrestricted.
		return c.dbClient.Service.Query().
			WithArea().
			WithTags().
			Order(ent.Asc(service.FieldID)).
			All(ctx)
	} else if len(areaName) != 0 && len(tagName) != 0 {
		// Query services restricted by area & tag.
		return c.dbClient.Service.Query().
			Where(
				service.And(
					service.HasAreaWith(area.NameEqualFold(areaName)),
					service.HasTagsWith(tag.NameEqualFold(tagName))),
			).
			WithArea().
			WithTags().
			Order(ent.Asc(service.FieldID)).
			All(ctx)
	} else if len(areaName) != 0 && len(tagName) == 0 {
		// Query services restricted by area only.
		return c.dbClient.Service.Query().
			Where(
				service.HasAreaWith(area.NameEqualFold(areaName)),
			).
			WithArea().
			WithTags().
			Order(ent.Asc(service.FieldID)).
			All(ctx)
	} else {
		// Query services restricted by tag only.
		return c.dbClient.Service.Query().
			Where(
				service.HasTagsWith(tag.NameEqualFold(tagName)),
			).
			WithArea().
			WithTags().
			Order(ent.Asc(service.FieldID)).
			All(ctx)
	}
}

func (c *Client) QueryTags(ctx context.Context) ([]*ent.Tag, error) {
	return c.dbClient.Tag.Query().Order(ent.Asc(tag.FieldName)).All(ctx)
}
