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

	err = db.AutoMigrate(&model.Question{}, &model.Examination{}, &model.ExamSection{}, &model.MultipleChoiceAnswer{}, &model.MultipleChoiceAnswerSubmission{})
	assert.NoError(t, err)

	return db
}

func createTestData(db *gorm.DB) (model.MultipleChoiceAnswerSubmission, uint) {
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

	mcSub := model.MultipleChoiceAnswerSubmission{
		QuestionID:   question.ID,
		SubmissionID: 1,
		ChoiceID:     choice.ID,
	}
	db.Create(&mcSub)

	return mcSub, choice.ID
}

func TestGetAllMultipleChoiceAnswerSubmissions(t *testing.T) {
	db := setupTestDBMCSubmission(t)
	ctrl := controller.NewMultipleChoiceAnswerSubmissionController(db)

	_, _ = createTestData(db)

	results, err := ctrl.GetAllMultipleChoiceAnswerSubmissions()
	assert.NoError(t, err)
	assert.Len(t, results, 1)
}

func TestGetMultipleChoiceAnswerSubmission(t *testing.T) {
	db := setupTestDBMCSubmission(t)
	ctrl := controller.NewMultipleChoiceAnswerSubmissionController(db)

	mcSub, _ := createTestData(db)

	result, err := ctrl.GetMultipleChoiceAnswerSubmission(mcSub.ID)
	assert.NoError(t, err)
	assert.Equal(t, mcSub.ID, result.ID)
}

func TestGetMultipleChoiceAnswerSubmissionsBySubmissionID(t *testing.T) {
	db := setupTestDBMCSubmission(t)
	ctrl := controller.NewMultipleChoiceAnswerSubmissionController(db)

	_, _ = createTestData(db)

	results, err := ctrl.GetMultipleChoiceAnswerSubmissionsBySubmissionID(1)
	assert.NoError(t, err)
	assert.Len(t, results, 1)
}

func TestUpdateMultipleChoiceAnswerSubmission(t *testing.T) {
	db := setupTestDBMCSubmission(t)
	ctrl := controller.NewMultipleChoiceAnswerSubmissionController(db)

	mcSub, _ := createTestData(db)

	newChoice := model.MultipleChoiceAnswer{
		QuestionID:  mcSub.QuestionID,
		AnswerLabel: "5",
		IsExpected:  false,
	}
	db.Create(&newChoice)

	mcSub.ChoiceID = newChoice.ID
	updated, err := ctrl.UpdateMultipleChoiceAnswerSubmission(&mcSub)
	assert.NoError(t, err)
	assert.Equal(t, newChoice.ID, updated.ChoiceID)
}

func TestDeleteMultipleChoiceAnswerSubmission(t *testing.T) {
	db := setupTestDBMCSubmission(t)
	ctrl := controller.NewMultipleChoiceAnswerSubmissionController(db)

	mcSub, _ := createTestData(db)

	deleted, err := ctrl.DeleteMultipleChoiceAnswerSubmission(mcSub.ID)
	assert.NoError(t, err)
	assert.Equal(t, mcSub.ID, deleted.ID)

	var check model.MultipleChoiceAnswerSubmission
	err = db.First(&check, mcSub.ID).Error
	assert.Error(t, err)
}
