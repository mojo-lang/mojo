package core

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestMediaType_IsSame(t *testing.T) {
	mt := &MediaType{
		Type:    "application",
		Subtype: "json",
	}

	assert.True(t, mt.IsSame(ApplicationJson))
}

func TestMediaTypeConstructorsAndParameters(t *testing.T) {
	for _, tt := range []struct {
		create func() *MediaType
		want   string
	}{
		{NewApplicationJson, ApplicationJson}, {NewApplicationOctetStream, ApplicationOctetStream}, {NewApplicationWwwFormUrlencoded, ApplicationWwwFormUrlencoded},
	} {
		mt := tt.create()
		require.NotNil(t, mt)
		assert.Equal(t, tt.want, mt.Format())
		assert.True(t, mt.IsSame(tt.want))
		assert.False(t, mt.IsSame("image/png"))
	}
	mt, err := NewMediaType("text/plain;charset=utf-8")
	require.NoError(t, err)
	assert.Equal(t, "utf-8", mt.Parameter.Value.GetString())
	require.NoError(t, mt.SetParameter("charset", "ascii"))
	assert.Equal(t, "text/plain;charset=ascii", mt.Format())
	require.NoError(t, mt.SetParameter("version", 3))
	assert.Equal(t, "version", mt.Parameter.Key)
	assert.Equal(t, int64(3), mt.Parameter.Value.GetInt64())
	previous := mt.Parameter
	assert.Error(t, mt.SetParameter("bad", make(chan int)))
	assert.Same(t, previous, mt.Parameter)
	_, err = NewMediaType("invalid")
	assert.Error(t, err)
	assert.False(t, (*MediaType)(nil).IsSame(ApplicationJson))
	assert.NoError(t, (*MediaType)(nil).SetParameter("x", "y"))
}
