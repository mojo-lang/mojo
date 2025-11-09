package core

import (
	"fmt"
	"math"
	"strconv"
)

func (x *Color) Format() string {
	if x != nil {
		buf := make([]byte, 0, 10)
		buf = append(buf, '#')

		buf = strconv.AppendUint(buf, uint64(x.Red), 16)
		buf = strconv.AppendUint(buf, uint64(x.Green), 16)
		buf = strconv.AppendUint(buf, uint64(x.Blue), 16)

		if x.Alpha != nil {
			alpha := uint64(math.Round(float64(x.GetAlphaValue() * math.MaxUint8)))
			buf = strconv.AppendUint(buf, alpha, 16)
		}
		return string(buf)
	}
	return ""
}

func (x *Color) ToString() string {
	return x.Format()
}

func (x *Color) Parse(value string) error {
	if x != nil && len(value) > 0 {
		if value[0] == '#' && len(value) >= 7 {
			if red, err := strconv.ParseInt(value[1:3], 16, 32); err != nil {
				return fmt.Errorf("failed to parse the color string: %s, error: %w", value, err)
			} else {
				x.Red = uint32(red)
			}

			if green, err := strconv.ParseInt(value[3:5], 16, 32); err != nil {
				return fmt.Errorf("failed to parse the color string: %s, error: %w", value, err)
			} else {
				x.Green = uint32(green)
			}

			if blue, err := strconv.ParseInt(value[5:7], 16, 32); err != nil {
				return fmt.Errorf("failed to parse the color string: %s, error: %w", value, err)
			} else {
				x.Blue = uint32(blue)
			}

			if len(value) == 9 {
				if alpha, err := strconv.ParseInt(value[7:9], 16, 32); err != nil {
					return fmt.Errorf("failed to parse the color string: %s, error: %w", value, err)
				} else {
					x.SetAlphaValue(float32(RoundNaive(float64(alpha)/255, 0.01)))
				}
			}
		} else {
			return fmt.Errorf("failed to parse the color string: %s", value)
		}
	}
	return nil
}
