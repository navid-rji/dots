package cli

import (
	"errors"
	"fmt"
	"strings"

	"github.com/navid-rji/dots/internal/config"
	"github.com/navid-rji/dots/internal/registry"
	"github.com/navid-rji/dots/internal/ui"

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

// suggestionPool is everything a user could legitimately type in place of
// <app>: the registry, dots' own config, and the subcommands.
func suggestionPool() []string {
	names := currentRegistry(loadedConfig).Names()
	names = append(names, "dots")
	for _, c := range rootCmd.Commands() {
		if !c.IsAvailableCommand() {
			continue
		}
		names = append(names, c.Name())
		names = append(names, c.Aliases...)
	}
	return names
}

// unknownAppHint turns a registry miss into the message the user actually
// sees. Errors go to stderr, so it styles for stderr.
func unknownAppHint(name string) error {
	s := ui.Err()
	suggestions := registry.SuggestFrom(suggestionPool(), name)

	var b strings.Builder
	fmt.Fprintf(&b, "unknown app %q", name)

	switch len(suggestions) {
	case 0:
		fmt.Fprintf(&b, "\n\nRun %s to see the apps dots knows about, or %s to teach it this one.",
			s.Render("dots list", ui.Bold),
			s.Render("dots add "+name+" <path>", ui.Bold))
	case 1:
		fmt.Fprintf(&b, "\n\nDid you mean %s?", s.Render(suggestions[0], ui.Bold))
	default:
		b.WriteString("\n\nDid you mean one of these?")
		for _, sug := range suggestions {
			fmt.Fprintf(&b, "\n    %s", s.Render(sug, ui.Bold))
		}
	}
	return errors.New(b.String())
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
