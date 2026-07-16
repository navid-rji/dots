package cli

import (
	"github.com/spf13/cobra"

	"github.com/navid-rji/dots/internal/buildinfo"
	"github.com/navid-rji/dots/internal/config"
)

var loadedConfig config.Config

var rootCmd = &cobra.Command{
	Use:           "dots [<app>]",
	Short:         "A CLI dotfile manager",
	Version:       buildinfo.Version(),
	SilenceErrors: true,
	PersistentPreRunE: func(cmd *cobra.Command, args []string) error {
		cmd.SilenceUsage = true // args/flags already validated; runtime errors past here shouldn't print usage
		completing := isCompletionRequest(cmd)

		cfg, found, err := config.Load()
		if err != nil {
			if completing {
				loadedConfig = config.Config{} // degrade to default, never fail a TAB
				return nil
			}
			return err
		}
		if !found && !completing {
			if cfg, err = firstRunSetup(); err != nil {
				return err
			}
		}
		loadedConfig = cfg
		return nil
	},
}

// Execute runs the root command. main.go calls this.
func Execute() error {
	return rootCmd.Execute()
}
