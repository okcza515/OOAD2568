// MEP-1003 Student Recruitment
package controller

import (
	"ModEd/core"
	"ModEd/recruit/model"
	"errors"

	"gorm.io/gorm"
)

type InterviewController struct {
	Base *core.BaseController[*model.Interview]
	DB   *gorm.DB
}

func NewInterviewController(db *gorm.DB) *InterviewController {
	return &InterviewController{
		DB:   db,
		Base: core.NewBaseController[*model.Interview](db),
	}
}

func (c *InterviewController) CreateInterview(interview *model.Interview) error {
	return c.Base.Insert(interview)
}

func (c *InterviewController) DeleteInterview(id uint) error {
	condition := map[string]interface{}{"interview_id": id}
	return c.Base.DeleteByCondition(condition)
}

func (c *InterviewController) GetFilteredInterviews(condition map[string]interface{}) ([]*model.Interview, error) {
	return c.Base.List(
		condition,
		"Instructor",
		"ApplicationReport",
		"ApplicationReport.Applicant",
	)
}

func (c *InterviewController) GetInterviewByApplicationReportID(reportID uint) ([]*model.Interview, error) {
	condition := map[string]interface{}{
		"application_report_id": reportID,
	}
	return c.Base.List(
		condition,
		"Instructor",
		"ApplicationReport",
		"ApplicationReport.Applicant",
		"ApplicationReport.ApplicationRound",
	)
}

func (c *InterviewController) SaveInterviewEvaluation(data *model.Interview) error {
	var interview model.Interview
	if err := c.DB.First(&interview, data.InterviewID).Error; err != nil {
		return errors.New("Notfound Interview ID")
	}

	interview.CriteriaScores = data.CriteriaScores
	interview.TotalScore = data.TotalScore
	interview.EvaluatedAt = data.EvaluatedAt
	interview.InterviewStatus = model.ApplicationStatus(data.InterviewStatus)

	return c.DB.Save(&interview).Error
}
