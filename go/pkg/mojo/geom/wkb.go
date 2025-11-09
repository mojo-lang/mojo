package geom

import (
	"bufio"
	"bytes"
	"encoding/binary"
	"errors"
	"fmt"
)

const (
	WKBPoint              uint32 = 1
	WKBLineString         uint32 = 2
	WKBPolygon            uint32 = 3
	WKBMultiPoint         uint32 = 4
	WKBMultiLineString    uint32 = 5
	WKBMultiPolygon       uint32 = 6
	WKBGeometryCollection uint32 = 7
)

func NewGeometryFromWKB(wkb []byte) (*Geometry, error) {
	d := &WKB{}
	return d.Decode(wkb)
}

func (x *Geometry) FromWKB(wkb []byte) error {
	d := &WKB{}
	g, e := d.Decode(wkb)
	if e != nil {
		return e
	}
	x.Geometry = g.Geometry
	return nil
}

func (x *Geometry) ToWKB() []byte {
	enc := &WKB{}
	return enc.Encode(x)
}

type WKB struct {
}

func (w *WKB) Decode(wkb []byte) (*Geometry, error) {
	return w.decode(bufio.NewReader(bytes.NewReader(wkb)))
}

func (w *WKB) Encode(geometry *Geometry) []byte {
	buffer := bytes.NewBuffer(nil)
	return w.encode(buffer, geometry).bytes(buffer)
}

func (w *WKB) decode(reader *bufio.Reader) (*Geometry, error) {
	var byteOrder binary.ByteOrder = binary.LittleEndian
	bo, err := reader.ReadByte()
	if err != nil {
		return nil, err
	}

	if bo == 0 {
		byteOrder = binary.BigEndian
	} else if bo > 1 {
		return nil, fmt.Errorf("invalid byte order: %d", bo)
	}

	var gType uint32
	err = binary.Read(reader, byteOrder, &gType)
	if err != nil {
		return nil, err
	}

	switch gType {
	case WKBPoint:
		return w.decodePoint(reader, byteOrder)
	case WKBLineString:
		return w.decodeLineString(reader, byteOrder)
	case WKBPolygon:
		return w.decodePolygon(reader, byteOrder)
	case WKBMultiPoint:
		return w.decodeMultiPoint(reader, byteOrder)
	case WKBMultiLineString:
		return w.decodeMultiLineString(reader, byteOrder)
	case WKBMultiPolygon:
		return w.decodeMultiPolygon(reader, byteOrder)
	case WKBGeometryCollection:
		return w.decodeGeometryCollection(reader, byteOrder)
	default:
		return nil, errors.New("unknown wkb type")
	}
}

func (w *WKB) decodePoint(reader *bufio.Reader, byteOrder binary.ByteOrder) (*Geometry, error) {
	var xy [2]float64
	err := binary.Read(reader, byteOrder, &xy)
	if err != nil {
		return nil, err
	}
	return NewPointGeometryFrom(xy[0], xy[1]), nil
}

func (w *WKB) decodeLineString(reader *bufio.Reader, byteOrder binary.ByteOrder) (*Geometry, error) {
	var n uint32
	if err := binary.Read(reader, byteOrder, &n); err != nil {
		return nil, err
	}

	pts := make([]*LngLat, n)
	for i := 0; i < int(n); i++ {
		var xy [2]float64
		if err := binary.Read(reader, byteOrder, &xy); err != nil {
			return nil, err
		}
		pts[i] = &LngLat{
			Longitude: xy[0],
			Latitude:  xy[1],
		}
	}

	return NewLineStringGeometry(pts...), nil
}

func (w *WKB) decodePolygon(reader *bufio.Reader, byteOrder binary.ByteOrder) (*Geometry, error) {
	var n uint32
	if err := binary.Read(reader, byteOrder, &n); err != nil {
		return nil, err
	}

	lines := make([]*LineString, 0, n)
	for i := 0; i < int(n); i++ {
		ring, err := w.decodeLineString(reader, byteOrder)
		if err != nil {
			return nil, err
		}

		lines = append(lines, ring.Geometry.(*Geometry_LineString).LineString)
	}

	return NewPolygonGeometry(lines...), nil
}

func (w *WKB) decodeMultiPoint(reader *bufio.Reader, byteOrder binary.ByteOrder) (*Geometry, error) {
	var n uint32
	if err := binary.Read(reader, byteOrder, &n); err != nil {
		return nil, err
	}

	pts := make([]*LngLat, n)
	var typ uint32
	for i := 0; i < int(n); i++ {
		b, err := reader.ReadByte()
		if err != nil {
			return nil, err
		}
		if b == 0 {
			byteOrder = binary.BigEndian
		} else if b > 1 {
			return nil, fmt.Errorf("invalid byte order: %d", b)
		}
		err = binary.Read(reader, byteOrder, &typ)
		if err != nil {
			return nil, err
		}

		p, err := w.decodePoint(reader, byteOrder)
		if err != nil {
			return nil, err
		}

		pts[i] = p.GetPoint().GetCoordinate()
	}

	return NewMultiPointGeometryFrom(pts...), nil
}

func (w *WKB) decodeMultiLineString(reader *bufio.Reader, byteOrder binary.ByteOrder) (*Geometry, error) {
	var n uint32
	if err := binary.Read(reader, byteOrder, &n); err != nil {
		return nil, err
	}

	lines := make([]*LineString, n)
	var typ uint32
	for i := 0; i < int(n); i++ {
		b, err := reader.ReadByte()
		if err != nil {
			return nil, err
		}
		if b == 0 {
			byteOrder = binary.BigEndian
		} else if b > 1 {
			return nil, fmt.Errorf("invalid byte order: %d", b)
		}

		err = binary.Read(reader, byteOrder, &typ)
		if err != nil {
			return nil, err
		}

		line, err := w.decodeLineString(reader, byteOrder)
		if err != nil {
			return nil, err
		}

		lines[i] = line.GetLineString()
	}

	return NewMultiLineStringGeometry(lines...), nil
}

func (w *WKB) decodeMultiPolygon(reader *bufio.Reader, byteOrder binary.ByteOrder) (*Geometry, error) {
	var n uint32
	if err := binary.Read(reader, byteOrder, &n); err != nil {
		return nil, err
	}

	polys := make([]*Polygon, n)
	var typ uint32
	for i := 0; i < int(n); i++ {
		b, err := reader.ReadByte()
		if err != nil {
			return nil, err
		}
		if b == 0 {
			byteOrder = binary.BigEndian
		} else if b > 1 {
			return nil, fmt.Errorf("invalid byte order: %d", b)
		}

		err = binary.Read(reader, byteOrder, &typ)
		if err != nil {
			return nil, err
		}

		poly, err := w.decodePolygon(reader, byteOrder)
		if err != nil {
			return nil, err
		}

		polys[i] = poly.GetPolygon()
	}

	return NewMultiPolygonGeometry(polys...), nil
}

func (w *WKB) decodeGeometryCollection(reader *bufio.Reader, byteOrder binary.ByteOrder) (*Geometry, error) {
	var n uint32
	if err := binary.Read(reader, byteOrder, &n); err != nil {
		return nil, err
	}

	gs := make([]*Geometry, n)
	for i := 0; i < int(n); i++ {
		g, err := w.decode(reader)
		if err != nil {
			return nil, err
		}

		gs[i] = g
	}

	return NewGeometryCollectionGeometry(gs...), nil
}

func (w *WKB) bytes(buffer *bytes.Buffer) []byte {
	return buffer.Bytes()
}

func (w *WKB) writeByteOrder(buffer *bytes.Buffer) *WKB {
	err := binary.Write(buffer, binary.LittleEndian, byte(1))
	if err != nil {
		panic(err)
	}
	return w
}

func (w *WKB) write(buffer *bytes.Buffer, values ...interface{}) *WKB {
	var err error
	for _, v := range values {
		err = binary.Write(buffer, binary.LittleEndian, v)
		if err != nil {
			panic(err)
		}
	}
	return w
}

func (w *WKB) writeLngLat(buffer *bytes.Buffer, ll *LngLat) *WKB {
	if ll != nil {
		return w.write(buffer, ll.Longitude, ll.Latitude)
	}
	return w
}

func (w *WKB) encodePoint(buffer *bytes.Buffer, p *Point) *WKB {
	if p != nil {
		w.writeByteOrder(buffer).write(buffer, WKBPoint)
		return w.writeLngLat(buffer, p.Coordinate)
	}
	return w
}

func (w *WKB) encodeMultiPoint(buffer *bytes.Buffer, points *MultiPoint) *WKB {
	pts := points.Points
	w.writeByteOrder(buffer).write(buffer, WKBMultiPoint, uint32(len(pts)))
	for _, p := range pts {
		w.encodePoint(buffer, p)
	}

	return w
}

func (w *WKB) encodeLineString(buffer *bytes.Buffer, line *LineString) *WKB {
	pts := line.Coordinates
	w.writeByteOrder(buffer).write(buffer, WKBLineString, uint32(len(pts)))
	for _, p := range pts {
		w.writeLngLat(buffer, p)
	}

	return w
}

func (w *WKB) encodeMultiLineString(buffer *bytes.Buffer, multiLine *MultiLineString) *WKB {
	lines := multiLine.LineStrings
	w.writeByteOrder(buffer).write(buffer, WKBMultiLineString, uint32(len(lines)))
	for _, line := range lines {
		w.encodeLineString(buffer, line)
	}

	return w
}

func (w *WKB) encodePolygon(buffer *bytes.Buffer, poly *Polygon) *WKB {
	rings := poly.LineStrings
	w.writeByteOrder(buffer).write(buffer, WKBPolygon, uint32(len(rings)))
	for _, ring := range rings {
		pts := ring.Coordinates
		w.write(buffer, uint32(len(pts)))
		for _, p := range pts {
			w.writeLngLat(buffer, p)
		}
	}

	return w
}

func (w *WKB) encodeMultiPolygon(buffer *bytes.Buffer, multiPolygon *MultiPolygon) *WKB {
	polys := multiPolygon.Polygons
	w.writeByteOrder(buffer).write(buffer, WKBMultiPolygon, uint32(len(polys)))
	for _, p := range polys {
		w.encodePolygon(buffer, p)
	}

	return w
}

func (w *WKB) encodeGeometryCollection(buffer *bytes.Buffer, collection *GeometryCollection) *WKB {
	gs := collection.Geometries
	w.writeByteOrder(buffer).write(buffer, WKBGeometryCollection, uint32(len(gs)))
	for _, g := range gs {
		w.encode(buffer, g)
	}

	return w
}

// TODO add internal error state for encoding
func (w *WKB) encode(buffer *bytes.Buffer, g *Geometry) *WKB {
	switch gt := g.Geometry.(type) {
	case *Geometry_Point:
		w.encodePoint(buffer, gt.Point)
	case *Geometry_MultiPoint:
		w.encodeMultiPoint(buffer, gt.MultiPoint)
	case *Geometry_LineString:
		w.encodeLineString(buffer, gt.LineString)
	case *Geometry_MultiLineString:
		w.encodeMultiLineString(buffer, gt.MultiLineString)
	case *Geometry_Polygon:
		w.encodePolygon(buffer, gt.Polygon)
	case *Geometry_MultiPolygon:
		w.encodeMultiPolygon(buffer, gt.MultiPolygon)
	case *Geometry_GeometryCollection:
		w.encodeGeometryCollection(buffer, gt.GeometryCollection)
	}

	return w
}
