package db

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestJSONValuerNilZeroAndContent(t *testing.T) {
	for _, value := range []any{nil, (*string)(nil), 0, "", false, []string(nil)} {
		require.NotPanics(t, func() {
			got, err := (JSONValuer{}).Value(value)
			require.NoError(t, err)
			require.Nil(t, got)
		})
	}
	got, err := (JSONValuer{}).Value(map[string]string{"name": "Ada"})
	require.NoError(t, err)
	require.JSONEq(t, "{\"name\":\"Ada\"}", got.(string))
	_, err = (JSONValuer{}).Value(make(chan int))
	require.Error(t, err)
}

func TestJSONScannerSourcesAndErrors(t *testing.T) {
	for _, src := range []any{"{\"name\":\"Ada\"}", []byte("{\"name\":\"Ada\"}")} {
		var got map[string]string
		require.NoError(t, (JSONScanner{}).Scan(&got, src))
		require.Equal(t, map[string]string{"name": "Ada"}, got)
	}
	got := map[string]string{"keep": "value"}
	require.NoError(t, (JSONScanner{}).Scan(&got, nil))
	require.Equal(t, map[string]string{"keep": "value"}, got)
	require.NoError(t, (JSONScanner{}).Scan(nil, "{}"))
	require.NoError(t, (JSONScanner{}).Scan((*map[string]string)(nil), "{}"))
	require.Error(t, (JSONScanner{}).Scan(&got, 42))
	require.Error(t, (JSONScanner{}).Scan(&got, "{"))
}
