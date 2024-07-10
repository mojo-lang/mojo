package templates

const EntityModel = `
// Code generated by mojo. DO NOT EDIT.
// Rerunning mojo will overwrite this file. {{if IsMojoPackage .PackageName}}
//
// Copyright 2021 Mojo-lang.org
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain m copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.{{end}}

package model

import (
	"sync"

	"github.com/ncraft-io/ncraft/go/pkg/ncraft/logs"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"

	{{.Import}}
)

var {{.LowerCamelName}} *{{.Name}}
var {{.LowerCamelName}}Once sync.Once

type {{.Name}} struct {
	*gorm.DB
}

func Get{{.Name}}() *{{.Name}} {
	{{.LowerCamelName}}Once.Do(func() {
		{{.LowerCamelName}} = New{{.Name}}()
	})

	return {{.LowerCamelName}}
}

func New{{.Name}}() (*{{.Name}}, error) {
	m := &{{.Name}}{
		DB: GetDB(),
	}
	if !m.DB.Config.DisableAutoMigrate || !d.Migrator().HasTable(&{{.GoPackageName}}.{{.Name}}{}) {
		if err := d.AutoMigrate(&{{.GoPackageName}}.{{.Name}}{}); err != nil {
			logs.ErrLog("Init {{.Name}} model err: ", err)
			panic(err)
	}
	}

	return m, nil 
}

func (m *{{.Name}}) Create(ctx context.Context, {{.LowerCamelName}} *{{.GoPackageName}}.{{.Name}}) (int64, error) {
	result := m.DB.Create({{.LowerCamelName}})
	return result.RowsAffected, result.Error
}

func (m *{{.Name}}) Get(ctx context.Context, id string) (*{{.GoPackageName}}.{{.Name}}, error) {
	{{.LowerCamelName}} := &{{.GoPackageName}}.{{.Name}}{}
	if err := m.DB.First(survey, "id = ?", id).Error; err != nil {
		return nil, err
	}
	return {{.LowerCamelName}}, nil
}

func (m *{{.Name}}) Delete(ctx context.Context, id string) (int64, error) {
	result := m.DB.Where("id = ?", uid).Delete(&{{.GoPackageName}}.{{.Name}}{})
	return result.RowsAffected, result.Error
}

func (m *{{.Name}}) Update(ctx context.Context, {{.LowerCamelName}} *{{.GoPackageName}}.{{.Name}}) (int64, error) {
	result := m.DB.Updates({{.LowerCamelName}})
	return result.RowsAffected, result.Error
}

func (m *{{.Name}}) List(ctx context.Context, query *db.Query) ([]*{{.GoPackageName}}.{{.Name}}, error) {
	var {{Plural .LowerCamelName}} []*{{.GoPackageName}}.{{.Name}}
	
	tx := m.WithContext(ctx)
	if err := query.Apply(m.DB); err != nil {
		return nil, err
	}
	return {{Plural .LowerCamelName}}, tx.Find(&{{Plural .LowerCamelName}}).Error
}

func (m *{{.Name}}) BatchCreate(ctx context.Context, {{Plural .LowerCamelName}} ...*{{.GoPackageName}}.{{.Name}}) (int64, error) {
	result := m.DB.CreateInBatches({{Plural .LowerCamelName}}, len({{Plural .LowerCamelName}}))
	return result.RowsAffected, result.Error
}

func (m *{{.Name}}) BatchGet(ctx context.Context, ids ...string) (*{{.GoPackageName}}.{{.Name}}, error) {
	var {{Plural .LowerCamelName}} []{{.GoPackageName}}.{{.Name}}
	if err := m.DB.Find(&{{Plural .LowerCamelName}}, "id = ?", ids).Error; err != nil {
		return nil, err
	}
	return {{Plural .LowerCamelName}}, nil
}

func (m *{{.Name}}) BatchDelete(ctx context.Context, ids ...string) (int64, error) {
	result := m.DB.Where("id = ?", ids).Delete(&{{.GoPackageName}}.{{.Name}}{})
	return result.RowsAffected, result.Error
}

func (m *{{.Name}}) BatchUpdate(ctx context.Context, {{Plural .LowerCamelName}} []*{{.GoPackageName}}.{{.Name}}) (int64, error) {
	result := m.DB.Updates({{Plural .LowerCamelName}}...)
	return result.RowsAffected, result.Error
}
`
