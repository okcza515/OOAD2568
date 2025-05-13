package testing

// MEP-1007

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

	err = db.AutoMigrate(&model.MultipleChoiceAnswer{})
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

	err := ctrl.Insert(mcAnswer)
	assert.NoError(t, err)

	var found model.MultipleChoiceAnswer
	err = db.First(&found, mcAnswer.ID).Error
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

	err := ctrl.Insert(mcAnswer1)
	assert.NoError(t, err)
	err = ctrl.Insert(mcAnswer2)
	assert.NoError(t, err)

	mcAnswers, err := ctrl.List(map[string]interface{}{})
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

	err := ctrl.Insert(mcAnswer)
	assert.NoError(t, err)

	found, err := ctrl.RetrieveByID(mcAnswer.ID)
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
	err := ctrl.Insert(mcAnswer)
	assert.NoError(t, err)
	assert.NotEqual(t, mcAnswer.ID, 0, "mcAnswerID should not be zero")

	mcAnswer.AnswerLabel = "C"
	mcAnswer.IsExpected = false
	err = ctrl.UpdateByID(mcAnswer)

	assert.NoError(t, err)

	assert.Equal(t, "C", mcAnswer.AnswerLabel, "AnswerLabel should be updated")
	assert.Equal(t, false, mcAnswer.IsExpected, "IsExpected should be updated")
}

func TestDeleteMultipleChoiceAnswer(t *testing.T) {
	db := setupTestDBMCAnswer(t)
	ctrl := controller.NewMultipleChoiceAnswerController(db)

	mcAnswer := &model.MultipleChoiceAnswer{
		QuestionID:  1,
		AnswerLabel: "A",
		IsExpected:  true,
	}
	err := ctrl.Insert(mcAnswer)
	assert.NoError(t, err)
	assert.NotEqual(t, mcAnswer.ID, 0, "mcAnswerID should not be zero")

	err = ctrl.DeleteByID(mcAnswer.ID)
	assert.NoError(t, err)

	var found model.MultipleChoiceAnswer
	err = db.First(&found, mcAnswer.ID).Error
	assert.Error(t, err)
}
