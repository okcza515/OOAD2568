package controller

import (
	"ModEd/core"
	"ModEd/project/model"
	"encoding/csv"
	"errors"
	"fmt"
	"os"
	"strconv"
	"time"

	"gorm.io/gorm"
)

type ReportController struct {
	*core.BaseController[*model.Report]
	db *gorm.DB
}

func NewReportController(db *gorm.DB) *ReportController {
	return &ReportController{
		db:             db,
		BaseController: core.NewBaseController[*model.Report](db),
	}
}

func (c *ReportController) UpdateReport(reportID uint, newDueDate time.Time) error {
	var report model.Report
	if err := c.db.First(&report, reportID).Error; err != nil {
		return fmt.Errorf("error retrieving report: %w", err)
	}
	report.DueDate = newDueDate
	if err := c.db.Save(&report).Error; err != nil {
		return fmt.Errorf("error updating report: %w", err)
	}

	return nil
}

func (c *ReportController) InsertReport(report model.Report) error {
	if !report.ReportType.IsValid() {
		return fmt.Errorf("invalid report type: %s", report.ReportType)
	}

	if report.DueDate.IsZero() {
		return fmt.Errorf("due date cannot be empty")
	}
	reportCopy := report
	return c.Insert(&reportCopy)
}

func (c *ReportController) GetFormattedReportList() ([]string, error) {
	reports, err := c.List(map[string]interface{}{})
	if err != nil {
		return nil, fmt.Errorf("failed to retrieve reports: %w", err)
	}

	if len(reports) == 0 {
		return nil, nil
	}

	var formattedReports []string
	for _, report := range reports {
		formattedReports = append(formattedReports, fmt.Sprintf(
			"Report ID: %d, Type: %s, DueDate: %s",
			report.ID, report.ReportType, report.DueDate.Format("2006-01-02"),
		))
	}

	return formattedReports, nil
}

func (c *ReportController) AddNewReport(seniorProjectID uint, reportType string, dueDate time.Time) error {

	report := model.Report{
		SeniorProjectId: seniorProjectID,
		ReportType:      model.ReportType(reportType),
		DueDate:         dueDate,
	}

	return c.InsertReport(report)
}

func (c *ReportController) SubmitReport(reportID uint) error {
	var report model.Report
	if err := c.db.First(&report, reportID).Error; err != nil {
		return fmt.Errorf("report not found: %w", err)
	}

	now := time.Now()
	report.SubmissionDate = &now
	if err := c.db.Save(&report).Error; err != nil {
		return fmt.Errorf("failed to submit report: %w", err)
	}
	return nil
}

func (c *ReportController) LoadReportsFromCSV(filePath string) error {
	file, err := os.Open(filePath)
	if err != nil {
		return fmt.Errorf("failed to open CSV file: %w", err)
	}
	defer file.Close()

	reader := csv.NewReader(file)
	rows, err := reader.ReadAll()
	if err != nil {
		return fmt.Errorf("failed to read CSV file: %w", err)
	}

	for _, row := range rows {
		if len(row) < 3 {
			return fmt.Errorf("invalid CSV format")
		}

		seniorProjectID, err := strconv.Atoi(row[0])
		if err != nil {
			return fmt.Errorf("invalid senior project ID: %w", err)
		}
		reportType := row[1]
		dueDate, err := time.Parse("2006-01-02", row[2])
		if err != nil {
			return fmt.Errorf("invalid due date: %w", err)
		}

		var seniorProject model.SeniorProject
		if err := c.db.First(&seniorProject, seniorProjectID).Error; err != nil {
			if errors.Is(err, gorm.ErrRecordNotFound) {
				seniorProject = model.SeniorProject{
					Model:     gorm.Model{ID: uint(seniorProjectID)},
					GroupName: fmt.Sprintf("Auto-Generated Group %d", seniorProjectID),
				}
				if err := c.db.Create(&seniorProject).Error; err != nil {
					return fmt.Errorf("failed to create senior project: %w", err)
				}
			} else {
				return fmt.Errorf("failed to retrieve senior project: %w", err)
			}
		}

		report := model.Report{
			SeniorProjectId: uint(seniorProjectID),
			ReportType:      model.ReportType(reportType),
			DueDate:         dueDate,
		}
		if err := c.db.Create(&report).Error; err != nil {
			return fmt.Errorf("failed to insert report: %w", err)
		}
	}

	return nil
}
