package testing

import (
	"ModEd/eval/controller"
	"ModEd/eval/model"
	"testing"

	"github.com/stretchr/testify/assert"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func setupTestDBTFAnswerSubmission(t *testing.T) *gorm.DB {
	db, err := gorm.Open(sqlite.Open(":memory:"), &gorm.Config{})
	assert.NoError(t, err)

	err = db.AutoMigrate(&model.TrueFalseAnswerSubmission{})
	assert.NoError(t, err)

	return db
}

func TestCreateTrueFalseAnswerSubmission(t *testing.T) {
	db := setupTestDBTFAnswerSubmission(t)
	ctrl := controller.NewTrueFalseAnswerSubmissionController(db)

	submission := &model.TrueFalseAnswerSubmission{
		QuestionID:    1,
		SubmissionID:  1,
		StudentAnswer: true,
	}
	err := ctrl.Insert(submission)
	assert.NoError(t, err)
	assert.NotZero(t, submission.ID)

	var found model.TrueFalseAnswerSubmission
	err = db.First(&found, submission.ID).Error
	assert.NoError(t, err)
	assert.Equal(t, submission.QuestionID, found.QuestionID)
	assert.Equal(t, submission.SubmissionID, found.SubmissionID)
	assert.Equal(t, submission.StudentAnswer, found.StudentAnswer)
}

func TestGetAllTrueFalseAnswerSubmissions(t *testing.T) {
	db := setupTestDBTFAnswerSubmission(t)
	ctrl := controller.NewTrueFalseAnswerSubmissionController(db)

	err := ctrl.Insert(&model.TrueFalseAnswerSubmission{
		QuestionID:    1,
		SubmissionID:  1,
		StudentAnswer: true,
	})
	assert.NoError(t, err)

	results, err := ctrl.List(nil, "Submission")
	assert.NoError(t, err)
	assert.Len(t, results, 1)
}

func TestGetTrueFalseAnswerSubmission(t *testing.T) {
	db := setupTestDBTFAnswerSubmission(t)
	ctrl := controller.NewTrueFalseAnswerSubmissionController(db)

	submission := &model.TrueFalseAnswerSubmission{
		QuestionID:    1,
		SubmissionID:  1,
		StudentAnswer: true,
	}
	err := ctrl.Insert(submission)
	assert.NoError(t, err)

	found, err := ctrl.RetrieveByID(submission.ID)
	assert.NoError(t, err)
	assert.Equal(t, submission.ID, found.ID)
	assert.Equal(t, submission.StudentAnswer, found.StudentAnswer)
}

func TestGetTrueFalseAnswerSubmissionsBySubmissionID(t *testing.T) {
	db := setupTestDBTFAnswerSubmission(t)
	ctrl := controller.NewTrueFalseAnswerSubmissionController(db)

	ctrl.Insert(&model.TrueFalseAnswerSubmission{QuestionID: 1, SubmissionID: 10, StudentAnswer: true})
	ctrl.Insert(&model.TrueFalseAnswerSubmission{QuestionID: 2, SubmissionID: 10, StudentAnswer: false})
	ctrl.Insert(&model.TrueFalseAnswerSubmission{QuestionID: 3, SubmissionID: 20, StudentAnswer: true})

	results, err := ctrl.List(map[string]interface{}{"submission_id": 10})
	assert.NoError(t, err)
	assert.Len(t, results, 2)
	for _, sub := range results {
		assert.Equal(t, uint(10), sub.SubmissionID)
	}
}

func TestUpdateTrueFalseAnswerSubmission(t *testing.T) {
	db := setupTestDBTFAnswerSubmission(t)
	ctrl := controller.NewTrueFalseAnswerSubmissionController(db)

	submission := &model.TrueFalseAnswerSubmission{
		QuestionID:    1,
		SubmissionID:  1,
		StudentAnswer: false,
	}
	err := ctrl.Insert(submission)
	assert.NoError(t, err)

	submission.StudentAnswer = true
	err = ctrl.UpdateByID(submission)
	assert.NoError(t, err)

	updated, err := ctrl.RetrieveByID(submission.ID)

	assert.NoError(t, err)
	assert.Equal(t, true, updated.StudentAnswer, "IsExpected should be false after update")
}

func TestDeleteTrueFalseAnswerSubmission(t *testing.T) {
	db := setupTestDBTFAnswerSubmission(t)
	ctrl := controller.NewTrueFalseAnswerSubmissionController(db)

	submission := &model.TrueFalseAnswerSubmission{
		QuestionID:    1,
		SubmissionID:  1,
		StudentAnswer: true,
	}
	err := ctrl.Insert(submission)
	assert.NoError(t, err)

	err = ctrl.DeleteByID(submission.ID)
	assert.NoError(t, err)

	_, err = ctrl.RetrieveByID(submission.ID)
	assert.Error(t, err)
}
