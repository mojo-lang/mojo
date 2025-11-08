package geom

import (
	"github.com/stretchr/testify/assert"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"os"
	"path/filepath"
	"testing"
)

type GeometryTest struct {
	Id       string
	Geometry *Geometry
}

// todo test mysql
func TestGeometry(t *testing.T) {
	//testGeometryWithDialector(t, postgres.Open("user=postgres password=postgres dbname=postgres port=15432 host=172.3.0.16 sslmode=disable TimeZone=Asia/Shanghai"))
	testGeometryWithDialector(t, sqlite.Open(filepath.Join(os.TempDir(), "gorm.db")))
}

func testGeometryWithDialector(t *testing.T, dialector gorm.Dialector) {
	db, err := gorm.Open(dialector, &gorm.Config{})
	if err != nil {
		return
	}

	_ = db.Migrator().DropTable(&GeometryTest{})
	err = db.AutoMigrate(&GeometryTest{})
	if err != nil {
		return
	}

	geometry := GeometryTest{
		Id:       "1",
		Geometry: NewPointGeometryFrom(121.1, 34.2),
	}
	db.Save(&geometry)

	var count int64
	g1 := &GeometryTest{}
	assert.NoError(t, db.Select("id", "geometry").First(g1).Count(&count).Error)
	assert.Equal(t, "1", g1.Id)
	assert.Equal(t, 121.1, g1.Geometry.GetPoint().GetLongitude())
	assert.Equal(t, 34.2, g1.Geometry.GetPoint().GetLatitude())

	g2 := &GeometryTest{}
	assert.NoError(t, db.First(g2).Count(&count).Error)
	assert.Equal(t, "1", g2.Id)
	assert.Equal(t, 121.1, g2.Geometry.GetPoint().GetLongitude())
	assert.Equal(t, 34.2, g2.Geometry.GetPoint().GetLatitude())
}
