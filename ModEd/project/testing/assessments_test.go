package testing

import (
	"ModEd/project/model"
	"testing"

	"gorm.io/gorm"
)

func TestListAllAssessments(t *testing.T) {
	db, assessmentCtrl, _, _, _, dbName := Init()
	t.Cleanup(func() { cleanup(dbName) })

	assessment := model.Assessment{}
	if err := db.Create(&assessment).Error; err != nil {
		t.Error(err)
		return
	}

	assessments, err := assessmentCtrl.ListAllAssessments()
	if err != nil || len(assessments) == 0 {
		t.Errorf("Expected assessments, got error: %v", err)
	}
}

func TestRetrieveAssessment(t *testing.T) {
	db, assessmentCtrl, _, _, _, dbName := Init()
	t.Cleanup(func() { cleanup(dbName) })

	assessment := model.Assessment{}
	if err := db.Create(&assessment).Error; err != nil {
		t.Fatalf("Failed to create assessment: %v", err)
	}

	res, err := assessmentCtrl.RetrieveAssessment(assessment.ID)
	if err != nil || res == nil {
		t.Errorf("Failed to retrieve assessment: %v", err)
	}
}

func TestInsertAssessment(t *testing.T) {
	_, assessmentCtrl, _, _, _, dbName := Init()
	t.Cleanup(func() { cleanup(dbName) })

	assessment := model.Assessment{}
	err := assessmentCtrl.InsertAssessment(&assessment)
	if err != nil {
		t.Errorf("Failed to insert assessment: %v", err)
	}
}

func TestDeleteAssessment(t *testing.T) {
	db, assessmentCtrl, _, _, _, dbName := Init()
	t.Cleanup(func() { cleanup(dbName) })

	assessment := model.Assessment{Model: gorm.Model{ID: 1}}
	db.Create(&assessment)

	err := assessmentCtrl.DeleteAssessment(1)
	if err != nil {
		t.Errorf("Failed to delete assessment: %v", err)
	}
}
