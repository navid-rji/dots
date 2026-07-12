package cli

import (
	"fmt"
	"sort"
	"strings"

	"github.com/spf13/cobra"
)

var listCustom bool

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
			fmt.Printf("%-12s %s\n", name, strings.Join(app.Paths, ", "))
		}
		return nil
	},
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
	listCmd.Flags().BoolVar(&listCustom, "custom", false, "Show only custom (uer-defined) apps, not built-in defaults")
	rootCmd.AddCommand(listCmd)
}
