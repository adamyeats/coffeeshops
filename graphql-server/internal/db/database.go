package db

import (
	"context"
	"fmt"

	"github.com/adamyeats/coffeeshops/graphql-server/internal/ent"
	"github.com/adamyeats/coffeeshops/graphql-server/pkg/config"
	_ "github.com/lib/pq"
)

type DB struct {
	Client *ent.Client
}

func New(ctx context.Context, config *config.DatabaseConfig) (*DB, error) {
	client, err := ent.Open("postgres", config.URL)

	if err != nil {
		return nil, fmt.Errorf("failed opening connection to Postgres: %v", err)
	}

	return &DB{Client: client}, nil
}

func (db *DB) Close() {
	db.Client.Close()
}
