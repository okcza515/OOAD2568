package model

import "github.com/go-playground/validator/v10"

type RequestLeaveInstructor struct {
	BaseLeaveRequest
	InstructorCode string `gorm:"not null" validate:"required"`
}

func (requestLeaveInstructor RequestLeaveInstructor) Validate() error {
	validate := validator.New()
	if err := validate.Struct(requestLeaveInstructor); err != nil {
		return err
	}
	return nil
}
