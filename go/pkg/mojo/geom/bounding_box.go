package geom

import (
	"errors"
	"fmt"
	jsoniter "github.com/json-iterator/go"
)

func NewBoundingBox(leftBottom *LngLat, rightTop *LngLat) *BoundingBox {
	if leftBottom == nil || rightTop == nil {
		return nil
	}
	return (&BoundingBox{LeftBottom: leftBottom, RightTop: leftBottom}).Extend(&BoundingBox{
		LeftBottom: rightTop,
		RightTop:   rightTop,
	})
}

// NewBoundingBoxFromString
// [[],[]]
// [left,bottom,right,top]
// left,bottom,right,top
func NewBoundingBoxFromString(str string) (*BoundingBox, error) {
	if len(str) == 0 {
		return nil, errors.New("empty input string")
	}

	var points [][]float64
	if err := jsoniter.UnmarshalFromString(str, &points); err == nil {
		if len(points) == 2 && ((len(points[0]) == 2 && len(points[1]) == 2) || (len(points[0]) == 3 && len(points[1]) == 3)) {
			return NewBoundingBox(NewLngLat(points[0]...), NewLngLat(points[1]...)), nil
		} else {
			return nil, errors.New(fmt.Sprintf("invalid bounding box format, %s", str))
		}
	}

	if str[0] != '[' {
		str = "[" + str + "]"
	}

	var pp []float64
	if err := jsoniter.UnmarshalFromString(str, &pp); err != nil {
		return nil, err
	}

	if len(pp) == 4 {
		return NewBoundingBox(NewLngLat(pp[0], pp[1]), NewLngLat(pp[2], pp[3])), nil
	} else if len(pp) == 6 {
		return NewBoundingBox(NewLngLat(pp[0], pp[1], pp[2]), NewLngLat(pp[3], pp[4], pp[5])), nil
	}

	return nil, errors.New(fmt.Sprintf("invalid bounding box format, %s", str))
}

func (x *BoundingBox) Equal(bb *BoundingBox) bool {
	if x != nil {
		if bb != nil {
			return x.LeftBottom.Equal(bb.LeftBottom) && x.RightTop.Equal(bb.RightTop)
		} else {
			return false
		}
	} else {
		if bb != nil {
			return false
		} else {
			return true
		}
	}
}

func (x *BoundingBox) Vertexes() []*LngLat {
	if x != nil {
		return []*LngLat{
			x.LeftBottom,
			{Longitude: x.RightTop.GetLongitude(), Latitude: x.LeftBottom.GetLatitude()},
			x.RightTop,
			{Longitude: x.LeftBottom.GetLongitude(), Latitude: x.RightTop.GetLatitude()},
		}
	}
	return nil
}

func (x *BoundingBox) Center() *LngLat {
	if x != nil {
		return &LngLat{
			Longitude: (x.LeftBottom.GetLongitude() + x.RightTop.GetLongitude()) / 2,
			Latitude:  (x.LeftBottom.GetLatitude() + x.RightTop.GetLatitude()) / 2,
		}
	}
	return nil
}

func (x *BoundingBox) ToPolygon() *Polygon {
	if x != nil {
		vertexes := x.Vertexes()
		vertexes = append(vertexes, x.LeftBottom)
		return NewPolygonFrom(vertexes...)
	}
	return nil
}

func (x *BoundingBox) ToGeometry() *Geometry {
	if x != nil {
		return NewPolygonGeometry(x.ToPolygon().LineStrings...)
	}
	return nil
}

func (x *BoundingBox) Vertex(index int) *LngLat {
	switch index {
	case 0:
		return x.LeftBottom
	case 1:
		return &LngLat{Longitude: x.RightTop.Longitude, Latitude: x.LeftBottom.Latitude}
	case 2:
		return x.RightTop
	case 3:
		return &LngLat{Longitude: x.LeftBottom.Longitude, Latitude: x.RightTop.Latitude}
	default:
		return nil
	}
}

func (x *BoundingBox) ExtendToPoint(point *LngLat) *BoundingBox {
	if x != nil && point != nil {
		if x.LeftBottom == nil || x.RightTop == nil {
			x.LeftBottom = NewLngLat(point.Longitude, point.Latitude)
			x.RightTop = NewLngLat(point.Longitude, point.Latitude)
		} else {
			if point.Longitude < x.LeftBottom.Longitude {
				x.LeftBottom.Longitude = point.Longitude
			}
			if point.Latitude < x.LeftBottom.Latitude {
				x.LeftBottom.Latitude = point.Latitude
			}
			if point.Longitude > x.RightTop.Longitude {
				x.RightTop.Longitude = point.Longitude
			}
			if point.Latitude > x.RightTop.Latitude {
				x.RightTop.Latitude = point.Latitude
			}
		}
	}
	return x
}

func (x *BoundingBox) Extend(box *BoundingBox) *BoundingBox {
	if box != nil {
		var minLat, minLng, maxLat, maxLng float64
		if x.LeftBottom.Latitude > box.LeftBottom.Latitude {
			minLat = box.LeftBottom.Latitude
		} else {
			minLat = x.LeftBottom.Latitude
		}

		if x.LeftBottom.Longitude > box.LeftBottom.Longitude {
			minLng = box.LeftBottom.Longitude
		} else {
			minLng = x.LeftBottom.Longitude
		}

		if x.RightTop.Latitude < box.RightTop.Latitude {
			maxLat = box.RightTop.Latitude
		} else {
			maxLat = x.RightTop.Latitude
		}

		if x.RightTop.Longitude < box.RightTop.Longitude {
			maxLng = box.RightTop.Longitude
		} else {
			maxLng = x.RightTop.Longitude
		}

		return &BoundingBox{LeftBottom: &LngLat{Longitude: minLng, Latitude: minLat},
			RightTop: &LngLat{Longitude: maxLng, Latitude: maxLat}}
	}
	return x
}

func (x *BoundingBox) ExtendDegree(lng float64, lat float64) *BoundingBox {
	minLng := x.LeftBottom.Longitude - lng
	minLat := x.LeftBottom.Latitude - lat

	maxLng := x.RightTop.Longitude + lng
	maxLat := x.RightTop.Latitude + lat

	return &BoundingBox{LeftBottom: &LngLat{Longitude: minLng, Latitude: minLat},
		RightTop: &LngLat{Longitude: maxLng, Latitude: maxLat}}
}

func (x *BoundingBox) CoordTransform(from, to SpatialReference) *BoundingBox {
	if x != nil {
		return &BoundingBox{
			LeftBottom: x.LeftBottom.CoordTransform(from, to),
			RightTop:   x.RightTop.CoordTransform(from, to),
		}
	}
	return x
}
