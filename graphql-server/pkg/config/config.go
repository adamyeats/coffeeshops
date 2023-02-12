package config

import (
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
)

// ServerConfig is the server config struct
type ServerConfig struct {
	// Port is the port the server will listen on
	// e.g. 8080. Note that this is not prefixed with a colon.
	Port string
}

// DatabaseConfig is the database config struct
type DatabaseConfig struct {
	// URL is the connection string for the database
	URL string
}

// Config is the main config struct
type Config struct {
	Server *ServerConfig
	DB     *DatabaseConfig
}

// FromEnv loads the config from the environment
func FromEnv() *Config {
	if err := godotenv.Load(); err != nil {
		log.Print("No .env file found")
	}

	port := os.Getenv("PORT")

	if port != "" {
		// set default port
		port = "8080"
	}

	sc := &ServerConfig{
		Port: fmt.Sprintf(":%s", port),
	}

	dbc := &DatabaseConfig{}

	return &Config{Server: sc, DB: dbc}
}
