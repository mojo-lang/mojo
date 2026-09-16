package compiler

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gorilla/mux"
	"github.com/stretchr/testify/require"
)

func TestCompilePathPatterns(t *testing.T) {
	for _, tc := range []struct {
		input, plain, route, valid, invalid string
		params                              map[string]bool
	}{
		{"/articles/{category}/{id:[0-9]+}", "/articles/{category}/{id}", "/articles/{category}/{id:[0-9]+}", "/articles/news/123", "/articles/news/abc", map[string]bool{"category": true, "id": true}},
		{`/items/{id.level:\d{2,4}}`, "/items/{id_level}", `/items/{id_level:\d{2,4}}`, "/items/123", "/items/1", map[string]bool{"id.level": true}},
		{"/files/{name:.+}", "/files/{name}", "/files/{name:.+}", "/files/a/b.txt", "/files/", map[string]bool{"name": true}},
		{"/files/{name:(?:images|docs)/.+}", "/files/{name}", "/files/{name:(?:images|docs)/.+}", "/files/docs/a.txt", "/files/other/a.txt", map[string]bool{"name": true}},
		{"/files/{name}", "/files/{name}", "/files/{name}", "/files/a.txt", "/files/a/b.txt", map[string]bool{"name": true}},
	} {
		t.Run(tc.input, func(t *testing.T) {
			plain, params := CompilePath(tc.input)
			require.Equal(t, tc.plain, plain)
			require.Equal(t, tc.params, params)
			route := CompileRoutePath(tc.input)
			require.Equal(t, tc.route, route)
			router := mux.NewRouter()
			require.NoError(t, router.Path(route).HandlerFunc(func(w http.ResponseWriter, r *http.Request) { w.WriteHeader(http.StatusNoContent) }).GetError())
			for _, c := range []struct {
				path   string
				status int
			}{{tc.valid, 204}, {tc.invalid, 404}} {
				w := httptest.NewRecorder()
				router.ServeHTTP(w, httptest.NewRequest("GET", c.path, nil))
				require.Equal(t, c.status, w.Code)
			}
		})
	}
}
