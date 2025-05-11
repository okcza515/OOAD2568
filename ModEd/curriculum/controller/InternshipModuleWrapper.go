// MEP-1009 Student Internship
package controller

import (
	"gorm.io/gorm"
)

type InternshipModuleWrapper struct {
	CurriculumController                 *CurriculumController
	InternshipApplication                *InternshipApplicationController
	Approved                             *ApprovedController
	Company                              *CompanyController
	InternStudent                        *InternStudentController
	GenericImport                        *GenericImportController
	InformationController                *InternshipInformationController
	InternshipMentorController           *InternshipMentorController
	InternshipResultEvaluationController *InternshipResultEvaluationController
	InternshipCriteriaController         *InternshipCriteriaController
}

func NewInternshipModuleWrapper(
	db *gorm.DB,
	curriculumController CurriculumControllerInterface,
) *InternshipModuleWrapper {
	return &InternshipModuleWrapper{
		InternshipApplication:                NewInternshipApplicationController(db),
		Approved:                             NewApprovedController(db),
		Company:                              NewCompanyController(db),
		InternStudent:                        NewInternStudentController(db),
		GenericImport:                        CreateGenericImportController(db),
		InformationController:                NewInternshipInformationController(db),
		InternshipMentorController:           NewInternshipMentorController(db),
		InternshipResultEvaluationController: NewInternshipResultEvaluationController(db),
		InternshipCriteriaController:         NewInternshipCriteriaController(db),
	}
}
