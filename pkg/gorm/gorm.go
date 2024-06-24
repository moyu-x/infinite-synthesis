package gorm

import (
	"github.com/rs/zerolog"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"

	"github.com/moyu-x/infinite-synthesis/pkg/config"
	gormZerolog "github.com/moyu-x/infinite-synthesis/third_party/gorm-zerolog"
)

func NewGORM(c *config.Bootstrap, l *zerolog.Logger) *gorm.DB {
	switch c.Database.Type {
	case "sqlite":
		return newSqlite3(c, l)
	default:
		panic("no database config")
	}
}

func newSqlite3(c *config.Bootstrap, l *zerolog.Logger) *gorm.DB {
	db, err := gorm.Open(sqlite.Open(c.Database.Uri), &gorm.Config{
		Logger: gormZerolog.NewWithLogger(l),
	})
	if err != nil {
		panic("failed to connect sqlite database")
	}
	return db
}
