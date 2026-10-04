package core

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestDuration_FromSeconds(t *testing.T) {
	duration := &Duration{}
	duration.FromSeconds(1.001)
	assert.Equal(t, int64(1), duration.Seconds)
	assert.Equal(t, int32(1000000), duration.Nanoseconds)
}

func TestDuration_ToSeconds(t *testing.T) {
	duration := Duration{Nanoseconds: 1000, Seconds: 10}
	assert.Equal(t, 10.000001, duration.ToSeconds())
}

func TestDurationConversions(t *testing.T) {
	for _, tt := range []struct {
		name     string
		duration time.Duration
		seconds  int64
		nanos    int32
	}{
		{"zero", 0, 0, 0}, {"positive", 1500 * time.Millisecond, 1, 500000000},
		{"negative", -1500 * time.Millisecond, -1, -500000000}, {"nanosecond", -time.Nanosecond, 0, -1},
		{"hours", 2 * time.Hour, 7200, 0},
	} {
		t.Run(tt.name, func(t *testing.T) {
			d := FromDuration(tt.duration)
			assert.Equal(t, tt.seconds, d.Seconds)
			assert.Equal(t, tt.nanos, d.Nanoseconds)
			assert.Equal(t, tt.duration, d.ToDuration())
			assert.Equal(t, tt.duration.Hours(), d.ToHours())
			assert.Equal(t, tt.duration.Minutes(), d.ToMinutes())
			assert.Equal(t, tt.duration.Seconds(), d.ToSeconds())
			assert.Equal(t, tt.duration.Nanoseconds(), d.ToNanoseconds())
			fromSeconds := NewDurationFromSeconds(tt.duration.Seconds())
			assert.Equal(t, tt.duration, fromSeconds.ToDuration())
			var fromFloat Duration
			require.NoError(t, fromFloat.FromFloat64(tt.duration.Seconds()))
			assert.Equal(t, tt.duration, fromFloat.ToDuration())
		})
	}
	assert.Equal(t, 2*time.Hour, Hours(2).ToDuration())
	assert.Equal(t, 3*time.Minute, Minute(3).ToDuration())
	assert.Equal(t, -4*time.Second, Second(-4).ToDuration())
	var absent *Duration
	assert.Nil(t, absent.FromDuration(time.Second))
	assert.Nil(t, absent.FromSeconds(1))
	assert.NoError(t, absent.FromFloat64(1))
	assert.Zero(t, absent.ToHours())
	assert.Zero(t, absent.ToMinutes())
	assert.Zero(t, absent.ToSeconds())
	assert.Zero(t, absent.ToNanoseconds())
}

func TestDurationCompareBoundaries(t *testing.T) {
	for _, tt := range []struct {
		left, right *Duration
		want        int
	}{
		{nil, nil, 0}, {nil, Second(0), -1}, {Second(0), nil, 1},
		{Second(1), Second(1), 0}, {Second(-1), Second(1), -1}, {Second(2), Second(1), 1},
		{FromDuration(time.Second + time.Nanosecond), Second(1), 1},
		{Second(1), FromDuration(time.Second + time.Nanosecond), -1},
	} {
		assert.Equal(t, tt.want, tt.left.Compare(tt.right))
	}
}
