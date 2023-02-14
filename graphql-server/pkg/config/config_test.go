package config_test

import (
	"os"
	"testing"

	"github.com/adamyeats/coffeeshops/graphql-server/pkg/config"
)

func TestFromEnv(t *testing.T) {
	os.Setenv("PORT", "8080")
	os.Setenv("DATABASE_URL", "test://user:password@host/dbname")

	// Ensure environment variables are cleaned up
	defer func() {
		os.Unsetenv("PORT")
		os.Unsetenv("DATABASE_URL")
	}()

	config := config.FromEnv()

	if config.Server.Port != ":8080" {
		t.Errorf("Expected server port :8080, but got %s", config.Server.Port)
	}

	if config.DB.URL != "test://user:password@host/dbname" {
		t.Errorf("Expected database URL test://user:password@host/dbname, but got %s", config.DB.URL)
	}
}
