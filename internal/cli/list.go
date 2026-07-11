package cli

import (
	"fmt"
	"strings"

	"github.com/spf13/cobra"
)

var listCmd = &cobra.Command{
	Use:     "list",
	Aliases: []string{"ls"},
	Short:   "List known apps and their config paths",
	RunE: func(cmd *cobra.Command, args []string) error {
		reg := currentRegistry(loadedConfig)
		for _, name := range reg.Names() {
			app, _ := reg.Resolve(name)
			fmt.Printf("%-12s %s\n", name, strings.Join(app.Paths, ", "))
		}
		return nil
	},
}

func init() {
	rootCmd.AddCommand(listCmd)
}
