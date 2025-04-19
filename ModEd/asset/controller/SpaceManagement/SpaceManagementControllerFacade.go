// MEP-1013
package spacemanagement

import (
	"errors"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

type SpaceManagementControllerFacade struct {
	Db                *gorm.DB
	AssetManagement   *AssetManagementController
	Booking           BookingController
	PermanentSchedule PermanentBookingController
	Room              RoomController
}

func CreateSpaceManagementControllerFacade() (*SpaceManagementControllerFacade, error) {
	db, err := gorm.Open(sqlite.Open("data/ModEd.bin"), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Silent), // Or logger.Error for only errors
	})
	if err != nil {
		return nil, err
	}

	InstrumentManagementAdapter := NewInstrumentManagementAdapter(db)
	SupplyManagementAdapter := NewSupplyManagementAdapter(db)

	AssetManagementController := NewAssetManagementController(InstrumentManagementAdapter, SupplyManagementAdapter)

	facade := SpaceManagementControllerFacade{Db: db}
	migrationController := MigrationController{db: db}
	err = migrationController.MigrateToDB()
	if err != nil {
		return nil, errors.New("failed to migrate schema")
	}
	facade.AssetManagement = AssetManagementController
	facade.Booking = BookingController{db: db}
	facade.PermanentSchedule = PermanentBookingController{db: db}
	facade.Room = RoomController{db: db}
	return &facade, nil

}
