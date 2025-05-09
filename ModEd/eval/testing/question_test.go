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

func setupTestDBQuestion(t *testing.T) *gorm.DB {
	db, err := gorm.Open(sqlite.Open(":memory:"), &gorm.Config{})
	assert.NoError(t, err)

	err = db.AutoMigrate(&model.Examination{}, &model.ExamSection{}, &model.Question{})
	assert.NoError(t, err)

	return db
}

func TestCreateQuestion(t *testing.T) {
	db := setupTestDBQuestion(t)

	ctrlExam := controller.NewExaminationController(db)
	ctrlExamSec := controller.NewExamSectionController(db)
	ctrlQuestion := controller.NewQuestionController(db)

	exam := &model.Examination{
		ExamName:     "Test Exam",
		CourseID:     1,
		InstructorID: 1,
		ExamStatus:   "DRAFT",
		Description:  "Sample description",
		StartDate:    time.Now(),
		EndDate:      time.Now().Add(2 * time.Hour),
	}
	_, err := ctrlExam.CreateExam(exam)
	assert.NoError(t, err)
	assert.NotZero(t, exam.ID)

	section := &model.ExamSection{
		ExamID:       exam.ID,
		SectionNo:    1,
		Description:  "Section 1",
		NumQuestions: 1,
		Score:        10.0,
	}
	_, err = ctrlExamSec.CreateExamSection(section)
	assert.NoError(t, err)
	assert.NotZero(t, section.ID)

	question := &model.Question{
		SectionID:      section.ID,
		Score:          5.0,
		ActualQuestion: "What is the capital of France?",
		QuestionType:   "MultipleChoiceQuestion",
	}
	err = ctrlQuestion.CreateQuestion(question)
	assert.NoError(t, err)
	assert.NotZero(t, question.ID)

	var found model.Question
	err = db.Preload("Section").First(&found, question.ID).Error
	assert.NoError(t, err)
	assert.Equal(t, question.ActualQuestion, found.ActualQuestion)
	assert.Equal(t, section.ID, found.SectionID)
	assert.Equal(t, exam.ID, found.Section.ExamID)
}

func TestUpdateQuestion(t *testing.T) {
	db := setupTestDBQuestion(t)

	ctrlExam := controller.NewExaminationController(db)
	ctrlExamSec := controller.NewExamSectionController(db)
	ctrlQuestion := controller.NewQuestionController(db)

	exam := &model.Examination{
		ExamName:     "Test Exam",
		CourseID:     1,
		InstructorID: 1,
		ExamStatus:   "DRAFT",
		StartDate:    time.Now(),
		EndDate:      time.Now().Add(2 * time.Hour),
	}
	_, err := ctrlExam.CreateExam(exam)
	assert.NoError(t, err)

	section := &model.ExamSection{
		ExamID:       exam.ID,
		SectionNo:    1,
		Description:  "Section 1",
		NumQuestions: 1,
		Score:        10.0,
	}
	_, err = ctrlExamSec.CreateExamSection(section)
	assert.NoError(t, err)

	question := &model.Question{
		SectionID:      section.ID,
		Score:          5.0,
		ActualQuestion: "Old Question",
		QuestionType:   "MULTIPLE_CHOICE",
	}
	err = ctrlQuestion.CreateQuestion(question)
	assert.NoError(t, err)

	question.ActualQuestion = "Updated Question"
	err = ctrlQuestion.UpdateQuestion(question.ID, question)
	assert.NoError(t, err)

	var updated model.Question
	err = db.First(&updated, question.ID).Error
	assert.NoError(t, err)
	assert.Equal(t, "Updated Question", updated.ActualQuestion)
}

func TestDeleteQuestion(t *testing.T) {
	db := setupTestDBQuestion(t)

	ctrlExam := controller.NewExaminationController(db)
	ctrlExamSec := controller.NewExamSectionController(db)
	ctrlQuestion := controller.NewQuestionController(db)

	exam := &model.Examination{
		ExamName:     "Test Exam",
		CourseID:     1,
		InstructorID: 1,
		ExamStatus:   "DRAFT",
		StartDate:    time.Now(),
		EndDate:      time.Now().Add(2 * time.Hour),
	}
	_, err := ctrlExam.CreateExam(exam)
	assert.NoError(t, err)

	section := &model.ExamSection{
		ExamID:       exam.ID,
		SectionNo:    1,
		Description:  "Section 1",
		NumQuestions: 1,
		Score:        10.0,
	}
	_, err = ctrlExamSec.CreateExamSection(section)
	assert.NoError(t, err)

	question := &model.Question{
		SectionID:      section.ID,
		Score:          5.0,
		ActualQuestion: "To be deleted",
		QuestionType:   "MultipleChoiceQuestion",
	}
	err = ctrlQuestion.CreateQuestion(question)
	assert.NoError(t, err)

	err = ctrlQuestion.DeleteQuestion(question.ID)
	assert.NoError(t, err)

	var found model.Question
	err = db.First(&found, question.ID).Error
	assert.Error(t, err)
	assert.Equal(t, gorm.ErrRecordNotFound, err)
}

func TestGetQuestionsByExamID(t *testing.T) {
	db := setupTestDBQuestion(t)

	ctrlExam := controller.NewExaminationController(db)
	ctrlExamSec := controller.NewExamSectionController(db)
	ctrlQuestion := controller.NewQuestionController(db)

	exam := &model.Examination{
		ExamName:     "Test Exam",
		CourseID:     1,
		InstructorID: 1,
		ExamStatus:   "DRAFT",
		StartDate:    time.Now(),
		EndDate:      time.Now().Add(2 * time.Hour),
	}
	_, err := ctrlExam.CreateExam(exam)
	assert.NoError(t, err)

	section := &model.ExamSection{
		ExamID:       exam.ID,
		SectionNo:    1,
		Description:  "Section 1",
		NumQuestions: 2,
		Score:        20.0,
	}
	_, err = ctrlExamSec.CreateExamSection(section)
	assert.NoError(t, err)

	q1 := &model.Question{
		SectionID:      section.ID,
		Score:          10.0,
		ActualQuestion: "What is 2 + 2?",
		QuestionType:   "ShortAnswerQuestion",
	}
	q2 := &model.Question{
		SectionID:      section.ID,
		Score:          10.0,
		ActualQuestion: "What is the capital of Spain?",
		QuestionType:   "MultipleChoiceQuestion",
	}
	err = ctrlQuestion.CreateQuestion(q1)
	assert.NoError(t, err)
	err = ctrlQuestion.CreateQuestion(q2)
	assert.NoError(t, err)

	questions, err := ctrlQuestion.GetQuestionsByExamID(exam.ID)
	assert.NoError(t, err)
	assert.Len(t, questions, 2)

	foundQuestions := map[string]bool{
		q1.ActualQuestion: false,
		q2.ActualQuestion: false,
	}
	for _, q := range questions {
		_, exists := foundQuestions[q.ActualQuestion]
		assert.True(t, exists)
		foundQuestions[q.ActualQuestion] = true
	}
}
