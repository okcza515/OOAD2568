// MEP-1009 Student Internship
package controller

import (
	"gorm.io/gorm"
)

type InternshipModuleWrapper struct {
	CurriculumController  *CurriculumController
	InternshipApplication *InternshipApplicationController
	Review                *ReviewController
	Report                *ReportController
	Approved              *ApprovedController
	Company               *CompanyController
	InternStudent         *InternStudentController
	GenericImport         *GenericImportController
}

func NewInternshipModuleWrapper(
	db *gorm.DB,
	curriculumController CurriculumControllerInterface,
) *InternshipModuleWrapper {
	return &InternshipModuleWrapper{
		InternshipApplication: NewInternshipApplicationController(db),
		Review:                NewReviewController(db),
		Report:                NewReportController(db),
		Approved:              NewApprovedController(db),
		Company:               NewCompanyController(db),
		InternStudent:         NewInternStudentController(db),
		GenericImport:         CreateGenericImportController(db),
	}
}
