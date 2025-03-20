package controller

import (
	model "ModEd/curriculum/model/Internship"
	"errors"

	"gorm.io/gorm"
)

type MigrationController struct {
	Db *gorm.DB
}

func (c *MigrationController) MigrateToDB() error {
	err := c.Db.AutoMigrate(
		&model.InternStudent{},
		&model.Company{},
		&model.InternshipSchedule{},
		&model.SupervisorReview{},
		//&model.InternshipApplication{},
	)
	if err != nil {
		return errors.New("err: migration failed")
	}

	return nil
}
