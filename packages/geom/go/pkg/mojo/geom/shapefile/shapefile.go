package shapefile

import (
	"fmt"
	"github.com/jonas-p/go-shp"
	"github.com/mojo-lang/mojo/packages/core/go/pkg/mojo/core"
	"github.com/mojo-lang/mojo/packages/geom/go/pkg/mojo/geom"
	"path"
	"strconv"
)

type Shapefile core.Options

func (s Shapefile) ReadAll(filename string) ([]*geom.Feature, error) {
	shape, err := shp.Open(filename)
	if err != nil {
		return nil, err
	}
	defer func() { _ = shape.Close() }()

	// fields from the attribute table (DBF)
	fields := shape.Fields()
	var features []*geom.Feature
	// loop through all features in the shapefile
	for shape.Next() {
		feat := geom.NewFeature(nil)

		n, p := shape.Shape()

		feat.Geometry, feat.Bbox, err = NewGeometry(p)
		if err != nil {
			return nil, err
		}

		for k, f := range fields {
			v := shape.ReadAttribute(n, k)
			switch f.Fieldtype {
			case 'N':
				number, err := strconv.ParseInt(v, 10, 64)
				if err != nil {
				} else {
					feat.SetInt64(f.String(), number)
				}
			case 'F':
				number, err := strconv.ParseFloat(v, 10)
				if err != nil {
				} else {
					feat.SetDouble(f.String(), number)
				}
			case 'C':
				feat.SetString(f.String(), v)
			case 'D':
				feat.SetString(f.String(), v)
			}
		}
		features = append(features, feat)
	}

	return features, nil
}

func (s Shapefile) WriteAll(feats []*geom.Feature, filename string) error {
	if len(feats) == 0 || len(filename) == 0 {
		return nil
	}

	if path.Ext(filename) != ".shp" {
		filename = filename + ".shp"
	}

	typ, err := GeometryType(feats[0].GetGeometry())
	if err != nil {
		return err
	}

	// create and open a shapefile for writing points
	shape, err := shp.Create(filename, typ)
	if err != nil {
		return err
	}
	defer shape.Close()

	// fields to write
	fields := []shp.Field{shp.StringField("id", 64)}
	var fieldNames []string

	for k, _ := range feats[0].Properties {
		var field shp.Field

		field = shp.StringField(k, 25)

		fields = append(fields, field)
	}

	// setup fields for attributes
	if err = shape.SetFields(fields); err != nil {
		return err
	}

	// write points and attributes
	for n, feat := range feats {
		if s, err := GeometryToShape(feat.GetGeometry(), feat.GetBbox()); err != nil {
			return err
		} else {
			shape.Write(s)
		}

		id := feat.Id.Format()
		if err = shape.WriteAttribute(n, 0, id); err != nil {
			return err
		}

		for i, name := range fieldNames {
			if name == "id" {
				continue
			}

			err = nil
			value, _ := feat.GetProperty(name)
			switch v := value.GetVal().(type) {
			case *core.Value_BoolVal:
			case *core.Value_PositiveVal:
			case *core.Value_NegativeVal:
			case *core.Value_DoubleVal:
			case *core.Value_StringVal:
				err = shape.WriteAttribute(n, i+1, v.StringVal)
			default:

			}

			if err != nil {
				return fmt.Errorf("failed to set the property (%s) of %s, err: %e", name, id, err)
			}
		}
	}

	return nil
}
