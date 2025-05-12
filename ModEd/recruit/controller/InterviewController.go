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

func GetApplicationStatus(db *gorm.DB, applicantID uint) (string, error) {
	var interview model.Interview

	err := db.Where("applicant_id = ?", applicantID).First(&interview).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return "", errors.New("ไม่พบข้อมูลสัมภาษณ์ของผู้สมัคร")
		}
		return "", err
	}

	return string(interview.InterviewStatus), nil
}

func (c *InterviewController) GetAllInterviews() ([]*model.Interview, error) {
	var interviews []*model.Interview
	err := c.DB.
		Preload("Instructor").
		Preload("ApplicationReport").
		Preload("ApplicationReport.Applicant").
		Find(&interviews).Error
	if err != nil {
		return nil, err
	}
	return interviews, nil
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
