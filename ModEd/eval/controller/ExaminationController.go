package controller

import (
	"errors"
	"time"

	"ModEd/eval/model"
	"gorm.io/gorm"
)

type IExaminationController interface {
	CreateExam(exam *model.Examination) error
	PublishExam(examID uint) error
	CloseExam(examID uint) error
	GetAllExams() ([]model.Examination, error)
	UpdateExam(id uint, exam *model.Examination) error
	DeleteExam(id uint) error
}

type ExaminationController struct {
	db *gorm.DB
}

func NewExaminationController(db *gorm.DB) *ExaminationController {
	return &ExaminationController{db: db}
}

func (c *ExaminationController) CreateExam(exam *model.Examination) error {
	if exam.Exam_name == "" || exam.InstructorID == 0 || exam.CourseID == 0 || exam.CurriculumID == 0 {
		return errors.New("missing required fields")
	}
	exam.ExamStatus = "draft"
	exam.Create_at = time.Now()

	return c.db.Create(exam).Error
}

func (c *ExaminationController) PublishExam(examID uint) error {
	var exam model.Examination
	if err := c.db.First(&exam, examID).Error; err != nil {
		return err
	}

	if time.Now().Before(exam.Start_date) {
		return errors.New("cannot publish exam before start date")
	}

	exam.ExamStatus = "published"
	return c.db.Save(&exam).Error
}

func (c *ExaminationController) CloseExam(examID uint) error {
	var exam model.Examination
	if err := c.db.First(&exam, examID).Error; err != nil {
		return err
	}

	if exam.ExamStatus != "published" {
		return errors.New("exam must be published before it can be closed")
	}	

	if time.Now().Before(exam.End_date) {
		return errors.New("cannot close exam before the end date")
	}

	exam.ExamStatus = "closed"
	return c.db.Save(&exam).Error
}

func (c *ExaminationController) GetAllExams() ([]model.Examination, error) {
	var exams []model.Examination
	if err := c.db.Find(&exams).Error; err != nil {
		return nil, err
	}
	return exams, nil
}

func (c *ExaminationController) UpdateExam(id uint, exam *model.Examination) error {
	return c.db.Model(&model.Examination{}).Where("id = ?", id).Updates(exam).Error
}

func (c *ExaminationController) DeleteExam(id uint) error {
	c.db.Where("ExamID = ?",id).Delete(&model.Question{})
	c.db.Where("ExamID = ?",id).Delete(&model.Answer{})
	c.db.Where("ExaminationID = ?",id).Delete(&model.Result{})
	return c.db.Where("id = ?", id).Delete(&model.Examination{}).Error
}
