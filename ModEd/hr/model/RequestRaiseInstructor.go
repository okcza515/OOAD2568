package model

import "github.com/go-playground/validator/v10"

type RequestRaiseInstructor struct {
	BaseStandardRequest
	InstructorCode string  `gorm:"not null" validate:"required"`
	TargetSalary   float64 `gorm:"not null" validate:"required"`
}

func (requestRaiseInstructor RequestRaiseInstructor) Validate() error {
	validate := validator.New()
	if err := validate.Struct(requestRaiseInstructor); err != nil {
		return err
	}
	return nil
}
