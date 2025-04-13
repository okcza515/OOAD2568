// MEP-1013
package spacemanagement

import (
	model "ModEd/asset/model/spacemanagement"
	"errors"

	"gorm.io/gorm"
)

type MigrationController struct {
	db *gorm.DB
}

func (c *MigrationController) MigrateToDB() error {
	err := c.db.AutoMigrate(
		&model.InstrumentManagement{},
		&model.SupplyManagement{},
		&model.Booking{},
		//It's not working at this moment, dependencies are not resolved
		// &model.PermanentSchedule{},
		&model.Room{},
	)
	if err != nil {
		return errors.New("migration failed")
	}
	return nil
}
