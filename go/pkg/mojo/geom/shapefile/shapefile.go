package shapefile

import (
	"archive/zip"
	"bytes"
	"fmt"
	"io"
	"math"
	"os"
	"path"
	"path/filepath"
	"strconv"
	"strings"

	"github.com/jonas-p/go-shp"
	"github.com/mojo-lang/mojo/go/pkg/mojo/core"
	"github.com/mojo-lang/mojo/go/pkg/mojo/geom"
)

const MaxPrecision = 8

// Shapefile
// max Precision  default is 8
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
					float, err := strconv.ParseFloat(v, 64)
					if err != nil {
					} else {
						if math.Abs(float64(int64(float))-float) < 1.0e-4 {
							feat.SetInt64(f.String(), int64(float))
						} else {
							feat.SetDouble(f.String(), float)
						}
					}
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

		if id, ok := feat.Properties["id"]; ok && id != nil {
			feat.SetStringId(id.GetString())
		}

		features = append(features, feat)
	}

	return features, nil
}

type Field struct {
	Type      string
	Name      string
	Size      int
	Precision int
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
	defer func() {
		shape.Close()
	}()

	for _, feat := range feats {
		if _, ok := feat.Properties["id"]; !ok {
			continue
		}

		id := feat.Id.Format()
		feat.Properties["id"] = core.NewStringValue(id)
	}

	// fields to write
	//fields := []shp.Field{shp.StringField("id", 64)}
	var fields []shp.Field
	var fieldNames []string
	originalFields := make(map[string]*Field)
	for i, feat := range feats {
		for k, v := range feat.Properties {
			field := originalFields[k]
			if field == nil {
				field = &Field{
					Name: k,
				}
				originalFields[k] = field
			}

			switch v.GetKind() {
			case core.ValueKind_VALUE_KIND_BOOLEAN:
			case core.ValueKind_VALUE_KIND_INTEGER:
				if len(field.Type) == 0 {
					field.Type = "integer"
				}
				iv := v.GetInt64()
				size := len(strconv.FormatInt(iv, 10)) + 1
				if size > field.Size {
					field.Size = size
				}
			case core.ValueKind_VALUE_KIND_NUMBER:
				if len(field.Type) == 0 || field.Type == "integer" {
					field.Type = "number"
				}
				fv := v.GetFloat64()
				precision := core.DecimalPrecision(fv)
				if precision > MaxPrecision {
					precision = MaxPrecision
					feat.Properties[k] = core.NewFloat64Value(core.TruncateTo(fv, precision))
				}
				if precision > field.Precision {
					field.Precision = precision
				}
				size := len(strconv.FormatFloat(fv, 'f', precision, 64)) + 1
				if size > field.Size {
					field.Size = size
				}
			case core.ValueKind_VALUE_KIND_STRING:
				if len(field.Type) == 0 {
					field.Type = "string"
				}
				size := len(v.GetString()) + 1
				if size > field.Size {
					field.Size = size
				}
			}
			if i == 0 {
				fieldNames = append(fieldNames, k)
			}
		}
	}

	for _, name := range fieldNames {
		field := originalFields[name]
		var f shp.Field
		switch field.Type {
		case "string":
			f = shp.StringField(name, uint8(field.Size))
		case "integer":
			f = shp.NumberField(name, uint8(field.Size))
		case "number":
			f = shp.FloatField(name, uint8(field.Size), uint8(field.Precision))
		}
		fields = append(fields, f)
	}

	// setup fields for attributes
	if err = shape.SetFields(fields); err != nil {
		return err
	}

	// write points and attributes
	for n, feat := range feats {
		if sp, err := GeometryToShape(feat.GetGeometry(), feat.GetBbox()); err != nil {
			return err
		} else {
			shape.Write(sp)
		}

		id := ""
		if feat.Id != nil {
			id = feat.Id.Format()
		}
		if len(id) == 0 {
			if val, ok := feat.Properties["id"]; ok && val != nil {
				id = val.GetString()
			}
		}

		for i, name := range fieldNames {
			err = nil
			value, _ := feat.GetProperty(name)
			field := fields[i]
			buf := bytes.Repeat([]byte{' '}, int(field.Size))
			switch v := value.GetVal().(type) {
			case *core.Value_BoolVal:
				vl := "false"
				if value.GetBool() {
					vl = "true"
				}
				copy(buf[:], vl)
			case *core.Value_PositiveVal, *core.Value_NegativeVal:
				copy(buf[:], strconv.Itoa(int(value.GetInt64())))
			case *core.Value_DoubleVal:
				copy(buf[:], strconv.FormatFloat(value.GetFloat64(), 'f', int(field.Precision), 64))
			case *core.Value_StringVal:
				copy(buf[:], v.StringVal)
			default:
				return fmt.Errorf("unsupported value type: %T for property (%s) of %s", v, name, id)
			}

			err = shape.WriteAttribute(n, i, string(buf))
			if err != nil {
				return fmt.Errorf("failed to set the property (%s) of %s, err: %e", name, id, err)
			}
		}
	}

	return nil
}

func (s Shapefile) WriteAllToBytes(feats []*geom.Feature, filename string) ([]byte, error) {
	dir, err := os.MkdirTemp("", "mojo-geom-shapefiles")
	if err != nil {
		return nil, err
	}
	fn := path.Join(dir, filename)
	err = core.CreateDir(path.Dir(fn))
	if err != nil {
		return nil, err
	}

	err = s.WriteAll(feats, fn)
	if err != nil {
		return nil, err
	}

	fp := strings.TrimSuffix(fn, ".shp")
	files := []string{
		fp + ".cpg",
		fp + ".dbf",
		fp + ".prj",
		fp + ".shp",
		fp + ".shx",
	}
	return ZipFiles(files)
}

func ZipFiles(filePaths []string) ([]byte, error) {
	buf := new(bytes.Buffer)
	zw := zip.NewWriter(buf)
	defer func() {
		if err := zw.Close(); err != nil {
		}
	}()

	for _, filePath := range filePaths {
		file, err := os.Open(filePath)
		if err != nil {
			if os.IsNotExist(err) {
				continue
			}

			return nil, err
		}

		fileInfo, err := file.Stat()
		if err != nil {
			_ = file.Close()
			return nil, err
		}

		header, err := zip.FileInfoHeader(fileInfo)
		if err != nil {
			_ = file.Close()
			return nil, err
		}

		header.Name = filepath.Base(filePath)
		header.Method = zip.Deflate

		zipFileWriter, err := zw.CreateHeader(header)
		if err != nil {
			_ = file.Close()
			return nil, err
		}

		if _, err = io.Copy(zipFileWriter, file); err != nil {
			_ = file.Close()
			return nil, err
		}

		_ = file.Close()
	}

	if err := zw.Close(); err != nil {
		return nil, err
	}

	return buf.Bytes(), nil
}
