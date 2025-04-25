package model

import "time"

type ExaminationBuilder struct {
	exam *Examination
}


func NewExaminationBuilder() *ExaminationBuilder {
	return &ExaminationBuilder{exam: &Examination{}}
}

type ExamOption func(*Examination)

func WithExamName(name string) ExamOption {
	return func(exam *Examination) {
		exam.Exam_name = name
	}
}

func WithInstructorID(instructorID uint) ExamOption {
	return func(exam *Examination) {
		exam.Instructor_id = instructorID
	}
}

func WithCourseId(courseId uint) ExamOption {
	return func(exam *Examination) {
		exam.CourseId = courseId
	}
}

func WithCurriculumId(curriculumId uint) ExamOption {
	return func(exam *Examination) {
		exam.CurriculumId = curriculumId
	}
}

func WithCriteria(criteria string) ExamOption {
	return func(exam *Examination) {
		exam.Criteria = criteria
	}
}

func WithDescription(description string) ExamOption {
	return func(exam *Examination) {
		exam.Description = description
	}
}

func WithExamDate(examDate time.Time) ExamOption {
	return func(exam *Examination) {
		exam.Exam_date = examDate
	}
}

func (b *ExaminationBuilder) Build(options ...ExamOption) *Examination {
	for _, option := range options {
		option(b.exam)
	}
	return b.exam
}