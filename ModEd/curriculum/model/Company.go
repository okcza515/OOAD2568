// MEP-1009 Student Internship
package model

import "ModEd/core"

type Company struct {
	core.BaseModel
	CompanyName    string `gorm:"type:varchar(255)"`
	CompanyAddress string `gorm:"type:varchar(255)"`
}
