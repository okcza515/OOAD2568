// MEP-1003 Student Recruitment
package cli

import (
	"ModEd/recruit/controller"
	"gorm.io/gorm"
)

type InstructorDependencies struct {
	ViewInterviewService     InstructorViewInterviewDetailsService
	EvaluateApplicantService InstructorEvaluateApplicantService
	ApplicantReportService   ApplicantReportService
	LoginCtrl                *controller.LoginController
	DB                       *gorm.DB
}
