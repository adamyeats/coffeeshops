package db_test

import (
	"context"
	"testing"

	_ "github.com/mattn/go-sqlite3"

	"github.com/adamyeats/coffeeshops/graphql-server/internal/db"
	"github.com/adamyeats/coffeeshops/graphql-server/internal/ent/enttest"
	"github.com/adamyeats/coffeeshops/graphql-server/pkg/config"
)

func TestNewDB(t *testing.T) {
	ctx := context.Background()
	uri := "sqlite3://file:ent?mode=memory&_fk=1"

	// use a temporary test database
	client := enttest.Open(t, "sqlite3", uri)
	defer client.Close()

	// set up the test database schema
	if err := client.Schema.Create(ctx); err != nil {
		t.Fatalf("Failed to create test schema: %v", err)
	}

	config := &config.DatabaseConfig{URL: uri}

	db, err := db.New(ctx, config)

	if err != nil {
		t.Fatalf("Failed to create new DB instance: %v", err)
	}

	if db.Client == nil {
		t.Error("DB client is nil")
	}

	// clean up
	t.Cleanup(func() {
		err := client.Close()

		if err == nil {
			t.Error("Failed to close DB client")
		}
	})

	if err := client.Close(); err != nil {
		t.Errorf("Failed to close test database: %v", err)
	}
}
