package config

import "time"

type Config struct {
	Typesense struct {
		URL    string
		APIKey string
	}

	OffloadAfter time.Duration
	SnapshotDir  string
	Port         string
}

func Default() Config {
	return Config{
		OffloadAfter: 6 * time.Hour,
		SnapshotDir:  "./snapshots",
		Port:         "8080",
	}
}
