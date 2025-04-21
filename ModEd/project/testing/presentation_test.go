package testing

import (
	"ModEd/project/model"
	"testing"

	"gorm.io/gorm"
)

func TestListAllPresentations(t *testing.T) {
	db, _, _, presentationCtrl, _, dbName := Init()
	t.Cleanup(func() { cleanup(dbName) })

	presentation := model.Presentation{}
	if err := db.Create(&presentation).Error; err != nil {
		t.Error(err)
		return
	}

	presentations, err := presentationCtrl.ListAllPresentations()
	if err != nil || len(presentations) == 0 {
		t.Errorf("Expected presentations, got error: %v", err)
	}
}

func TestRetrievePresentation(t *testing.T) {
	db, _, _, presentationCtrl, _, dbName := Init()
	t.Cleanup(func() { cleanup(dbName) })

	presentation := model.Presentation{}
	if err := db.Create(&presentation).Error; err != nil {
		t.Fatalf("Failed to create presentation: %v", err)
	}

	res, err := presentationCtrl.RetrievePresentation(presentation.ID)
	if err != nil || res == nil {
		t.Errorf("Failed to retrieve presentation: %v", err)
	}
}

func TestInsertPresentation(t *testing.T) {
	_, _, _, presentationCtrl, _, dbName := Init()
	t.Cleanup(func() { cleanup(dbName) })

	presentation := model.Presentation{}
	err := presentationCtrl.InsertPresentation(&presentation)
	if err != nil {
		t.Errorf("Failed to insert presentation: %v", err)
	}
}

func TestDeletePresentation(t *testing.T) {
	db, _, _, presentationCtrl, _, dbName := Init()
	t.Cleanup(func() { cleanup(dbName) })

	presentation := model.Presentation{Model: gorm.Model{ID: 1}}
	db.Create(&presentation)

	err := presentationCtrl.DeletePresentation(1)
	if err != nil {
		t.Errorf("Failed to delete presentation: %v", err)
	}
}
