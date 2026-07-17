package cli

import (
	"errors"
	"fmt"
	"path/filepath"

	"github.com/spf13/cobra"

	"github.com/navid-rji/dots/internal/config"
	"github.com/navid-rji/dots/internal/editor"
	"github.com/navid-rji/dots/internal/paths"
	"github.com/navid-rji/dots/internal/picker"
	"github.com/navid-rji/dots/internal/registry"
	"github.com/navid-rji/dots/internal/ui"
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

	rootCmd.ValidArgsFunction = completeApps
	_ = rootCmd.RegisterFlagCompletionFunc("editor", cobra.NoFileCompletions)

	rootCmd.RunE = func(cmd *cobra.Command, args []string) error {
		if len(args) == 0 {
			if !picker.Available() {
				s := ui.Err()
				fmt.Fprintf(cmd.ErrOrStderr(), "%s\n\n %s\n\n", "The interactive picker needs fzf, which isn't on your PATH.", s.Render("brew install fzf", ui.Bold))
				return cmd.Help()
			}
			name, err := pickApp()
			if errors.Is(err, picker.ErrAborted) {
				return nil
			}
			if err != nil {
				return err
			}
			return openApp(name)
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
		if errors.Is(err, registry.ErrUnknownApp) {
			return "", unknownAppHint(name)
		}
		return "", err
	}
	if len(app.Paths) == 0 {
		return "", fmt.Errorf("%q has no configured paths", name)
	}

	// TODO: Handle multiple paths with an interactive picker
	return paths.Expand(app.Paths[0]) // NOTE: just the first path for now
}

func pickApp() (string, error) {
	reg := currentRegistry(loadedConfig)

	names := append(reg.Names(), "dots")
	lines := make([]string, 0, len(names))
	byLine := make(map[string]string, len(names))

	for _, name := range names {
		if isReserved(name) && name != "dots" {
			continue
		}
		desc := ""
		if name == "dots" {
			desc = "dots' own config"
		} else if app, err := reg.Resolve(name); err == nil && len(app.Paths) > 0 {
			desc = app.Paths[0]
		}
		line := fmt.Sprintf("%-14s %s", name, desc)
		lines = append(lines, line)
		byLine[line] = name
	}

	line, err := picker.Pick(lines, "--style=full", "--no-multi", "--layout=reverse", "--height=~40%", "--highlight-line", "--border", "--prompt=dots > ", "--input-label= search ", "--list-label= apps ")
	if err != nil {
		return "", err
	}
	return byLine[line], nil
}
