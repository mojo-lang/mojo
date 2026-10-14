package http

import (
	"bytes"
	"compress/gzip"
	"encoding/json"
	"io"
	nethttp "net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestClientMethodsAndJSON(t *testing.T) {
	server := httptest.NewServer(nethttp.HandlerFunc(func(w nethttp.ResponseWriter, r *nethttp.Request) {
		body, _ := io.ReadAll(r.Body)
		_ = json.NewEncoder(w).Encode(map[string]string{"method": r.Method, "body": string(body), "type": r.Header.Get("Content-Type"), "token": r.Header.Get("X-Token")})
	}))
	defer server.Close()
	for _, method := range []Method{Method_METHOD_POST, Method_METHOD_PUT, Method_METHOD_GET} {
		t.Run(method.Format(), func(t *testing.T) {
			var result map[string]string
			err := Do(method, server.URL, NewHeaders().Set("X-Token", "secret"), map[string]string{"name": "Ada"}, &result)
			require.NoError(t, err)
			require.Equal(t, method.Format(), result["method"])
			require.JSONEq(t, "{\"name\":\"Ada\"}", result["body"])
			require.Equal(t, "application/json", result["type"])
			require.Equal(t, "secret", result["token"])
		})
	}
	for _, call := range []func(string, interface{}, interface{}) error{Post, Put} {
		var response map[string]string
		require.NoError(t, call(server.URL, map[string]string{"name": "Ada"}, &response))
		require.JSONEq(t, "{\"name\":\"Ada\"}", response["body"])
	}
	require.NoError(t, Delete(server.URL))
}

func TestClientReadsRawAndGzipResponses(t *testing.T) {
	for _, compressed := range []bool{false, true} {
		t.Run(map[bool]string{false: "plain", true: "gzip"}[compressed], func(t *testing.T) {
			payload := []byte("binary\x00payload")
			server := httptest.NewServer(nethttp.HandlerFunc(func(w nethttp.ResponseWriter, r *nethttp.Request) {
				if compressed {
					w.Header().Set("Content-Encoding", "gzip")
					gz := gzip.NewWriter(w)
					_, _ = gz.Write(payload)
					_ = gz.Close()
				} else {
					_, _ = w.Write(payload)
				}
			}))
			defer server.Close()
			var got []byte
			require.NoError(t, Do(Method_METHOD_GET, server.URL, NewHeaders().Set("Accept-Encoding", "gzip"), nil, &got))
			require.Equal(t, payload, got)
			got, err := GetContent(server.URL)
			require.NoError(t, err)
			require.Equal(t, payload, got)
		})
	}
}

func TestClientResponseFailures(t *testing.T) {
	for _, tt := range []struct {
		name, encoding, body string
		status               int
		wantError            bool
	}{
		{"bad gzip", "gzip", "invalid compressed data", 200, true},
		{"bad JSON", "", "{", 200, true},
		{"HTTP error", "", "{}", 503, true},
		{"empty", "", "", 204, false},
		{"valid JSON", "", "{\"ok\":true}", 200, false},
	} {
		t.Run(tt.name, func(t *testing.T) {
			server := httptest.NewServer(nethttp.HandlerFunc(func(w nethttp.ResponseWriter, r *nethttp.Request) {
				w.Header().Set("Content-Encoding", tt.encoding)
				w.WriteHeader(tt.status)
				_, _ = io.WriteString(w, tt.body)
			}))
			defer server.Close()
			var got map[string]bool
			require.NotPanics(t, func() {
				err := Do(Method_METHOD_GET, server.URL, NewHeaders().Set("Accept-Encoding", "gzip"), nil, &got)
				if tt.wantError {
					require.Error(t, err)
				} else {
					require.NoError(t, err)
				}
			})
			if tt.name == "valid JSON" {
				require.True(t, got["ok"])
			}
		})
	}
	var got any
	require.Error(t, Get("://invalid", &got))
	require.Error(t, Post("http://unused", make(chan int), &got))
	require.Equal(t, "short", logContent([]byte("short")))
	require.Equal(t, 1027, len(logContent(bytes.Repeat([]byte{'a'}, 1025))))
}
