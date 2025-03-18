package testing

import (
	"ModEd/project/controller"
	"ModEd/project/model"
	"os"
	"testing"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func Init() (*gorm.DB, controller.IAssessmentController, string) {
	dbName := "test.db"
	db, _ := gorm.Open(sqlite.Open(dbName), &gorm.Config{})
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

	controller := controller.NewAssessmentController(db)

	return db, controller, dbName
}

func cleanup(dbName string) {
	os.Remove(dbName)
}

func TestListAllAssessments(t *testing.T) {
	db, ctrl, dbName := Init()
	t.Cleanup(func() { cleanup(dbName) })

	assessment := model.Assessment{}
	if err := db.Create(&assessment).Error; err != nil {
		t.Error(err)
		return
	}

	assessments, err := ctrl.ListAllAssessments()
	if err != nil || len(assessments) == 0 {
		t.Errorf("Expected assessments, got error: %v", err)
	}
}

func TestRetrieveAssessment(t *testing.T) {
	db, ctrl, dbName := Init()
	t.Cleanup(func() { cleanup(dbName) })

	assessment := model.Assessment{}
	if err := db.Create(&assessment).Error; err != nil {
		t.Fatalf("Failed to create assessment: %v", err)
	}

	res, err := ctrl.RetrieveAssessment(assessment.ID)
	if err != nil || res == nil {
		t.Errorf("Failed to retrieve assessment: %v", err)
	}
}

func TestInsertAssessment(t *testing.T) {
	_, ctrl, dbName := Init()
	t.Cleanup(func() { cleanup(dbName) })

	assessment := model.Assessment{}
	err := ctrl.InsertAssessment(&assessment)
	if err != nil {
		t.Errorf("Failed to insert assessment: %v", err)
	}
}

func TestDeleteAssessment(t *testing.T) {
	db, ctrl, dbName := Init()
	t.Cleanup(func() { cleanup(dbName) })

	assessment := model.Assessment{Model: gorm.Model{ID: 1}}
	db.Create(&assessment)

	err := ctrl.DeleteAssessment(1)
	if err != nil {
		t.Errorf("Failed to delete assessment: %v", err)
	}
}
