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
