package testing

import (
	"ModEd/project/controller"
	"ModEd/project/model"
	"errors"
	"testing"

	"gorm.io/gorm"
)

func TestListAllAdvisorScores(t *testing.T) {
	db, _, _, _, _, dbName := Init()
	t.Cleanup(func() { cleanup(dbName) })

	score := model.ScoreAssignmentAdvisor{}
	if err := db.Create(&score).Error; err != nil {
		t.Fatalf("Failed to create score: %v", err)
	}

	advisorCtrl := controller.NewScoreAdvisorController(db)
	scores, err := advisorCtrl.ListAllAdvisorScores("assignment")
	if err != nil || len(scores.([]model.ScoreAssignmentAdvisor)) == 0 {
		t.Errorf("Expected scores, got error: %v", err)
	}
}

func TestRetrieveAdvisorScore(t *testing.T) {
	db, _, _, _, _, dbName := Init()
	t.Cleanup(func() { cleanup(dbName) })

	score := model.ScoreAssignmentAdvisor{}
	if err := db.Create(&score).Error; err != nil {
		t.Fatalf("Failed to create score: %v", err)
	}

	advisorCtrl := controller.NewScoreAdvisorController(db)
	res, err := advisorCtrl.RetrieveAdvisorScore("assignment", score.ID)
	if err != nil || res == nil {
		t.Errorf("Failed to retrieve score: %v", err)
	}
}

func TestInsertAdvisorScore(t *testing.T) {
	db, _, _, _, _, dbName := Init()
	t.Cleanup(func() { cleanup(dbName) })

	score := model.ScoreAssignmentAdvisor{}
	advisorCtrl := controller.NewScoreAdvisorController(db)
	err := advisorCtrl.InsertAdvisorScore(&score)
	if err != nil {
		t.Errorf("Failed to insert score: %v", err)
	}
}

func TestUpdateAdvisorScore(t *testing.T) {
	db, _, _, _, _, dbName := Init()
	t.Cleanup(func() { cleanup(dbName) })

	score := model.ScoreAssignmentAdvisor{Score: 80}
	if err := db.Create(&score).Error; err != nil {
		t.Fatalf("Failed to create score: %v", err)
	}

	score.Score = 90
	advisorCtrl := controller.NewScoreAdvisorController(db)
	err := advisorCtrl.UpdateAdvisorScore(&score)
	if err != nil {
		t.Errorf("Failed to update score: %v", err)
	}

	var updatedScore model.ScoreAssignmentAdvisor
	if err := db.First(&updatedScore, score.ID).Error; err != nil {
		t.Fatalf("Failed to retrieve updated score: %v", err)
	}

	expectedScore := 90.0
	if updatedScore.Score != expectedScore {
		t.Errorf("Expected score %f but got %f", expectedScore, updatedScore.Score)
	}
}

func TestDeleteAdvisorScore(t *testing.T) {
	db, _, _, _, _, dbName := Init()
	t.Cleanup(func() { cleanup(dbName) })

	score := model.ScoreAssignmentAdvisor{}
	if err := db.Create(&score).Error; err != nil {
		t.Fatalf("Failed to create score: %v", err)
	}

	advisorCtrl := controller.NewScoreAdvisorController(db)
	err := advisorCtrl.DeleteAdvisorScore(&model.ScoreAssignmentAdvisor{}, score.ID)
	if err != nil {
		t.Errorf("Failed to delete score: %v", err)
	}

	var deletedScore model.ScoreAssignmentAdvisor
	err = db.First(&deletedScore, score.ID).Error
	if !errors.Is(err, gorm.ErrRecordNotFound) {
		t.Errorf("Expected record to be deleted, but it still exists")
	}
}
