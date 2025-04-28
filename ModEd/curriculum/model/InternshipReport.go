// MEP-1009 Student Internship
package model

import "ModEd/core"

type InternshipReport struct {
	core.BaseModel
	ReportScore int `gorm:"type:int"`
}
