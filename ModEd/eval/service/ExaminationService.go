package service

import (
	"errors"
	"time"

	"ModEd/eval/model"
	"gorm.io/gorm"
)

type ExaminationService struct {
	db *gorm.DB
}

func NewExaminationService(db *gorm.DB) *ExaminationService {
	return &ExaminationService{db: db}
}

// Create new exam
func (s *ExaminationService) CreateExam(exam *model.Examination) error {
	if exam.Exam_name == "" || exam.Instructor_id == 0 || exam.CourseId == 0 || exam.CurriculumId == 0 {
		return errors.New("missing required fields")
	}
	exam.ExamStatus = "draft"
	exam.Create_at = time.Now()

	return s.db.Create(exam).Error
}

// Publish exam
func (s *ExaminationService) PublishExam(examID uint) error {
	var exam model.Examination
	if err := s.db.First(&exam, examID).Error; err != nil {
		return err
	}

	if time.Now().Before(exam.Start_date) {
		return errors.New("cannot publish exam before start date")
	}

	exam.ExamStatus = "published"
	return s.db.Save(&exam).Error
}

// Close exam
func (s *ExaminationService) CloseExam(examID uint) error {
	var exam model.Examination
	if err := s.db.First(&exam, examID).Error; err != nil {
		return err
	}

	if exam.ExamStatus != "published" {
		return errors.New("exam must be published before it can be closed")
	}	

	if time.Now().Before(exam.End_date) {
		return errors.New("cannot close exam before the end date")
	}

	exam.ExamStatus = "closed"
	return s.db.Save(&exam).Error
}

// List all exams
func (s *ExaminationService) GetAllExams() ([]model.Examination, error) {
	var exams []model.Examination
	if err := s.db.Find(&exams).Error; err != nil {
		return nil, err
	}
	return exams, nil
}

// Update exam
func (s *ExaminationService) UpdateExam(id uint, exam *model.Examination) error {
	return s.db.Model(&model.Examination{}).Where("id = ?", id).Updates(exam).Error
}

// Delete exam
func (s *ExaminationService) DeleteExam(id uint) error {
	return s.db.Where("id = ?", id).Delete(&model.Examination{}).Error
}