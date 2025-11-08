package lang

import (
	"strings"

	"github.com/mojo-lang/mojo/packages/core/go/pkg/mojo/core"
)

func NewPackageRequirementVersion(str string) *Package_Requirement_Version {
	v := &Package_Requirement_Version{}
	v.Type = Package_Requirement_Version_TYPE_CARET

	if strings.Contains(str, "*") {
		v.Type = Package_Requirement_Version_TYPE_WILDCARD
		v.Range = parseWildcard(str)
	} else {
		switch str[0] {
		case '0', '1', '2', '3', '4', '5', '6', '7', '8', '9':
			v.Type = Package_Requirement_Version_TYPE_CARET
			v.Range = parseCaret(str)
		case '^':
			v.Type = Package_Requirement_Version_TYPE_CARET
			v.Range = parseCaret(str[1:])
		case '~':
			v.Type = Package_Requirement_Version_TYPE_TILDE
			v.Range = parseTilde(str[1:])
		case '>', '<', '=':
			v.Type = Package_Requirement_Version_TYPE_COMPARISON
			v.Range = parseComparison(str)
		}
	}

	return v
}

func firstNonZero(v *core.Version) int {
	if v.Major > 0 || v.Level == 1 {
		return 0
	} else if v.Minor > 0 || v.Level == 2 {
		return 1
	} else if v.Major > 0 {
		return 2
	}
	return 2
}

func parseWildcard(v string) *core.VersionRange {
	_ = v
	return nil
}

func parseCaret(v string) *core.VersionRange {
	vr := &core.VersionRange{
		EndIncluded: true,
	}

	version, _ := core.ParseVersion(v)

	vr.Start = version
	vr.End = &core.Version{
		Major: vr.Start.Major,
		Minor: vr.Start.Minor,
		Patch: vr.Start.Patch,
	}

	if i := firstNonZero(version); i >= 0 && i < 3 {
		values := []*uint64{&vr.End.Major, &vr.End.Minor, &vr.End.Patch}
		*values[i]++
		for j := i + 1; j < 3; j++ {
			*values[j] = 0
		}
		vr.EndIncluded = false
	}
	return vr
}

func parseTilde(v string) *core.VersionRange {
	_ = v
	return nil
}

func parseComparison(version string) *core.VersionRange {
	segments := strings.Split(version, ",")
	for _, segment := range segments {
		segment = strings.TrimSpace(segment)
		if strings.HasPrefix(segment, ">=") {
		} else if strings.HasPrefix(segment, ">") {
		} else if strings.HasPrefix(segment, "<=") {
		} else if strings.HasPrefix(segment, "<") {
		}
	}

	return nil
}
