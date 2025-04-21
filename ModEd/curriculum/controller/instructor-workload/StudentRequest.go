package controller

import (
	model "ModEd/curriculum/model/instructor-workload"

	"gorm.io/gorm"

	"errors"
)

type StudentRequestController struct {
	DB *gorm.DB
}

func (src *StudentRequestController) CreateStudentRequest(studentRequest model.StudentRequest) error {
	return src.DB.Create(&studentRequest).Error
}

func (src *StudentRequestController) GetStudentRequestsByInstructorId(instructorId uint) ([]model.StudentRequest, error) {
	var studentRequests []model.StudentRequest
	err := src.DB.Where("instructor_id = ?", instructorId).Find(&studentRequests).Error
	if err != nil {
		return nil, err
	}
	return studentRequests, nil
}

func (src *StudentRequestController) ReviewStudentRequest(id uint, review string, comment string) error {
	var studentRequest model.StudentRequest
	err := src.DB.First(&studentRequest, id).Error
	if err != nil {
		return err
	}
	if review != "accept" && review != "reject" {
		return errors.New("invalid review value, must be 'accept' or 'reject'")
	}
	studentRequest.Review = review
	studentRequest.Comment = comment
	return src.DB.Save(&studentRequest).Error
}
