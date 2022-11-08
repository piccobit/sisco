package db

import (
	"context"
	"errors"
	"fmt"

	"sisco/ent"
	"sisco/ent/area"
	"sisco/ent/tag"
)

func (c *Client) CreateService(ctx context.Context, serviceName string, areaName string, description string, protocol string, host string, port string, serviceTags []string) error {
	if ok, err := c.dbClient.Area.Query().Where(area.Name(areaName)).Exist(ctx); !ok || err != nil {
		return errors.New(fmt.Sprintf("area %s not found", areaName))
	}

	var tagEntries []*ent.Tag

	for _, tagName := range serviceTags {
		var err error

		t, _ := c.dbClient.Tag.Query().Where(tag.Name(tagName)).Only(ctx)
		if t == nil {
			t, err = c.dbClient.Tag.Create().SetName(tagName).Save(ctx)
			if err != nil {
				return err
			}
		}

		tagEntries = append(tagEntries, t)
	}

	s, err := c.dbClient.Service.Create().
		SetName(serviceName).
		SetDescription(description).
		SetProtocol(protocol).
		SetHost(host).
		SetPort(port).
		AddTags(tagEntries...).
		Save(ctx)

	_, err = c.dbClient.Area.Update().Where(area.Name(areaName)).AddServices(s).Save(ctx)
	if err != nil {
		return err
	}

	return err
}
