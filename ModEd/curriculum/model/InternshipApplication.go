// MEP-1009 Student Internship
package model

import (
	"ModEd/core"
	"time"
)

type InternshipApplication struct {
	core.BaseModel
	TurninDate            time.Time        `gorm:"not null"`
	ApprovalAdvisorStatus ApprovedStatus   `gorm:"type:varchar(20)"`
	ApprovalCompanyStatus ApprovedStatus   `gorm:"type:varchar(20)"`
	AdvisorCode           uint             `gorm:"not null"`
	Advisor               Advisor          `gorm:"foreignKey:AdvisorCode;references:ID"`
	CompanyId             uint             `gorm:"not null"`
	Company               Company          `gorm:"foreignKey:CompanyId;references:ID"`
	InternshipReportId    uint             `gorm:"not null"`
	InternshipReport      InternshipReport `gorm:"foreignKey:InternshipReportId;references:ID"`
	SupervisorReviewId    uint             `gorm:"not null"`
	SupervisorReview      SupervisorReview `gorm:"foreignKey:SupervisorReviewId;references:ID"`
	StudentCode           string           `gorm:"type:varchar(255);not null"`
	Student               InternStudent    `gorm:"foreignKey:StudentCode;references:StudentCode"`
}
