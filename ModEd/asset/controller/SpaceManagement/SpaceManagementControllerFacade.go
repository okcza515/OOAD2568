// MEP-1013
package spacemanagement

import (
	"errors"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

type SpaceManagementControllerFacade struct {
	Db                			*gorm.DB
	InstrumentManagement   		InstrumentManagementController
	SupplyManagementController 	SupplyManagementController
	Booking           			BookingController
	PermanentSchedule 			PermanentScheduleController
	Room              			RoomController
}

func CreateSpaceManagementControllerFacade() (*SpaceManagementControllerFacade, error) {
	db, err := gorm.Open(sqlite.Open("data/ModEd.bin"), &gorm.Config{})
	if err != nil {
		return nil, err
	}
	facade := SpaceManagementControllerFacade{Db: db}
	migrationController := MigrationController{db: db}
	err = migrationController.MigrateToDB()
	if err != nil {
		return nil, errors.New("Failed to migrate schema")
	}
	facade.InstrumentManagement = InstrumentManagementController{db: db}
	facade.SupplyManagementController = SupplyManagementController{db: db}
	facade.Booking = BookingController{db: db}
	facade.PermanentSchedule = PermanentScheduleController{db: db}
	facade.Room = RoomController{db: db}
	return &facade, nil

}
