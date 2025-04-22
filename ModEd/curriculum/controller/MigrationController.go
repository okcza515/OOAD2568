package controller

import (
	// modelCommon "ModEd/common/model"
	modelCommon "ModEd/common/model"
	"ModEd/curriculum/model"
	modelCurriculum "ModEd/curriculum/model"
	modelInternShip "ModEd/curriculum/model"

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
		&model.WILProjectCourse{},
		&model.WILProjectClass{},
		&model.WILProjectMember{},
		&model.WILProjectApplication{},
		&model.WILProject{},
		&model.IndependentStudy{},
		&modelInternShip.InternStudent{},
		&modelInternShip.Company{},
		&modelInternShip.SupervisorReview{},
		&modelCommon.Student{},
		&modelInternShip.InternshipReport{},
		&modelInternShip.InternshipApplication{},
	)
	if err != nil {
		return errors.Wrap(err, "failed to migrate to db")
	}

	return nil
}

func (c *MigrationController) DropAllTables() error {
	err := c.Db.Migrator().DropTable(
		&modelCurriculum.Curriculum{},
		&modelCurriculum.Class{},
		&modelCurriculum.Course{},
		&model.WILProjectCourse{},
		&model.WILProjectClass{},
		&model.WILProjectMember{},
		&model.WILProjectApplication{},
		&model.WILProject{},
		&model.IndependentStudy{},
		&modelInternShip.InternStudent{},
		&modelInternShip.Company{},
		&modelInternShip.SupervisorReview{},
		&modelCommon.Student{},
		&modelInternShip.InternshipReport{},
		&modelInternShip.InternshipApplication{},
	)
	if err != nil {
		return err
	}

	return nil
}
