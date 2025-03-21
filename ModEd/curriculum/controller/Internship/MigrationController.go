package controller

import (
	modelCommon "ModEd/common/model"
	model "ModEd/curriculum/model/Internship"
	"errors"

	"gorm.io/gorm"
)

type MigrationController struct {
	Db *gorm.DB
}

// Comment from group 2 maybe move all of this migrate to "ModEd/curriculum/controller/MigrationController.go"
func (c *MigrationController) MigrateToDB() error {
	err := c.Db.AutoMigrate(
		&model.InternStudent{},
		&model.Company{},
		&model.InternshipSchedule{},
		&model.SupervisorReview{},
		&modelCommon.Student{},
		&model.InternshipReport{},
		&model.InternshipApplication{},
	)
	if err != nil {
		return errors.New("err: migration failed")
	}

	return nil
}
