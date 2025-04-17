package testing

import (
	"ModEd/project/controller"
	"ModEd/project/model"
	"os"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func Init() (*gorm.DB, *controller.AssessmentController, *controller.AssignmentController, *controller.PresentationController, *controller.ReportController, string) {
	dbName := "test.db"
	db, _ := gorm.Open(sqlite.Open(dbName), &gorm.Config{})
	db = db.Debug()
	db.Exec("PRAGMA foreign_keys = ON;")

	if err := db.AutoMigrate(
		&model.Advisor{},
		&model.Assessment{},
		&model.AssessmentCriteria{},
		&model.Assignment{},
		&model.Committee{},
		&model.GroupMember{},
		&model.Presentation{},
		&model.Progress{},
		&model.Report{},
		&model.ScoreAssessmentAdvisor{},
		&model.ScoreAssessmentCommittee{},
		&model.ScoreAssignmentAdvisor{},
		&model.ScoreAssignmentCommittee{},
		&model.ScorePresentationAdvisor{},
		&model.ScorePresentationCommittee{},
		&model.ScoreReportAdvisor{},
		&model.ScoreReportCommittee{},
		&model.SeniorProject{},
	); err != nil {
		panic(err)
	}

	assessmentController := controller.NewAssessmentController(db)
	assignmentController := controller.NewAssignmentController(db)
	presentationController := controller.NewPresentationController(db)
	reportController := controller.NewReportController(db)

	return db, assessmentController.(*controller.AssessmentController), assignmentController.(*controller.AssignmentController), presentationController, reportController, dbName
}

func cleanup(dbName string) {
	os.Remove(dbName)
}
