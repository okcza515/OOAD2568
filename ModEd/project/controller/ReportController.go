package controller

import (
	"ModEd/core"
	"ModEd/project/model"
	"encoding/csv"
	"fmt"
	"os"
	"strconv"
	"time"

	"gorm.io/gorm"
)

type ReportController struct {
	*core.BaseController
	db *gorm.DB
}

func NewReportController(db *gorm.DB) *ReportController {
	return &ReportController{
		db:             db,
		BaseController: core.NewBaseController("reports", db),
	}
}

func (c *ReportController) ListAllReports() ([]model.Report, error) {
	records, err := c.List(map[string]interface{}{})
	if err != nil {
		return nil, err
	}

	var reports []model.Report
	for _, record := range records {
		report, ok := record.(*model.Report)
		if !ok {
			return nil, fmt.Errorf("failed to cast record to model.Report")
		}
		reports = append(reports, *report)
	}

	return reports, nil
}

func (c *ReportController) RetrieveReport(id uint) (*model.Report, error) {
	var report model.Report
	if err := c.db.First(&report, id).Error; err != nil {
		return nil, err
	}
	return &report, nil
}

func (c *ReportController) InsertReport(report model.Report) error {
	if !report.ReportType.IsValid() {
		return fmt.Errorf("invalid report type: %s", report.ReportType)
	}

	if report.DueDate.IsZero() {
		return fmt.Errorf("due date cannot be empty")
	}
	return c.Insert(&report)
}

func (c *ReportController) UpdateReport(report *model.Report) error {
	return c.UpdateByID(report)
}

func (c *ReportController) DeleteReport(id uint) error {
	return c.DeleteByID(id)
}

func (c *ReportController) GetFormattedReportList() ([]string, error) {
	reports, err := c.ListAllReports()
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
		return fmt.Errorf("failed to read CSV: %w", err)
	}

	if err := c.db.Exec("DELETE FROM reports").Error; err != nil {
		return fmt.Errorf("failed to clear reports table: %w", err)
	}

	for _, row := range rows {
		if len(row) < 3 {
			continue
		}

		seniorProjectID, err := strconv.Atoi(row[0])
		if err != nil {
			return fmt.Errorf("invalid SeniorProjectId: %w", err)
		}

		dueDate, err := time.Parse("2006-01-02", row[2])
		if err != nil {
			return fmt.Errorf("invalid DueDate format: %w", err)
		}

		report := model.Report{
			SeniorProjectId: uint(seniorProjectID),
			ReportType:      model.ReportType(row[1]),
			DueDate:         dueDate,
		}

		if err := c.db.Create(&report).Error; err != nil {
			return fmt.Errorf("failed to insert report: %w", err)
		}
	}

	return nil
}
