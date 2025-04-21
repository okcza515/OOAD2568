package testing

import (
	"ModEd/project/controller"
	"ModEd/project/model"
	"errors"
	"testing"

	"gorm.io/gorm"
)

func TestListAllAdvisorScores(t *testing.T) {
	db, _, _, _, _, _, _, dbName := Init()
	t.Cleanup(func() { cleanup(dbName) })

	score := model.ScoreAssignmentAdvisor{
		AssignmentId: 1,
		AdvisorId:    2,
		Score:        85.0,
	}
	if err := db.Create(&score).Error; err != nil {
		t.Fatalf("Failed to create score: %v", err)
	}

	advisorCtrl := controller.NewScoreAdvisorController(db)
	scores, err := advisorCtrl.ListAllAdvisorScores("assignment")
	if err != nil {
		t.Errorf("Expected scores, got error: %v", err)
		return
	}

	scoreList, ok := scores.(*[]model.ScoreAssignmentAdvisor)
	if !ok || len(*scoreList) == 0 {
		t.Errorf("Expected scores, got none or incorrect type")
	}
}

func TestRetrieveAdvisorScore(t *testing.T) {
	db, _, _, _, _, _, _, dbName := Init()
	t.Cleanup(func() { cleanup(dbName) })

	score := model.ScoreAssignmentAdvisor{
		AssignmentId: 1,
		AdvisorId:    2,
		Score:        90.0,
	}
	if err := db.Create(&score).Error; err != nil {
		t.Fatalf("Failed to create score: %v", err)
	}

	t.Logf("Created ScoreAssignmentAdvisor with ID: %d", score.ID)

	advisorCtrl := controller.NewScoreAdvisorController(db)
	res, err := advisorCtrl.RetrieveAdvisorScore("assignment", score.ID)
	if err != nil {
		t.Fatalf("Failed to retrieve score: %v", err)
	}

	retrievedScore, ok := res.(*model.ScoreAssignmentAdvisor)
	if !ok {
		t.Fatalf("Failed to cast retrieved score to *model.ScoreAssignmentAdvisor")
	}

	if retrievedScore.Score != score.Score {
		t.Errorf("Expected score %v, got %v", score.Score, retrievedScore.Score)
	}
}

func TestInsertAdvisorScore(t *testing.T) {
	db, _, _, _, _, _, _, dbName := Init()
	t.Cleanup(func() { cleanup(dbName) })

	score := &model.ScoreAssignmentAdvisor{
		AssignmentId: 1,
		AdvisorId:    2,
		Score:        95.0,
	}
	advisorCtrl := controller.NewScoreAdvisorController(db)
	err := advisorCtrl.InsertAdvisorScore(score)
	if err != nil {
		t.Errorf("Failed to insert score: %v", err)
	}

	var insertedScore model.ScoreAssignmentAdvisor
	if err := db.First(&insertedScore, score.ID).Error; err != nil {
		t.Fatalf("Failed to retrieve inserted score: %v", err)
	}

	if insertedScore.Score != score.Score {
		t.Errorf("Expected score %v, got %v", score.Score, insertedScore.Score)
	}
}

func TestUpdateAdvisorScore(t *testing.T) {
	db, _, _, _, _, _, _, dbName := Init()
	t.Cleanup(func() { cleanup(dbName) })

	score := &model.ScoreAssignmentAdvisor{
		AssignmentId: 1,
		AdvisorId:    2,
		Score:        80.0,
	}
	if err := db.Create(score).Error; err != nil {
		t.Fatalf("Failed to create score: %v", err)
	}

	score.Score = 100.0
	advisorCtrl := controller.NewScoreAdvisorController(db)
	err := advisorCtrl.UpdateAdvisorScore("assignment", score)
	if err != nil {
		t.Errorf("Failed to update score: %v", err)
	}

	var updatedScore model.ScoreAssignmentAdvisor
	if err := db.First(&updatedScore, score.ID).Error; err != nil {
		t.Fatalf("Failed to retrieve updated score: %v", err)
	}

	if updatedScore.Score != score.Score {
		t.Errorf("Expected score %v, got %v", score.Score, updatedScore.Score)
	}
}

func TestDeleteAdvisorScore(t *testing.T) {
	db, _, _, _, _, _, _, dbName := Init()
	t.Cleanup(func() { cleanup(dbName) })

	score := &model.ScoreAssignmentAdvisor{
		AssignmentId: 1,
		AdvisorId:    2,
		Score:        75.0,
	}
	if err := db.Create(score).Error; err != nil {
		t.Fatalf("Failed to create score: %v", err)
	}

	advisorCtrl := controller.NewScoreAdvisorController(db)
	err := advisorCtrl.DeleteAdvisorScore("assignment", score)
	if err != nil {
		t.Errorf("Failed to delete score: %v", err)
	}

	var deletedScore model.ScoreAssignmentAdvisor
	err = db.First(&deletedScore, score.ID).Error
	if !errors.Is(err, gorm.ErrRecordNotFound) {
		t.Errorf("Expected record to be deleted, but it still exists")
	}
}

func TestValidateAdvisorScore(t *testing.T) {
	score := &model.ScoreAssignmentAdvisor{
		AssignmentId: 1,
		AdvisorId:    2,
		Score:        110.0, // Invalid score
	}
	err := score.Validate()
	if err == nil {
		t.Errorf("Expected validation error for invalid score, got nil")
	}

	score.Score = 90.0 // Valid score
	err = score.Validate()
	if err != nil {
		t.Errorf("Expected no validation error for valid score, got %v", err)
	}
}
