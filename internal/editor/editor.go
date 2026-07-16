package editor

import (
	"fmt"
	"os"
	"os/exec"
	"strings"
	"syscall"

	"github.com/google/shlex"
)

// Open launches the user's editor on the given file path.
// On success it never returns.
func Open(path, editor string) error {
	command := resolveCommand(editor)

	parts, err := shlex.Split(command)
	if err != nil {
		return fmt.Errorf("parsing editor command %q: %w", command, err)
	}
	if len(parts) == 0 {
		return fmt.Errorf("no editor configured")
	}

	parts = applyPath(parts, path)

	binary, err := exec.LookPath(parts[0])
	if err != nil {
		return fmt.Errorf("editor %q not found: %w", parts[0], err)
	}

	// Replace the current process image. Execution stops here on success.
	return syscall.Exec(binary, parts, os.Environ())
}

// resolveCommand picks the editor command from config, then the environment.
func resolveCommand(editor string) string {
	if editor != "" {
		return editor
	}
	if v := os.Getenv("VISUAL"); v != "" {
		return v
	}
	if e := os.Getenv("EDITOR"); e != "" {
		return e
	}
	return "vi"
}

// applyPath substitutes {} with the path, or appends it if no {} is present.
func applyPath(parts []string, path string) []string {
	replaced := false
	for i, p := range parts {
		if strings.Contains(p, "{}") {
			parts[i] = strings.ReplaceAll(p, "{}", path)
			replaced = true
		}
	}
	if !replaced {
		parts = append(parts, path)
	}
	return parts
}
