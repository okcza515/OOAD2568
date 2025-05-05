package testing

import (
	"ModEd/project/controller"
	"ModEd/project/model"
	"os"
	"testing"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func Init() (*gorm.DB, *controller.AssessmentController, *controller.AssessmentCriteriaController, *controller.AssessmentCriteriaLinkController, *controller.AssignmentController, *controller.PresentationController, *controller.ReportController, string) {
	dbName := "test.db"
	db, _ := gorm.Open(sqlite.Open(dbName), &gorm.Config{})
	db = db.Debug()
	db.Exec("PRAGMA foreign_keys = ON;")

	if err := db.AutoMigrate(
		&model.SeniorProject{},
		&model.GroupMember{},
		&model.Advisor{},
		&model.Assessment{},
		&model.AssessmentCriteria{},
		&model.Assignment{},
		&model.Committee{},
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
	); err != nil {
		panic(err)
	}

	assessmentController := controller.NewAssessmentController(db)
	assignmentController := controller.NewAssignmentController(db)
	presentationController := controller.NewPresentationController(db)
	reportController := controller.NewReportController(db)
	assessmentCriteriaController := controller.NewAssessmentCriteriaController(db)
	assessmentCriteriaLinkController := controller.NewAssessmentCriteriaLinkController(db)

	return db, assessmentController, assessmentCriteriaController, assessmentCriteriaLinkController, assignmentController, presentationController, reportController, dbName
}

func cleanup(dbName string) {
	os.Remove(dbName)
}

func TestDatabase(t *testing.T) {
	Init()
}
