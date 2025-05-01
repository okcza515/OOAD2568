// MEP-1013
package controller

import (
	"ModEd/core"
	"ModEd/core/migration"
	"errors"

	"gorm.io/gorm"
)

type SpaceManagementControllerManager struct {
	db *gorm.DB
	// InstrumentManagement InstrumentManagementController
	// SupplyManagement     SupplyManagementController
	// Booking              BookingController
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
	manager.Room = NewRoomController()

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
