package cli

import (
	"github.com/spf13/cobra"

	"github.com/navid-rji/dots/internal/buildinfo"
	"github.com/navid-rji/dots/internal/config"
)

var loadedConfig config.Config

var rootCmd = &cobra.Command{
	Use:     "dots [<app>]",
	Short:   "Open any app's config file in your editor",
	Version: buildinfo.Version(),
	Long: `dots opens an app's config file in your editor, so you don't have to
remember whether it lives in ~/.config, ~/.local/share, or straight in $HOME.

It ships with best-guess paths for 90+ well-known tools. Anything you register
yourself wins over those, and use_defaults = false drops the built-ins
entirely. dots' own config is TOML at ~/.config/dots/config.toml — "dots dots"
opens it.`,
	Example: `  # Open a config
  dots nvim
  dots zsh

  # Print the path instead of opening it
  dots -p nvim
  wc -l "$(dots -p git)"

  # Open the containing folder, or use a different editor just this once
  dots nvim --dir
  dots zsh -e code

  # Teach dots about a new app, then check what's actually on disk
  dots add hyprpaper ~/.config/hypr/hyprpaper.conf
  dots list --check`,
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
