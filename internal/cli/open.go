package cli

import (
	"fmt"

	"github.com/spf13/cobra"

	"github.com/navid-rji/dots/internal/config"
	"github.com/navid-rji/dots/internal/editor"
	"github.com/navid-rji/dots/internal/paths"
)

func init() {
	rootCmd.Args = cobra.MaximumNArgs(1)
	rootCmd.RunE = func(cmd *cobra.Command, args []string) error {
		if len(args) == 0 {
			return cmd.Help() // TODO: TUI goes here
		}
		return openApp(args[0])
	}
}

func openApp(name string) error {
	reg := currentRegistry(loadedConfig)
	// `dots dots` opens dots' own config -> built-in shortcut, not a registered app.
	if name == "dots" {
		path, err := config.Path()
		if err != nil {
			return err
		}
		return editor.Open(loadedConfig, path)
	}

	app, err := reg.Resolve(name)
	if err != nil {
		return err
	}
	if len(app.Paths) == 0 {
		return fmt.Errorf("%q has no configured paths", name)
	}

	// TODO: Handle multiple paths with an interactive picker
	path, err := paths.Expand(app.Paths[0]) // NOTE: just the first path for now
	if err != nil {
		return err
	}
	return editor.Open(loadedConfig, path)
}
