package registry

import (
	"slices"
	"testing"
)

func TestSuggestFrom(t *testing.T) {
	pool := []string{"nvim", "nu", "npm", "zsh", "git", "ghostty", "list"}

	tests := []struct {
		name  string
		typed string
		want  []string
	}{
		// npm and nvim are both one edit from "nvm"; ties break alphabetically
		{"single typo", "nvm", []string{"npm", "nvim"}},
		{"transposition", "gti", []string{"git"}},
		{"prefix beats distance", "ghost", []string{"ghostty"}},
		{"command name", "lst", []string{"list"}},
		{"short word stays strict", "zsr", []string{"zsh"}},
		{"nothing close", "kubernetes", nil},
		{"empty", "", nil},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := SuggestFrom(pool, tt.typed)
			if !slices.Equal(got, tt.want) {
				t.Errorf("SuggestFrom(pool, %q) = %v, want %v", tt.typed, got, tt.want)
			}
		})
	}
}

func TestSuggestFromCaps(t *testing.T) {
	pool := []string{"nvim", "nvim1", "nvim2", "nvim3", "nvim4"}
	if got := SuggestFrom(pool, "nvim0"); len(got) != maxSuggestions {
		t.Errorf("SuggestFrom returned %d suggestions, want %d", len(got), maxSuggestions)
	}
}

func TestEditDistance(t *testing.T) {
	tests := []struct {
		a, b string
		want int
	}{
		{"nvim", "nvim", 0},
		{"nvm", "nvim", 1}, // insert
		{"nvim", "nvm", 1}, // delete
		{"zsr", "zsh", 1},  // substitute
		{"gti", "git", 1},  // transpose: one edit, not two
		{"ghostty", "git", 5},
		{"", "git", 3},
		{"git", "", 3},
		{"", "", 0},
		{"café", "cafe", 1}, // runes, not bytes
	}
	for _, tt := range tests {
		if got := editDistance(tt.a, tt.b); got != tt.want {
			t.Errorf("editDistance(%q, %q) = %d, want %d", tt.a, tt.b, got, tt.want)
		}
	}
}
