package controller

import (
	"fmt"
	"reflect"

	"gorm.io/gorm"
)

type AdminLoginStrategy struct {
	DB *gorm.DB
}

func (s *AdminLoginStrategy) ApplyLogin(req LoginRequest, model interface{}) (bool, error) {
	modelVal := reflect.ValueOf(model)
	if modelVal.Kind() != reflect.Ptr || modelVal.IsNil() {
		return false, fmt.Errorf("model must be a non-nil pointer")
	}

	usernameField := modelVal.Elem().FieldByName("Username")
	passwordField := modelVal.Elem().FieldByName("Password")

	if !usernameField.IsValid() || !passwordField.IsValid() {
		return false, fmt.Errorf("model must have Username and Password fields")
	}

	// ค้นหาผู้ใช้ในฐานข้อมูล
	err := s.DB.Where("username = ?", req.Username).First(model).Error
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return false, nil
		}
		return false, err
	}

	// ตรวจสอบรหัสผ่าน
	if passwordField.String() != req.Password {
		return false, nil
	}
	return true, nil
}
