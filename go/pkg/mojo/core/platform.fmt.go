package core

import (
	"fmt"
	"path"
	"regexp"
	"strings"
)

func ParsePlatform(value string) (*Platform, error) {
	p := &Platform{}
	if err := p.Parse(value); err != nil {
		return nil, err
	}
	return p, nil
}

func (x *Platform) Format() string {
	if x != nil && x.Os > 0 {
		firstPart := ""
		if x.Architecture > 0 {
			if len(x.Variant) > 0 {
				firstPart = path.Join(x.Os.Format(), x.Architecture.Format(), x.Variant)
			} else {
				firstPart = path.Join(x.Os.Format(), x.Architecture.Format())
			}
		} else {
			firstPart = x.Os.Format()
		}
		if len(x.OsName) > 0 {
			name := x.OsName
			if len(x.OsVersion) > 0 {
				name = path.Join(x.OsName, x.OsVersion)
			}
			return strings.Join([]string{firstPart, name}, "-")
		}
		return firstPart
	}
	return ""
}

func (x *Platform) ToString() string {
	return x.Format()
}

var variantFormat = regexp.MustCompile(`^[a-z\d]+$`)
var labelFormat = regexp.MustCompile(`^[a-zA-Z\d \._]+$`)

func (x *Platform) Parse(value string) error {
	if x != nil && len(value) > 0 {
		parts := strings.Split(value, "-")
		if len(parts) > 2 {
			return fmt.Errorf("invalid platfrom format: %s", value)
		}

		if len(parts) > 0 {
			segments := strings.Split(parts[0], "/")
			if len(segments) > 3 {
				return fmt.Errorf("invalid platfrom format: %s", value)
			}
			if len(segments) > 0 {
				if err := x.Os.Parse(segments[0]); err != nil {
					return fmt.Errorf("failed to parse platfrom in os part, error: %w", err)
				}
			}
			if len(segments) > 1 {
				if err := x.Architecture.Parse(segments[1]); err != nil {
					return fmt.Errorf("failed to parse platfrom in architecture part, error: %w", err)
				}
			}
			if len(segments) == 3 {
				variant := segments[2]
				if len(variant) > 0 {
					if variantFormat.MatchString(variant) {
						x.Variant = variant
					} else {
						return fmt.Errorf("invalid variant part in platform, variant: %s, should be `^[a-z0-9]+$`", variant)
					}
				}
			}
		}
		if len(parts) > 1 {
			segments := strings.Split(parts[1], "/")
			if len(segments) > 0 {
				name := segments[0]
				if labelFormat.MatchString(name) {
					x.OsName = name
				} else {
					return fmt.Errorf("invalid os name part in platform, os name: %s, should be `^[a-zA-Z0-9 \\._]+$`", name)
				}
			}
			if len(segments) > 1 {
				version := segments[1]
				if labelFormat.MatchString(version) {
					x.OsVersion = version
				} else {
					return fmt.Errorf("invalid os version part in platform, os version: %s, should be `^[a-zA-Z0-9 \\._]+$`", version)
				}
			}
		}
	}
	return nil
}
