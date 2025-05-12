// MEP-1003 Student Recruitment
package cli

import (
	"ModEd/recruit/controller"
	"ModEd/recruit/model"
	"fmt"
)

type AdminShowApplicationReportsService interface {
	GetApplicationReport(applicantID uint) (*model.ApplicationReport, error)
	GetAllApplicationReports() ([]*model.ApplicationReport, error)
	GetApplicationReportsByStatus(status string) ([]*model.ApplicationReport, error)
	DisplayOnlyApplicationReport([]*model.ApplicationReport)
	DisplayOnlyInterviews([]*model.ApplicationReport)
	DisplayReport([]*model.ApplicationReport)
}

type adminShowApplicationReportsService struct {
	ApplicantReportCtrl *controller.ApplicationReportController
	InterviewCtrl       *controller.InterviewController
}

func NewAdminShowApplicationReportsService(applicantReportCtrl *controller.ApplicationReportController, interviewCtrl *controller.InterviewController) AdminShowApplicationReportsService {
	return &adminShowApplicationReportsService{
		ApplicantReportCtrl: applicantReportCtrl,
		InterviewCtrl:       interviewCtrl,
	}
}

func (s *adminShowApplicationReportsService) GetApplicationReport(applicantionReportID uint) (*model.ApplicationReport, error) {
	report := controller.ApplicationReport{
		Controller: s.ApplicantReportCtrl,
	}

	var condition map[string]interface{}
	condition = map[string]interface{}{
		"application_report_id": applicantionReportID,
	}

	filteredData, err := report.GetFilteredReport(condition)
	if err != nil {
		println("can't get report")
		return nil, err
	}
	return filteredData[0], nil
}

func (s *adminShowApplicationReportsService) GetAllApplicationReports() ([]*model.ApplicationReport, error) {
	report := controller.ApplicationReport{Controller: s.ApplicantReportCtrl}
	return report.GetFilteredReport(nil)
}

func (s *adminShowApplicationReportsService) GetApplicationReportsByStatus(status string) ([]*model.ApplicationReport, error) {
	report := controller.ApplicationReport{Controller: s.ApplicantReportCtrl}
	condition := map[string]interface{}{
		"application_statuses": status,
	}
	return report.GetFilteredReport(condition)
}

func (s *adminShowApplicationReportsService) DisplayOnlyApplicationReport(reports []*model.ApplicationReport) {
	reportDisplay := controller.ApplicationReport{Controller: s.ApplicantReportCtrl}
	converted := make([]model.ApplicationReport, len(reports))
	for i, r := range reports {
		if r != nil {
			converted[i] = *r
		}
	}
	reportDisplay.DisplayReport(converted)
}

func (s *adminShowApplicationReportsService) DisplayOnlyInterviews(reports []*model.ApplicationReport) {
	i := 0
	for _, r := range reports {
		if r != nil && r.ApplicationStatuses == "Interview" {
			i++
			interview := controller.InterviewReport{Controller: s.InterviewCtrl}
			condition := map[string]interface{}{
				"application_report_id": r.ApplicationReportID,
			}
			interviews, err := interview.GetFilteredReport(condition)
			if err != nil || len(interviews) == 0 {
				fmt.Println("\nNo interview report found for applicant:", r.ApplicantID)
				continue
			}
			fmt.Printf("\nInterview #%d\n", i+1)
			interview.DisplayReport(interviews)
		}
	}
}

func (s *adminShowApplicationReportsService) DisplayReport(reports []*model.ApplicationReport) {
	reportDisplay := controller.ApplicationReport{
		Controller: s.ApplicantReportCtrl,
	}
	converted := make([]model.ApplicationReport, len(reports))

	for i, r := range reports {
		if r != nil {
			converted[i] = *r
		}

		if r.ApplicationStatuses == "Interview" {
			interview := controller.InterviewReport{
				Controller: s.InterviewCtrl,
			}

			condition := map[string]interface{}{
				"application_report_id": r.ApplicationReportID,
			}

			interviews, err := interview.GetFilteredReport(condition)
			if err != nil || len(interviews) == 0 {
				println("No interview report found.")
				continue
			}

			reportDisplay.DisplayReport(converted)
			fmt.Println("\n===== Interview Report =====")
			interview.DisplayReport(interviews)
		}
	}

}
