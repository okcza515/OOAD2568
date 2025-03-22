package model

import (
	modelWorkload "ModEd/curriculum/model/instructor-workload"
	"time"

	"gorm.io/gorm"
)

type InternshipApplication struct {
	gorm.Model
	TurninDate            time.Time                    `gorm:"not null"`
	ApprovalAdvisorStatus bool                         `gorm:"not null"`
	ApprovalCompanyStatus bool                         `gorm:"not null"`
	AdvisorCode           uint                         `gorm:"not null"`
	Advisor               modelWorkload.StudentAdvisor `gorm:"foreignKey:AdvisorCode;references:ID"`
	CompanyId             uint                         `gorm:"not null"`
	Company               Company                      `gorm:"foreignKey:CompanyId;references:ID"`
	InternshipReportId    uint                         `gorm:"not null"`
	InternshipReport      InternshipReport             `gorm:"foreignKey:InternshipReportId;references:ID"`
	SupervisorReviewId    uint                         `gorm:"not null"`
	SupervisorReview      SupervisorReview             `gorm:"foreignKey:SupervisorReviewId;references:ID"`
	InternshipScheduleId  uint                         `gorm:"not null"`
	InternshipSchedule    InternshipSchedule           `gorm:"foreignKey:InternshipScheduleId;references:ID"`
}
