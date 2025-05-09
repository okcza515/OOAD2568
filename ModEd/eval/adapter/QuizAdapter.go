// MEP-1006
package adapter

import (
	"ModEd/eval/controller"
	"ModEd/eval/model"

	"gorm.io/gorm"
)

type QuizAdapter struct {
	quiz *model.Quiz
}

func NewQuizAdapter(quiz *model.Quiz) *QuizAdapter {
	return &QuizAdapter{quiz: quiz}
}

func (a *QuizAdapter) ToExamination() *model.Examination {
	return &model.Examination{
		ExamName:     a.quiz.Title,
		Description:  a.quiz.Description,
		ExamStatus:   model.ExamStatus(a.quiz.Status),
		StartDate:    a.quiz.StartDate,
		EndDate:      a.quiz.EndDate,
		InstructorID: a.quiz.InstructorID,
		CourseID:     a.quiz.CourseID,
		Attempt:      uint(a.quiz.Attempts),
	}
}

func (a *QuizAdapter) FromExamination(exam *model.Examination) {
	a.quiz.Title = exam.ExamName
	a.quiz.Description = exam.Description
	a.quiz.Status = string(exam.ExamStatus)
	a.quiz.StartDate = exam.StartDate
	a.quiz.EndDate = exam.EndDate
	a.quiz.InstructorID = exam.InstructorID
	a.quiz.CourseID = exam.CourseID
	a.quiz.Attempts = int(exam.Attempt)
}

func (a *QuizAdapter) GetQuiz() *model.Quiz {
	return a.quiz
}

type QuizControllerAdapter struct {
	examController     *controller.ExaminationController
	questionController *controller.QuestionController
}

func NewQuizControllerAdapter(db *gorm.DB) *QuizControllerAdapter {
	return &QuizControllerAdapter{
		examController:     controller.NewExaminationController(db),
		questionController: controller.NewQuestionController(db),
	}
}

func (a *QuizControllerAdapter) CreateQuiz(quiz *model.Quiz) (*model.Quiz, error) {
	exam := &model.Examination{
		ExamName:     quiz.Title,
		Description:  quiz.Description,
		ExamStatus:   model.ExamStatus(quiz.Status),
		StartDate:    quiz.StartDate,
		EndDate:      quiz.EndDate,
		InstructorID: quiz.InstructorID,
		CourseID:     quiz.CourseID,
		Attempt:      uint(quiz.Attempts),
	}

	examID, err := a.examController.CreateExam(exam)
	if err != nil {
		return nil, err
	}

	quiz.ID = examID
	return quiz, nil
}

func (a *QuizControllerAdapter) UpdateQuiz(quiz *model.Quiz) (*model.Quiz, error) {
	exam := &model.Examination{
		ExamName:     quiz.Title,
		Description:  quiz.Description,
		ExamStatus:   model.ExamStatus(quiz.Status),
		StartDate:    quiz.StartDate,
		EndDate:      quiz.EndDate,
		InstructorID: quiz.InstructorID,
		CourseID:     quiz.CourseID,
		Attempt:      uint(quiz.Attempts),
	}

	updatedExam, err := a.examController.UpdateExam(exam)
	if err != nil {
		return nil, err
	}

	quiz.ID = updatedExam.ID
	return quiz, nil
}

func (a *QuizControllerAdapter) DeleteQuiz(id uint) error {
	exam := &model.Examination{
		ExamStatus: model.Hidden,
	}
	_, err := a.examController.UpdateExam(exam)
	return err
}

func (a *QuizControllerAdapter) GetQuiz(id uint) (*model.Quiz, error) {
	exam, err := a.examController.GetExam(id)
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
		CourseID:     exam.CourseID,
		Attempts:     int(exam.Attempt),
	}

	return quiz, nil
}

func (a *QuizControllerAdapter) GetAllQuizzes() ([]*model.Quiz, error) {
	exams, err := a.examController.GetAllExams()
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
			CourseID:     exam.CourseID,
			Attempts:     int(exam.Attempt),
		}
		quizzes = append(quizzes, quiz)
	}

	return quizzes, nil
}
