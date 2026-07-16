package ui

import (
	"os"
	"testing"
)

func TestRender(t *testing.T) {
	on := Style{on: true}
	off := Style{on: false}

	if got, want := on.Render("x", Bold, Red), "\033[1;31mx\033[0m"; got != want {
		t.Errorf("Render = %q, want %q", got, want)
	}
	if got := off.Render("x", Bold, Red); got != "x" {
		t.Errorf("Render with color off = %q, want %q", got, "x")
	}
	if got := on.Render("x"); got != "x" {
		t.Errorf("Render with no attrs = %q, want %q", got, "x")
	}
	if got := on.Render("", Bold); got != "" {
		t.Errorf("Render of empty string = %q, want %q", got, "")
	}
}

func TestColorEnabled(t *testing.T) {
	tests := []struct {
		name string
		env  map[string]string
		f    *os.File
		want bool
	}{
		{"NO_COLOR wins over force", map[string]string{"NO_COLOR": "1", "CLICOLOR_FORCE": "1"}, nil, false},
		{"force without a tty", map[string]string{"CLICOLOR_FORCE": "1"}, nil, true},
		{"force=0 is not a force", map[string]string{"CLICOLOR_FORCE": "0"}, nil, false},
		{"dumb terminal", map[string]string{"TERM": "dumb"}, nil, false},
		{"not a tty", nil, nil, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			for _, k := range []string{"NO_COLOR", "CLICOLOR_FORCE", "TERM"} {
				t.Setenv(k, tt.env[k]) // t.Setenv restores after the test
			}
			if got := colorEnabled(tt.f); got != tt.want {
				t.Errorf("colorEnabled() = %v, want %v", got, tt.want)
			}
		})
	}
}
