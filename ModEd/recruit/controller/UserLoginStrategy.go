// MEP-1003 Student Recruitment
package controller

import (
	"gorm.io/gorm"
)

type UserIDLoginStrategy struct {
	DB *gorm.DB
}

func (s *UserIDLoginStrategy) ApplyLogin(req LoginRequest, model interface{}) (bool, error) {

	err := s.DB.First(model, "id = ?", req.ID).Error
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return false, nil
		}
		return false, err
	}
	return true, nil
}
