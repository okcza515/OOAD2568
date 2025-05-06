// MEP-1013
package controller

import (
	"ModEd/asset/model"
	"ModEd/core"
	"ModEd/core/migration"
	"errors"

	"gorm.io/gorm"
)

type SpaceManagementControllerManager struct {
	db                   *gorm.DB
	InstrumentManagement InstrumentManagementInterface
	SupplyManagement     SupplyManagementInterface
	Booking              BookingControllerInterface
	PermanentSchedule    PermanentBookingControllerInterface
	Room                 RoomControllerInterface
}

var spaceManagementInstance *SpaceManagementControllerManager

func GetSpaceManagementInstance(db *gorm.DB) *SpaceManagementControllerManager {
	if spaceManagementInstance != nil {
		return spaceManagementInstance
	}
	spaceManagementInstance, err := NewSpaceManagementControllerManager(db)
	if err != nil {
		panic("failed to create SpaceManagementControllerManager instance")
	}
	return spaceManagementInstance
}

func NewSpaceManagementControllerManager(db *gorm.DB) (*SpaceManagementControllerManager, error) {
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
	manager.PermanentSchedule = NewPermanentBookingController()
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
	err := manager.ResetDatabase()
	if err != nil {
		return err
	}

	err = migration.GetInstance().AddSeedData(("data/asset/Room.JSON"), &[]model.Room{}).
		AddSeedData(("data/asset/Booking.JSON"), &[]model.Booking{}).
		AddSeedData(("data/asset/PermanentSchedule.JSON"), &[]model.PermanentSchedule{}).
		//Add more seed data as needed
		LoadSeedData()
	if err != nil {
		return err
	}

	return nil
}
