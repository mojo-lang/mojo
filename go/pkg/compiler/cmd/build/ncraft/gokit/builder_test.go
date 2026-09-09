package gokit

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestUnsupportedBuildType(t *testing.T) {
	for _, buildType := range []string{"sidecar", "ncraft.sidecar", "unknown"} {
		t.Run(buildType, func(t *testing.T) {
			// Reject removed or unknown types before reading the package or
			// falling through to service generation.
			b := Builder{Type: buildType}
			require.ErrorContains(t, b.Build(), "unsupported ncraft build type")
		})
	}
}
