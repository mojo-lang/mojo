package core

import (
	"testing"

	jsoniter "github.com/json-iterator/go"
	"github.com/stretchr/testify/assert"
)

var platforms = map[string]*Platform{
	`"linux/arm64"`: {
		Os:           OS_OS_LINUX,
		Architecture: Architecture_ARCHITECTURE_ARM64,
	},
}

func TestPlatformStringCodec_Decode(t *testing.T) {
	for k, p := range platforms {
		platform := &Platform{}
		err := jsoniter.UnmarshalFromString(k, platform)
		assert.NoError(t, err)
		assert.Equal(t, p.Os, platform.Os)
		assert.Equal(t, p.Architecture, platform.Architecture)
	}
}

type PlatformTest struct {
	Platform *Platform `json:"platform"`
	Number   int32     `json:"number"`
	Version  *Version  `json:"version"`
}

func TestPlatformStringCodec_Decode2(t *testing.T) {
	const pt = `{"platform": "x86/ubuntu 20.04", "number": 56, "version":"1.2.3"}`

	pT := &PlatformTest{}
	err := jsoniter.ConfigDefault.UnmarshalFromString(pt, pT)
	assert.Error(t, err)
}
