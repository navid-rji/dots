package main

import (
	"fmt"
	"os"

	"github.com/navid-rji/dots/internal/cli"
)

func main() {
	if err := cli.Execute(); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}
