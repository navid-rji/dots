package cli

import (
	"errors"
	"fmt"

	"github.com/spf13/cobra"

	"github.com/navid-rji/dots/internal/config"
	"github.com/navid-rji/dots/internal/registry"
)

var clearCmd = &cobra.Command{
	Use:     "clear <app>",
	Args:    cobra.ExactArgs(1),
	Short:   "Clear an app's custom path, reverting to its default if one exists",
	Example: "  dots clear ghostty",
	RunE: func(cmd *cobra.Command, args []string) error {
		// TODO: how to handle dots config update
		name := args[0]

		if isReserved(name) {
			return fmt.Errorf("%q is a reserved name and has no custom entry to clear", name)
		}

		if _, existed := loadedConfig.Apps[name]; !existed {
			return fmt.Errorf("no custom entry for %q", name)
		}

		delete(loadedConfig.Apps, name)
		if err := config.Save(loadedConfig); err != nil {
			return err
		}

		// Ask the registry what happens now: if the app still resolves, a
		// default surfaced; if not, it's gone entirely

		reg := currentRegistry(loadedConfig)
		app, err := reg.Resolve(name)
		if errors.Is(err, registry.ErrUnknownApp) {
			fmt.Printf("cleared %q\n", name)
			return nil
		}
		if err != nil {
			return err // a real error, not just "not found"
		}

		fmt.Printf("cleared %q, reverted to default (%s)\n", name, app.Paths[0])
		return nil
	},
}

func init() {
	rootCmd.AddCommand(clearCmd)
}
