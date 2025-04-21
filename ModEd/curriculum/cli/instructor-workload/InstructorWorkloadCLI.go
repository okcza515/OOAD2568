package main

import (
	controller "ModEd/curriculum/controller/instructor-workload"
	instructorWorkloadModel "ModEd/curriculum/model/instructor-workload"
	projectModel "ModEd/project/model"
	"fmt"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func main() {
	var (
		database string
		// path     string
	)

	// flag.StringVar(&database, "database", "", "Path of SQLite Database.")
	// flag.StringVar(&path, "path", "", "Path to CSV or JSON for student registration.")
	// flag.Parse()

	connector, err := gorm.Open(sqlite.Open(database), &gorm.Config{})
	// if err != nil {
	// 	panic("failed to connect database")
	// }

	// if _, err := os.Stat(path); errors.Is(err, os.ErrNotExist) {
	// 	fmt.Printf("*** Error: %s does not exist.\n", path)
	// 	return
	// }

	// migrationController := controller.NewMigrationController(connector)
	// err = migrationController.MigrateToDB()
	// if err != nil {
	// 	panic("err: migration failed")
	// }

	fmt.Println("Migration successful")

	projectController := controller.NewProjectController(connector)
	criteria := []projectModel.AssessmentCriteria{
		{AssessmentCriteriaId: 1, CriteriaName: "Quality of Work"},
		{AssessmentCriteriaId: 2, CriteriaName: "Timeliness"},
	}

	eval := &instructorWorkloadModel.ProjectEvaluation{
		GroupId:        1,
		AssignmentId:   1,
		AssignmentType: "presentation",
		Score:          0,
		Comment:        "Good job!",
	}

	err = projectController.CreateEvaluation(eval, "presentation", criteria)
	if err != nil {
		fmt.Println("Error creating evaluation:", err)
	} else {
		fmt.Println("Evaluation created successfully")
	}
}
