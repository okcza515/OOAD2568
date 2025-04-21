//MEP-1009 Student Internship
package model

import (
commonModel "ModEd/common/model"

"gorm.io/gorm"
)

type Advisor struct {
gorm.Model
AdvisorId int                    `gorm:"not null;index"`
Advisor   commonModel.Instructor `gorm:"foreignKey:AdvisorId"`
Students  []commonModel.Student  `gorm:"many2many:student_advisor_students;"`
}