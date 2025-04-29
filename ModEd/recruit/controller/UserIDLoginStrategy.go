package controller

import (
	"ModEd/recruit/model"
	"fmt"

	"gorm.io/gorm"
)

type UserIDLoginStrategy struct {
	DB *gorm.DB
}

func (s *UserIDLoginStrategy) CheckUsername(username string) (bool, error) {
	return false, fmt.Errorf("UserID login does not support username login")
}

func (s *UserIDLoginStrategy) CheckUsernameAndPassword(username, password string) (bool, error) {
	return false, fmt.Errorf("UserID login does not support username/password login")
}

func (s *UserIDLoginStrategy) CheckID(id string) (bool, error) {
	var applicant model.Applicant
	err := s.DB.First(&applicant, "id = ?", id).Error
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return false, nil
		}
		return false, err
	}
	return true, nil
}
