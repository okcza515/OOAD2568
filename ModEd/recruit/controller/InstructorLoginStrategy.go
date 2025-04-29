package controller

import (
	"ModEd/recruit/model"

	"gorm.io/gorm"
)

type InstructorLoginStrategy struct {
	DB *gorm.DB
}

func (s *InstructorLoginStrategy) ApplyLogin(req LoginRequest) (bool, error) {
	var instructor model.Instructor
	err := s.DB.Where("username = ?", req.Username).First(&instructor).Error
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return false, nil
		}
		return false, err
	}
	return true, nil
}
