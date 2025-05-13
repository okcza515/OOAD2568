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

func setupTestDBShortAnswer(t *testing.T) *gorm.DB {
	db, err := gorm.Open(sqlite.Open(":memory:"), &gorm.Config{})
	assert.NoError(t, err)

	err = db.AutoMigrate(&model.Exam{}, &model.ExamSection{}, &model.Question{}, &model.ShortAnswer{})
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

	err := ctrl.Insert(shortAnswer)
	assert.NoError(t, err)
	assert.NotZero(t, shortAnswer.ID)

	var found model.ShortAnswer
	err = db.First(&found, shortAnswer.ID).Error
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

	err := ctrl.Insert(shortAnswer1)
	assert.NoError(t, err)
	err = ctrl.Insert(shortAnswer2)
	assert.NoError(t, err)

	shortAnswers, err := ctrl.List(map[string]interface{}{})
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

	err := ctrl.Insert(shortAnswer)
	assert.NoError(t, err)

	found, err := ctrl.RetrieveByID(shortAnswer.ID)
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
	err := ctrl.Insert(shortAnswer)
	assert.NoError(t, err)
	assert.NotEqual(t, shortAnswer.ID, 0, "shortAnswerID should not be zero")

	shortAnswer.ExpectedAnswer = "Updated answer"
	err = ctrl.UpdateByID(shortAnswer)
	assert.NoError(t, err)

	assert.Equal(t, "Updated answer", shortAnswer.ExpectedAnswer, "ExpectedAnswer should be updated")
}

func TestDeleteShortAnswer(t *testing.T) {
	db := setupTestDBShortAnswer(t)
	ctrl := controller.NewShortAnswerController(db)

	shortAnswer := &model.ShortAnswer{
		QuestionID:     1,
		ExpectedAnswer: "Delete me",
	}
	err := ctrl.Insert(shortAnswer)
	assert.NoError(t, err)
	assert.NotEqual(t, shortAnswer.ID, 0, "shortAnswer.ID should not be zero")

	err = ctrl.DeleteByID(shortAnswer.ID)
	assert.NoError(t, err)

	var found model.ShortAnswer
	err = db.First(&found, shortAnswer.ID).Error
	assert.Error(t, err, "Expected error when retrieving deleted shortAnswer")
}
