package cli

import (
	"github.com/navid-rji/dots/internal/config"
	"github.com/navid-rji/dots/internal/registry"

	"github.com/spf13/cobra"
)

func currentRegistry(cfg config.Config) *registry.Registry {
	return registry.New(cfg)
}

func isReserved(name string) bool {
	if name == "dots" {
		return true
	}
	for _, c := range rootCmd.Commands() {
		if c.Name() == name || c.HasAlias(name) {
			return true
		}
	}
	return false
}

// isCompletionRequest reports whether cmd is part of the shell-completion
// machinery (`dots __complete ...` or `dots completion <shell>`). THese runs must
// never prompt or touch disk
func isCompletionRequest(cmd *cobra.Command) bool {
	for c := cmd; c != nil; c = c.Parent() {
		if c.Name() == cobra.ShellCompRequestCmd || c.Name() == "completion" {
			return true
		}
	}
	return false
}
