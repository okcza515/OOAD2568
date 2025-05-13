package model

import "github.com/go-playground/validator/v10"

type RequestLeaveStudent struct {
	BaseLeaveRequest
	StudentCode string `gorm:"not null" validate:"required"`
}

func (requestLeaveStudent RequestLeaveStudent) Validate() error {
	validate := validator.New()
	if err := validate.Struct(requestLeaveStudent); err != nil {
		return err
	}
	return nil
}
