package cli

import (
	"fmt"
	"path/filepath"

	"github.com/spf13/cobra"

	"github.com/navid-rji/dots/internal/config"
	"github.com/navid-rji/dots/internal/editor"
	"github.com/navid-rji/dots/internal/paths"
)

var (
	openPrint, openDir bool
	openEditor         string
)

func init() {
	rootCmd.Args = cobra.MaximumNArgs(1)
	rootCmd.Flags().BoolVarP(&openPrint, "print", "p", false, "Print the resolved path instead of opening it")
	rootCmd.Flags().BoolVar(&openDir, "dir", false, "Open the parent directory of the resolved path instead of the file itself")
	rootCmd.Flags().StringVarP(&openEditor, "editor", "e", "", "Override the editor to use for opening files")
	rootCmd.RunE = func(cmd *cobra.Command, args []string) error {
		if len(args) == 0 {
			if openPrint {
				return fmt.Errorf("--print needs an app name")
			}
			return cmd.Help() // TODO: interactive fzf goes here
		}
		return openApp(args[0])
	}
}

func openApp(name string) error {
	path, err := resolvePath(name)
	if err != nil {
		return err
	}
	if openDir {
		path = filepath.Dir(path)
	}
	if openPrint {
		fmt.Println(path)
		return nil
	}
	editor_cmd := loadedConfig.Editor
	if openEditor != "" {
		editor_cmd = openEditor
	}
	return editor.Open(path, editor_cmd)
}

func resolvePath(name string) (string, error) {
	if name == "dots" {
		return config.Path()
	}

	app, err := currentRegistry(loadedConfig).Resolve(name)
	if err != nil {
		return "", err
	}
	if len(app.Paths) == 0 {
		return "", fmt.Errorf("%q has no configured paths", name)
	}

	// TODO: Handle multiple paths with an interactive picker
	return paths.Expand(app.Paths[0]) // NOTE: just the first path for now
}
