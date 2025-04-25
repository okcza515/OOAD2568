package model

type QuestionBuilder struct {
	question Question
}

func NewQuestionBuilder() *QuestionBuilder {
	return &QuestionBuilder{question: Question{}}
}

func (qb *QuestionBuilder) ID(id uint) *QuestionBuilder {
	qb.question.ID = id
	return qb
}

func (qb *QuestionBuilder) ExamID(examID uint) *QuestionBuilder {
	qb.question.Exam_id = examID
	return qb
}

func (qb *QuestionBuilder) Examination(exam Examination) *QuestionBuilder {
	qb.question.Examination = exam
	return qb
}

func (qb *QuestionBuilder) Detail(detail string) *QuestionBuilder {
	qb.question.Question_detail = detail
	return qb
}

func (qb *QuestionBuilder) Type(qType QuestionType) *QuestionBuilder {
	qb.question.Question_type = qType
	return qb
}

func (qb *QuestionBuilder) CorrectAnswer(answer string) *QuestionBuilder {
	qb.question.Correct_answer = answer
	return qb
}

func (qb *QuestionBuilder) Score(score float64) *QuestionBuilder {
	qb.question.Score = score
	return qb
}

func (qb *QuestionBuilder) Build() Question {
	return qb.question
}
