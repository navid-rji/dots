package registry

import (
	"errors"
	"fmt"
	"sort"

	"github.com/navid-rji/dots/internal/config"
)

// Registry is the merged view of defaults + the user's overrides.
type Registry struct {
	apps map[string]config.App
}

// New builds a Registry by layering the user's config over the defaults.
func New(cfg config.Config) *Registry {
	merged := defaults()
	for name, app := range cfg.Apps {
		merged[name] = app // user entry overrides the default
	}
	return &Registry{apps: merged}
}

// ErrUnknownApp is returned when an app isn't in the registry.
var ErrUnknownApp = errors.New("unknown app")

// Resolve looks up one app by name.
func (r *Registry) Resolve(name string) (config.App, error) {
	app, ok := r.apps[name]
	if !ok {
		return config.App{}, fmt.Errorf("%q: %w", name, ErrUnknownApp)
	}
	return app, nil
}

// Names returns the app names in stable, sorted order.
func (r *Registry) Names() []string {
	names := make([]string, 0, len(r.apps))
	for name := range r.apps {
		names = append(names, name)
	}
	sort.Strings(names)
	return names
}
