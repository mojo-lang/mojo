package geom

import (
    "github.com/golang/geo/s2"
    jsoniter "github.com/json-iterator/go"
    "github.com/stretchr/testify/assert"
    "strings"
    "testing"
)

const (
    tokensNoHole  = "35ad83095,35ad83097,35ad83099,35ad830a3,35ad830a5,35ad830ad,35ad830af,35ad830b1,35ad830b3,35ad830b5,35ad830b7,35ad830b9,35ad830bb,35ad830bd,35ad830bf,35ad830c1,35ad830c3,35ad830c5,35ad830c7,35ad830c9,35ad830cb,35ad830cd,35ad830cf,35ad830d1,35ad830d3,35ad830d5,35ad830d7,35ad830d9,35ad830db,35ad830dd,35ad830df,35ad8312b,35ad836d5,35ad836d7,35ad83729,35ad8372b,35ad8372d,35ad8372f,35ad83733,35ad83735,35ad8374b,35ad8374d"
    tokensHasHole = "35ad83095,35ad83097,35ad83099,35ad830a3,35ad830a5,35ad830ad,35ad830af,35ad830b1,35ad830b3,35ad830bb,35ad830bd,35ad830bf,35ad830c1,35ad830c3,35ad830c5,35ad830c7,35ad830cd,35ad830cf,35ad830d1,35ad830d3,35ad830d5,35ad830d7,35ad830d9,35ad830db,35ad830dd,35ad830df,35ad8312b,35ad836d5,35ad836d7,35ad83729,35ad8372b,35ad8372d,35ad8372f,35ad83733,35ad83735,35ad8374b,35ad8374d"
)

func TestPolygon_CellCovering(t *testing.T) {
    p := &Polygon{}
    err := jsoniter.UnmarshalFromString(polygonWithHole, p)
    assert.NoError(t, err)

    cells := p.CellCovering(16, 16, 100)
    assert.Equal(t, true, cellIdEqual(cells, tokensNoHole))
}

func TestPolygon_CellCoveringExceptHoles(t *testing.T) {
    p := &Polygon{}
    err := jsoniter.UnmarshalFromString(polygonWithHole, p)
    assert.NoError(t, err)

    cells := p.CellCoveringExceptHoles(16, 16, 200)
    assert.Equal(t, true, cellIdEqual(cells, tokensHasHole))
}

func TestPolygonRegion_ContainsCell(t *testing.T) {
    p := &Polygon{}
    err := jsoniter.UnmarshalFromString(polygonWithHole, p)
    assert.NoError(t, err)
    pr := newPolygonRegion(p)

    cellId := (&LngLat{Longitude: 121.768212, Latitude: 31.029051}).CellIDWithLevel(16)
    cell := s2.CellFromCellID(cellId)

    c := pr.ContainsCell(cell)
    i := pr.IntersectsCell(cell)
    assert.True(t, c || i)
}

func TestPolygon_GeoHashCovering(t *testing.T) {
}

func cellIdEqual(ids []s2.CellID, idStr string) bool {
    tokens := strings.Split(idStr, ",")
    tokenMap := make(map[string]bool, len(tokens))
    for _, token := range tokens {
        tokenMap[token] = true
    }

    for _, id := range ids {
        token := id.ToToken()
        if v, ok := tokenMap[token]; !ok || !v {
            return false
        }
        tokenMap[token] = false
    }
    return true
}
