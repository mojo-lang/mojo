package http

import (
    "bytes"
    "fmt"
    "strconv"
    "strings"
)

const httpPrefix = "HTTP/"

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
        if strings.HasPrefix(version, httpPrefix) {
            version = strings.TrimPrefix(version, httpPrefix)
        }

        segments := strings.Split(version, ".")
        if len(segments) > 1 {
            if minor, err := strconv.Atoi(segments[1]); err != nil {
                return fmt.Errorf("failed to parse version in minor (%s) part, error: %w", segments[1], err)
            } else {
                x.Minor = int32(minor)
            }
        }
        if major, err := strconv.Atoi(segments[0]); err != nil {
            return fmt.Errorf("failed to parse version in major (%s) part, error: %w", segments[0], err)
        } else {
            x.Major = int32(major)
        }
    }
    return nil
}

func (x *Version) Format() string {
    if x != nil {
        buffer := bytes.Buffer{}
        buffer.WriteString(httpPrefix)
        buffer.WriteString(strconv.FormatInt(int64(x.Major), 10))

        if x.Major == 1 {
            buffer.WriteByte('.')
            buffer.WriteString(strconv.FormatInt(int64(x.Minor), 10))
        }

        return buffer.String()
    }
    return ""
}
