package db

import (
	"context"
)

func (c *Client) CreateArea(ctx context.Context, area string, description string) error {
	_, err := c.dbClient.Area.Create().
		SetName(area).
		SetDescription(description).
		Save(ctx)

	return err
}