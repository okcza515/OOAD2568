package model

import (
	"time"
	"gorm.io/gorm"
)

type InternshipApplication struct {
	gorm.Model
	InternshipApplicationId int           `gorm:"primaryKey autoIncrement"`
	Company                 Company       `gorm:"not null"`
	Mentor                  string        `gorm:"not null"`
	Advisor                 string        `gorm:"not null"`
	Student                 InternStudent `gorm:"foreignKey:internI"`
	TurninDate              time.Time     `gorm:"not null"`
	ApprovalAdvisorStatus   bool          `gorm:"not null"`
	ApprovalCompanyStatus   bool          `gorm:"not null"`
}
