// MEP-1003 Student Recruitment
package cli

import (
	"ModEd/recruit/controller"
	"ModEd/recruit/model"
	"fmt"
	"time"
)

type AdminScheduleInterviewService interface {
	ScheduleInterview(instructorID uint, applicationReportID uint, scheduledTime string) error
}

type adminScheduleInterviewService struct {
	interviewCtrl           *controller.InterviewController
	applicationReportCtrl   *controller.ApplicationReportController
}

func NewAdminScheduleInterviewService(interviewCtrl *controller.InterviewController, applicationReportCtrl *controller.ApplicationReportController) AdminScheduleInterviewService {
	return &adminScheduleInterviewService{
		interviewCtrl:         interviewCtrl,
		applicationReportCtrl: applicationReportCtrl,
	}
}

func (s *adminScheduleInterviewService) ScheduleInterview(instructorID uint, applicationReportID uint, scheduledTime string) error {
	applicationReport, err := s.applicationReportCtrl.GetApplicationReportByID(applicationReportID)
	if err != nil {
		return fmt.Errorf("failed to retrieve ApplicationReport: %w", err)
	}

	if applicationReport.ApplicationStatuses != "Pending" {
		return fmt.Errorf("you cannot assign interview details at this stage. Current status: %s", applicationReport.ApplicationStatuses)
	}

	scheduledTimeParsed, err := time.Parse("2006-01-02 15:04:05", scheduledTime)
	if err != nil {
		return fmt.Errorf("invalid date format. Use YYYY-MM-DD HH:MM:SS")
	}

	interview := &model.Interview{
		InstructorID:         instructorID,
		ApplicationReportID:  applicationReportID,
		ScheduledAppointment: scheduledTimeParsed,
		CriteriaScores:       "",
		TotalScore:           0,
		EvaluatedAt:          time.Time{},
		InterviewStatus:      model.Pending,
	}
	
	err = s.interviewCtrl.CreateInterview(interview)
	if err != nil {
		return fmt.Errorf("failed to create interview: %w", err)
	}

	err = s.applicationReportCtrl.UpdateApplicationStatus(applicationReport.ApplicationReportID, model.InterviewStage)
	if err != nil {
		return fmt.Errorf("failed to update status: %w", err)
	}

	return nil
}
