package db

import (
	"fmt"
	"log"

	"sisco/ent"
	"sisco/internal/cfg"
)

type Client struct {
	dbClient *ent.Client
}

func New() (*Client, error) {
	var dbURL string

	dbType := cfg.Config.DBType

	switch dbType {
	case "postgres":
		dbURL = fmt.Sprintf("postgres://%s:%s@%s:%d/%s?sslmode=%s",
			cfg.Config.DBUser,
			cfg.Config.DBPassword,
			cfg.Config.DBHost,
			cfg.Config.DBPort,
			cfg.Config.DBName,
			cfg.Config.DBSSLMode,
		)
	case "mysql":
		dbURL = fmt.Sprintf("mysql://%s:%s@%s:%d/%s",
			cfg.Config.DBUser,
			cfg.Config.DBPassword,
			cfg.Config.DBHost,
			cfg.Config.DBPort,
			cfg.Config.DBName,
		)
	default:
		log.Fatalf("unknown database type: %s", dbType)
	}

	c, err := ent.Open(dbType, dbURL)
	if err != nil {
		return nil, err
	}

	dbc := Client{
		dbClient: c,
	}

	return &dbc, nil
}

func (c *Client) Close() {
	c.dbClient.Close()
}
