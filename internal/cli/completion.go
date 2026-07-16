package cli

import (
	"strings"

	"github.com/spf13/cobra"
)

// completeNames filters names by prefix and annotates each with its first
// configured path (zsh and fish render that as a description column).
// Reserved names are dropped
func completeNames(names []string, toComplete string) []cobra.Completion {
	reg := currentRegistry(loadedConfig)

	comps := make([]cobra.Completion, 0, len(names)+1)
	for _, name := range names {
		if !strings.HasPrefix(name, toComplete) || isReserved(name) {
			continue
		}

		desc := ""
		if app, err := reg.Resolve(name); err == nil && len(app.Paths) > 0 {
			desc = app.Paths[0]
		}
		comps = append(comps, cobra.CompletionWithDesc(name, desc))
	}
	return comps
}

// completeApps compeletes the <app> argument of `dots <app>`.
func completeApps(cmd *cobra.Command, args []string, toComplete string) ([]cobra.Completion, cobra.ShellCompDirective) {
	if len(args) != 0 {
		return nil, cobra.ShellCompDirectiveNoFileComp
	}
	comps := completeNames(currentRegistry(loadedConfig).Names(), toComplete)
	if strings.HasPrefix("dots", toComplete) {
		comps = append(comps, cobra.CompletionWithDesc("dots", "dots' own config"))
	}
	return comps, cobra.ShellCompDirectiveNoFileComp
}

func completeClear(cmd *cobra.Command, args []string, toComplete string) ([]cobra.Completion, cobra.ShellCompDirective) {
	if len(args) != 0 {
		return nil, cobra.ShellCompDirectiveNoFileComp
	}
	return completeNames(customNames(), toComplete), cobra.ShellCompDirectiveNoFileComp
}

func completeUpdate(cmd *cobra.Command, args []string, toComplete string) ([]cobra.Completion, cobra.ShellCompDirective) {
	switch len(args) {
	case 0:
		return completeNames(currentRegistry(loadedConfig).Names(), toComplete), cobra.ShellCompDirectiveNoFileComp
	case 1:
		return nil, cobra.ShellCompDirectiveDefault // shell does its normal path completion
	default:
		return nil, cobra.ShellCompDirectiveNoFileComp
	}
}

func completeAdd(cmd *cobra.Command, args []string, toComplete string) ([]cobra.Completion, cobra.ShellCompDirective) {
	switch len(args) {
	case 0:
		// add refuses anything that already resolves, so there is nothing to complete here
		return cobra.AppendActiveHelp(nil, "Type a name that isn't registered yet"), cobra.ShellCompDirectiveNoFileComp
	case 1:
		return nil, cobra.ShellCompDirectiveDefault // shell does its normal path completion
	default:
		return nil, cobra.ShellCompDirectiveNoFileComp
	}
}
