package cli

import (
	"errors"
	"fmt"

	"github.com/spf13/cobra"

	"github.com/navid-rji/dots/internal/config"
	"github.com/navid-rji/dots/internal/registry"
)

var addCmd = &cobra.Command{
	Use:     "add <app> <path>",
	Aliases: []string{"a"},
	Args:    cobra.ExactArgs(2),
	Short:   "Add a path to an app's config",
	Example: "  dots add nvim ~/.config/nvim",
	RunE: func(cmd *cobra.Command, args []string) error {
		// TODO: how to handle dots config
		reg := currentRegistry(loadedConfig)
		name := args[0]
		path := args[1]

		_, err := reg.Resolve(name)
		if err == nil {
			// app already resolves to something -> refuse
			// TODO: later add path to path list
			return fmt.Errorf("%q already exists - use `dots update` to change it", name)
		}
		if !errors.Is(err, registry.ErrUnknownApp) {
			return err // a real error, not just "not found"
		}

		if loadedConfig.Apps == nil {
			loadedConfig.Apps = make(map[string]config.App)
		}
		loadedConfig.Apps[name] = config.App{Paths: []string{path}}

		if err := config.Save(loadedConfig); err != nil {
			return err
		}

		fmt.Printf("added %s -> %s\n", name, path)

		// NOTE: could also check wether the path exists
		// and print a warnining if it doesn't.

		return nil
	},
}

func init() {
	rootCmd.AddCommand(addCmd)
}
