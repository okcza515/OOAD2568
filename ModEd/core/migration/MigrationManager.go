package migration

// Wrote By : MEP-1010, MEP-1012

import (
	"ModEd/core"
	"fmt"
	"github.com/cockroachdb/errors"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

var instance = newMigrationManager()

type MigrationManager struct {
	DB                   *gorm.DB
	err                  error
	pathDB               string
	models               []interface{}
	migrationStrategyMap map[core.ModuleOptionEnum]MigrationStrategy
}

func GetInstance() *MigrationManager {
	return instance
}

func newMigrationManager() *MigrationManager {
	migrationMap := make(map[core.ModuleOptionEnum]MigrationStrategy)

	// To use the core migration module you need to create your own migration strategy
	// Then come here to replace `nil` with your model here to register
	migrationMap[core.MODULE_ASSET] = &AssetMigrationStrategy{}
	migrationMap[core.MODULE_PROCUREMENT] = nil
	migrationMap[core.MODULE_SPACEMANAGEMENT] = &SpaceManagementMigrationStrategy{}
	migrationMap[core.MODULE_COMMON] = nil
	migrationMap[core.MODULE_CURRICULUM] = &CurriculumMigrationStrategy{}
	migrationMap[core.MODULE_INSTRUCTOR] = &InstructorWorkloadMigrationStrategy{}
	migrationMap[core.MODULE_INTERNSHIP] = &InternshipMigrationStrategy{}
	migrationMap[core.MODULE_WILPROJECT] = &WILProjectMigrationStrategy{}
	migrationMap[core.MODULE_QUIZ] = nil
	migrationMap[core.MODULE_EVAL] = nil
	migrationMap[core.MODULE_HR] = nil
	migrationMap[core.MODULE_PROJECT] = nil
	migrationMap[core.MODULE_RECRUIT] = nil

	return &MigrationManager{
		migrationStrategyMap: migrationMap,
	}
}

func (m *MigrationManager) SetPathDB(pathDB string) *MigrationManager {
	m.pathDB = pathDB
	return m
}

func (m *MigrationManager) BuildDB() (*gorm.DB, error) {
	if m.err != nil {
		return nil, m.err
	}

	defaultPath := "data/ModEd.bin"
	dbPath := defaultPath

	if m.pathDB != "" {
		dbPath = m.pathDB
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
	strategy, ok := m.migrationStrategyMap[module]
	fmt.Println(strategy)
	fmt.Println(ok)
	fmt.Println(len(m.migrationStrategyMap))
	if !ok || strategy == nil {
		panic(fmt.Sprintf("err: module '%v' migration is not implemented", string(module)))
	}

	m.models = append(m.models, strategy.GetModels()...)
	return m
}

func (m *MigrationManager) migrateToDB() error {
	var modelsToMigrate []interface{}
	for i := range m.models {
		if m.DB.Migrator().HasTable(m.models[i]) {
			continue
		}

		modelsToMigrate = append(modelsToMigrate, m.models[i])
	}

	err := m.DB.AutoMigrate(modelsToMigrate...)
	if err != nil {
		return errors.Wrap(err, "failed to migrate to db")
	}
	return nil
}

func (m *MigrationManager) DropAllTables() error {

	if m.DB == nil {
		return errors.New("db not initialize")
	}

	err := m.DB.Migrator().DropTable(m.models...)
	if err != nil {
		return err
	}
	return nil
}
