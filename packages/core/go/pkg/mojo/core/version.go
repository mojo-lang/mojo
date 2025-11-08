package core

import (
	"regexp"
)

const VersionTypeName = "Version"
const VersionTypeFullName = "mojo.core.Version"

// IsVersionTag
// v1
// v1alpha
// v1beta
func IsVersionTag(version string) bool {
	matched, _ := regexp.MatchString(`v[0-9]+[a-zA-z]*[0-9]*`, version)
	return matched
}

func NewVersion(major int, minor int, patch int) *Version {
	return &Version{
		Major: uint64(major),
		Minor: uint64(minor),
		Patch: uint64(patch),
		Level: 3,
	}
}

func (x *Version) IsEmpty() bool {
	return x == nil || (x.Major == 0 && x.Minor == 0 && x.Patch == 0)
}
