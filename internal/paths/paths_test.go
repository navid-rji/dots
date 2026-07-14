package paths

import (
	"path/filepath"
	"testing"
)

func TestExpand(t *testing.T) {
	home := t.TempDir()
	t.Setenv("HOME", home)
	t.Setenv("DOTS_TEST_TOOLS", "/opt/tools")

	tests := []struct {
		name string
		in   string
		want string
	}{
		{
			name: "expands leading tilde slash to home",
			in:   "~/.config/nvim/init.lua",
			want: filepath.Join(home, ".config", "nvim", "init.lua"),
		},
		{
			name: "expands environment variable",
			in:   "$DOTS_TEST_TOOLS/bin",
			want: "/opt/tools/bin",
		},
		{
			name: "leaves plain absolute path unchanged",
			in:   "/etc/hosts",
			want: "/etc/hosts",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := Expand(tt.in)
			if err != nil {
				t.Fatalf("Expand(%q) returned error: %v", tt.in, err)
			}
			if got != tt.want {
				t.Errorf("Expand(%q) = %q, want %q", tt.in, got, tt.want)
			}
		})
	}
}

func TestCollapse(t *testing.T) {
	home := t.TempDir()
	t.Setenv("HOME", home)

	tests := []struct {
		name string
		in   string
		want string
	}{
		{
			name: "replaces home prefix with tilde",
			in:   filepath.Join(home, ".zshrc"),
			want: filepath.Join("~", ".zshrc"),
		},
		{
			name: "exact home becomes tilde",
			in:   home,
			want: "~",
		},
		{
			name: "path outside home stays absolute",
			in:   "/etc/hosts",
			want: "/etc/hosts",
		},
		{
			name: "sibling dir sharing home as string prefix is not collapsed",
			in:   home + "-backup/.zshrc",
			want: home + "-backup/.zshrc",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := Collapse(tt.in)
			if err != nil {
				t.Fatalf("Collapse(%q) returned error: %v", tt.in, err)
			}
			if got != tt.want {
				t.Errorf("Collapse(%q) = %q, want %q", tt.in, got, tt.want)
			}
		})
	}
}
