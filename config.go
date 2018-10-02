package main

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"

	"gopkg.in/urfave/cli.v1"
)

// Config is a set of options for junior.
type Config struct {
	// TrailingSlash dictates whether or not to enforce a trailing slash using
	// a HTTP 301 "Permanently Moved" redirect.
	//
	// If "true", will redirect "/subpath" to "/subpath/".
	// If "false", will redirect "/subpath/" to "/subpath".
	// If any other value, will not enforce any rule; no redirects will be
	// performed.
	TrailingSlash string `env:"TRAILING_SLASH"`

	// NotFound is the location for the file that will be served when a page
	// cannot be found.
	//
	// If it is not absolute, it will be interpreted relative to RootDir.
	//
	// If it does not exist, a default 404 page will be served.
	NotFound string `env:"NOT_FOUND" envDefault:"404.html"`

	Port    string `env:"PORT"`
	RootDir string `env:"ROOT_DIR,required"`
}

// ReadConfig reads a Config from a cli.Context.
func ReadConfig(ctx *cli.Context) *Config {
	return &Config{
		Port:          ctx.String("port"),
		TrailingSlash: ctx.String("trailing-slash"),
		RootDir:       ctx.String("root-dir"),
		NotFound:      ctx.String("not-found"),
	}
}

// DefaultFile is the default file of a directory that will be returned if
// the path has no file extension.
const DefaultFile = "index.html"

// Validate returns an error if cfg is invalid.
func (cfg *Config) Validate() error {
	// Check for cfg.RootDir existence.
	if _, err := os.Stat(cfg.RootDir); err != nil {
		if os.IsNotExist(err) {
			return fmt.Errorf("cfg.RootDir \"%s\" does not exist", err)
		}
		return fmt.Errorf("error while checking RootDir: %v", err)
	}

	// Ensure cfg.NotFound is relative to cfg.RootDir, and is a file.
	if !filepath.IsAbs(cfg.NotFound) {
		cfg.NotFound = filepath.Join(cfg.RootDir, cfg.NotFound)
	}
	if strings.IndexRune(filepath.Base(cfg.NotFound), '.') == 0 {
		cfg.NotFound = filepath.Join(cfg.RootDir, DefaultFile)
	}
	return nil
}
