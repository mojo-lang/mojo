package geom

import (
    "github.com/stretchr/testify/assert"
    "testing"
)

func TestPoint_CellCovering(t *testing.T) {
    point := NewPointFrom(121.123, 24.234)
    cells := point.CellCovering(13, 13, 1)

    for _, cell := range cells {
        level := cell.Level()
        token := cell.ToToken()

        assert.Equal(t, 13, level)
        assert.Equal(t, "3468efd4", token)
    }
}

func TestPoint_GeoHashCovering(t *testing.T) {
    point := NewPointFrom(121.531974, 31.224385)
    hashes := point.GeoHashCovering(7, 7, 7)

    print(hashes)
}
