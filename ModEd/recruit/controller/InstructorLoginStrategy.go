package controller

import (
	"ModEd/recruit/model"

	"gorm.io/gorm"
)

type InstructorLoginStrategy struct {
	DB *gorm.DB
}

func (s *InstructorLoginStrategy) CheckUsername(username string) (bool, error) {
	var instructor model.Instructor
	err := s.DB.Where("username = ?", username).First(&instructor).Error
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return false, nil
		}
		return false, err
	}
	return true, nil
}
