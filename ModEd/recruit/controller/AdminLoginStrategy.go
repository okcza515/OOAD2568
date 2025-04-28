package controller

import (
	"ModEd/recruit/model"

	"gorm.io/gorm"
)

type AdminLoginStrategy struct {
	DB *gorm.DB
}

func (s *AdminLoginStrategy) CheckUsername(username string) (bool, error) {
	var admin model.Admin
	err := s.DB.Where("username = ?", username).First(&admin).Error
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return false, nil
		}
		return false, err
	}
	return true, nil
}
