// MEP-1009 Student Internship
package model

import (
	commonModel "ModEd/common/model"
	"ModEd/core"
)

type InternStudent struct {
	core.BaseModel
	InternStatus InternStatus        `gorm:"type:varchar(20)"`
	StudentCode  string              `gorm:"type:varchar(255);not null;unique" csv:"student_code"`
	Student      commonModel.Student `gorm:"foreignKey:StudentCode;references:StudentCode"`
}
