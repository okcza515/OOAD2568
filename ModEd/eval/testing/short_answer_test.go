package testing

import (
	"ModEd/eval/controller"
	"ModEd/eval/model"
	"testing"

	"github.com/stretchr/testify/assert"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func setupTestDBShortAnswer(t *testing.T) *gorm.DB {
	db, err := gorm.Open(sqlite.Open(":memory:"), &gorm.Config{})
	assert.NoError(t, err)

	// Use AutoMigrate for ShortAnswer model
	err = db.AutoMigrate(&model.Examination{}, &model.ExamSection{}, &model.Question{}, &model.ShortAnswer{})
	assert.NoError(t, err)

	return db
}

func TestCreateShortAnswer(t *testing.T) {
	db := setupTestDBShortAnswer(t)

	ctrl := controller.NewShortAnswerController(db)

	shortAnswer := &model.ShortAnswer{
		QuestionID:     1,
		ExpectedAnswer: "This is a short answer",
	}

	shortAnswerID, err := ctrl.CreateShortAnswer(shortAnswer)
	assert.NoError(t, err)
	assert.NotZero(t, shortAnswerID)

	var found model.ShortAnswer
	err = db.First(&found, shortAnswerID).Error
	assert.NoError(t, err)
	assert.Equal(t, shortAnswer.QuestionID, found.QuestionID)
	assert.Equal(t, shortAnswer.ExpectedAnswer, found.ExpectedAnswer)
}

func TestGetAllShortAnswers(t *testing.T) {
	db := setupTestDBShortAnswer(t)

	ctrl := controller.NewShortAnswerController(db)

	shortAnswer1 := &model.ShortAnswer{
		QuestionID:     1,
		ExpectedAnswer: "Answer 1",
	}
	shortAnswer2 := &model.ShortAnswer{
		QuestionID:     2,
		ExpectedAnswer: "Answer 2",
	}

	_, err := ctrl.CreateShortAnswer(shortAnswer1)
	assert.NoError(t, err)
	_, err = ctrl.CreateShortAnswer(shortAnswer2)
	assert.NoError(t, err)

	shortAnswers, err := ctrl.GetAllShortAnswers()
	assert.NoError(t, err)
	assert.Len(t, shortAnswers, 2)
}

func TestGetShortAnswer(t *testing.T) {
	db := setupTestDBShortAnswer(t)

	ctrl := controller.NewShortAnswerController(db)

	shortAnswer := &model.ShortAnswer{
		QuestionID:     1,
		ExpectedAnswer: "Short answer",
	}

	shortAnswerID, err := ctrl.CreateShortAnswer(shortAnswer)
	assert.NoError(t, err)

	found, err := ctrl.GetShortAnswer(shortAnswerID)
	assert.NoError(t, err)
	assert.Equal(t, shortAnswer.QuestionID, found.QuestionID)
	assert.Equal(t, shortAnswer.ExpectedAnswer, found.ExpectedAnswer)
}

func TestUpdateShortAnswer(t *testing.T) {
	db := setupTestDBShortAnswer(t)
	ctrl := controller.NewShortAnswerController(db)

	shortAnswer := &model.ShortAnswer{
		QuestionID:     1,
		ExpectedAnswer: "Original answer",
	}
	shortAnswerID, err := ctrl.CreateShortAnswer(shortAnswer)
	assert.NoError(t, err)
	assert.NotEqual(t, shortAnswerID, 0, "shortAnswerID should not be zero")

	shortAnswer.ExpectedAnswer = "Updated answer"
	updatedShortAnswer, err := ctrl.UpdateShortAnswer(shortAnswer)
	assert.NoError(t, err)

	assert.Equal(t, "Updated answer", updatedShortAnswer.ExpectedAnswer, "ExpectedAnswer should be updated")
}

func TestDeleteShortAnswer(t *testing.T) {
	db := setupTestDBShortAnswer(t)
	ctrl := controller.NewShortAnswerController(db)

	shortAnswer := &model.ShortAnswer{
		QuestionID:     1,
		ExpectedAnswer: "Delete me",
	}
	shortAnswerID, err := ctrl.CreateShortAnswer(shortAnswer)
	assert.NoError(t, err)
	assert.NotEqual(t, shortAnswerID, 0, "shortAnswerID should not be zero")

	deletedShortAnswer, err := ctrl.DeleteShortAnswer(shortAnswerID)
	assert.NoError(t, err)
	assert.Equal(t, shortAnswerID, deletedShortAnswer.ID)

	var found model.ShortAnswer
	err = db.First(&found, shortAnswerID).Error
	assert.Error(t, err)
}
