package testing

import (
	"ModEd/project/model"
	"testing"
	"time"
)

func TestInsertPresentation(t *testing.T) {
	db, _, _, _, _, presentationCtrl, _, dbName := Init()
	t.Cleanup(func() { cleanup(dbName) })

	seniorProject := model.SeniorProject{GroupName: "Test Group"}
	if err := db.Create(&seniorProject).Error; err != nil {
		t.Fatalf("Failed to create senior project: %v", err)
	}

	presentation := model.Presentation{
		PresentationType: model.PresentationTypeProposal,
		Date:             time.Now(),
		SeniorProjectId:  seniorProject.ID,
	}
	err := presentationCtrl.InsertPresentation(&presentation)
	if err != nil {
		t.Errorf("Failed to insert presentation: %v", err)
	}
}

func TestListAllPresentations(t *testing.T) {
	db, _, _, _, _, presentationCtrl, _, dbName := Init()
	t.Cleanup(func() { cleanup(dbName) })

	seniorProject := model.SeniorProject{GroupName: "Test Group"}
	if err := db.Create(&seniorProject).Error; err != nil {
		t.Fatalf("Failed to create senior project: %v", err)
	}

	presentation := model.Presentation{
		PresentationType: model.PresentationTypeProposal,
		Date:             time.Now(),
		SeniorProjectId:  seniorProject.ID,
	}
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
	db, _, _, _, _, presentationCtrl, _, dbName := Init()
	t.Cleanup(func() { cleanup(dbName) })

	seniorProject := model.SeniorProject{GroupName: "Test Group"}
	if err := db.Create(&seniorProject).Error; err != nil {
		t.Fatalf("Failed to create senior project: %v", err)
	}

	presentation := model.Presentation{
		PresentationType: model.PresentationTypeProposal,
		Date:             time.Now(),
		SeniorProjectId:  seniorProject.ID,
	}
	if err := db.Create(&presentation).Error; err != nil {
		t.Fatalf("Failed to create presentation: %v", err)
	}

	res, err := presentationCtrl.RetrievePresentation(presentation.ID)
	if err != nil || res == nil {
		t.Errorf("Failed to retrieve presentation: %v", err)
	}
}

func TestDeletePresentation(t *testing.T) {
	db, _, _, _, _, presentationCtrl, _, dbName := Init()
	t.Cleanup(func() { cleanup(dbName) })

	seniorProject := model.SeniorProject{GroupName: "Test Group"}
	if err := db.Create(&seniorProject).Error; err != nil {
		t.Fatalf("Failed to create senior project: %v", err)
	}

	presentation := model.Presentation{
		PresentationType: model.PresentationTypeProposal,
		Date:             time.Now(),
		SeniorProjectId:  seniorProject.ID,
	}
	if err := db.Create(&presentation).Error; err != nil {
		t.Fatalf("Failed to create presentation: %v", err)
	}

	err := presentationCtrl.DeletePresentation(presentation.ID)
	if err != nil {
		t.Errorf("Failed to delete presentation: %v", err)
	}
}
