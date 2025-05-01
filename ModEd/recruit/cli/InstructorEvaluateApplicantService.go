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
	EvaluateApplicant(applicationReportID uint, roundName string) error
}

type instructorEvaluateApplicantService struct {
	DB            *gorm.DB
	InterviewCtrl *controller.InterviewController
}

func NewInstructorEvaluateApplicantService(db *gorm.DB) InstructorEvaluateApplicantService {
	return &instructorEvaluateApplicantService{
		DB:            db,
		InterviewCtrl: controller.NewInterviewController(db),
	}
}

func (s *instructorEvaluateApplicantService) HasPermissionToEvaluate(instructorID, applicationReportID uint) (bool, error) {
	interview, err := s.InterviewCtrl.GetInterviewByApplicationReportID(applicationReportID)
	if err != nil {
		return false, err
	}
	return interview.InstructorID == instructorID, nil
}

func (s *instructorEvaluateApplicantService) EvaluateApplicant(applicationReportID uint, roundName string) error {

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

	interviewModel, err := s.InterviewCtrl.GetInterviewByApplicationReportID(applicationReportID)
	if err != nil {
		return fmt.Errorf("failed to get interview record: %w", err)
	}

	interview := &model.Interview{
		InterviewID:          interviewModel.InterviewID,
		InstructorID:         interviewModel.InstructorID,
		ApplicationReportID:  interviewModel.ApplicationReportID,
		ScheduledAppointment: interviewModel.ScheduledAppointment,
		TotalScore:           totalScore,
		EvaluatedAt:          time.Now(),
		InterviewStatus:      model.Pending,
	}

	err = interview.SetCriteriaScores(scores)
	if err != nil {
		return fmt.Errorf("failed to set criteria scores: %w", err)
	}

	return s.InterviewCtrl.SaveInterviewEvaluation(interview)
}
