package model

import "github.com/go-playground/validator/v10"

type RequestResignationStudent struct {
	BaseStandardRequest
	StudentCode string `gorm:"type:text;default:'';not null" validate:"required"`
}

func (requestResignationStudent RequestResignationStudent) Validate() error {
	validate := validator.New()
	if err := validate.Struct(requestResignationStudent); err != nil {
		return err
	}
	return nil
}
