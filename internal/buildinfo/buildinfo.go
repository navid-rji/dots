package buildinfo

import "runtime/debug"

// version is injexted at build time via:
//
//	-ldflags "-X github.com/navid-rji/dots/internal/buildinfo.version=1.2.3"
var version string

// Version reports the build version, preferring an injected value,
// then Go's embedded module version, then a fallback.
func Version() string {
	if version != "" {
		return version // Homebrew / release builds
	}
	if info, ok := debug.ReadBuildInfo(); ok {
		if v := info.Main.Version; v != "" && v != "(devel)" {
			return v // go install ...@v1.2.3 embeds this automatically
		}
	}
	return "dev"
}
