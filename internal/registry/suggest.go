package registry

import (
	"cmp"
	"slices"
	"strings"
)

// maxSuggestions caps how many names an unknown-app error offers.
const maxSuggestions = 3

// Suggest returns registry names similar to name, closest first.
func (r *Registry) Suggest(name string) []string {
	return SuggestFrom(r.Names(), name)
}

// SuggestFrom ranks candidates against what the user typed, closest first. A
// candidate qualifies if it is within a small edit distance of name, or if
// name is a prefix of it. Callers that know about names outside the registry
// (dots' own config, the subcommands) pass their own candidate list.
func SuggestFrom(candidates []string, name string) []string {
	if name == "" {
		return nil
	}
	typed := strings.ToLower(name)

	type scored struct {
		name string
		dist int
	}
	hits := make([]scored, 0, maxSuggestions)

	for _, c := range candidates {
		cand := strings.ToLower(c)
		dist := editDistance(typed, cand)

		switch {
		case dist <= maxDistance(typed):
			// close enough on spelling alone
		case len(typed) >= 2 && strings.HasPrefix(cand, typed):
			dist = 0 // a typed prefix is a stronger signal than any edit
		default:
			continue
		}
		hits = append(hits, scored{name: c, dist: dist})
	}

	slices.SortFunc(hits, func(a, b scored) int {
		if d := cmp.Compare(a.dist, b.dist); d != 0 {
			return d
		}
		return strings.Compare(a.name, b.name) // stable, and matches `list` order
	})

	out := make([]string, 0, maxSuggestions)
	for _, h := range hits {
		if len(out) == maxSuggestions {
			break
		}
		out = append(out, h.name)
	}
	return out
}

// maxDistance scales the tolerated edit distance with the length of what was typed
func maxDistance(typed string) int {
	if len(typed) <= 3 {
		return 1
	}
	return 2
}

// editDistance returns the Damerau-Levenshtein distance (optimal string alignment)
// between a and b: insert, delete, substitute and transpose each cost one.
func editDistance(a, b string) int {
	ar, br := []rune(a), []rune(b)
	prev2 := make([]int, len(br)+1)
	prev := make([]int, len(br)+1)
	curr := make([]int, len(br)+1)

	for j := range prev {
		prev[j] = j
	}
	for i := 1; i <= len(ar); i++ {
		curr[0] = i
		for j := 1; j <= len(br); j++ {
			cost := 1
			if ar[i-1] == br[j-1] {
				cost = 0
			}
			curr[j] = min(prev[j]+1, curr[j-1]+1, prev[j-1]+cost)
			if i > 1 && j > 1 && ar[i-1] == br[j-2] && ar[i-2] == br[j-1] {
				curr[j] = min(curr[j], prev2[j-2]+1) // swapped pair: one edit
			}
		}
		prev2, prev, curr = prev, curr, prev2
	}
	return prev[len(br)]
}
