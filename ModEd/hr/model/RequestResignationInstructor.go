package model

import "github.com/go-playground/validator/v10"

type RequestResignationInstructor struct {
	BaseStandardRequest
	InstructorCode string `gorm:"not null" validate:"required"`
}

func (requestResignationInstructor RequestResignationInstructor) Validate() error {
	validate := validator.New()
	if err := validate.Struct(requestResignationInstructor); err != nil {
		return err
	}
	return nil
}
