package config

import (
	"errors"
	"os"
	"path/filepath"

	"github.com/pelletier/go-toml/v2"
)

type Config struct {
	Editor      string         `toml:"editor"`
	UseDefaults *bool          `toml:"use_defaults,omitempty"`
	Apps        map[string]App `toml:"apps"`
}

type App struct {
	Paths []string `toml:"paths"`
}

// DefaultsEnabled reports wether the built-in app registry should be
// layerd in. Absent on disk (nil) means "defer to the code default".
func (c Config) DefaultsEnabled() bool {
	return c.UseDefaults == nil || *c.UseDefaults
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

// Save writes the config atomically
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

	// Temp file in the SAME dir so the final rename stays on one filesystem
	tmp, err := os.CreateTemp(dir, "config-*toml.tmp")
	if err != nil {
		return err
	}
	tmpName := tmp.Name()
	defer os.Remove(tmpName) // no-op after a successfull rename; cleans up on failure

	if _, err := tmp.Write(data); err != nil {
		tmp.Close()
		return err
	}
	if err := tmp.Sync(); err != nil { // flush to disk before swapping it in
		tmp.Close()
		return err
	}
	if err := tmp.Close(); err != nil {
		return err
	}
	if err := os.Chmod(tmpName, 0o644); err != nil { // CreateTemp makes 0o600
		return err
	}

	path := filepath.Join(dir, "config.toml")
	return os.Rename(tmpName, path) // atomic swap
}
