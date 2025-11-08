package xlsx

import (
	"fmt"
	"github.com/mojo-lang/mojo/packages/core/go/pkg/logs"
	"github.com/mojo-lang/mojo/packages/core/go/pkg/mojo/core"
	"github.com/mojo-lang/mojo/packages/geom/go/pkg/mojo/geom"
	"github.com/xuri/excelize/v2"
	"path/filepath"
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

	var titles []string
	for i, feat := range feats {
		for j, title := range titles {
			cell, err := excelize.CoordinatesToCellName(i+2, j+1)
			if err != nil {
				return err
			}
			var row []interface{}

			switch title {
			case idField:
			case geomField:
				row = append(row, feat.GetGeometry().ToWKT())
			default:
			}

			if err = f.SetSheetRow(sheetName, cell, &row); err != nil {
				return err
			}
		}
	}

	// Save spreadsheet by the given path.
	if err := f.SaveAs(filename); err != nil {
		return err
	}
	return nil
}
