package utils

import (
	"github.com/cockroachdb/errors"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

type GormConfig struct {
	DBPath string
	Config *gorm.Config
}

func NewGormSqlite(config *GormConfig) (*gorm.DB, error) {
	db, err := gorm.Open(sqlite.Open(config.DBPath), config.Config)
	if err != nil {
		return nil, errors.Wrap(err, "failed to connect to gorm sqlite")
	}
	return db, nil
}
