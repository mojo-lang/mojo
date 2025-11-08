package core

import (
	"bytes"
	"fmt"
	"strconv"
	"strings"
)

func ParseVersion(version string) (*Version, error) {
	v := &Version{}
	err := v.Parse(version)
	if err != nil {
		return nil, err
	}
	return v, nil
}

func (x *Version) Parse(version string) error {
	if x != nil && len(version) > 0 {
		// if v or V prefixed, then remove the v tag.
		if version[0] == 'v' || version[0] == 'V' {
			version = version[1:]
		}
		if len(version) == 0 {
			return nil
		}

		segments := strings.Split(version, "+")
		if len(segments) > 1 {
			x.Builds = strings.Split(segments[1], ".")
		}

		segments = strings.Split(segments[0], "-")
		if len(segments) > 1 {
			segments[1] = strings.Join(segments[1:], "-")
			segments = segments[:2]
			x.PreReleases = strings.Split(segments[1], ".")
		}

		segments = strings.Split(segments[0], ".")
		x.Level = int32(len(segments))
		if x.Level > 0 {
			major, err := strconv.Atoi(segments[0])
			if err != nil {
				return fmt.Errorf("failed to parse version in major (%s) part, error: %w", segments[0], err)
			} else {
				x.Major = uint64(major)
			}
		}
		if x.Level > 1 {
			minor, err := strconv.Atoi(segments[1])
			if err != nil {
				return fmt.Errorf("failed to parse version in minor (%s) part, error: %w", segments[1], err)
			} else {
				x.Minor = uint64(minor)
			}
		}
		if x.Level > 2 {
			patch, err := strconv.Atoi(segments[2])
			if err != nil {
				return fmt.Errorf("failed to parse version in patch (%s) part, error: %w", segments[2], err)
			} else {
				x.Patch = uint64(patch)
			}
		}
	}

	return nil
}

func (x *Version) Format() string {
	if x != nil {
		buffer := bytes.Buffer{}
		buffer.WriteString(strconv.FormatUint(x.Major, 10))

		if x.Level == 0 || x.Level > 1 {
			buffer.WriteByte('.')
			buffer.WriteString(strconv.FormatUint(x.Minor, 10))
		}
		if x.Level == 0 || x.Level > 2 {
			buffer.WriteByte('.')
			buffer.WriteString(strconv.FormatUint(x.Patch, 10))
		}

		if len(x.PreReleases) > 0 {
			buffer.WriteByte('-')
			buffer.WriteString(strings.Join(x.PreReleases, "."))
		}

		if len(x.Builds) > 0 {
			buffer.WriteByte('+')
			buffer.WriteString(strings.Join(x.PreReleases, "."))
		}

		return buffer.String()
	}
	return ""
}

func (x *Version) ToString() string {
	return x.Format()
}
