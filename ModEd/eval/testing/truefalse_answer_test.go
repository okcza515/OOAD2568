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

func setupTestDBTFAnswer(t *testing.T) *gorm.DB {
	db, err := gorm.Open(sqlite.Open(":memory:"), &gorm.Config{})
	assert.NoError(t, err)

	err = db.AutoMigrate(&model.Exam{}, &model.ExamSection{}, &model.Question{}, &model.TrueFalseAnswer{})
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

	err := ctrl.Insert(tfAnswer)
	assert.NoError(t, err)
	assert.NotZero(t, tfAnswer.ID)

	var found model.TrueFalseAnswer
	err = db.First(&found, tfAnswer.ID).Error
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

	err := ctrl.Insert(tfAnswer1)
	assert.NoError(t, err)
	assert.NotZero(t, tfAnswer1.ID)
	err = ctrl.Insert(tfAnswer2)
	assert.NoError(t, err)
	assert.NotZero(t, tfAnswer2.ID)

	tfAnswers, err := ctrl.List(map[string]interface{}{})
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

	err := ctrl.Insert(tfAnswer)
	assert.NoError(t, err)
	assert.NotZero(t, tfAnswer.ID)

	found, err := ctrl.RetrieveByID(tfAnswer.ID)
	assert.NoError(t, err)
	assert.Equal(t, tfAnswer.QuestionID, found.QuestionID)
	assert.Equal(t, tfAnswer.IsExpected, found.IsExpected)
}

func TestUpdateTrueFalseAnswer(t *testing.T) {
	db := setupTestDBTFAnswer(t)
	ctrl := controller.NewTrueFalseAnswerController(db)

	tfAnswer := &model.TrueFalseAnswer{
		QuestionID: 1,
		IsExpected: true,
	}
	err := ctrl.Insert(tfAnswer)
	assert.NoError(t, err)
	assert.NotEqual(t, tfAnswer.ID, 0, "tfAnswerId should not be zero")

	tfAnswer.IsExpected = false
	err = ctrl.UpdateByID(tfAnswer)
	assert.NoError(t, err)
	assert.Equal(t, false, tfAnswer.IsExpected, "IsExpected should be false after update")
}

func TestDeleteTrueFalseAnswer(t *testing.T) {
	db := setupTestDBTFAnswer(t)
	ctrl := controller.NewTrueFalseAnswerController(db)

	tfAnswer := &model.TrueFalseAnswer{
		QuestionID: 1,
		IsExpected: true,
	}
	err := ctrl.Insert(tfAnswer)
	assert.NoError(t, err)
	assert.NotEqual(t, tfAnswer.ID, 0, "tfAnswer.ID should not be zero")

	err = ctrl.DeleteByID(tfAnswer.ID)
	assert.NoError(t, err)

	var found model.TrueFalseAnswer
	err = db.First(&found, tfAnswer.ID).Error
	assert.Error(t, err, "Expected error when retrieving deleted record")
}
