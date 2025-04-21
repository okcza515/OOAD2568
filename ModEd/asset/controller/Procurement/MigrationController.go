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

func (c *MigrationController) migrateToDB() error {
	err := c.db.AutoMigrate(&model.InstrumentRequest{})

	if err != nil {
		return errors.New("err: migration failed")
	}

	return nil
}
