package db

import (
	"context"
	"time"

	"sisco/ent/token"
	"sisco/internal/auth"
	"sisco/internal/cfg"
	"sisco/internal/ldapconn"
)

func (c *Client) GetSecretToken(ctx context.Context, user string, password string) (string, bool, error) {
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
