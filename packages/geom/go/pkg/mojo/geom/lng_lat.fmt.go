package geom

import (
    "fmt"
    jsoniter "github.com/json-iterator/go"
    "strings"
)

func ParseLngLat(lngLat string) (*LngLat, error) {
    ll := &LngLat{}
    err := ll.Parse(lngLat)
    if err != nil {
        return nil, err
    }
    return ll, nil
}

func (x *LngLat) Format() string {
    if x == nil || x.IsEmpty() {
        return ""
    }
    if x.Altitude != 0 {
        return fmt.Sprintf("%f,%f,%f", x.Longitude, x.Latitude, x.Altitude)
    } else {
        return fmt.Sprintf("%f,%f", x.Longitude, x.Latitude)
    }
}

func (x *LngLat) Parse(value string) error {
    if x == nil {
        return nil
    }

    value = strings.TrimSpace(value)
    if !strings.HasPrefix(value, "[") {
        value = "[" + value + "]"
    }
    var points []float64
    err := jsoniter.ConfigFastest.UnmarshalFromString(value, &points)
    if err != nil {
        return err
    }
    if len(points) < 2 {
        return fmt.Errorf("invalid lnglat string format: %s", value)
    }

    x.Longitude = points[0]
    x.Latitude = points[1]
    if len(points) > 2 {
        x.Altitude = points[2]
    }
    return nil
}
