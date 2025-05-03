// MEP-1013
package controller

import (
	"ModEd/asset/model"
	"ModEd/core"
	"ModEd/core/migration"
	"ModEd/utils/deserializer"
	"errors"

	"gorm.io/gorm"
)

type SpaceManagementControllerManager struct {
	db *gorm.DB
	InstrumentManagement 	InstrumentManagementInterface
	SupplyManagement     	SupplyManagementInterface
	Booking              	BookingControllerInterface
	// PermanentSchedule    PermanentBookingController
	Room RoomControllerInterface
}

var spaceManagementInstance *SpaceManagementControllerManager

func GetSpaceManagementInstance(db *gorm.DB) *SpaceManagementControllerManager {
	if spaceManagementInstance != nil {
		return spaceManagementInstance
	}
	spaceManagementInstance, err := NewSpaceManagementControllerFacade(db)
	if err != nil {
		panic("failed to create SpaceManagementControllerManager instance")
	}
	return spaceManagementInstance
}

func NewSpaceManagementControllerFacade(db *gorm.DB) (*SpaceManagementControllerManager, error) {
	// db := migration.GetInstance().DB
	if db == nil {
		return nil, errors.New("db not initialized")
	}

	manager := &SpaceManagementControllerManager{
		db: db,
	}
	// facade.InstrumentManagement = InstrumentManagementController{db: db}
	// facade.SupplyManagement = SupplyManagementController{db: db}
	// facade.Booking = BookingController{db: db}
	// facade.PermanentSchedule = *NewPermanentBookingController(db, core.NewBaseController[model.PermanentSchedule](db))
	manager.Booking = NewBookingController()
	manager.Room = NewRoomController()
	manager.InstrumentManagement = NewInstrumentManagementController()
	manager.SupplyManagement = NewSupplyManagementController()
	return manager, nil
}

func (manager *SpaceManagementControllerManager) ResetDatabase() error {
	db := migration.GetInstance().DropAllTables()
	if db != nil {
		return errors.New("db not initialized")
	}

	_, err := migration.GetInstance().MigrateModule(core.MODULE_SPACEMANAGEMENT).BuildDB()
	if err != nil {
		return errors.New("failed to migrate database")
	}

	return nil
}

func (manager *SpaceManagementControllerManager) LoadSeedData() error {
	seedData := map[string]interface{}{
		"Room":    &[]model.Room{},
		"Booking": &[]model.Booking{},
	}
	for filename, m := range seedData {
		fd, err := deserializer.NewFileDeserializer("data/asset/" + filename + ".JSON")
		if err != nil {
			return err
		}

		err = fd.Deserialize(m)
		if err != nil {
			return err
		}

		result := migration.GetInstance().DB.Create(m)
		if result.Error != nil {
			return result.Error
		}
	}
	return nil
}

func (manager *SpaceManagementControllerManager) ResetDB() error {
	err := migration.GetInstance().DropAllTables()
	if err != nil {
		return err
	}

	_, err = migration.GetInstance().MigrateModule(core.MODULE_SPACEMANAGEMENT).BuildDB()
	if err != nil {
		return err
	}
	return nil
}
