package migration

import (
	model "ModEd/eval/model"

	"github.com/cockroachdb/errors"

	"gorm.io/gorm"
)

type MigrationController struct {
	db *gorm.DB
}

func NewMigrationController(db *gorm.DB) *MigrationController {
	return &MigrationController{db: db}
}

func (c *MigrationController) MigrateToDB() error {
	err := c.db.AutoMigrate(
		&model.Examination{},
		&model.Question{},
	)

	if err != nil {
		return errors.New("fail to Migrate to DB")
	}
	return nil
}
