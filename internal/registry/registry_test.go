package registry

import (
	"errors"
	"maps"
	"slices"
	"testing"

	"github.com/navid-rji/dots/internal/config"
)

func TestResolve(t *testing.T) {
	tests := []struct {
		name      string
		cfg       config.Config
		lookup    string
		wantPaths []string
		wantErr   error
	}{
		{
			name:      "default app resolves",
			cfg:       config.Config{},
			lookup:    "zsh",
			wantPaths: []string{"~/.zshrc"},
		},
		{
			name: "user override beats default",
			cfg: config.Config{Apps: map[string]config.App{
				"zsh": {Paths: []string{"~/custom/zshrc"}},
			}},
			lookup:    "zsh",
			wantPaths: []string{"~/custom/zshrc"},
		},
		{
			name: "user-only app resolves",
			cfg: config.Config{Apps: map[string]config.App{
				"alacritty": {Paths: []string{"~/.config/alacritty/alacritty.toml"}},
			}},
			lookup:    "alacritty",
			wantPaths: []string{"~/.config/alacritty/alacritty.toml"},
		},
		{
			name:    "unknown app returns ErrUnknownApp",
			cfg:     config.Config{},
			lookup:  "does-not-exist",
			wantErr: ErrUnknownApp,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			app, err := New(tt.cfg).Resolve(tt.lookup)
			if tt.wantErr != nil {
				if !errors.Is(err, tt.wantErr) {
					t.Fatalf("Resolve(%q) error = %v, want errors.Is(err, %v)", tt.lookup, err, tt.wantErr)
				}
				return
			}
			if err != nil {
				t.Fatalf("Resolve(%q) returned error: %v", tt.lookup, err)
			}
			if !slices.Equal(app.Paths, tt.wantPaths) {
				t.Errorf("Resolve(%q).Paths = %v, want %v", tt.lookup, app.Paths, tt.wantPaths)
			}
		})
	}
}

func TestNames(t *testing.T) {
	cfg := config.Config{Apps: map[string]config.App{
		"alacritty": {Paths: []string{"~/.config/alacritty/alacritty.toml"}},
		"zsh":       {Paths: []string{"~/custom/zshrc"}}, // override must not duplicate the default entry
	}}

	want := slices.Sorted(maps.Keys(defaults()))
	want = append(want, "alacritty")
	slices.Sort(want)

	got := New(cfg).Names()
	if !slices.Equal(got, want) {
		t.Errorf("Names() = %v, want sorted %v", got, want)
	}
}
