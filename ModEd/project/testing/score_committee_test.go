package testing

import (
	"ModEd/project/controller"
	"ModEd/project/model"
	"errors"
	"testing"

	"gorm.io/gorm"
)

func TestListAllCommitteeScores(t *testing.T) {
	db, _, _, _, _, _, _, dbName := Init()
	t.Cleanup(func() { cleanup(dbName) })

	score := model.ScoreAssignmentCommittee{
		AssignmentId: 1,
		CommitteeId:  2,
		Score:        85.0,
	}
	if err := db.Create(&score).Error; err != nil {
		t.Fatalf("Failed to create score: %v", err)
	}

	committeeCtrl := controller.NewScoreCommitteeController(db)
	scores, err := committeeCtrl.ListAllCommitteeScores("assignment")
	if err != nil {
		t.Errorf("Expected scores, got error: %v", err)
		return
	}

	scoreList, ok := scores.(*[]model.ScoreAssignmentCommittee)
	if !ok || len(*scoreList) == 0 {
		t.Errorf("Expected scores, got none or incorrect type")
	}
}

func TestRetrieveCommitteeScore(t *testing.T) {
	db, _, _, _, _, _, _, dbName := Init()
	t.Cleanup(func() { cleanup(dbName) })

	score := model.ScoreAssignmentCommittee{
		AssignmentId: 1,
		CommitteeId:  2,
		Score:        90.0,
	}
	if err := db.Create(&score).Error; err != nil {
		t.Fatalf("Failed to create score: %v", err)
	}

	t.Logf("Created ScoreAssignmentCommittee with ID: %d", score.ID)

	committeeCtrl := controller.NewScoreCommitteeController(db)
	res, err := committeeCtrl.RetrieveCommitteeScore("assignment", score.ID)
	if err != nil {
		t.Fatalf("Failed to retrieve score: %v", err)
	}

	retrievedScore, ok := res.(*model.ScoreAssignmentCommittee)
	if !ok {
		t.Fatalf("Failed to cast retrieved score to *model.ScoreAssignmentCommittee")
	}

	if retrievedScore.Score != score.Score {
		t.Errorf("Expected score %v, got %v", score.Score, retrievedScore.Score)
	}
}

func TestInsertCommitteeScore(t *testing.T) {
	db, _, _, _, _, _, _, dbName := Init()
	t.Cleanup(func() { cleanup(dbName) })

	score := &model.ScoreAssignmentCommittee{
		AssignmentId: 1,
		CommitteeId:  2,
		Score:        95.0,
	}
	committeeCtrl := controller.NewScoreCommitteeController(db)
	err := committeeCtrl.InsertCommitteeScore(score)
	if err != nil {
		t.Errorf("Failed to insert score: %v", err)
	}

	var insertedScore model.ScoreAssignmentCommittee
	if err := db.First(&insertedScore, score.ID).Error; err != nil {
		t.Fatalf("Failed to retrieve inserted score: %v", err)
	}

	if insertedScore.Score != score.Score {
		t.Errorf("Expected score %v, got %v", score.Score, insertedScore.Score)
	}
}

func TestUpdateCommitteeScore(t *testing.T) {
	db, _, _, _, _, _, _, dbName := Init()
	t.Cleanup(func() { cleanup(dbName) })

	score := &model.ScoreAssignmentCommittee{
		AssignmentId: 1,
		CommitteeId:  2,
		Score:        80.0,
	}
	if err := db.Create(score).Error; err != nil {
		t.Fatalf("Failed to create score: %v", err)
	}

	score.Score = 100.0
	committeeCtrl := controller.NewScoreCommitteeController(db)
	err := committeeCtrl.UpdateCommitteeScore("assignment", score)
	if err != nil {
		t.Errorf("Failed to update score: %v", err)
	}

	var updatedScore model.ScoreAssignmentCommittee
	if err := db.First(&updatedScore, score.ID).Error; err != nil {
		t.Fatalf("Failed to retrieve updated score: %v", err)
	}

	if updatedScore.Score != score.Score {
		t.Errorf("Expected score %v, got %v", score.Score, updatedScore.Score)
	}
}

func TestDeleteCommitteeScore(t *testing.T) {
	db, _, _, _, _, _, _, dbName := Init()
	t.Cleanup(func() { cleanup(dbName) })

	score := &model.ScoreAssignmentCommittee{
		AssignmentId: 1,
		CommitteeId:  2,
		Score:        75.0,
	}
	if err := db.Create(score).Error; err != nil {
		t.Fatalf("Failed to create score: %v", err)
	}

	committeeCtrl := controller.NewScoreCommitteeController(db)
	err := committeeCtrl.DeleteCommitteeScore("assignment", score)
	if err != nil {
		t.Errorf("Failed to delete score: %v", err)
	}

	var deletedScore model.ScoreAssignmentCommittee
	err = db.First(&deletedScore, score.ID).Error
	if !errors.Is(err, gorm.ErrRecordNotFound) {
		t.Errorf("Expected record to be deleted, but it still exists")
	}
}

func TestValidateCommitteeScore(t *testing.T) {
	score := &model.ScoreAssignmentCommittee{
		AssignmentId: 1,
		CommitteeId:  2,
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
