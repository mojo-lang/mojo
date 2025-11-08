package geom

import (
	"bufio"
	"bytes"
	"fmt"
	"io"
	"strconv"
	"strings"
)

const Empty = "EMPTY"

func NewGeometryFromWKT(wkt string) (*Geometry, error) {
	d := &WKT{}
	return d.Decode(wkt)
}

func (x *Geometry) FromWKT(wkt string) error {
	if x != nil {
		g, err := NewGeometryFromWKT(wkt)
		if err != nil {
			return err
		}

		x.Geometry = g.Geometry
	}
	return nil
}

func (x *Geometry) ToWKT() string {
	enc := &WKT{}
	return enc.Encode(x)
}

type WKT struct {
}

func (w *WKT) Decode(wkt string) (*Geometry, error) {
	return w.decode(&reader{
		Reader: bufio.NewReader(strings.NewReader(wkt)),
	})
}

func (w *WKT) Encode(geometry *Geometry) string {
	buffer := bytes.NewBuffer(nil)
	return w.encode(buffer, geometry).string(buffer)
}

func (w *WKT) decodePoint(r *reader) (*Geometry, error) {
	pts, err := r.readPoints()
	if err != nil {
		return nil, err
	}

	switch len(pts) {
	case 0:
		return nil, r.newError("POINT", "cannot be empty")
	case 1:
		p := &Geometry{}
		_ = p.SetValue(&Point{Coordinate: pts[0]})
		return p, nil
	default:
		return nil, r.newError("POINT", "too many points %d", len(pts))
	}
}

func (w *WKT) decodeMultiPoint(r *reader) (*Geometry, error) {
	pts, err := r.readPoints()
	if err != nil {
		return nil, err
	}

	return NewMultiPointGeometryFrom(pts...), nil
}

func (w *WKT) decodeLineString(r *reader) (*Geometry, error) {
	pts, err := r.readPoints()
	if err != nil {
		return nil, err
	}

	if len(pts) < 2 {
		return nil, r.newError("LINESTRING", "not enough points %d", len(pts))
	}

	return NewLineStringGeometry(pts...), nil
}

func (w *WKT) decodeMultiLineString(r *reader) (*Geometry, error) {
	lines, err := r.readLines()
	if err != nil {
		return nil, err
	}

	if len(lines) < 1 {
		return nil, r.newError("MULTILINESTRING", "not enough lines %d", len(lines))
	}

	lineStrings := make([]*LineString, 0, len(lines))
	for i, v := range lines {
		if len(v) < 2 {
			return nil, r.newError("MULTILINESTRING", "not enough points in LINESTRING[%d], %d", i, len(v))
		}
		lineStrings = append(lineStrings, &LineString{Coordinates: v})
	}

	return NewMultiLineStringGeometry(lineStrings...), nil
}

func (w *WKT) decodePolygon(r *reader) (*Geometry, error) {
	lines, err := r.readLines()
	if err != nil {
		return nil, err
	}

	if len(lines) < 1 {
		return nil, r.newError("POLYGON", "not enough lines %d", len(lines))
	}

	lineStrings := make([]*LineString, 0, len(lines))
	for i, v := range lines {
		n := len(v)
		if n < 4 {
			return nil, r.newError("POLYGON", "not enough points in linear-ring[%d], %d", i, len(v))
		}

		// part of the spec
		if !v[0].Equal(v[n-1]) {
			return nil, r.newError("POLYGON", "linear-ring[%d] not closed", i)
		}
		lineStrings = append(lineStrings, &LineString{Coordinates: v})
	}

	return NewPolygonGeometry(lineStrings...), nil
}

func (w *WKT) decodeMultiPolygon(r *reader) (*Geometry, error) {
	polys, err := r.readPolys()
	if err != nil {
		return nil, err
	}

	if len(polys) < 1 {
		return nil, r.newError("MULTIPOLYGON", "not enough polygons %d", len(polys))
	}

	polygons := make([]*Polygon, 0, len(polys))
	for ii, vv := range polys {
		rings := make([]*LineString, 0, len(vv))
		for i, v := range vv {
			n := len(v)
			if n < 4 {
				return nil, r.newError("MULTIPOLYGON", "not enough points in polygon[%d] linear-ring[%d], %d", ii, i, len(v))
			}

			// part of the spec
			if !v[0].Equal(v[n-1]) {
				return nil, r.newError("MULTIPOLYGON", "polygon[%d] linear-ring[%v] not closed", i, ii)
			}
			rings = append(rings, &LineString{Coordinates: v})
		}
		polygons = append(polygons, &Polygon{LineStrings: rings})
	}

	return NewMultiPolygonGeometry(polygons...), err
}

func (w *WKT) decodeGeometryCollection(r *reader) (*Geometry, error) {
	b, err := r.pop()
	if err != nil {
		return nil, err
	}
	if b != '(' {
		return nil, r.expect("(")
	}
	_, err = r.skipSpaces()
	if err != nil {
		return nil, err
	}

	gs := make([]*Geometry, 0, 4)
	var geo *Geometry
	for b, err = r.pop(); b != ')' && err == nil; b, err = r.pop() {
		r.unPop()

		geo, err = w.decode(r)
		if err != nil {
			return nil, err
		}
		gs = append(gs, geo)

		_, err = r.skipSpaces()
		if err != nil {
			return nil, err
		}

		b, err = r.pop()
		if err != nil {
			return nil, err
		}
		switch b {
		case ')':
			r.unPop()
		case ',':
			//noop
			_, err = r.skipSpaces()
			if err != nil {
				return nil, err
			}
		default:
			return nil, r.expect("',' or ')'")
		}
	}

	if err != nil {
		return nil, err
	}

	if len(gs) < 1 {
		return nil, r.newError("GEOMETRYCOLLECTION", "not enough geometries %d", len(gs))
	}

	if b != ')' {
		panic("unreachable")
	}

	return NewGeometryCollectionGeometry(gs...), nil
}

func (w *WKT) decode(r *reader) (*Geometry, error) {
	tag, err := r.readTag() // remove leading whitespaces
	if err != nil {
		return nil, err
	}

	_, err = r.skipSpaces()
	if err != nil {
		return nil, err
	}

	switch strings.ToLower(tag) {
	case "point":
		return w.decodePoint(r)
	case "multipoint":
		return w.decodeMultiPoint(r)
	case "linestring":
		return w.decodeLineString(r)
	case "multilinestring":
		return w.decodeMultiLineString(r)
	case "polygon":
		return w.decodePolygon(r)
	case "multipolygon":
		return w.decodeMultiPolygon(r)
	case "geometrycollection":
		return w.decodeGeometryCollection(r)
	default:
		return nil, r.newError("GEOMETRY", "unknown type %q", tag)
	}
}

func (w *WKT) writeLngLat(buffer *bytes.Buffer, p *LngLat) *WKT {
	buffer.Grow(24)

	buffer.WriteString(strconv.FormatFloat(p.Longitude, 'f', -1, 64))
	buffer.WriteByte(' ')
	buffer.WriteString(strconv.FormatFloat(p.Latitude, 'f', -1, 64))

	return w
}

func (w *WKT) writeLngLats(buffer *bytes.Buffer, points ...*LngLat) *WKT {
	n := len(points)
	buffer.Grow(n * 24)

	buffer.WriteByte('(')
	for i, p := range points {
		if i > 0 {
			buffer.WriteByte(',')
		}
		w.writeLngLat(buffer, p)
	}
	buffer.WriteByte(')')

	return w
}

func (w *WKT) writeLines(buffer *bytes.Buffer, lines ...*LineString) *WKT {
	n := len(lines)
	buffer.Grow(n * 64)

	buffer.WriteByte('(')
	for i, line := range lines {
		if i > 0 {
			buffer.WriteByte(',')
		}
		w.writeLngLats(buffer, line.Coordinates...)
	}
	buffer.WriteByte(')')

	return w
}

func (w *WKT) encodePoint(buffer *bytes.Buffer, p *Point) *WKT {
	buffer.Grow(32)
	buffer.WriteString("POINT ")
	if p.GetCoordinate().IsEmpty() {
		buffer.WriteString(Empty)
		return w
	}
	buffer.WriteByte('(')
	w.writeLngLat(buffer, p.Coordinate)
	buffer.WriteByte(')')

	return w
}

func (w *WKT) encodeMultiPoint(buffer *bytes.Buffer, multiPoint *MultiPoint) *WKT {
	n := len(multiPoint.Points)
	buffer.Grow(n * 24)

	buffer.WriteString("MULTIPOINT ")
	if n == 0 {
		buffer.WriteString(Empty)
	} else {
		w.writeLngLats(buffer, multiPoint.GetLngLats()...)
	}

	return w
}

func (w *WKT) encodeLineString(buffer *bytes.Buffer, line *LineString) *WKT {
	n := len(line.Coordinates)

	buffer.Grow(n * 24)
	buffer.WriteString("LINESTRING ")
	if n == 0 {
		buffer.WriteString(Empty)
	} else {
		w.writeLngLats(buffer, line.Coordinates...)
	}

	return w
}

func (w *WKT) encodeMultiLineString(buffer *bytes.Buffer, multiLineString *MultiLineString) *WKT {
	n := len(multiLineString.LineStrings)
	buffer.Grow(n * 64)

	buffer.WriteString("MULTILINESTRING ")
	if n == 0 {
		buffer.WriteString(Empty)
	} else {
		w.writeLines(buffer, multiLineString.LineStrings...)
	}

	return w
}

func (w *WKT) encodePolygon(buffer *bytes.Buffer, polygon *Polygon) *WKT {
	n := len(polygon.LineStrings)
	buffer.Grow(n * 64)

	buffer.WriteString("POLYGON ")
	if n == 0 {
		buffer.WriteString(Empty)
	} else {
		w.writeLines(buffer, polygon.LineStrings...)
	}

	return w
}

func (w *WKT) encodeMultiPolygon(buffer *bytes.Buffer, multiPolygon *MultiPolygon) *WKT {
	n := len(multiPolygon.Polygons)
	buffer.Grow(n * 256)

	buffer.WriteString("MULTIPOLYGON ")
	buffer.WriteByte('(')
	for i, polygon := range multiPolygon.Polygons {
		if i > 0 {
			buffer.WriteByte(',')
		}
		w.writeLines(buffer, polygon.LineStrings...)
	}
	buffer.WriteByte(')')

	return w
}

func (w *WKT) encodeGeometryCollection(buffer *bytes.Buffer, collection *GeometryCollection) *WKT {
	gs := collection.Geometries
	n := len(gs)
	if n == 0 {
		buffer.WriteString("GEOMETRYCOLLECTION EMPTY")
		return w
	}

	buffer.Grow(n * 256)

	buffer.WriteString("GEOMETRYCOLLECTION (")

	for i, g := range gs {
		if i > 0 {
			buffer.WriteByte(',')
		}
		w.encode(buffer, g)
	}

	buffer.WriteByte(')')

	return w
}

func (w *WKT) encode(buffer *bytes.Buffer, geometry *Geometry) *WKT {
	switch gt := geometry.Geometry.(type) {
	case *Geometry_Point:
		return w.encodePoint(buffer, gt.Point)
	case *Geometry_MultiPoint:
		return w.encodeMultiPoint(buffer, gt.MultiPoint)
	case *Geometry_LineString:
		return w.encodeLineString(buffer, gt.LineString)
	case *Geometry_MultiLineString:
		return w.encodeMultiLineString(buffer, gt.MultiLineString)
	case *Geometry_Polygon:
		return w.encodePolygon(buffer, gt.Polygon)
	case *Geometry_MultiPolygon:
		return w.encodeMultiPolygon(buffer, gt.MultiPolygon)
	case *Geometry_GeometryCollection:
		return w.encodeGeometryCollection(buffer, gt.GeometryCollection)
	}
	return w
}

func (w *WKT) string(buffer *bytes.Buffer) string {
	return buffer.String()
}

type reader struct {
	*bufio.Reader
	row    int
	column int
}

func (r *reader) peek() (byte, error) {
	bts, err := r.Peek(1)
	if err != nil {
		return ' ', err
	}
	return bts[0], nil
}

func (r *reader) pop() (byte, error) {
	b, err := r.ReadByte()
	if err == io.EOF {
		return b, r.newError("UNEXPECTED", "eof")
	}

	if b == '\n' || b == '\r' {
		r.row++
		r.column = 0
	} else {
		r.column++
	}

	return b, err
}

func (r *reader) unPop() error {
	r.column--
	if r.column < 0 {
		r.column = 0
		r.row--
	}

	return r.UnreadByte()
}

// skipSpaces skips whitespaces and returns
// true if any characters were read
func (r *reader) skipSpaces() (bool, error) {
	var b byte
	var err error
	var skip bool

	for b, err = r.pop(); err == nil && isSpace(b); b, err = r.pop() {
		skip = true
	}

	if err != nil {
		return skip, err
	}

	r.unPop()

	return skip, nil
}

func (r *reader) readTag() (string, error) {
	var tok strings.Builder
	tok.Grow(20)

	var err error
	var b byte

	_, err = r.skipSpaces()
	if err != nil {
		return "", err
	}

	for b, err = r.pop(); isAlpha(b) && err == nil; b, err = r.pop() {
		if b < 'a' {
			b += 'a' - 'A'
		}
		tok.WriteByte(b)
	}

	if err != nil {
		return "", err
	}

	r.unPop()

	return tok.String(), nil
}

func (r *reader) readFloat() (float64, error) {
	var tok strings.Builder
	var err error
	var b byte

	tok.Grow(12)
	for b, err = r.pop(); err == nil && isNumeric(b); b, err = r.pop() {
		tok.WriteByte(b)
	}

	if err != nil {
		return 0, err
	}

	r.unPop()

	ret, err := strconv.ParseFloat(tok.String(), 64)
	if err != nil {
		return 0, r.newError("float", "cannot parse %q", tok.String())
	}

	return ret, nil
}

// readPoint reads a space separated tuple of floats, the inside
// of a wkt POINT
func (r *reader) readPoint() (*LngLat, error) {
	lon, err := r.readFloat()
	if err != nil {
		return nil, err
	}

	skip, err := r.skipSpaces()
	if err != nil {
		return nil, err
	}
	if !skip {
		return nil, r.expect("WHITESPACE")
	}

	lat, err := r.readFloat()

	return &LngLat{
		Longitude: lon,
		Latitude:  lat,
	}, err
}

// readPoints (x y,x1 y1,...)
func (r *reader) readPoints() ([]*LngLat, error) {
	b, err := r.pop()
	if err != nil {
		return nil, err
	}
	if b != '(' {
		return nil, r.expect("'('")
	}
	_, err = r.skipSpaces()
	if err != nil {
		return nil, err
	}

	pts := make([]*LngLat, 0, 8)
	for {
		pt, err := r.readPoint()
		if err != nil {
			return nil, err
		}
		pts = append(pts, pt)

		_, err = r.skipSpaces()
		if err != nil {
			return nil, err
		}

		b, err := r.pop()
		if err != nil {
			return nil, err
		}

		switch b {
		case ',':
			_, err = r.skipSpaces()
			if err != nil {
				return nil, err
			}
			continue
		case ')':
			return pts, nil
		default:
			return nil, r.expect("',' or ')'")
		}
	}
}

func (r *reader) readLines() ([][]*LngLat, error) {
	b, err := r.pop()
	if err != nil {
		return nil, err
	}

	if b != '(' {
		return nil, r.expect("(")
	}

	_, err = r.skipSpaces()
	if err != nil {
		return nil, err
	}

	b, err = r.pop()
	if err != nil {
		return nil, err
	}
	if b == ')' {
		return nil, nil
	}
	r.unPop()

	lines := make([][]*LngLat, 0, 4)

	var pts []*LngLat
	for {
		pts, err = r.readPoints()
		if err != nil {
			return nil, err
		}

		lines = append(lines, pts)

		_, err = r.skipSpaces()
		if err != nil {
			return nil, err
		}

		b, err = r.pop()
		if err != nil {
			return nil, err
		}

		switch b {
		case ',':
			_, err = r.skipSpaces()
			if err != nil {
				return nil, err
			}

			continue
		case ')':
			return lines, nil
		default:
			return nil, r.expect("',' or ')'")
		}
	}
}

func (r *reader) readPolys() ([][][]*LngLat, error) {
	b, err := r.pop()
	if err != nil {
		return nil, err
	}

	if b != '(' {
		return nil, r.expect("(")
	}

	_, err = r.skipSpaces()
	if err != nil {
		return nil, err
	}

	b, err = r.pop()
	if err != nil {
		return nil, err
	}
	if b == ')' {
		return nil, nil
	}
	r.unPop()

	polys := make([][][]*LngLat, 0, 4)
	for {
		lines, err := r.readLines()
		if err != nil {
			return nil, err
		}

		polys = append(polys, lines)

		_, err = r.skipSpaces()
		if err != nil {
			return nil, err
		}

		b, err = r.pop()
		if err != nil {
			return nil, err
		}

		switch b {
		case ',':
			_, err = r.skipSpaces()
			if err != nil {
				return nil, err
			}
			continue
		case ')':
			return polys, nil
		default:
			return nil, r.expect("',' or ')'")
		}
	}
}

func (r *reader) expect(chars string) error {
	r.unPop()
	b, err := r.pop()
	if err != nil {
		// this shouldn't happen
		return err
	}

	return r.newError(
		"expect",
		"one of `%q` got %q",
		chars,
		b,
	)
}

func (r *reader) newError(typ string, format string, v ...interface{}) error {
	return &parseError{
		row:  r.row,
		char: r.column,
		typ:  typ,
		why:  fmt.Sprintf(format, v...),
	}
}

type parseError struct {
	row  int
	char int

	typ string
	why string
}

func (e *parseError) Error() string {
	return fmt.Sprintf("syntax error at <line %d, char %d>: %s : %s", e.row+1, e.char+1, e.typ, e.why)
}

func isAlpha(b byte) bool {
	return (b >= 'a' && b <= 'z') || (b >= 'A' && b <= 'Z')
}

func isNumeric(b byte) bool {
	return (b >= '0' && b <= '9') ||
		b == '-' ||
		b == '.' ||
		// b == ',' || // technically part of the spec,
		// but even postgis does not support it
		b == 'E'
}

func isSpace(b byte) bool {
	if b == '\t' ||
		b == '\n' ||
		b == '\v' ||
		b == '\f' ||
		b == '\r' ||
		b == ' ' {
		return true
	}

	return false
}
