// MEP-1003 Student Recruitment
package cli

import (
	common "ModEd/common/controller"
	commonModel "ModEd/common/model"
	"ModEd/recruit/controller"
	"ModEd/recruit/model"
	"errors"
	"fmt"
	"time"
)

type ConfirmedApplicantToStudentService interface {
	TransferConfirmedApplicants() error
}

type confirmedApplicantToStudentService struct {
	applicationReportCtrl *controller.ApplicationReportController
	studentCtrl           *common.StudentController
}

func NewConfirmedApplicantToStudentService(applicationReportCtrl *controller.ApplicationReportController, studentCtrl *common.StudentController) *confirmedApplicantToStudentService {
	return &confirmedApplicantToStudentService{
		applicationReportCtrl: applicationReportCtrl,
		studentCtrl:           studentCtrl,
	}
}

func (s *confirmedApplicantToStudentService) TransferConfirmedApplicants() error {
	confirmedCondition := map[string]interface{}{
		"application_statuses": model.Confirmed,
	}
	reports, err := s.applicationReportCtrl.GetFilteredApplication(confirmedCondition)
	if err != nil {
		return fmt.Errorf("failed to retrieve confirmed applications: %v", err)
	}

	activeStatus := commonModel.ACTIVE
	var students []commonModel.Student
	for _, report := range reports {
		if report.ApplicationStatuses == model.Confirmed {
			applicant := report.Applicant
			var studentcode string
			if report.Program.String() == "Regular" {
				studentcode = fmt.Sprintf("680705010%02d", applicant.ApplicantID)
			} else {
				studentcode = fmt.Sprintf("680705340%02d", applicant.ApplicantID)
			}
			student := commonModel.Student{
				StudentCode: studentcode,
				FirstName:   applicant.FirstName,
				LastName:    applicant.LastName,
				Email:       applicant.Email,
				StartDate:   time.Now(),
				BirthDate:   applicant.BirthDate,
				Program:     *report.Program,
				Department:  report.Department.Name,
				Status:      &activeStatus,
			}
			students = append(students, student)
			s.applicationReportCtrl.UpdateApplicationStatus(report.ApplicationReportID, model.Student)
		}
	}

	if len(students) == 0 {
		return errors.New("no confirmed applicants to transfer")
	}

	return s.studentCtrl.Register(students)
}
