package testing

import (
	"ModEd/eval/controller"
	"ModEd/eval/model"
	"testing"

	"github.com/stretchr/testify/assert"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func setupTestDBShortAnswerSubmission(t *testing.T) *gorm.DB {
	db, err := gorm.Open(sqlite.Open(":memory:"), &gorm.Config{})
	assert.NoError(t, err)

	err = db.AutoMigrate(&model.ShortAnswerSubmission{})
	assert.NoError(t, err)

	return db
}

func TestCreateShortAnswerSubmission(t *testing.T) {
	db := setupTestDBShortAnswerSubmission(t)
	ctrl := controller.NewShortAnswerSubmissionController(db)

	submission := &model.ShortAnswerSubmission{
		QuestionID:    1,
		SubmissionID:  1,
		StudentAnswer: "London",
	}

	err := ctrl.Insert(submission)
	assert.NoError(t, err)
	assert.NotZero(t, submission.ID, "ID should not be zero after insertion")

	var found model.ShortAnswerSubmission
	err = db.First(&found, submission.ID).Error
	assert.NoError(t, err)
	assert.Equal(t, submission.QuestionID, found.QuestionID)
	assert.Equal(t, submission.SubmissionID, found.SubmissionID)
	assert.Equal(t, submission.StudentAnswer, found.StudentAnswer)
}

func TestGetAllShortAnswerSubmissions(t *testing.T) {
	db := setupTestDBShortAnswerSubmission(t)
	ctrl := controller.NewShortAnswerSubmissionController(db)

	submission1 := &model.ShortAnswerSubmission{
		QuestionID:    1,
		SubmissionID:  1,
		StudentAnswer: "London",
	}
	submission2 := &model.ShortAnswerSubmission{
		QuestionID:    2,
		SubmissionID:  2,
		StudentAnswer: "Berlin",
	}

	err := ctrl.Insert(submission1)
	assert.NoError(t, err)
	assert.NotZero(t, submission1.ID, "ID should not be zero after insertion")

	err = ctrl.Insert(submission2)
	assert.NoError(t, err)
	assert.NotZero(t, submission2.ID, "ID should not be zero after insertion")

	submissions, err := ctrl.List(map[string]interface{}{})
	assert.NoError(t, err)
	assert.Len(t, submissions, 2)
}

func TestGetShortAnswerSubmission(t *testing.T) {
	db := setupTestDBShortAnswerSubmission(t)
	ctrl := controller.NewShortAnswerSubmissionController(db)

	submission := &model.ShortAnswerSubmission{
		QuestionID:    1,
		SubmissionID:  1,
		StudentAnswer: "London",
	}

	err := ctrl.Insert(submission)
	assert.NoError(t, err)
	assert.NotZero(t, submission.ID, "ID should not be zero after insertion")

	var found model.ShortAnswerSubmission
	err = db.First(&found, submission.ID).Error
	assert.NoError(t, err)
	assert.Equal(t, submission.QuestionID, found.QuestionID)
	assert.Equal(t, submission.StudentAnswer, found.StudentAnswer)
}

func TestUpdateShortAnswerSubmission(t *testing.T) {
	db := setupTestDBShortAnswerSubmission(t)
	ctrl := controller.NewShortAnswerSubmissionController(db)

	submission := &model.ShortAnswerSubmission{
		QuestionID:    1,
		SubmissionID:  1,
		StudentAnswer: "London",
	}

	err := ctrl.Insert(submission)
	assert.NoError(t, err)
	assert.NotZero(t, submission.ID, "ID should not be zero after insertion")

	submission.StudentAnswer = "Berlin"
	err = ctrl.UpdateByID(submission)
	assert.NoError(t, err)

	var found model.ShortAnswerSubmission
	err = db.First(&found, submission.ID).Error
	assert.NoError(t, err)
	assert.Equal(t, "Berlin", found.StudentAnswer)
}

func TestDeleteShortAnswerSubmission(t *testing.T) {
	db := setupTestDBShortAnswerSubmission(t)
	ctrl := controller.NewShortAnswerSubmissionController(db)

	submission := &model.ShortAnswerSubmission{
		QuestionID:    1,
		SubmissionID:  1,
		StudentAnswer: "London",
	}

	err := ctrl.Insert(submission)
	assert.NoError(t, err)

	err = ctrl.DeleteByID(submission.ID)
	assert.NoError(t, err)

	var found model.ShortAnswerSubmission
	err = db.First(&found, submission.ID).Error
	assert.Error(t, err)
}
