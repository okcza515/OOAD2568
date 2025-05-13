package model

import "ModEd/core/validation"

type RequestResignationStudent struct {
	BaseStandardRequest
	StudentCode string `gorm:"type:text;default:'';not null" validation:"studentId"`
}

func (r *RequestResignationStudent) Validate() error {
	modelValidator := validation.NewModelValidator()

	if err := modelValidator.ModelValidate(r); err != nil {
		return err
	}

	return nil
}
