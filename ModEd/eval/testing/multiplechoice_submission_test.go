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

func setupTestDBMCSubmission(t *testing.T) *gorm.DB {
	db, err := gorm.Open(sqlite.Open(":memory:"), &gorm.Config{})
	assert.NoError(t, err)

	err = db.AutoMigrate(&model.Exam{}, &model.ExamSection{}, &model.Question{}, &model.MultipleChoiceAnswer{}, &model.MultipleChoiceAnswerSubmission{})
	assert.NoError(t, err)

	return db
}

func createTestDataMC(db *gorm.DB) (*model.Question, *model.MultipleChoiceAnswer) {
	exam := model.Exam{
		ExamName:     "Test Exam",
		InstructorID: 1,
		ClassID:      1,
		Description:  "desc",
		ExamStatus:   model.Draft,
		Attempt:      1,
		StartDate:    time.Now(),
		EndDate:      time.Now(),
	}
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
		ActualQuestion: "What is 2 + 2?",
		QuestionType:   model.MultipleChoiceQuestion,
	}
	db.Create(&question)

	choice := model.MultipleChoiceAnswer{
		QuestionID:  question.ID,
		AnswerLabel: "4",
		IsExpected:  true,
	}
	db.Create(&choice)

	return &question, &choice
}

func TestCreateMultipleChoiceAnswerSubmission(t *testing.T) {
	db := setupTestDBMCSubmission(t)
	ctrl := controller.NewMultipleChoiceAnswerSubmissionController(db)

	question, choice := createTestDataMC(db)

	submission := &model.MultipleChoiceAnswerSubmission{
		QuestionID:   question.ID,
		SubmissionID: 1,
		ChoiceID:     choice.ID,
	}
	err := ctrl.Insert(submission)
	assert.NoError(t, err)
	assert.NotEqual(t, submission.ID, 0, "submission.ID should not be zero")

	var found model.MultipleChoiceAnswerSubmission
	err = db.First(&found, submission.ID).Error
	assert.NoError(t, err)
	assert.Equal(t, submission.ChoiceID, found.ChoiceID)
	assert.Equal(t, submission.SubmissionID, found.SubmissionID)
	assert.Equal(t, submission.QuestionID, found.QuestionID)
}

func TestGetAllMultipleChoiceAnswerSubmissions(t *testing.T) {
	db := setupTestDBMCSubmission(t)
	ctrl := controller.NewMultipleChoiceAnswerSubmissionController(db)

	_, _ = createTestDataMC(db)

	submission := model.MultipleChoiceAnswerSubmission{
		QuestionID:   1,
		SubmissionID: 1,
		ChoiceID:     1,
	}
	db.Create(&submission)

	results, err := ctrl.List(map[string]interface{}{})
	assert.NoError(t, err)
	assert.Len(t, results, 1)
}

func TestGetMultipleChoiceAnswerSubmissionByID(t *testing.T) {
	db := setupTestDBMCSubmission(t)
	ctrl := controller.NewMultipleChoiceAnswerSubmissionController(db)

	question, choice := createTestDataMC(db)

	submission := model.MultipleChoiceAnswerSubmission{
		QuestionID:   question.ID,
		SubmissionID: 1,
		ChoiceID:     choice.ID,
	}
	db.Create(&submission)

	results, err := ctrl.List(map[string]interface{}{"submission_id": submission.SubmissionID})
	assert.NoError(t, err)
	assert.Len(t, results, 1)
	assert.Equal(t, submission.ID, results[0].ID)
}

func TestGetMultipleChoiceAnswerSubmissionsBySubmissionID(t *testing.T) {
	db := setupTestDBMCSubmission(t)
	ctrl := controller.NewMultipleChoiceAnswerSubmissionController(db)

	question, choice := createTestDataMC(db)

	submission := model.MultipleChoiceAnswerSubmission{
		QuestionID:   question.ID,
		SubmissionID: 1,
		ChoiceID:     choice.ID,
	}
	db.Create(&submission)

	result, err := ctrl.RetrieveByCondition(map[string]interface{}{"submission_id": submission.SubmissionID})
	assert.NoError(t, err)
	assert.Equal(t, submission.ID, result.ID)
}

func TestUpdateMultipleChoiceAnswerSubmission(t *testing.T) {
	db := setupTestDBMCSubmission(t)
	ctrl := controller.NewMultipleChoiceAnswerSubmissionController(db)

	question, oldChoice := createTestDataMC(db)

	submission := &model.MultipleChoiceAnswerSubmission{
		QuestionID:   question.ID,
		SubmissionID: 1,
		ChoiceID:     oldChoice.ID,
	}
	db.Create(submission)

	newChoice := model.MultipleChoiceAnswer{
		QuestionID:  question.ID,
		AnswerLabel: "5",
		IsExpected:  false,
	}
	db.Create(&newChoice)

	submission.ChoiceID = newChoice.ID
	err := ctrl.UpdateByID(submission)
	assert.NoError(t, err)

	var updated model.MultipleChoiceAnswerSubmission
	err = db.First(&updated, submission.ID).Error
	assert.NoError(t, err)
	assert.Equal(t, newChoice.ID, updated.ChoiceID)
}

func TestDeleteMultipleChoiceAnswerSubmission(t *testing.T) {
	db := setupTestDBMCSubmission(t)
	ctrl := controller.NewMultipleChoiceAnswerSubmissionController(db)

	question, choice := createTestDataMC(db)

	submission := &model.MultipleChoiceAnswerSubmission{
		QuestionID:   question.ID,
		SubmissionID: 1,
		ChoiceID:     choice.ID,
	}
	db.Create(submission)

	err := ctrl.DeleteByID(submission.ID)
	assert.NoError(t, err)

	var found model.MultipleChoiceAnswerSubmission
	err = db.First(&found, submission.ID).Error
	assert.Error(t, err)
	assert.Equal(t, gorm.ErrRecordNotFound, err)
}
