package core

import (
	"testing"

	jsoniter "github.com/json-iterator/go"
	"github.com/stretchr/testify/assert"
)

func TestVersionStringCodec_Decode(t *testing.T) {
	version := &Version{}
	jsoniter.UnmarshalFromString(`"1.2.3"`, version)
	assert.Equal(t, uint64(1), version.Major)
}

func TestVersionStringCodec_Encode(t *testing.T) {
	version := &Version{
		Major: 1,
		Minor: 2,
		Patch: 3,
	}

	str, err := jsoniter.MarshalToString(version)
	assert.NoError(t, err)
	assert.Equal(t, `"1.2.3"`, str)
}

type VersionWrap struct {
	Version *Version `json:"version"`
}

func init() {
	RegisterJSONFieldEncoder("core.VersionWrap", "Version", &VersionStructCodec{IsFieldPointer: true})
	RegisterJSONFieldDecoder("core.VersionWrap", "Version", &VersionStructCodec{IsFieldPointer: true})
}

const versionWrapJson = `{"version":{"major":1,"minor":2,"patch":3,"level":3}}`

func TestVersionStructCodec_Encode(t *testing.T) {
	json, err := jsoniter.ConfigDefault.MarshalToString(&VersionWrap{Version: NewVersion(1, 2, 3)})
	assert.NoError(t, err)
	assert.Equal(t, versionWrapJson, json)
}

func TestVersionStructCodec_Decode(t *testing.T) {
	v := &VersionWrap{}
	err := jsoniter.ConfigDefault.UnmarshalFromString(versionWrapJson, v)
	assert.NoError(t, err)
	assert.Equal(t, uint64(1), v.Version.Major)
	assert.Equal(t, uint64(2), v.Version.Minor)
}
