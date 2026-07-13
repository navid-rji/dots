package cli

import (
	"fmt"
	"os"
	"sort"

	"github.com/spf13/cobra"
	"golang.org/x/term"

	"github.com/navid-rji/dots/internal/paths"
)

var listCustom, listCheck bool

var listCmd = &cobra.Command{
	Use:     "list",
	Aliases: []string{"ls"},
	Short:   "List known apps and their config paths",
	RunE: func(cmd *cobra.Command, args []string) error {
		reg := currentRegistry(loadedConfig)

		names := reg.Names()
		if listCustom {
			names = customNames()
		}

		for _, name := range names {
			app, _ := reg.Resolve(name)
			if len(app.Paths) == 0 {
				fmt.Printf("%-12s\n", name)
				continue
			}
			for i, p := range app.Paths {
				label := name
				if i > 0 {
					label = "" // continuation line: blank the name column
				}
				if listCheck {
					fmt.Printf("%s %-12s %s\n", existMarker(p), label, p)
				} else {
					fmt.Printf("%-12s %s\n", label, p)
				}
			}
		}
		return nil
	},
}

// existMarker stats the expanded path and returns a ✓/✗ marker.
func existMarker(storedPath string) string {
	ok := false
	if p, err := paths.Expand(storedPath); err == nil {
		_, statErr := os.Stat(p)
		ok = statErr == nil
	}

	mark := "✗"
	if ok {
		mark = "✓"
	}
	if !useColor() {
		return mark
	}

	if ok {
		return "\033[32m" + mark + "\033[0m" // green
	}
	return "\033[31m" + mark + "\033[0m" // red
}

func useColor() bool {
	if os.Getenv("NO_COLOR") != "" {
		return false
	}
	return term.IsTerminal(int(os.Stdout.Fd()))
}

// customNames returns the user-overlaid app names, sorted.
func customNames() []string {
	names := make([]string, 0, len(loadedConfig.Apps))
	for name := range loadedConfig.Apps {
		names = append(names, name)
	}
	sort.Strings(names)
	return names
}

func init() {
	listCmd.Flags().BoolVar(&listCustom, "custom", false, "Show only custom (user-defined) apps, not built-in defaults")
	listCmd.Flags().BoolVar(&listCheck, "check", false, "Show whether each config file exists on disk")
	rootCmd.AddCommand(listCmd)
}
