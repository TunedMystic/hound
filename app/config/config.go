package config

import (
	"os"
	"path/filepath"
)

// Config ...
type Config struct {
	BaseDir      string
	DatabaseName string
}

// GetConfig ...
func GetConfig() *Config {
	c := Config{}
	c.BaseDir = filepath.Dir(os.Getenv("BASE_DIR"))
	c.DatabaseName = os.Getenv("DATABASE_NAME")
	return &c
}
