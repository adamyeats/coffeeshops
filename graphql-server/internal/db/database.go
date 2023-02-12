package db

import (
	"context"
	"fmt"
	"os"

	"github.com/adamyeats/coffeeshops/graphql-server/internal/ent"
	_ "github.com/lib/pq"
)

type DB struct {
	Client *ent.Client
}

func New(ctx context.Context) (*DB, error) {
	client, err := ent.Open("postgres", os.Getenv("DATABASE_URL"))

	if err != nil {
		return nil, fmt.Errorf("failed opening connection to Postgres: %v", err)
	}

	return &DB{Client: client}, nil
}

func (db *DB) Close() {
	db.Client.Close()
}
