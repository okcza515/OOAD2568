package controller

import (
	"ModEd/recruit/model"
	"fmt"

	"gorm.io/gorm"
)

type UsernamePasswordLoginStrategy struct {
	DB *gorm.DB
}

func (s *UsernamePasswordLoginStrategy) CheckUsername(username string) (bool, error) {
	var admin model.Admin
	err := s.DB.First(&admin, "username = ?", username).Error
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return false, nil
		}
		return false, err
	}
	return true, nil
}

func (s *UsernamePasswordLoginStrategy) CheckUsernameAndPassword(username, password string) (bool, error) {
	var admin model.Admin
	err := s.DB.First(&admin, "username = ? AND password = ?", username, password).Error
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return false, nil
		}
		return false, err
	}
	return true, nil
}

func (s *UsernamePasswordLoginStrategy) CheckID(id string) (bool, error) {
	return false, fmt.Errorf("Username/Password strategy does not support ID login")
}
