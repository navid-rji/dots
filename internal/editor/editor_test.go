package editor

import (
	"slices"
	"testing"
)

func TestApplyPath(t *testing.T) {
	tests := []struct {
		name  string
		parts []string
		path  string
		want  []string
	}{
		{
			name:  "substitutes {} placeholder",
			parts: []string{"code", "--wait", "{}"},
			path:  "/tmp/config.toml",
			want:  []string{"code", "--wait", "/tmp/config.toml"},
		},
		{
			name:  "substitutes {} inside a larger argument",
			parts: []string{"sh", "-c", "vim {} +1"},
			path:  "/tmp/config.toml",
			want:  []string{"sh", "-c", "vim /tmp/config.toml +1"},
		},
		{
			name:  "substitutes every {} occurrence",
			parts: []string{"diff", "{}", "{}.bak"},
			path:  "/tmp/config.toml",
			want:  []string{"diff", "/tmp/config.toml", "/tmp/config.toml.bak"},
		},
		{
			name:  "appends path when no {} is present",
			parts: []string{"vim", "-n"},
			path:  "/tmp/config.toml",
			want:  []string{"vim", "-n", "/tmp/config.toml"},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := applyPath(slices.Clone(tt.parts), tt.path)
			if !slices.Equal(got, tt.want) {
				t.Errorf("applyPath(%v, %q) = %v, want %v", tt.parts, tt.path, got, tt.want)
			}
		})
	}
}

func TestResolveCommand(t *testing.T) {
	tests := []struct {
		name      string
		cfgEditor string
		visual    string
		editor    string
		want      string
	}{
		{
			name:      "config editor wins over environment",
			cfgEditor: "hx",
			visual:    "code --wait",
			editor:    "nano",
			want:      "hx",
		},
		{
			name:   "empty config falls back to VISUAL",
			visual: "code --wait",
			editor: "nano",
			want:   "code --wait",
		},
		{
			name:   "unset VISUAL falls back to EDITOR",
			editor: "nano",
			want:   "nano",
		},
		{
			name: "nothing set defaults to vi",
			want: "vi",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Setenv("VISUAL", tt.visual)
			t.Setenv("EDITOR", tt.editor)
			got := resolveCommand(tt.cfgEditor)
			if got != tt.want {
				t.Errorf("resolveCommand() = %q, want %q", got, tt.want)
			}
		})
	}
}
