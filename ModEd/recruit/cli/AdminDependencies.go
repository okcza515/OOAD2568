// MEP-1003 Student Recruitment
package cli

import (
	"ModEd/recruit/controller"
)

type AdminDependencies struct {
	ApplicantController       *controller.ApplicantController
	ApplicationReportCtrl     *controller.ApplicationReportController
	InterviewCtrl             *controller.InterviewController
	AdminCtrl                 *controller.AdminController
	LoginCtrl                 *controller.LoginController
	AdminInterviewService     AdminInterviewService
}

