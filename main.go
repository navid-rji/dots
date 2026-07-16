package main

import (
	"fmt"
	"os"

	"github.com/navid-rji/dots/internal/cli"
	"github.com/navid-rji/dots/internal/ui"
)

func main() {
	if err := cli.Execute(); err != nil {
		s := ui.Err()
		fmt.Fprintf(os.Stderr, "%s %v\n", s.Render("error:", ui.Bold, ui.Red), err)
		os.Exit(1)
	}
}
