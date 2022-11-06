package db

import (
	"context"
	"time"

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
