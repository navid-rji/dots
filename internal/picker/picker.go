package picker

import (
	"errors"
	"fmt"
	"os"
	"os/exec"
	"strings"
)

// ErrUnavailable reports that fzf is not on PATH.
var ErrUnavailable = errors.New("fzf not found on PATH")

// ErrAborted reports thta the user dismissed the picker, or nothing matched.
var ErrAborted = errors.New("no selection")

// Available reports whether fzf can be run.
func Available() bool {
	_, err := exec.LookPath("fzf")
	return err == nil
}

// Pick runs fzf over lines and return the chosen line verbatim.
func Pick(lines []string, opts ...string) (string, error) {
	bin, err := exec.LookPath("fzf")
	if err != nil {
		return "", ErrUnavailable
	}

	cmd := exec.Command(bin, opts...)
	cmd.Stdin = strings.NewReader(strings.Join(lines, "\n"))
	cmd.Stderr = os.Stderr // fzf draws its UI on /dev/tty; this only carries its error messages

	out, err := cmd.Output()
	if err != nil {
		var exitErr *exec.ExitError
		if errors.As(err, &exitErr) {
			switch exitErr.ExitCode() {
			case 1, 130: // no match / Ctr-C or Esc
				return "", ErrAborted
			}
		}
		return "", fmt.Errorf("running fzf: %w", err)
	}
	return strings.TrimRight(string(out), "\r\n"), nil
}
