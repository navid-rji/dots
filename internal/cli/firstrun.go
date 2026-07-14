package cli

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"golang.org/x/term"

	"github.com/navid-rji/dots/internal/config"
)

func firstRunSetup() (config.Config, error) {
	cfg := config.Config{} // the config we will build and return

	if !term.IsTerminal(int(os.Stdin.Fd())) {
		return cfg, nil // non-interactive: skip the prompt, proceed with defaults
	}

	fmt.Println("Welcome to dots! Let's do a quick one-time setup.")
	fmt.Println()
	fmt.Println("Which command should open your dotfiles?")
	fmt.Println("  Examples:   nvim   |   code   |   emacsclient -c {} -n")
	fmt.Println("  Put {} where the file path should go. If you leave {} out,")
	fmt.Println("  the path is appended at the end. Press Enter to use $EDITOR.")
	fmt.Print("> ")

	scanner := bufio.NewScanner(os.Stdin)

	// editor prompt
	scanner.Scan()                              // read one line
	editor := strings.TrimSpace(scanner.Text()) // Text() strips the newline
	cfg.Editor = editor

	// defaults prompt
	useDefaults := askYesNo(scanner, "Include the built-in app defaults? [Y/n] ", true)
	cfg.UseDefaults = &useDefaults

	// Save only the editor and the defaults choice - apps come from code
	// defaults, not this file.
	if err := config.Save(cfg); err != nil {
		return cfg, err
	}

	path, _ := config.Path()
	fmt.Printf("\nSaved to %s - you're all set.\n\n", path)
	return cfg, nil
}

func askYesNo(scanner *bufio.Scanner, question string, def bool) bool {
	for {
		fmt.Print(question)
		if !scanner.Scan() {
			return def // EOF (piped/closed stdin) -> accept the default
		}
		switch strings.ToLower(strings.TrimSpace(scanner.Text())) {
		case "":
			return def
		case "y", "yes":
			return true
		case "n", "no":
			return false
		default:
			fmt.Println("Please answer 'y' or 'n'.")
		}
	}
}
