package server_test

import (
	"context"
	"net/http"
	"testing"

	"github.com/adamyeats/coffeeshops/graphql-server/internal/server"
	"github.com/adamyeats/coffeeshops/graphql-server/pkg/config"
)

func TestNew(t *testing.T) {
	// Create a test context
	ctx := context.Background()

	// Create a test config
	testConfig := &config.Config{
		Server: &config.ServerConfig{
			Port: ":8080",
		},
		DB: &config.DatabaseConfig{
			URL: "postgresql://localhost/mydatabase",
		},
	}

	s, err := server.New(ctx, testConfig)

	// Create a test server instance
	if err != nil {
		t.Errorf("New() returned an error: %v", err)
	}

	if s == nil {
		t.Error("New() returned a nil server")
	}

	go func() {
		if err := s.Serve(); err != nil {
			t.Errorf("Serve() returned an error: %v", err)
		}
	}()

	// Send a test request to the server
	resp, err := http.Get("http://localhost:8080/")

	if err != nil {
		t.Errorf("Error making request to server: %v", err)
	}

	if resp.StatusCode != http.StatusOK {
		t.Errorf("Expected status code %v, but got %v", http.StatusOK, resp.StatusCode)
	}

	// Stop the server
	if err := s.Stop(ctx); err != nil {
		t.Errorf("Stop() returned an error: %v", err)
	}
}
