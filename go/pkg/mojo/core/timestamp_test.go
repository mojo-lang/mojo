package core

import (
	"testing"

	"github.com/json-iterator/go"
	"github.com/stretchr/testify/assert"
)

const (
	TimestampString1       = "2019-10-13T19:52:25.000+08:00"
	TimestampString2       = "2019-10-13T19:53:25.000+08:00"
	TimestampString1Json   = `"2019-10-13T19:52:25.000+08:00"`
	TimestampString1Number = "1570967545"
	Timestamp1             = int64(1570967545)
)

var exampleFormats = []string{
	`"2019-10-13T19:52:25+08:00"`,
	`"2019-10-13 19:52:25+0800"`,
	`"2019-10-13 19:52:25 08:00"`,
	`"2019-10-13 11:52:25"`,
	// `"2019/10/13 19:52:25 CST"`,
	`"2019/10/13 11:52:25"`,
	`"2019/10/13 11:52:25+0000"`, // UTC 时间，等同于北京时间 2019/10/13 19:52:25+0800
	`"2019/10/13 11:52:25+00:00"`,
	`"2019-10-13 11:52:25+0000"`,
	`"2019-10-13 11:52:25+00:00"`,
	`"2019-10-13 18:52:25 07:00"`, // test escaped '+'
	`"1570967545"`,
	`"1570967545000"`,
	`1570967545`,
}

func TestTimestampUnmarshal(t *testing.T) {
	var ts Timestamp
	jsoniter.Unmarshal([]byte(TimestampString1Json), &ts)

	assert.Equal(t, Timestamp1, ts.Seconds)
	assert.Equal(t, int32(0), ts.Nanoseconds)

	jsoniter.Unmarshal([]byte(TimestampString1Number), &ts)

	assert.Equal(t, Timestamp1, ts.Seconds)
	assert.Equal(t, int32(0), ts.Nanoseconds)
}

func TestTimestampMarshal(t *testing.T) {
	ts := &Timestamp{Seconds: 1570967545, Nanoseconds: 0}
	str, _ := jsoniter.ConfigDefault.MarshalToString(ts)
	assert.Equal(t, TimestampString1Json, str)
}

func TestTimestampParse(t *testing.T) {
	var ts Timestamp
	for i := range exampleFormats {
		ts.Reset()
		jsoniter.Unmarshal([]byte(exampleFormats[i]), &ts)
		assert.Equal(t, Timestamp1, ts.Seconds, "%s seconds not equal", exampleFormats[i])
		assert.Equal(t, int32(0), ts.Nanoseconds, "%s nanos not equal", exampleFormats[i])
	}
}

func TestTimestamp_Equal(t *testing.T) {
	t1, _ := ParseTimestamp(TimestampString1)
	t12, _ := ParseTimestamp(TimestampString1)
	t2, _ := ParseTimestamp(TimestampString2)

	assert.True(t, t1.Equal(t12))
	assert.False(t, t1.Equal(t2))
}

func TestTimestamp_Sub(t *testing.T) {
	t1, _ := ParseTimestamp(TimestampString1)
	t2, _ := ParseTimestamp(TimestampString2)

	assert.Equal(t, 1.0, t2.Sub(t1).ToMinutes())
	assert.Equal(t, -1.0, t1.Sub(t2).ToMinutes())
}

func TestTimestamp_After(t *testing.T) {
	t1, _ := ParseTimestamp(TimestampString1)
	t2, _ := ParseTimestamp(TimestampString2)

	assert.True(t, t2.After(t1))
}

func TestTimestamp_Before(t *testing.T) {
	t1, _ := ParseTimestamp(TimestampString1)
	t2, _ := ParseTimestamp(TimestampString2)

	assert.True(t, t1.Before(t2))
}
