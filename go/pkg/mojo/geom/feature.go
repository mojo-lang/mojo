package geom

import (
	"errors"
	"github.com/mojo-lang/mojo/go/pkg/mojo/core"
)

func NewFeature(geometry *Geometry) *Feature {
	feature := &Feature{}
	feature.Geometry = geometry
	feature.Properties = make(map[string]*core.Value)
	return feature
}

func NewPointFeatureFrom(coordinates ...float64) *Feature {
	return NewFeature(NewPointGeometryFrom(coordinates...))
}

func NewPointFeature(ll *LngLat) *Feature {
	return NewFeature(NewPointGeometry(ll))
}

func NewLineStringFeature(coordinates ...*LngLat) *Feature {
	return NewFeature(NewLineStringGeometry(coordinates...))
}

func NewPolygonFeatureFrom(coordinates ...*LngLat) *Feature {
	return NewFeature(NewPolygonGeometryFrom(coordinates...))
}

func (x *Feature) SetIntId(id uint64) {
	if x != nil {
		if x.Id == nil {
			x.Id = &core.Id{}
		}

		x.Id.SetInt(id)
	}
}

func (x *Feature) SetStringId(id string) {
	if x != nil {
		if x.Id == nil {
			x.Id = &core.Id{}
		}

		x.Id.SetString(id)
	}
}

func (x *Feature) SetProperty(key string, val *core.Value) *Feature {
	if x != nil && x.Properties != nil {
		x.Properties[key] = val
	}
	return x
}

func (x *Feature) SetBool(key string, val bool) *Feature {
	return x.SetProperty(key, core.NewBoolValue(val))
}

func (x *Feature) SetInt64(key string, val int64) *Feature {
	return x.SetProperty(key, core.NewInt64Value(val))
}

func (x *Feature) SetString(key string, val string) *Feature {
	return x.SetProperty(key, core.NewStringValue(val))
}

func (x *Feature) SetDouble(key string, val float64) *Feature {
	return x.SetProperty(key, core.NewFloat64Value(val))
}

func (x *Feature) GetPoint() *Point {
	if x != nil && x.Geometry != nil {
		return x.Geometry.GetPoint()
	}
	return nil
}

func (x *Feature) GetLineString() *LineString {
	if x != nil && x.Geometry != nil {
		return x.Geometry.GetLineString()
	}
	return nil
}

func (x *Feature) GetPolygon() *Polygon {
	if x != nil && x.Geometry != nil {
		return x.Geometry.GetPolygon()
	}
	return nil
}

func (x *Feature) GetProperty(key string) (*core.Value, error) {
	if x != nil && x.Properties != nil {
		v, ok := x.Properties[key]
		if ok {
			return v, nil
		} else {
			return nil, errors.New("NotFound")
		}
	}

	return nil, errors.New("feature is nil")
}

func (x *Feature) ToObject() *core.Object {
	return x.ToObjectWith("id", "geometry")
}

func (x *Feature) ToObjectWith(idName string, geometryName string) *core.Object {
	if x == nil || len(idName) == 0 || len(geometryName) == 0 {
		return nil
	}

	obj := core.NewObject()
	hasId := false
	if x.Id != nil {
		obj.SetString(idName, x.Id.Format())
		hasId = true
	}
	obj.SetObject(geometryName, x.Geometry.ToObject())

	for k, v := range x.Properties {
		if hasId && k == idName {
			continue
		}
		obj.SetValue(k, v)
	}
	return obj
}

func NewFeatureFromObject(obj *core.Object) *Feature {
	return NewFeatureFromObjectWith(obj, "id", "geometry")
}

func NewFeatureFromObjectWith(obj *core.Object, idName string, geometryName string) *Feature {
	if obj == nil || len(idName) == 0 || len(geometryName) == 0 {
		return nil
	}

	geometry, _ := NewGeometryFrom(obj.GetValue(geometryName).GetObject())
	if geometry == nil {
		return nil
	}

	feat := NewFeature(geometry)
	if id := obj.GetValue(idName); id != nil {
		switch id.GetKind() {
		case core.ValueKind_VALUE_KIND_INTEGER:
			feat.SetIntId(id.GetUint64())
		case core.ValueKind_VALUE_KIND_STRING:
			feat.SetStringId(id.GetString())
		}
	}

	for k, v := range obj.Vals {
		if k == idName && feat.Id != nil {
			continue
		}
		if k == geometryName {
			continue
		}

		feat.Properties[k] = v
	}
	return feat
}

func (x *Feature) GetIntProperty(key string) (int64, error) {
	val, err := x.GetProperty(key)
	if err != nil {
		return 0, err
	}
	return val.GetInt64(), nil
}

func (x *Feature) GetStringProperty(key string) (string, error) {
	val, err := x.GetProperty(key)
	if err != nil {
		return "", err
	}
	return val.GetString(), nil
}

func (x *Feature) GetBoolProperty(key string) (bool, error) {
	val, err := x.GetProperty(key)
	if err != nil {
		return false, err
	}
	return val.GetBool(), nil
}

func (x *Feature) GetFloat64Property(key string) (float64, error) {
	val, err := x.GetProperty(key)
	if err != nil {
		return 0, err
	}
	return val.GetFloat64(), nil
}

func (x *Feature) GetDoubleProperty(key string) (float64, error) {
	return x.GetFloat64Property(key)
}

func (x *Feature) BoundingBox() *BoundingBox {
	if x != nil {
		if x.Bbox == nil {
			x.Bbox = x.Geometry.BoundingBox()
		}
		return x.Bbox
	}
	return nil
}

func (x *Feature) CoordTransform(from, to SpatialReference) *Feature {
	if x == nil {
		return x
	}
	return &Feature{
		Type:             x.Type,
		Id:               x.Id,
		Geometry:         x.Geometry.CoordTransform(from, to),
		Bbox:             x.Bbox.CoordTransform(from, to),
		SpatialReference: x.SpatialReference,
		Properties:       x.Properties,
	}
}
