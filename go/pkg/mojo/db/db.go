package db

import (
	dm "github.com/smilextay/gorm-dm"
	"gorm.io/driver/mysql"
	"gorm.io/driver/postgres"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

const (
	PostgresDriverName   = "postgres"
	MysqlDriverName      = "mysql"
	SqliteDriverName     = "sqlite"
	DmDriverName         = "dm"
	ClickHouseDriverName = "clickhouse"
)

type DB struct {
	*gorm.DB
	Config *Config
}

func New(cfg *Config) *DB {
	db := &DB{
		Config: cfg,
	}

	config := func(gdb *gorm.DB) *gorm.DB {
		if db.Config.Debug {
			gdb = gdb.Debug()
		}

		return gdb
	}

	if db.Config.Driver == PostgresDriverName {
		if d, err := gorm.Open(postgres.Open(db.Config.Dsn), &gorm.Config{}); err != nil {
			return nil
		} else {
			db.DB = d
			db.DB = config(d)
			return db
		}
	} else if db.Config.Driver == MysqlDriverName {
		if d, err := gorm.Open(mysql.New(mysql.Config{DSN: db.Config.Dsn, DefaultStringSize: uint(db.Config.DefaultStringSize)}), &gorm.Config{}); err != nil {
			return nil
		} else {
			db.DB = config(d)
			return db
		}
	} else if db.Config.Driver == SqliteDriverName {
		if d, err := gorm.Open(sqlite.Open(db.Config.Dsn), &gorm.Config{}); err != nil {
			return nil
		} else {
			db.DB = d
			db.DB = config(d)
			return db
		}
	} else if db.Config.Driver == DmDriverName {
		if d, err := gorm.Open(dm.New(dm.Config{DSN: db.Config.Dsn, DefaultStringSize: uint(db.Config.DefaultStringSize)}), &gorm.Config{
			DisableForeignKeyConstraintWhenMigrating: true,
		}); err != nil {
			return nil
		} else {
			db.DB = config(d)
			return db
		}
	} else if db.Config.Driver == ClickHouseDriverName {

	}
	return nil
}

func (d *DB) Tx(tx *gorm.DB) *DB {
	if d != nil && tx != nil {
		return &DB{DB: tx, Config: d.Config}
	}
	return d
}

func (d *DB) Model(value interface{}) *DB {
	if d != nil && value != nil {
		return &DB{DB: d.DB.Model(value), Config: d.Config}
	}
	return d
}
