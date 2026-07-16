package ui

import (
	"os"
	"strings"

	"golang.org/x/term"
)

// Attr is an ANSI SGR attribute.
type Attr string

const (
	Bold  Attr = "1"
	Dim   Attr = "2"
	Red   Attr = "31"
	Green Attr = "32"
)

// Style renders text for one output stream. The zero value emits plain text.
type Style struct{ on bool }

// For returns the Style appropriate for f.
func For(f *os.File) Style { return Style{on: colorEnabled(f)} }

// Out and Err are the styles for the standard streams.
func Out() Style { return For(os.Stdout) }
func Err() Style { return For(os.Stderr) }

// Enabled reports whether s emits escape codes.
func (s Style) Enabled() bool { return s.on }

// Render wraps v in attrs, or returns it untouched when color is off.
func (s Style) Render(v string, attrs ...Attr) string {
	if !s.on || v == "" || len(attrs) == 0 {
		return v
	}
	codes := make([]string, len(attrs))
	for i, a := range attrs {
		codes[i] = string(a)
	}
	return "\033[" + strings.Join(codes, ";") + "m" + v + "\033[0m"
}

// colorEnabled applies the usual precedence: NO_COLOR wins over everything,
// then an explicit force, then the dumb-terminal and isatty checks.
func colorEnabled(f *os.File) bool {
	if os.Getenv("NO_COLOR") != "" {
		return false
	}
	if v := os.Getenv("CLICOLOR_FORCE"); v != "" && v != "0" {
		return true
	}
	if os.Getenv("TERM") == "dumb" {
		return false
	}
	if f == nil {
		return false
	}
	return term.IsTerminal(int(f.Fd()))
}
