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

func setupTestDBTFAnswerSubmission(t *testing.T) *gorm.DB {
	db, err := gorm.Open(sqlite.Open(":memory:"), &gorm.Config{})
	assert.NoError(t, err)

	err = db.AutoMigrate(&model.Question{}, &model.Examination{}, &model.ExamSection{}, &model.TrueFalseAnswerSubmission{})
	assert.NoError(t, err)

	return db
}

func createTestDataTF(db *gorm.DB) (model.TrueFalseAnswerSubmission, uint) {
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
		ActualQuestion: "Is 2 + 2 = 4?",
		QuestionType:   model.TrueFalseQuestion,
	}
	db.Create(&question)

	tfAnsSub := model.TrueFalseAnswerSubmission{
		QuestionID:    question.ID,
		SubmissionID:  1,
		StudentAnswer: true,
	}
	db.Create(&tfAnsSub)

	return tfAnsSub, question.ID
}

func TestGetAllTrueFalseAnswerSubmissions(t *testing.T) {
	db := setupTestDBTFAnswerSubmission(t)
	ctrl := controller.NewTrueFalseAnswerSubmissionController(db)

	_, _ = createTestDataTF(db)

	results, err := ctrl.GetAllTrueFalseAnswerSubmissions()
	assert.NoError(t, err)
	assert.Len(t, results, 1)
}

func TestGetTrueFalseAnswerSubmission(t *testing.T) {
	db := setupTestDBTFAnswerSubmission(t)
	ctrl := controller.NewTrueFalseAnswerSubmissionController(db)

	tfAnsSub, _ := createTestDataTF(db)

	result, err := ctrl.GetTrueFalseAnswerSubmission(tfAnsSub.ID)
	assert.NoError(t, err)
	assert.Equal(t, tfAnsSub.ID, result.ID)
}

func TestGetTrueFalseAnswerSubmissionsBySubmissionID(t *testing.T) {
	db := setupTestDBTFAnswerSubmission(t)
	ctrl := controller.NewTrueFalseAnswerSubmissionController(db)

	_, _ = createTestDataTF(db)

	results, err := ctrl.GetTrueFalseAnswerSubmissionsBySubmissionID(1)
	assert.NoError(t, err)
	assert.Len(t, results, 1)
}

func TestUpdateTrueFalseAnswerSubmission(t *testing.T) {
	db := setupTestDBTFAnswerSubmission(t)
	ctrl := controller.NewTrueFalseAnswerSubmissionController(db)

	tfAnsSub, _ := createTestDataTF(db)

	tfAnsSub.StudentAnswer = false
	updated, err := ctrl.UpdateTrueFalseAnswerSubmission(&tfAnsSub)
	assert.NoError(t, err)
	assert.False(t, updated.StudentAnswer)
}

func TestDeleteTrueFalseAnswerSubmission(t *testing.T) {
	db := setupTestDBTFAnswerSubmission(t)
	ctrl := controller.NewTrueFalseAnswerSubmissionController(db)

	tfAnsSub, _ := createTestDataTF(db)

	deleted, err := ctrl.DeleteTrueFalseAnswerSubmission(tfAnsSub.ID)
	assert.NoError(t, err)
	assert.Equal(t, tfAnsSub.ID, deleted.ID)

	var check model.TrueFalseAnswerSubmission
	err = db.First(&check, tfAnsSub.ID).Error
	assert.Error(t, err)
}
