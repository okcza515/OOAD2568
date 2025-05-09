package testing

import (
	"ModEd/eval/controller"
	"ModEd/eval/model"
	"testing"

	"github.com/stretchr/testify/assert"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func setupTestDBMCAnswer(t *testing.T) *gorm.DB {
	db, err := gorm.Open(sqlite.Open(":memory:"), &gorm.Config{})
	assert.NoError(t, err)

	err = db.AutoMigrate(&model.Examination{}, &model.ExamSection{}, &model.Question{}, &model.MultipleChoiceAnswer{})
	assert.NoError(t, err)

	return db
}

func TestCreateMultipleChoiceAnswer(t *testing.T) {
	db := setupTestDBMCAnswer(t)

	ctrl := controller.NewMultipleChoiceAnswerController(db)

	mcAnswer := &model.MultipleChoiceAnswer{
		QuestionID:  1,
		AnswerLabel: "A",
		IsExpected:  true,
	}

	mcAnswerID, err := ctrl.CreateMultipleChoiceAnswer(mcAnswer)
	assert.NoError(t, err)
	assert.NotZero(t, mcAnswerID)

	var found model.MultipleChoiceAnswer
	err = db.First(&found, mcAnswerID).Error
	assert.NoError(t, err)
	assert.Equal(t, mcAnswer.QuestionID, found.QuestionID)
	assert.Equal(t, mcAnswer.AnswerLabel, found.AnswerLabel)
	assert.Equal(t, mcAnswer.IsExpected, found.IsExpected)
}

func TestGetAllMultipleChoiceAnswers(t *testing.T) {
	db := setupTestDBMCAnswer(t)

	ctrl := controller.NewMultipleChoiceAnswerController(db)

	mcAnswer1 := &model.MultipleChoiceAnswer{
		QuestionID:  1,
		AnswerLabel: "A",
		IsExpected:  true,
	}
	mcAnswer2 := &model.MultipleChoiceAnswer{
		QuestionID:  2,
		AnswerLabel: "B",
		IsExpected:  false,
	}

	_, err := ctrl.CreateMultipleChoiceAnswer(mcAnswer1)
	assert.NoError(t, err)
	_, err = ctrl.CreateMultipleChoiceAnswer(mcAnswer2)
	assert.NoError(t, err)

	mcAnswers, err := ctrl.GetAllMultipleChoiceAnswers()
	assert.NoError(t, err)
	assert.Len(t, mcAnswers, 2)
}

func TestGetMultipleChoiceAnswer(t *testing.T) {
	db := setupTestDBMCAnswer(t)

	ctrl := controller.NewMultipleChoiceAnswerController(db)

	mcAnswer := &model.MultipleChoiceAnswer{
		QuestionID:  1,
		AnswerLabel: "A",
		IsExpected:  true,
	}

	mcAnswerID, err := ctrl.CreateMultipleChoiceAnswer(mcAnswer)
	assert.NoError(t, err)

	found, err := ctrl.GetMultipleChoiceAnswer(mcAnswerID)
	assert.NoError(t, err)
	assert.Equal(t, mcAnswer.QuestionID, found.QuestionID)
	assert.Equal(t, mcAnswer.AnswerLabel, found.AnswerLabel)
	assert.Equal(t, mcAnswer.IsExpected, found.IsExpected)
}

func TestUpdateMultipleChoiceAnswer(t *testing.T) {
	db := setupTestDBMCAnswer(t)
	ctrl := controller.NewMultipleChoiceAnswerController(db)

	mcAnswer := &model.MultipleChoiceAnswer{
		QuestionID:  1,
		AnswerLabel: "A",
		IsExpected:  true,
	}
	mcAnswerID, err := ctrl.CreateMultipleChoiceAnswer(mcAnswer)
	assert.NoError(t, err)
	assert.NotEqual(t, mcAnswerID, 0, "mcAnswerID should not be zero")

	mcAnswer.AnswerLabel = "C"
	mcAnswer.IsExpected = false
	updatedMcAnswer, err := ctrl.UpdateMultipleChoiceAnswer(mcAnswer)
	assert.NoError(t, err)

	assert.Equal(t, "C", updatedMcAnswer.AnswerLabel, "AnswerLabel should be updated")
	assert.Equal(t, false, updatedMcAnswer.IsExpected, "IsExpected should be updated")
}

func TestDeleteMultipleChoiceAnswer(t *testing.T) {
	db := setupTestDBMCAnswer(t)
	ctrl := controller.NewMultipleChoiceAnswerController(db)

	mcAnswer := &model.MultipleChoiceAnswer{
		QuestionID:  1,
		AnswerLabel: "A",
		IsExpected:  true,
	}
	mcAnswerID, err := ctrl.CreateMultipleChoiceAnswer(mcAnswer)
	assert.NoError(t, err)
	assert.NotEqual(t, mcAnswerID, 0, "mcAnswerID should not be zero")

	deletedMcAnswer, err := ctrl.DeleteMultipleChoiceAnswer(mcAnswerID)
	assert.NoError(t, err)
	assert.Equal(t, mcAnswerID, deletedMcAnswer.ID)

	var found model.MultipleChoiceAnswer
	err = db.First(&found, mcAnswerID).Error
	assert.Error(t, err)
}
