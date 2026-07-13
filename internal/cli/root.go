package cli

import (
	"github.com/spf13/cobra"

	"github.com/navid-rji/dots/internal/buildinfo"
	"github.com/navid-rji/dots/internal/config"
)

var loadedConfig config.Config

var rootCmd = &cobra.Command{
	Use:     "dots",
	Short:   "A CLI dotfile manager",
	Version: buildinfo.Version(),
	PersistentPreRunE: func(cmd *cobra.Command, args []string) error {
		cfg, found, err := config.Load()
		if err != nil {
			return err
		}
		if !found {
			cfg, err = firstRunSetup()
			if err != nil {
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
