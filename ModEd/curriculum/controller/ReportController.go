// MEP-1009 Student Internship
package controller

import (
	"ModEd/curriculum/model"
	"fmt"

	"gorm.io/gorm"
)

type ReportController struct {
	*BaseScoreController[model.InternshipReport]
	Fetcher *Fetcher
}

func NewReportController(connector *gorm.DB) *ReportController {
	return &ReportController{
		BaseScoreController: &BaseScoreController[model.InternshipReport]{Connector: connector},
		Fetcher:             NewFetcher(connector),
	}
}

func (rc *ReportController) UpdateReportScore(studentID string, score int) error {
	reportID, err := rc.Fetcher.FetchIDByStudentID(studentID)
	if err != nil {
		return fmt.Errorf("failed to fetch report ID for student '%s': %w", studentID, err)
	}

	scoreFields := map[string]interface{}{
		"ReportScore": score,
	}

	return rc.UpdateScoreByID(reportID, scoreFields)
}
