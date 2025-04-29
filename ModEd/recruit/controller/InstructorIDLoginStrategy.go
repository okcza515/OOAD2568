package controller

import (
	"ModEd/recruit/model"
	"fmt"

	"gorm.io/gorm"
)

type InstructorIDLoginStrategy struct {
	DB *gorm.DB
}

func (s *InstructorIDLoginStrategy) CheckUsername(username string) (bool, error) {
	return false, fmt.Errorf("InstructorID login does not support username login")
}

func (s *InstructorIDLoginStrategy) CheckUsernameAndPassword(username, password string) (bool, error) {
	return false, fmt.Errorf("InstructorID login does not support username/password login")
}

func (s *InstructorIDLoginStrategy) CheckID(id string) (bool, error) {
	var instructor model.Instructor
	err := s.DB.First(&instructor, "id = ?", id).Error
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return false, nil
		}
		return false, err
	}
	return true, nil
}
