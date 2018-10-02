package main

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/caarlos0/env"
)

// Config is a set of options for junior.
type Config struct {
	Port            string `env:"PORT" envDefault:"80"`
	TrailingSlashes bool   `env:"TRAILING_SLASHES" envDefault:"false"`
	RootDir         string `env:"ROOT_DIR,required"`
	Has404          bool
}

// NFName (short for "not found name") is the name of the 404 file.
const NFName = "404.html"

// ReadConfig reads a Config from the system environment.
func ReadConfig() (*Config, error) {
	var (
		cfg = new(Config)
		err = env.Parse(cfg)
	)

	if err != nil {
		return nil, fmt.Errorf("junior: failed to read Config: %v", err)
	}
	if err = cfg.Validate(); err != nil {
		return nil, fmt.Errorf("junior: invalid config: %v", err)
	}

	if _, err := os.Stat(filepath.Join(cfg.RootDir, NFName)); err != nil {
		if !os.IsNotExist(err) {
			return nil, fmt.Errorf("junior: error while checking for \"%s\": %v",
				NFName, err)
		}
	} else {
		cfg.Has404 = true
	}

	return cfg, nil
}

// Validate returns an error if cfg is invalid.
func (cfg *Config) Validate() error {
	if _, err := os.Stat(cfg.RootDir); err != nil {
		if os.IsNotExist(err) {
			return fmt.Errorf("cfg.RootDir \"%s\" does not exist", err)
		}
		return fmt.Errorf("error while checking RootDir: %v", err)
	}
	return nil
}
