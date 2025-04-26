package controller

import (
	// modelCommon "ModEd/common/model"
	modelCommon "ModEd/common/model"
	modelCurriculum "ModEd/curriculum/model"
	modelInstructorWorkload "ModEd/curriculum/model"
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
		&modelCommon.Student{},
		&modelCurriculum.Curriculum{},
		&modelCurriculum.Class{},
		&modelCurriculum.Course{},
		&modelCurriculum.WILProjectCourse{},
		&modelCurriculum.WILProjectClass{},
		&modelCurriculum.WILProjectMember{},
		&modelCurriculum.WILProjectApplication{},
		&modelCurriculum.WILProject{},
		&modelCurriculum.IndependentStudy{},
		&modelInternShip.InternStudent{},
		&modelInternShip.Company{},
		&modelInternShip.SupervisorReview{},
		&modelInternShip.InternshipReport{},
		&modelInternShip.InternshipApplication{},
		&modelInstructorWorkload.ClassLecture{},
		&modelInstructorWorkload.ClassMaterial{},
		&modelInstructorWorkload.CoursePlan{},
		&modelInstructorWorkload.StudentAdvisor{},
		&modelInstructorWorkload.StudentRequest{},
		&modelInstructorWorkload.Meeting{},
		&modelInstructorWorkload.ProjectEvaluation{},
	)
	if err != nil {
		return errors.Wrap(err, "failed to migrate to db")
	}

	return nil
}

func (c *MigrationController) DropAllTables() error {
	err := c.Db.Migrator().DropTable(
		&modelCommon.Student{},
		&modelCurriculum.Curriculum{},
		&modelCurriculum.Class{},
		&modelCurriculum.Course{},
		&modelCurriculum.WILProjectCourse{},
		&modelCurriculum.WILProjectClass{},
		&modelCurriculum.WILProjectMember{},
		&modelCurriculum.WILProjectApplication{},
		&modelCurriculum.WILProject{},
		&modelCurriculum.IndependentStudy{},
		&modelInternShip.InternStudent{},
		&modelInternShip.Company{},
		&modelInternShip.SupervisorReview{},
		&modelInternShip.InternshipReport{},
		&modelInternShip.InternshipApplication{},
		&modelInstructorWorkload.ClassLecture{},
		&modelInstructorWorkload.ClassMaterial{},
		&modelInstructorWorkload.CoursePlan{},
		&modelInstructorWorkload.StudentAdvisor{},
		&modelInstructorWorkload.StudentRequest{},
		&modelInstructorWorkload.Meeting{},
		&modelInstructorWorkload.ProjectEvaluation{},
	)
	if err != nil {
		return err
	}

	return nil
}
