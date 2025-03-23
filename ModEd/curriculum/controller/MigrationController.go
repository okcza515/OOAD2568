package controller

import (
	// modelCommon "ModEd/common/model"
	modelCurriculum "ModEd/curriculum/model"
	// modelInternShip "ModEd/curriculum/model/Internship"

	"github.com/cockroachdb/errors"
	"gorm.io/gorm"
)

type MigrationController struct {
	Db *gorm.DB
}

func NewMigrationController(db *gorm.DB) *MigrationController {
	return &MigrationController{Db: db}
}

func (c *MigrationController) MigrateToDB() error {
	err := c.Db.AutoMigrate(
		&modelCurriculum.Curriculum{},
		&modelCurriculum.Class{},
		&modelCurriculum.Course{},
		// &modelInternShip.InternStudent{},
		// &modelInternShip.Company{},
		// &modelInternShip.InternshipSchedule{},
		// &modelInternShip.SupervisorReview{},
		// &modelCommon.Student{},
		// &modelInternShip.InternshipReport{},
		// &modelInternShip.InternshipApplication{},
	)
	if err != nil {
		return errors.Wrap(err, "failed to migrate to db")
	}

	return nil
}
