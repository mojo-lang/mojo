package geom

import (
    "github.com/golang/geo/s2"
    "github.com/mmcloughlin/geohash"
    "math"
)

// CellCovering Notice: Holes in polygon are covered too!
func (x *Polygon) CellCovering(minLevel, maxLevel, maxCells int) s2.CellUnion {
    rc := &s2.RegionCoverer{MinLevel: minLevel, MaxLevel: maxLevel, MaxCells: maxCells}

    var points []s2.Point
    for _, ll := range x.LineStrings[0].Coordinates {
        points = append(points, s2.PointFromLatLng(s2.LatLngFromDegrees(ll.Latitude, ll.Longitude)))
    }

    region := s2.LoopFromPoints(points)
    //region.Normalize()

    //n := region.IsNormalized()
    //s := region.Sign()
    //t := region.TurningAngle()

    if region.Area() > 2*math.Pi {
        region.Invert()
    }
    return rc.Covering(region)
}

// CellCoveringExceptHoles Notice: Holes in polygon are not covered!
func (x *Polygon) CellCoveringExceptHoles(minLevel, maxLevel, maxCells int) s2.CellUnion {
    rc := &s2.RegionCoverer{MinLevel: minLevel, MaxLevel: maxLevel, MaxCells: maxCells}
    pr := newPolygonRegion(x)
    return rc.Covering(pr)
}

func (x *Polygon) GeoHashCovering(minLevel, maxLevel, maxHashes int) []string {
    return x.GeoHashFractionCovering(minLevel, maxLevel, maxHashes, 0)
}

func (x *Polygon) GeoHashFractionCovering(minLevel, maxLevel, maxHashes int, fraction float32) []string {
    cells := x.CellCovering(18, 18, 10000)

    geohashes := make(map[string]int)
    for _, cell := range cells {
        center := CellCenter(cell)
        hash := geohash.EncodeWithPrecision(center.Latitude, center.Longitude, uint(maxLevel))
        geohashes[hash] += 1
    }

    hashes := make([]string, 0, len(geohashes))
    threshold := int(fraction * 16)
    for hash, v := range geohashes {
        if v > threshold {
            hashes = append(hashes, hash)
        }
    }

    // use the max hash for all skip
    if len(hashes) == 0 {
        maxV := 0
        maxHash := ""
        for hash, v := range geohashes {
            if v > maxV {
                maxV = v
                maxHash = hash
            }
        }
        hashes = append(hashes, maxHash)
    }

    return hashes
}

type polygonRegion struct {
    loops []*s2.Loop
}

func newPolygonRegion(p *Polygon) polygonRegion {
    pr := polygonRegion{}
    for _, lineString := range p.LineStrings {
        var points []s2.Point
        for _, ll := range lineString.Coordinates {
            points = append(points, s2.PointFromLatLng(s2.LatLngFromDegrees(ll.Latitude, ll.Longitude)))
        }

        loop := s2.LoopFromPoints(points)
        if loop.Area() > 2*math.Pi {
            loop.Invert()
        }
        pr.loops = append(pr.loops, loop)
    }
    return pr
}

// CapBound returns a bounding spherical cap. This is not guaranteed to be exact.
func (p polygonRegion) CapBound() s2.Cap {
    return p.loops[0].CapBound()
}

// RectBound returns a bounding latitude-longitude rectangle that contains
// the region. The bounds are not guaranteed to be tight.
func (p polygonRegion) RectBound() s2.Rect {
    return p.loops[0].RectBound()
}

// ContainsCell reports whether the region completely contains the given region.
func (p polygonRegion) ContainsCell(c s2.Cell) bool {
    if !p.loops[0].ContainsCell(c) {
        return false
    }

    for _, loop := range p.loops[1:] {
        if loop.ContainsCell(c) {
            return false
        } else if loop.IntersectsCell(c) {
            return false
        }
    }
    return true
}

// IntersectsCell reports whether the region intersects the given cell or
func (p polygonRegion) IntersectsCell(c s2.Cell) bool {
    if !p.loops[0].IntersectsCell(c) {
        return false
    }

    for _, loop := range p.loops[1:] {
        if loop.ContainsCell(c) {
            return false
        }
    }
    return true
}

// ContainsPoint reports whether the region contains the given point or not.
// The point should be unit length, although some implementations may relax
// this restriction.
func (p polygonRegion) ContainsPoint(point s2.Point) bool {
    if !p.loops[0].ContainsPoint(point) {
        return false
    }

    for _, loop := range p.loops[1:] {
        if loop.ContainsPoint(point) {
            return false
        }
    }
    return true
}

// CellUnionBound returns a small collection of CellIDs whose union covers
// the region. The cells are not sorted, may have redundancies (such as cells
// that contain other cells), and may cover much more area than necessary.
//
// This method is not intended for direct use by client code. Clients
// should typically use Covering, which has options to control the size and
// accuracy of the covering. Alternatively, if you want a fast covering and
// don't care about accuracy, consider calling FastCovering (which returns a
// cleaned-up version of the covering computed by this method).
//
// CellUnionBound implementations should attempt to return a small
// covering (ideally 4 cells or fewer) that covers the region and can be
// computed quickly. The result is used by RegionCoverer as a starting
// point for further refinement.
func (p polygonRegion) CellUnionBound() []s2.CellID {
    return p.loops[0].CellUnionBound()
}
