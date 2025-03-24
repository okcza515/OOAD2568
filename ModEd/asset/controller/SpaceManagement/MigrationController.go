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
		&model.AssetManagement{},
		&model.Booking{},
		&model.PermanentSchedule{},
		&model.Room{},
	)
	if err != nil {
		return errors.New("Migration failed")
	}
	return nil
}
