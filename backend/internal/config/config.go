// Package config reads a node's settings from its environment.
package config

import (
	"errors"
	"os"
)

// Config is everything a node needs to start.
type Config struct {
	// Addr is the address the HTTP API listens on.
	Addr string

	// DatabaseURL points at the PostgreSQL instance holding this node's data.
	DatabaseURL string
}

// Load reads the configuration, refusing to start rather than falling back to a
// guess. A node that silently starts against the wrong database is worse than
// one that does not start.
func Load() (Config, error) {
	cfg := Config{
		Addr:        envOr("ADDR", ":8080"),
		DatabaseURL: os.Getenv("DATABASE_URL"),
	}

	if cfg.DatabaseURL == "" {
		return Config{}, errors.New("DATABASE_URL is not set")
	}

	return cfg, nil
}

func envOr(key, fallback string) string {
	if v := os.Getenv(key); v != "" {
		return v
	}
	return fallback
}
