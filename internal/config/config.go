package config

import (
	"log"
	"os"
	"time"
)

type Config struct {
	Typesense struct {
		URL    string
		APIKey string
	}

	Port         string
	OffloadAfter time.Duration
	SnapshotDir  string
}

func Load() Config {
	var c Config

	// Server
	c.Port = getEnv("PORT", "8080")

	// Typesense
	c.Typesense.URL = getEnv("TYPESENSE_URL", "http://localhost:8108")
	c.Typesense.APIKey = getEnv("TYPESENSE_API_KEY", "xyz")

	// Defaults
	c.OffloadAfter = 6 * time.Hour
	c.SnapshotDir = "./snapshots"

	// Validate (IMPORTANT)
	if c.Typesense.URL == "" {
		log.Fatal("TYPESENSE_URL is required (example: http://localhost:8108)")
	}
	if c.Typesense.APIKey == "" {
		log.Fatal("TYPESENSE_API_KEY is required")
	}

	return c
}

func getEnv(key, fallback string) string {
	if v := os.Getenv(key); v != "" {
		return v
	}
	return fallback
}
