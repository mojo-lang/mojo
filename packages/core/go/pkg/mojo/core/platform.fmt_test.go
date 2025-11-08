package core

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

const platformStr = "linux/amd64-Ubuntu/14.02"

var platformStruct = &Platform{
	Architecture: Architecture_ARCHITECTURE_AMD64,
	Variant:      "",
	Os:           OS_OS_LINUX,
	OsName:       "Ubuntu",
	OsVersion:    "14.02",
}

func TestPlatform_Format(t *testing.T) {
	assert.Equal(t, platformStr, platformStruct.Format())
}

func TestPlatform_Format_1(t *testing.T) {
	assert.Equal(t, "linux", NewPlatform(OS_OS_LINUX, Architecture_ARCHITECTURE_UNSPECIFIED, "").Format())
}

func TestPlatform_Parse(t *testing.T) {
	platform := &Platform{}
	err := platform.Parse(platformStr)
	assert.NoError(t, err)
	assert.Equal(t, platformStruct.Architecture, platform.Architecture)
	assert.Equal(t, platformStruct.Os, platform.Os)
	assert.Equal(t, platformStruct.OsName, platform.OsName)
	assert.Equal(t, platformStruct.OsVersion, platform.OsVersion)
}

func TestParsePlatform(t *testing.T) {
	_, err := ParsePlatform("ubuntu/20.04/linux/amd64")
	assert.Error(t, err)
}

func TestParsePlatform_1(t *testing.T) {
	_, err := ParsePlatform("ubuntu/20.04-linux/amd64")
	assert.Error(t, err)

	_, err = ParsePlatform("linux/amd64/V8")
	assert.Error(t, err)

	_, err = ParsePlatform("linux/amd64-ubuntu*12")
	assert.Error(t, err)
}
