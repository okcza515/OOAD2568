// MEP-1006
package controller

import (
	"ModEd/eval/model"

	"gorm.io/gorm"
)

type QuizAdapter struct {
	quiz *model.Quiz
}

func NewQuizAdapter(quiz *model.Quiz) *QuizAdapter {
	return &QuizAdapter{quiz: quiz}
}

func (a *QuizAdapter) ToExamination() *model.Exam {
	return &model.Exam{
		ExamName:     a.quiz.Title,
		Description:  a.quiz.Description,
		ExamStatus:   model.ExamStatus(a.quiz.Status),
		StartDate:    a.quiz.StartDate,
		EndDate:      a.quiz.EndDate,
		InstructorID: a.quiz.InstructorID,
		ClassID:      a.quiz.CourseID,
		Attempt:      uint(a.quiz.Attempts),
	}
}

func (a *QuizAdapter) FromExamination(exam *model.Exam) {
	a.quiz.Title = exam.ExamName
	a.quiz.Description = exam.Description
	a.quiz.Status = string(exam.ExamStatus)
	a.quiz.StartDate = exam.StartDate
	a.quiz.EndDate = exam.EndDate
	a.quiz.InstructorID = exam.InstructorID
	a.quiz.CourseID = exam.ClassID
	a.quiz.Attempts = uint(exam.Attempt)
}

func (a *QuizAdapter) GetQuiz() *model.Quiz {
	return a.quiz
}

type QuizControllerAdapter struct {
	examController     *ExamController
	questionController *QuestionController
}

func NewQuizControllerAdapter(db *gorm.DB) *QuizControllerAdapter {
	return &QuizControllerAdapter{
		examController:     NewExamController(db),
		questionController: NewQuestionController(db),
	}
}

func (a *QuizControllerAdapter) CreateQuiz(quiz *model.Quiz) (*model.Quiz, error) {
	exam := &model.Exam{
		ExamName:     quiz.Title,
		Description:  quiz.Description,
		ExamStatus:   model.ExamStatus(quiz.Status),
		StartDate:    quiz.StartDate,
		EndDate:      quiz.EndDate,
		InstructorID: quiz.InstructorID,
		ClassID:      quiz.CourseID,
		Attempt:      uint(quiz.Attempts),
	}

	err := a.examController.Insert(exam)
	if err != nil {
		return nil, err
	}

	quiz.ID = exam.ID
	return quiz, nil
}

func (a *QuizControllerAdapter) UpdateQuiz(quiz *model.Quiz) (*model.Quiz, error) {
	exam := &model.Exam{
		ExamName:     quiz.Title,
		Description:  quiz.Description,
		ExamStatus:   model.ExamStatus(quiz.Status),
		StartDate:    quiz.StartDate,
		EndDate:      quiz.EndDate,
		InstructorID: quiz.InstructorID,
		ClassID:      quiz.CourseID,
		Attempt:      uint(quiz.Attempts),
	}

	err := a.examController.UpdateByID(exam)
	if err != nil {
		return nil, err
	}

	quiz.ID = exam.ID
	return quiz, nil
}

func (a *QuizControllerAdapter) DeleteQuiz(id uint) error {
	exam := &model.Exam{
		ExamStatus: model.Hidden,
	}
	return a.examController.UpdateByID(exam)
}

func (a *QuizControllerAdapter) GetQuiz(id uint) (*model.Quiz, error) {
	exam, err := a.examController.RetrieveByID(id)
	if err != nil {
		return nil, err
	}

	quiz := &model.Quiz{
		Title:        exam.ExamName,
		Description:  exam.Description,
		Status:       string(exam.ExamStatus),
		StartDate:    exam.StartDate,
		EndDate:      exam.EndDate,
		InstructorID: exam.InstructorID,
		CourseID:     exam.ClassID,
		Attempts:     uint(exam.Attempt),
	}

	return quiz, nil
}

func (a *QuizControllerAdapter) GetAllQuizzes() ([]*model.Quiz, error) {
	exams, err := a.examController.List(nil)
	if err != nil {
		return nil, err
	}

	var quizzes []*model.Quiz
	for _, exam := range exams {
		quiz := &model.Quiz{
			Title:        exam.ExamName,
			Description:  exam.Description,
			Status:       string(exam.ExamStatus),
			StartDate:    exam.StartDate,
			EndDate:      exam.EndDate,
			InstructorID: exam.InstructorID,
			CourseID:     exam.ClassID,
			Attempts:     uint(exam.Attempt),
		}
		quizzes = append(quizzes, quiz)
	}

	return quizzes, nil
}
