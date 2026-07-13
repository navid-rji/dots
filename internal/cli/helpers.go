package cli

import (
	"github.com/navid-rji/dots/internal/config"
	"github.com/navid-rji/dots/internal/registry"
)

func currentRegistry(cfg config.Config) *registry.Registry {
	return registry.New(cfg)
}

func isReserved(name string) bool {
	return name == "dots"
}
