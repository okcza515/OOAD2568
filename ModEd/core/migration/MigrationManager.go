package migration

// Wrote By : MEP-1010, MEP-1012

import (
	"ModEd/core"
	"github.com/cockroachdb/errors"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

var instance *MigrationManager = &MigrationManager{}

type MigrationManager struct {
	DB     *gorm.DB
	err    error
	pathDB string
	Models []interface{}
}

func GetInstance() *MigrationManager {
	return instance
}

func (m *MigrationManager) SetPathDB(pathDB string) *MigrationManager {
	m.pathDB = pathDB
	return m
}

func (m *MigrationManager) BuildDB() (*gorm.DB, error) {
	if m.err != nil {
		return nil, m.err
	}

	dbPath := ""
	defaultPath := "data/ModEd.bin"

	if m.pathDB == "" {
		dbPath = defaultPath
	}

	db, err := gorm.Open(sqlite.Open(dbPath), &gorm.Config{})
	if err != nil {
		return nil, errors.New("err: failed to connect database")
	}

	m.DB = db

	err = m.migrateToDB()
	if err != nil {
		return nil, err
	}

	return m.DB, nil
}

func (m *MigrationManager) MigrateModule(module core.ModuleOptionEnum) *MigrationManager {
	var strategy MigrationStrategy

	switch module {
	case core.MODULE_ASSET:
		strategy = &AssetMigrationStrategy{}
	case core.MODULE_PROCUREMENT:
		panic("not implemented")
	case core.MODULE_SPACEMANAGEMENT:
		panic("not implemented")
	case core.MODULE_COMMON:
		panic("not implemented")
	case core.MODULE_CURRICULUM:
		strategy = &CurriculumMigrationStrategy{}
	case core.MODULE_INSTRUCTOR:
		panic("not implemented")
	case core.MODULE_INTERNSHIP:
		panic("not implemented")
	case core.MODULE_WILPROJECT:
		strategy = &WILProjectMigrationStrategy{}
	case core.MODULE_QUIZ:
		panic("not implemented")
	case core.MODULE_EVAL:
		panic("not implemented")
	case core.MODULE_HR:
		panic("not implemented")
	case core.MODULE_PROJECT:
		panic("not implemented")
	case core.MODULE_RECRUIT:
		panic("not implemented")
	default:
		return m
	}

	m.Models = append(m.Models, strategy.GetModels()...)
	return m
}

func (m *MigrationManager) migrateToDB() error {
	err := m.DB.AutoMigrate(m.Models...)
	if err != nil {
		return errors.Wrap(err, "failed to migrate to db")
	}
	return nil
}

func (m *MigrationManager) DropAllTables() error {

	if m.DB == nil {
		return errors.New("db not initialize")
	}

	err := m.DB.Migrator().DropTable(m.Models...)
	if err != nil {
		return err
	}
	return nil
}
