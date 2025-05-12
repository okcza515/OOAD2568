// MEP-1003 Student Recruitment
package cli

import (
	"ModEd/recruit/controller"
	"ModEd/recruit/model"
	"bufio"
	"fmt"
	"os"
	"strconv"
	"time"

	"gorm.io/gorm"
)

type InstructorEvaluateApplicantService interface {
	HasPermissionToEvaluate(instructorID, applicationReportID uint) (bool, error)
	EvaluateApplicant(applicationReportID uint, roundName, facultyName, departmentName string) error
	DetermineInterviewStatus(roundName, facultyName, departmentName string, totalScore float64) (model.ApplicationStatus, error)
}

type instructorEvaluateApplicantService struct {
	DB                    *gorm.DB
	InterviewCtrl         *controller.InterviewController
	InterviewCriteriaCtrl *controller.InterviewCriteriaCtrl
	ApplicationReportCtrl *controller.ApplicationReportController
}

func NewInstructorEvaluateApplicantService(db *gorm.DB, interviewCreiteriaCtrl *controller.InterviewCriteriaCtrl, applicationReportCtrl *controller.ApplicationReportController) InstructorEvaluateApplicantService {
	return &instructorEvaluateApplicantService{
		DB:                    db,
		InterviewCtrl:         controller.NewInterviewController(db),
		InterviewCriteriaCtrl: interviewCreiteriaCtrl,
		ApplicationReportCtrl: applicationReportCtrl,
	}
}

func (s *instructorEvaluateApplicantService) HasPermissionToEvaluate(instructorID, applicationReportID uint) (bool, error) {
	interviews, err := s.InterviewCtrl.GetInterviewByApplicationReportID(applicationReportID)
	if err != nil {
		return false, err
	}
	if len(interviews) == 0 {
		return false, fmt.Errorf("no interview found for application report ID %d", applicationReportID)
	}
	interview := interviews[0]
	return interview.InstructorID == instructorID, nil
}

func (s *instructorEvaluateApplicantService) EvaluateApplicant(applicationReportID uint, roundName, facultyName, departmentName string) error {

	strat, err := controller.GetStrategyByRoundName(roundName)
	if err != nil {
		return fmt.Errorf("failed to get evaluation strategy: %w", err)
	}

	scores := make(map[string]float64)
	for _, criterion := range strat.GetCriteria() {
		fmt.Printf("Enter score for %s: ", criterion)
		scanner := bufio.NewScanner(os.Stdin)
		scanner.Scan()
		scoreStr := scanner.Text()
		score, err := strconv.ParseFloat(scoreStr, 64)
		if err != nil || score < 0 || score > 10 {
			return fmt.Errorf("invalid score for %s", criterion)
		}
		scores[criterion] = score
	}

	totalScore, err := strat.Evaluate(scores)
	if err != nil {
		return err
	}

	interviewModels, err := s.InterviewCtrl.GetInterviewByApplicationReportID(applicationReportID)
	if err != nil {
		return fmt.Errorf("failed to get interview record: %w", err)
	}
	if len(interviewModels) == 0 {
		return fmt.Errorf("no interview found for application report ID %d", applicationReportID)
	}
	interviewModel := interviewModels[0]

	interview := &model.Interview{
		InterviewID:          interviewModel.InterviewID,
		InstructorID:         interviewModel.InstructorID,
		ApplicationReportID:  interviewModel.ApplicationReportID,
		ScheduledAppointment: interviewModel.ScheduledAppointment,
		TotalScore:           totalScore,
		EvaluatedAt:          time.Now(),
		InterviewStatus:      model.Evaluated,
	}

	err = interview.SetCriteriaScores(scores)
	if err != nil {
		return fmt.Errorf("failed to set criteria scores: %w", err)
	}

	status, err := s.DetermineInterviewStatus(roundName, facultyName, departmentName, totalScore)
	if err != nil {
		return fmt.Errorf("failed to determine interview status: %w", err)
	}

	err = s.ApplicationReportCtrl.UpdateApplicationStatus(applicationReportID, status)
	if err != nil {
		return fmt.Errorf("failed to update interview status: %w", err)
	}

	err = s.InterviewCtrl.SaveInterviewEvaluation(interview)
	if err != nil {
		return fmt.Errorf("failed to save interview evaluation: %w", err)
	}
	return nil
}

func (s *instructorEvaluateApplicantService) DetermineInterviewStatus(roundName, facultyName, departmentName string, totalScore float64) (model.ApplicationStatus, error) {
	allCriteria, err := s.InterviewCriteriaCtrl.GetFullInterviewCriteria()
	if err != nil {
		return "", fmt.Errorf("failed to get interview criteria: %w", err)
	}

	for _, c := range allCriteria {
		if c.ApplicationRound.RoundName == roundName &&
			c.Faculty.Name == facultyName &&
			c.Department.Name == departmentName {

			if totalScore >= c.PassingScore {
				return model.Accepted, nil
			}
			return model.Rejected, nil
		}
	}

	return "", fmt.Errorf("criteria not found")

}
