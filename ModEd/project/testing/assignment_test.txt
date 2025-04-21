package testing

import (
	"ModEd/project/model"
	"testing"

	"gorm.io/gorm"
)

func TestListAllAssignments(t *testing.T) {
	db, _, assignmentCtrl, _, _, dbName := Init()
	t.Cleanup(func() { cleanup(dbName) })

	assignment := model.Assignment{}
	if err := db.Create(&assignment).Error; err != nil {
		t.Error(err)
		return
	}

	assignments, err := assignmentCtrl.ListAllAssignments()
	if err != nil || len(assignments) == 0 {
		t.Errorf("Expected assignments, got error: %v", err)
	}
}

func TestRetrieveAssignment(t *testing.T) {
	db, _, assignmentCtrl, _, _, dbName := Init()
	t.Cleanup(func() { cleanup(dbName) })

	assignment := model.Assignment{}
	if err := db.Create(&assignment).Error; err != nil {
		t.Fatalf("Failed to create assignment: %v", err)
	}

	res, err := assignmentCtrl.RetrieveAssignment(assignment.ID)
	if err != nil || res == nil {
		t.Errorf("Failed to retrieve assignment: %v", err)
	}
}

func TestInsertAssignment(t *testing.T) {
	_, _, assignmentCtrl, _, _, dbName := Init()
	t.Cleanup(func() { cleanup(dbName) })

	assignment := model.Assignment{}
	err := assignmentCtrl.InsertAssignment(&assignment)
	if err != nil {
		t.Errorf("Failed to insert assignment: %v", err)
	}
}

func TestDeleteAssignment(t *testing.T) {
	db, _, assignmentCtrl, _, _, dbName := Init()
	t.Cleanup(func() { cleanup(dbName) })

	assignment := model.Assignment{Model: gorm.Model{ID: 1}}
	db.Create(&assignment)

	err := assignmentCtrl.DeleteAssignment(1)
	if err != nil {
		t.Errorf("Failed to delete assignment: %v", err)
	}
}
