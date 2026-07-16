package cli

import (
	"fmt"

	"github.com/spf13/cobra"

	"github.com/navid-rji/dots/internal/config"
	"github.com/navid-rji/dots/internal/paths"
)

var updateCmd = &cobra.Command{
	Use:               "update <app> <path>",
	Args:              cobra.ExactArgs(2),
	ValidArgsFunction: completeUpdate,
	Short:             "Change an app's config path, overwriting any existing one",
	Example:           "  dots update hyprland ~/.config/hypr/other.conf",
	RunE: func(cmd *cobra.Command, args []string) error {
		// TODO: how to handle dots config update
		name := args[0]
		if isReserved(name) {
			return fmt.Errorf("%q is a reserved name; use `dots dots` to edit dots' own config", name)
		}

		path, err := paths.Collapse(args[1])
		if err != nil {
			return err
		}

		if loadedConfig.Apps == nil {
			loadedConfig.Apps = make(map[string]config.App)
		}

		_, existed := loadedConfig.Apps[name]
		loadedConfig.Apps[name] = config.App{Paths: []string{path}}

		if err := config.Save(loadedConfig); err != nil {
			return err
		}

		if existed {
			fmt.Printf("updated %s -> %s\n", name, path)
		} else {
			fmt.Printf("set %s -> %s\n", name, path)
		}

		return nil
	},
}

func init() {
	rootCmd.AddCommand(updateCmd)
}
