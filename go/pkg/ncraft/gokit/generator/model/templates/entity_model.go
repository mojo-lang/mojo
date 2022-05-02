package templates

const EntityModel = `
package model

import (
	"sync"

	"github.com/ncraft-io/ncraft/go/pkg/ncraft/logs"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"

	{{.Import}}
)

var {{.LowerCamelName}}Model *{{.Name}}Model
var {{.LowerCamelName}}ModelOnce sync.Once

type {{.Name}}Model struct {
	DB *gorm.DB
}

func Get{{.Name}}Model() *{{.Name}}Model {
	{{.LowerCamelName}}ModelOnce.Do(func() {
		postgresConfig, err := NewPostgresConfig()
		if err != nil {
			logs.ELog("Init postgres config err: ", err)
			panic(err)
		}

		if surveyModel == nil {
			surveyModel, err = New{{.Name}}Model(postgresConfig.URL)
			if err != nil {
				logs.ELog("Init survey model err: ", err)
				panic(err)
			}
		}
	})

	return surveyModel
}

func New{{.Name}}Model(dsn string) (*{{.Name}}Model, error) {
	if db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{
		DisableForeignKeyConstraintWhenMigrating: true,
	}); err != nil {
		return nil, err
	}

	if err := db.AutoMigrate(&{{.GoPackageName}}.{{.Name}}{}); err != nil {
		return nil, err
	}

	return &{{.Name}}Model{
		DB: db,
	}, nil
}

func (m *{{.Name}}Model) Create({{.LowerCamelName}} *{{.GoPackageName}}.{{.Name}}) (int64, error) {
	result := m.DB.Create({{.LowerCamelName}})
	return result.RowsAffected, result.Error
}

func (m *{{.Name}}Model) Get(id string) (*{{.GoPackageName}}.{{.Name}}, error) {
	{{.LowerCamelName}} := &{{.GoPackageName}}.{{.Name}}{}
	if err := m.DB.First(survey, "id = ?", id).Error; err != nil {
		return nil, err
	}
	return {{.LowerCamelName}}, nil
}

func (m *{{.Name}}Model) Delete(id string) (int64, error) {
	result := m.DB.Where("id = ?", uid).Delete(&{{.GoPackageName}}.{{.Name}}{})
	return result.RowsAffected, result.Error
}

func (m *{{.Name}}Model) Update({{.LowerCamelName}} *{{.GoPackageName}}.{{.Name}}) (int64, error) {
	result := m.DB.Updates({{.LowerCamelName}})
	return result.RowsAffected, result.Error
}

func (m *{{.Name}}Model) List(filter string, fieldMask FieldMask) ([]*{{.GoPackageName}}.{{.Name}}, error) {
	var {{Plural .LowerCamelName}} []*{{.GoPackageName}}.{{.Name}}
	if err := m.DB.Find(&{{Plural .LowerCamelName}}).Error; err != nil {
		return nil, err
	}
	return {{Plural .LowerCamelName}}, nil
}

func (m *{{.Name}}Model) BatchCreate({{Plural .LowerCamelName}} ...*{{.GoPackageName}}.{{.Name}}) (int64, error) {
	result := m.DB.CreateInBatches({{Plural .LowerCamelName}}, len({{Plural .LowerCamelName}}))
	return result.RowsAffected, result.Error
}

func (m *{{.Name}}Model) BatchGet(ids ...string) (*{{.GoPackageName}}.{{.Name}}, error) {
	var {{Plural .LowerCamelName}} []{{.GoPackageName}}.{{.Name}}
	if err := m.DB.Find(&{{Plural .LowerCamelName}}, "id = ?", ids).Error; err != nil {
		return nil, err
	}
	return {{Plural .LowerCamelName}}, nil
}

func (m *{{.Name}}Model) BatchDelete(ids ...string) (int64, error) {
	result := m.DB.Where("id = ?", ids).Delete(&{{.GoPackageName}}.{{.Name}}{})
	return result.RowsAffected, result.Error
}

func (m *{{.Name}}Model) BatchUpdate({{Plural .LowerCamelName}} []*{{.GoPackageName}}.{{.Name}}) (int64, error) {
	result := m.DB.Updates({{Plural .LowerCamelName}}...)
	return result.RowsAffected, result.Error
}
`
