package controller

import (
	"ModEd/recruit/model"
	"gorm.io/gorm"
)

// AdminLoginStrategy implements LoginStrategy for admin login
type AdminLoginStrategy struct {
	DB *gorm.DB
}

// CheckUsername checks if an admin with the given username exists
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

// CheckUsernameAndPassword checks if the given username and password match for an admin
func (s *AdminLoginStrategy) CheckUsernameAndPassword(username, password string) (bool, error) {
	var admin model.Admin
	err := s.DB.Where("username = ?", username).First(&admin).Error
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return false, nil
		}
		return false, err
	}

	// Assuming hashed passwords are stored, compare using a method like bcrypt
	if admin.Password != password {
		return false, nil
	}
	return true, nil
}
