package model

import (
	"gorm.io/gorm"
	"time"
)

type InternshipApplication struct {
	gorm.Model
	InternshipApplicationId int           `gorm:"primaryKey autoIncrement"`
	Student                 InternStudent `gorm:"foreignKey:InternStudentId"`
	Company                 Company       `gorm:"foreignKey:CompanyId"`
	Advisor                 string        `gorm:"not null"`
	TurninDate              time.Time     `gorm:"not null"`
	ApprovalAdvisorStatus   bool          `gorm:"not null"`
	ApprovalCompanyStatus   bool          `gorm:"not null"`
}
