// MEP-1014
package procurement

import (
	model "ModEd/asset/model/Procurement"
	"errors"

	"gorm.io/gorm"
)

type MigrationController struct {
	db *gorm.DB
}

func (c *MigrationController) MigrateToDB() error {
	err := c.db.AutoMigrate(
		&model.ItemRequest{},
	)
	if err != nil {
		return errors.New("migration failed")
	}
	return nil
}
