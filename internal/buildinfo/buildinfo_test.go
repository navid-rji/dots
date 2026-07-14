package buildinfo

import "testing"

func TestVersionNonEmpty(t *testing.T) {
	if v := Version(); v == "" {
		t.Error("Version() = \"\", want a non-empty string")
	}
}
