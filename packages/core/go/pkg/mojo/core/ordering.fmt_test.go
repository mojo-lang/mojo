package core

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestOrdering_Format(t *testing.T) {
	type OrderingCase struct {
		Name    string
		AltName string
		Ordering
	}

	var cases = []OrderingCase{
		{"name", "name", Ordering{Orders: []*Ordering_Order{{Field: "name"}}}},
		{"name asc", "name   asc", Ordering{Orders: []*Ordering_Order{{Field: "name", Sort: Ordering_SORT_ASC}}}},
		{"name asc, count desc", "name asc   , count desc", Ordering{Orders: []*Ordering_Order{{Field: "name", Sort: Ordering_SORT_ASC}, {Field: "count", Sort: Ordering_SORT_DESC}}}},
		{"name, count desc", "name    , count desc", Ordering{Orders: []*Ordering_Order{{Field: "name"}, {Field: "count", Sort: Ordering_SORT_DESC}}}},
	}

	for _, c := range cases {
		o, err := ParseOrdering(c.AltName)
		assert.NoError(t, err)

		assert.Equal(t, len(c.Orders), len(o.Orders))
		for i := range o.Orders {
			assert.Equal(t, c.Orders[i].Field, o.Orders[i].Field)
			assert.Equal(t, c.Orders[i].Sort, o.Orders[i].Sort)
		}

		assert.Equal(t, c.Name, c.Format())
	}
}
