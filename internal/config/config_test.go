package config

import (
	"path/filepath"
	"reflect"
	"testing"
)

func TestDir(t *testing.T) {
	tests := []struct {
		name    string
		dotsDir string
		xdg     string
		home    string
		want    string
	}{
		{
			name:    "DOTS_CONFIG_DIR wins over everything",
			dotsDir: "/explicit/dots",
			xdg:     "/xdg",
			home:    "/home/u",
			want:    "/explicit/dots",
		},
		{
			name: "XDG_CONFIG_HOME used when DOTS_CONFIG_DIR is unset",
			xdg:  "/xdg",
			home: "/home/u",
			want: filepath.Join("/xdg", "dots"),
		},
		{
			name: "falls back to ~/.config/dots",
			home: "/home/u",
			want: filepath.Join("/home/u", ".config", "dots"),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Setenv("DOTS_CONFIG_DIR", tt.dotsDir)
			t.Setenv("XDG_CONFIG_HOME", tt.xdg)
			t.Setenv("HOME", tt.home)
			got, err := Dir()
			if err != nil {
				t.Fatalf("Dir() returned error: %v", err)
			}
			if got != tt.want {
				t.Errorf("Dir() = %q, want %q", got, tt.want)
			}
		})
	}
}

func TestSaveLoadRoundTrip(t *testing.T) {
	// Point at a directory that does not exist yet so Save's MkdirAll is exercised too.
	dir := filepath.Join(t.TempDir(), "nested", "dots")
	t.Setenv("DOTS_CONFIG_DIR", dir)

	want := Config{
		Editor: "code --wait {}",
		Apps: map[string]App{
			"zsh":  {Paths: []string{"~/.zshrc", "~/.zprofile"}},
			"nvim": {Paths: []string{"~/.config/nvim/init.lua"}},
		},
	}

	if err := Save(want); err != nil {
		t.Fatalf("Save() returned error: %v", err)
	}

	got, found, err := Load()
	if err != nil {
		t.Fatalf("Load() returned error: %v", err)
	}
	if !found {
		t.Fatal("Load() found = false, want true after Save")
	}
	if !reflect.DeepEqual(got, want) {
		t.Errorf("Load() = %+v, want %+v", got, want)
	}

	t.Run("save over existing file replaces contents", func(t *testing.T) {
		want.Editor = "hx"
		if err := Save(want); err != nil {
			t.Fatalf("second Save() returned error: %v", err)
		}
		got, found, err := Load()
		if err != nil || !found {
			t.Fatalf("Load() after second Save = (found %v, err %v), want (true, nil)", found, err)
		}
		if !reflect.DeepEqual(got, want) {
			t.Errorf("Load() = %+v, want %+v", got, want)
		}
	})

	t.Run("no temp files left behind", func(t *testing.T) {
		leftovers, err := filepath.Glob(filepath.Join(dir, "*.tmp"))
		if err != nil {
			t.Fatalf("globbing for temp files: %v", err)
		}
		if len(leftovers) != 0 {
			t.Errorf("Save() left temp files behind: %v", leftovers)
		}
	})
}

func TestLoadMissingFile(t *testing.T) {
	t.Setenv("DOTS_CONFIG_DIR", t.TempDir())

	cfg, found, err := Load()
	if err != nil {
		t.Fatalf("Load() returned error: %v", err)
	}
	if found {
		t.Error("Load() found = true, want false for missing file")
	}
	if !reflect.DeepEqual(cfg, Config{}) {
		t.Errorf("Load() = %+v, want zero Config", cfg)
	}
}
