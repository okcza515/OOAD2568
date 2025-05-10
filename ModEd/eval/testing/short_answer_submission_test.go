package testing

import (
	"ModEd/eval/controller"
	"ModEd/eval/model"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func setupTestDBShortAnswerSubmission(t *testing.T) *gorm.DB {
	db, err := gorm.Open(sqlite.Open(":memory:"), &gorm.Config{})
	assert.NoError(t, err)

	err = db.AutoMigrate(&model.Question{}, &model.Examination{}, &model.ExamSection{}, &model.ShortAnswerSubmission{})
	assert.NoError(t, err)

	return db
}

func createTestDataShortAnswerSubmission(db *gorm.DB) (model.ShortAnswerSubmission, uint) {
	exam := model.Examination{ExamName: "Test Exam", InstructorID: 1, CourseID: 1, Description: "desc", ExamStatus: model.Draft, Attempt: 1, StartDate: time.Now(), EndDate: time.Now()}
	db.Create(&exam)

	section := model.ExamSection{
		ExamID:       exam.ID,
		SectionNo:    1,
		Description:  "Test Section",
		NumQuestions: 1,
		Score:        1,
	}
	db.Create(&section)

	question := model.Question{
		SectionID:      section.ID,
		Score:          1,
		ActualQuestion: "What is the capital of France?",
		QuestionType:   model.ShortAnswerQuestion,
	}
	db.Create(&question)

	submission := model.ShortAnswerSubmission{
		QuestionID:    question.ID,
		SubmissionID:  1,
		StudentAnswer: "Paris",
	}
	db.Create(&submission)

	return submission, question.ID
}

func TestCreateShortAnswerSubmission(t *testing.T) {
	db := setupTestDBShortAnswerSubmission(t)
	ctrl := controller.NewShortAnswerSubmissionController(db)

	submission := &model.ShortAnswerSubmission{
		QuestionID:    1,
		SubmissionID:  1,
		StudentAnswer: "London",
	}

	id, err := ctrl.CreateShortAnswerSubmission(submission)
	assert.NoError(t, err)
	assert.NotZero(t, id)
}

func TestGetAllShortAnswerSubmissions(t *testing.T) {
	db := setupTestDBShortAnswerSubmission(t)
	ctrl := controller.NewShortAnswerSubmissionController(db)

	_, _ = createTestDataShortAnswerSubmission(db)

	results, err := ctrl.GetAllShortAnswerSubmissions()
	assert.NoError(t, err)
	assert.Len(t, results, 1)
}

func TestGetShortAnswerSubmission(t *testing.T) {
	db := setupTestDBShortAnswerSubmission(t)
	ctrl := controller.NewShortAnswerSubmissionController(db)

	submission, _ := createTestDataShortAnswerSubmission(db)

	result, err := ctrl.GetShortAnswerSubmission(submission.ID)
	assert.NoError(t, err)
	assert.Equal(t, submission.ID, result.ID)
}

func TestGetShortAnswerSubmissionsBySubmissionID(t *testing.T) {
	db := setupTestDBShortAnswerSubmission(t)
	ctrl := controller.NewShortAnswerSubmissionController(db)

	submission, _ := createTestDataShortAnswerSubmission(db)

	results, err := ctrl.GetShortAnswerSubmissionsBySubmissionID(1)
	assert.NoError(t, err)
	assert.Len(t, results, 1)
	assert.Equal(t, submission.ID, results[0].ID)
}

func TestUpdateShortAnswerSubmission(t *testing.T) {
	db := setupTestDBShortAnswerSubmission(t)
	ctrl := controller.NewShortAnswerSubmissionController(db)

	submission, _ := createTestDataShortAnswerSubmission(db)

	submission.StudentAnswer = "Berlin"
	updated, err := ctrl.UpdateShortAnswerSubmission(&submission)
	assert.NoError(t, err)
	assert.Equal(t, "Berlin", updated.StudentAnswer)
}

func TestDeleteShortAnswerSubmission(t *testing.T) {
	db := setupTestDBShortAnswerSubmission(t)
	ctrl := controller.NewShortAnswerSubmissionController(db)

	submission, _ := createTestDataShortAnswerSubmission(db)

	deleted, err := ctrl.DeleteShortAnswerSubmission(submission.ID)
	assert.NoError(t, err)
	assert.Equal(t, submission.ID, deleted.ID)

	var check model.ShortAnswerSubmission
	err = db.First(&check, submission.ID).Error
	assert.Error(t, err)
}
