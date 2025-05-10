package testing

import (
	"ModEd/eval/controller"
	"ModEd/eval/model"
	"testing"

	"github.com/stretchr/testify/assert"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func setupTestDBTFAnswer(t *testing.T) *gorm.DB {
	db, err := gorm.Open(sqlite.Open(":memory:"), &gorm.Config{})
	assert.NoError(t, err)

	err = db.AutoMigrate(&model.Examination{}, &model.ExamSection{}, &model.Question{}, &model.TrueFalseAnswer{})
	assert.NoError(t, err)

	return db
}

func TestCreateTrueFalseAnswer(t *testing.T) {
	db := setupTestDBTFAnswer(t)

	ctrl := controller.NewTrueFalseAnswerController(db)

	tfAnswer := &model.TrueFalseAnswer{
		QuestionID: 1,
		IsExpected: true,
	}

	tfAnswerID, err := ctrl.CreateTrueFalseAnswer(tfAnswer)
	assert.NoError(t, err)
	assert.NotZero(t, tfAnswerID)

	var found model.TrueFalseAnswer
	err = db.First(&found, tfAnswerID).Error
	assert.NoError(t, err)
	assert.Equal(t, tfAnswer.QuestionID, found.QuestionID)
	assert.Equal(t, tfAnswer.IsExpected, found.IsExpected)
}

func TestGetAllTrueFalseAnswers(t *testing.T) {
	db := setupTestDBTFAnswer(t)

	ctrl := controller.NewTrueFalseAnswerController(db)

	tfAnswer1 := &model.TrueFalseAnswer{
		QuestionID: 1,
		IsExpected: true,
	}
	tfAnswer2 := &model.TrueFalseAnswer{
		QuestionID: 2,
		IsExpected: false,
	}

	_, err := ctrl.CreateTrueFalseAnswer(tfAnswer1)
	assert.NoError(t, err)
	_, err = ctrl.CreateTrueFalseAnswer(tfAnswer2)
	assert.NoError(t, err)

	tfAnswers, err := ctrl.GetAllTrueFalseAnswers()
	assert.NoError(t, err)
	assert.Len(t, tfAnswers, 2)
}

func TestGetTrueFalseAnswer(t *testing.T) {
	db := setupTestDBTFAnswer(t)

	ctrl := controller.NewTrueFalseAnswerController(db)

	tfAnswer := &model.TrueFalseAnswer{
		QuestionID: 1,
		IsExpected: true,
	}

	tfAnswerID, err := ctrl.CreateTrueFalseAnswer(tfAnswer)
	assert.NoError(t, err)

	found, err := ctrl.GetTrueFalseAnswer(tfAnswerID)
	assert.NoError(t, err)
	assert.Equal(t, tfAnswer.QuestionID, found.QuestionID)
	assert.Equal(t, tfAnswer.IsExpected, found.IsExpected)
}

func TestUpdateTrueFalseAnswer(t *testing.T) {
	db := setupTestDBTFAnswer(t)
	ctrl := controller.NewTrueFalseAnswerController(db)

	tfAnswer := &model.TrueFalseAnswer{
		QuestionID: 1,
		IsExpected: false,
	}
	tfAnswerId, err := ctrl.CreateTrueFalseAnswer(tfAnswer)
	assert.NoError(t, err)
	assert.NotEqual(t, tfAnswerId, 0, "tfAnswerId should not be zero")

	tfAnswer.IsExpected = true
	updatedTfAnswer, err := ctrl.UpdateTrueFalseAnswer(tfAnswer)
	assert.NoError(t, err)

	assert.Equal(t, true, updatedTfAnswer.IsExpected, "IsExpected should be true after update")
}

func TestDeleteTrueFalseAnswer(t *testing.T) {
	db := setupTestDBTFAnswer(t)
	ctrl := controller.NewTrueFalseAnswerController(db)

	tfAnswer := &model.TrueFalseAnswer{
		QuestionID: 1,
		IsExpected: true,
	}
	tfAnswerId, err := ctrl.CreateTrueFalseAnswer(tfAnswer)
	assert.NoError(t, err)
	assert.NotEqual(t, tfAnswerId, 0, "tfAnswerId should not be zero")

	deletedTfAnswer, err := ctrl.DeleteTrueFalseAnswer(tfAnswerId)
	assert.NoError(t, err)
	assert.Equal(t, tfAnswerId, deletedTfAnswer.ID)

	var found model.TrueFalseAnswer
	err = db.First(&found, tfAnswerId).Error
	assert.Error(t, err)
}
