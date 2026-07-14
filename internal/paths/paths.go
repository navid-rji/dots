package paths

import (
	"os"
	"path/filepath"
	"strings"
)

// Expand turns a leading ~ and any $VARS into real values
func Expand(p string) (string, error) {
	if strings.HasPrefix(p, "~/") {
		home, err := os.UserHomeDir()
		if err != nil {
			return "", err
		}
		p = filepath.Join(home, p[2:])
	}
	return os.ExpandEnv(p), nil
}

// Collapse replaces a leading home directory with ~ for portable storage.
func Collapse(p string) (string, error) {
	home, err := os.UserHomeDir()
	if err != nil {
		return "", err
	}
	if p == home {
		return "~", nil
	}
	if rest, ok := strings.CutPrefix(p, home+string(os.PathSeparator)); ok {
		return filepath.Join("~", rest), nil
	}
	return p, nil // not under home -> leave absolute paths alone
}
