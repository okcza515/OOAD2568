// MEP-1013
package controller

import (
	"ModEd/asset/model"
	"ModEd/core"
	"ModEd/core/migration"

	"gorm.io/gorm"
)

type SpaceManagementControllerFacade struct {
	Db                	   *gorm.DB
	InstrumentManagement    InstrumentManagementController
	SupplyManagement		SupplyManagementController
	Booking           		BookingController
	PermanentSchedule 		PermanentBookingController
	Room              		RoomController
}

func NewSpaceManagementControllerFacade() (*SpaceManagementControllerFacade, error) {
	db, err := migration.GetInstance().MigrateModule(core.MODULE_SPACEMANAGEMENT).BuildDB()
	if err != nil {
		return nil, err
	}


	facade := SpaceManagementControllerFacade{Db: db}
	facade.InstrumentManagement = InstrumentManagementController{db:db}
	facade.SupplyManagement = SupplyManagementController{db:db}
	facade.Booking = BookingController{db: db}
	facade.PermanentSchedule = PermanentBookingController{db: db}
	facade.Room = RoomController{db: db, BaseController: core.NewBaseController[model.Room](db)}
	return &facade, nil

}
