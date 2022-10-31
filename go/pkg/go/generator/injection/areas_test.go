package injection

import (
	"go/token"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSort(t *testing.T) {
	areas := []Area{
		&TextArea{
			Start:   100,
			End:     105,
			Content: "testdata-2",
		},
		&TextArea{
			Start:   90,
			End:     95,
			Content: "testdata-1",
		},
	}

	Sort(areas)
	assert.Equal(t, token.Pos(90), areas[0].GetStart())
}
