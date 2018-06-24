package config

import (
	"os"
)

func loadStagingConfig(cfg *Config) {
	cfg.DatabaseHost = os.Getenv("DATABASE_HOST")
	cfg.DatabaseName = os.Getenv("DATABASE_NAME")
	cfg.DatabaseSSLMode = os.Getenv("DATABASE_SSL_MODE")
	cfg.DatabaseUser = os.Getenv("DATABASE_USER")
	cfg.Environment = "staging"
}
