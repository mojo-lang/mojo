package xlsx

import (
	"errors"
	"fmt"
	"github.com/mojo-lang/mojo/go/pkg/logs"
	"github.com/mojo-lang/mojo/go/pkg/mojo/core"
	"github.com/mojo-lang/mojo/go/pkg/mojo/geom"
	"github.com/xuri/excelize/v2"
	"os"
	"path"
	"path/filepath"
	"sort"
	"strconv"
	"strings"
)

type File core.Options

func (x File) ReadAll(filename string) ([]*geom.Feature, error) {
	f, err := excelize.OpenFile(filename)
	if err != nil {
		return nil, err
	}

	defer func() {
		if err = f.Close(); err != nil {
			logs.Warnw("failed to close the file", "error", err)
		}
	}()

	sheetName := core.Options(x).GetString("sheet_name")
	idField := core.Options(x).GetString("id_field")
	idType := core.Options(x).GetString("id_type")
	geomField := core.Options(x).GetString("geometry_field")

	if len(sheetName) == 0 {
		sheets := f.GetSheetList()
		if len(sheets) > 0 {
			sheetName = sheets[0]
		} else {
			return nil, fmt.Errorf("the excel file %s is empty", filename)
		}
	}

	if len(idField) == 0 {
		idField = "id"
	}
	if len(geomField) == 0 {
		geomField = "geometry"
	}

	lngField := ""
	latField := ""
	fields := strings.Split(geomField, ",")
	if len(fields) > 1 {
		if len(fields) == 2 {
			lngField = fields[0]
			latField = fields[1]
		} else {
			return nil, fmt.Errorf("invalid geometry field name (%s) for the excel file %s is empty", geomField, filename)
		}
	}

	rows, err := f.GetRows(sheetName)
	if err != nil {
		return nil, err
	}

	var features []*geom.Feature
	var title []string
	idIndex := -1
	geomIndex := -1
	lngIndex := -1
	latIndex := -1
	for i, row := range rows {
		if i == 0 {
			for j, cell := range row {
				cell = strings.TrimSpace(cell)
				title = append(title, cell)

				if idField == "id" && strings.ToLower(cell) == idField {
					idIndex = j
				} else if geomField == "geometry" && strings.ToLower(cell) == geomField {
					geomIndex = j
				} else {
					switch cell {
					case idField:
						idIndex = j
					case geomField:
						geomIndex = j
					case lngField:
						lngIndex = j
					case latField:
						latIndex = j
					}
				}
			}
			continue
		} else {
			if geomIndex < 0 && (lngIndex < 0 || latIndex < 0) {
				return nil, fmt.Errorf("not found the geometry field name (%s) for the excel file %s", geomField, filename)
			}
			if idIndex < 0 {
				return nil, fmt.Errorf("not found the field name (%s) for the excel file %s", idField, filename)
			}
		}

		feature := geom.NewFeature(nil)
		lnglat := &geom.LngLat{}
		for j, cell := range row {
			key := title[j]
			cell = strings.TrimSpace(cell)

			switch j {
			case idIndex:
				if idType == "integer" {
					if iv, err := strconv.ParseInt(cell, 10, 64); err != nil {
						return nil, fmt.Errorf("the id type is integer but failed to parse %s, error: %e", cell, err)
					} else {
						feature.Id = core.NewIntId(uint64(iv))
					}
				} else {
					feature.Id = core.NewStringId(cell)
				}
			case geomIndex:
				feature.Geometry, err = geom.NewGeometryFromWKT(cell)
			case lngIndex:
				if fv, err := strconv.ParseFloat(cell, 64); err == nil {
					lnglat.Longitude = fv
				}
			case latIndex:
				if fv, err := strconv.ParseFloat(cell, 64); err == nil {
					lnglat.Latitude = fv
				}
			}

			if iv, err := strconv.ParseInt(cell, 10, 64); err == nil {
				feature.SetInt64(key, iv)
			} else if fv, err := strconv.ParseFloat(cell, 64); err == nil {
				feature.SetDouble(key, fv)
			} else {
				feature.SetString(key, cell)
			}
		}

		if feature.Geometry == nil {
			if !lnglat.IsEmpty() {
				feature.Geometry = geom.NewPointGeometry(lnglat)
			} else {
				return nil, fmt.Errorf("should set the geometry fild name for the excel file %s is empty", filename)
			}
		}

		features = append(features, feature)
	}

	return features, nil
}

func (x File) WriteAll(feats []*geom.Feature, filename string) error {
	if len(feats) == 0 {
		return errors.New("feats is empty")
	}
	if len(filename) == 0 {
		return errors.New("no filename")
	}

	f := excelize.NewFile()
	defer func() {
		if err := f.Close(); err != nil {
			logs.Warnw("failed to close the file", "error", err)
		}
	}()

	if filepath.Ext(filename) != ".xlsx" {
		filename = filename + ".xlsx"
	}

	sheetName := core.Options(x).GetString("sheet_name")
	idField := core.Options(x).GetString("id_field")
	geomField := core.Options(x).GetString("geometry_field")

	if len(sheetName) == 0 {
		sheetName = "Sheet1"
	}
	if len(idField) == 0 {
		idField = "id"
	}
	if len(geomField) == 0 {
		geomField = "geometry"
	}

	var titles []string
	for k, _ := range feats[0].Properties {
		titles = append(titles, k)
	}
	sort.Strings(titles)

	ts := []string{idField}
	for _, t := range titles {
		if t != idField {
			ts = append(ts, t)
		}
	}
	titles = append(ts, geomField)
	_ = SetRow(f, sheetName, 1, 1, &titles)

	for i, feat := range feats {
		var row []interface{}
		for _, title := range titles {
			switch title {
			case idField:
				row = append(row, feat.Id.Format())
			case geomField:
				row = append(row, feat.GetGeometry().ToWKT())
			default:
				val := ""
				if v, _ := feat.GetProperty(title); v != nil {
					switch v.GetKind() {
					case core.ValueKind_VALUE_KIND_STRING:
						val = v.GetString()
					case core.ValueKind_VALUE_KIND_INTEGER:
						val = strconv.Itoa(int(v.GetInt64()))
					case core.ValueKind_VALUE_KIND_NUMBER:
						val = strconv.FormatFloat(v.GetFloat64(), 'f', -1, 64)
					}
				}
				row = append(row, val)
			}
		}
		if err := SetRow(f, sheetName, i+2, 1, &row); err != nil {
			return err
		}
	}

	// Save spreadsheet by the given path.
	if err := f.SaveAs(filename); err != nil {
		return err
	}
	return nil
}

func SetValue(file *excelize.File, sheet string, row int, col int, value interface{}) error {
	cell, err := excelize.CoordinatesToCellName(row, col)
	if err != nil {
		return err
	}
	return file.SetCellValue(sheet, cell, value)
}

func SetRow(file *excelize.File, sheet string, row int, col int, data interface{}) error {
	cell, err := excelize.CoordinatesToCellName(col, row)
	if err != nil {
		return err
	}
	return file.SetSheetRow(sheet, cell, data)
}

func (x File) WriteAllToBytes(feats []*geom.Feature, filename string) ([]byte, error) {
	dir, err := os.MkdirTemp("", "mojo-geom-xlsx")
	if err != nil {
		return nil, err
	}
	fn := path.Join(dir, filename)
	err = core.CreateDir(path.Dir(fn))
	if err != nil {
		return nil, err
	}

	err = x.WriteAll(feats, fn)
	if err != nil {
		return nil, err
	}

	return os.ReadFile(fn)
}
