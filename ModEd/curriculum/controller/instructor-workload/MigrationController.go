package controller

import (
	modelInstructureWorkLoad "ModEd/curriculum/model/instructor-workload"
	"fmt"

	"github.com/cockroachdb/errors"
	"gorm.io/gorm"
)

type MigrationController struct {
	Db *gorm.DB
}

func NewMigrationController(db *gorm.DB) *MigrationController {
	return &MigrationController{Db: db}
}

// Comment from group 2 maybe move all of this migrate to "ModEd/curriculum/controller/MigrationController.go"
func (c *MigrationController) MigrateToDB() error {
	err := c.Db.AutoMigrate(
		// &modelCurriculum.Curriculum{},
		// &modelCurriculum.Class{},
		// &modelCurriculum.Course{},
		// Instructor Workload
		&modelInstructureWorkLoad.AssignedCourse{},
		&modelInstructureWorkLoad.ClassMaterial{},
		&modelInstructureWorkLoad.ClassSchedule{},
		&modelInstructureWorkLoad.CourseUpdateRequest{},
		&modelInstructureWorkLoad.CourseNameUpdate{},
		&modelInstructureWorkLoad.CoursePrerequisiteUpdate{},
		&modelInstructureWorkLoad.ProjectEvaluation{},
		&modelInstructureWorkLoad.StudentAdvisor{},
		&modelInstructureWorkLoad.StudentRequest{},
	)
	if err != nil {
		return errors.Wrap(err, "failed to migrate to db")
	}

	// Log success or any additional info
	fmt.Println("Migration completed successfully")
	return nil
}
