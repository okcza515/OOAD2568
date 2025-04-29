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
	InternStudent         *InternStudentController
	GenericImport         *GenericImportController
}

func NewInternshipModuleWrapper(
	db *gorm.DB,
	curriculumController CurriculumControllerInterface,
) *InternshipModuleWrapper {
	return &InternshipModuleWrapper{
		InternshipApplication: CreateInternshipApplicationController(db),
		Review:                CreateReviewController(db),
		Report:                CreateReportController(db),
		Approved:              CreateApprovedController(db),
		InternStudent:         CreateInternStudentController(db),
		GenericImport:         CreateGenericImportController(db),
	}
}
