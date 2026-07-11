package config

import (
	"errors"
	"os"
	"path/filepath"

	"github.com/pelletier/go-toml/v2"
)

type Config struct {
	Editor string         `toml:"editor"`
	Apps   map[string]App `toml:"apps"`
}

type App struct {
	Paths []string `toml:"paths"`
}

// Dir resolves the configuration directory for the application.
func Dir() (string, error) {
	if d := os.Getenv("DOTS_CONFIG_DIR"); d != "" {
		return d, nil
	}
	if x := os.Getenv("XDG_CONFIG_HOME"); x != "" {
		return filepath.Join(x, "dots"), nil
	}
	home, err := os.UserHomeDir()
	if err != nil {
		return "", err
	}
	return filepath.Join(home, ".config", "dots"), nil
}

// Path returns the full path to the config file.
func Path() (string, error) {
	dir, err := Dir()
	if err != nil {
		return "", err
	}
	return filepath.Join(dir, "config.toml"), nil
}

// Load reads the configuration from the config file and returns it.
func Load() (Config, bool, error) {
	var cfg Config

	path, err := Path()
	if err != nil {
		return cfg, false, err
	}

	data, err := os.ReadFile(path)
	if errors.Is(err, os.ErrNotExist) {
		return cfg, false, nil // no file yet -> empty config, not an error
	}
	if err != nil {
		return cfg, false, err
	}

	if err := toml.Unmarshal(data, &cfg); err != nil {
		return cfg, false, err
	}
	return cfg, true, nil
}

// Save writes the given configuration to the config file.
func Save(cfg Config) error {
	// TODO: does not preserve comments etc. Fix this later
	dir, err := Dir()
	if err != nil {
		return err
	}
	if err := os.MkdirAll(dir, 0o755); err != nil { // create dir tree; no-op if present
		return err
	}

	data, err := toml.Marshal(cfg)
	if err != nil {
		return err
	}

	path := filepath.Join(dir, "config.toml")
	return os.WriteFile(path, data, 0o644) // 0o644 = rw-r--r-- owner writes, all read
}
