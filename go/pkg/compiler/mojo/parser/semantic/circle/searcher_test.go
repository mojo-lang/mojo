package circle

import (
	"sort"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSearchCycle(t *testing.T) {
	files := map[string]*fileNode{
		"a": {name: "a", dependencies: []string{"c", "b"}},
		"b": {name: "b", dependencies: []string{"d", "e"}},
		"c": {name: "c", dependencies: []string{"e"}},
		"d": {name: "d", dependencies: []string{"e"}},
		"e": {name: "e", dependencies: []string{"d", "a"}},
		"f": {name: "f", dependencies: []string{"e"}},
	}

	circles := NewSearcher().Search(files)
	sort.Strings(circles[0])
	assert.ElementsMatch(t, []string{"a", "b", "c", "d", "e"}, circles[0])
}
