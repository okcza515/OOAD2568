package testing

import (
	"ModEd/project/model"
	"errors"
	"os"
	"testing"
	"time"

	"gorm.io/gorm"
)

func TestInsertReport(t *testing.T) {
	db, _, _, _, _, _, reportCtrl, dbName := Init()
	t.Cleanup(func() { cleanup(dbName) })

	seniorProject := model.SeniorProject{
		GroupName: "Test Group",
	}
	if err := db.Create(&seniorProject).Error; err != nil {
		t.Fatalf("Failed to create senior project: %v", err)
	}

	report := model.Report{
		SeniorProjectId: seniorProject.ID,
		ReportType:      model.ReportTypeProposal,
		DueDate:         time.Now().AddDate(0, 1, 0),
	}

	if !report.ReportType.IsValid() {
		t.Fatalf("Invalid ReportType: %s", report.ReportType)
	}

	err := reportCtrl.InsertReport(report)
	if err != nil {
		t.Errorf("Failed to insert report: %v", err)
	}
}

func TestRetrieveReport(t *testing.T) {
	db, _, _, _, _, _, reportCtrl, dbName := Init()
	t.Cleanup(func() { cleanup(dbName) })

	seniorProject := model.SeniorProject{
		GroupName: "Test Group",
	}
	if err := db.Create(&seniorProject).Error; err != nil {
		t.Fatalf("Failed to create senior project: %v", err)
	}

	report := model.Report{
		SeniorProjectId: seniorProject.ID,
		ReportType:      model.ReportTypeProposal,
		DueDate:         time.Now().AddDate(0, 1, 0),
	}
	if err := db.Create(&report).Error; err != nil {
		t.Fatalf("Failed to create report: %v", err)
	}

	res, err := reportCtrl.RetrieveByID(report.ID)
	if err != nil || res == nil {
		t.Errorf("Failed to retrieve report: %v", err)
	}
}

func TestListAllReports(t *testing.T) {
	db, _, _, _, _, _, reportCtrl, dbName := Init()
	t.Cleanup(func() { cleanup(dbName) })

	seniorProject := model.SeniorProject{
		GroupName: "Test Group",
	}
	if err := db.Create(&seniorProject).Error; err != nil {
		t.Fatalf("Failed to create senior project: %v", err)
	}

	report := model.Report{
		SeniorProjectId: seniorProject.ID,
		ReportType:      model.ReportTypeProposal,
		DueDate:         time.Now().AddDate(0, 1, 0),
	}
	if err := db.Create(&report).Error; err != nil {
		t.Fatalf("Failed to create report: %v", err)
	}

	reports, err := reportCtrl.ListAllReports()
	if err != nil {
		t.Errorf("Expected no error, got: %v", err)
	}
	if len(reports) == 0 {
		t.Errorf("Expected reports, got none")
	}
}

func TestUpdateReport(t *testing.T) {
	db, _, _, _, _, _, reportCtrl, dbName := Init()
	t.Cleanup(func() { cleanup(dbName) })

	seniorProject := model.SeniorProject{
		GroupName: "Test Group",
	}
	if err := db.Create(&seniorProject).Error; err != nil {
		t.Fatalf("Failed to create senior project: %v", err)
	}

	report := model.Report{
		SeniorProjectId: seniorProject.ID,
		ReportType:      model.ReportTypeProposal,
		DueDate:         time.Now().AddDate(0, 1, 0),
	}
	if err := db.Create(&report).Error; err != nil {
		t.Fatalf("Failed to create report: %v", err)
	}

	report.ReportType = "Final"
	err := reportCtrl.UpdateByID(&report)
	if err != nil {
		t.Errorf("Failed to update report: %v", err)
	}

	var updatedReport model.Report
	if err := db.First(&updatedReport, report.ID).Error; err != nil {
		t.Fatalf("Failed to retrieve updated report: %v", err)
	}

	if updatedReport.ReportType != "Final" {
		t.Errorf("Expected ReportType to be 'Final', got '%s'", updatedReport.ReportType)
	}
}

func TestDeleteReport(t *testing.T) {
	db, _, _, _, _, _, reportCtrl, dbName := Init()
	t.Cleanup(func() { cleanup(dbName) })

	seniorProject := model.SeniorProject{
		GroupName: "Test Group",
	}
	if err := db.Create(&seniorProject).Error; err != nil {
		t.Fatalf("Failed to create senior project: %v", err)
	}

	report := model.Report{
		SeniorProjectId: seniorProject.ID,
		ReportType:      model.ReportTypeProposal,
		DueDate:         time.Now().AddDate(0, 1, 0),
	}
	if err := db.Create(&report).Error; err != nil {
		t.Fatalf("Failed to create report: %v", err)
	}

	err := reportCtrl.DeleteByID(report.ID)
	if err != nil {
		t.Errorf("Failed to delete report: %v", err)
	}

	var deletedReport model.Report
	err = db.First(&deletedReport, report.ID).Error
	if !errors.Is(err, gorm.ErrRecordNotFound) {
		t.Errorf("Expected record to be deleted, but it still exists")
	}
}

func TestSubmitReport(t *testing.T) {
	db, _, _, _, _, _, reportCtrl, dbName := Init()
	t.Cleanup(func() { cleanup(dbName) })

	seniorProject := model.SeniorProject{
		GroupName: "Test Group",
	}
	if err := db.Create(&seniorProject).Error; err != nil {
		t.Fatalf("Failed to create senior project: %v", err)
	}

	report := model.Report{
		SeniorProjectId: seniorProject.ID,
		ReportType:      model.ReportTypeProposal,
		DueDate:         time.Now().AddDate(0, 1, 0),
	}
	if err := db.Create(&report).Error; err != nil {
		t.Fatalf("Failed to create report: %v", err)
	}

	err := reportCtrl.SubmitReport(report.ID)
	if err != nil {
		t.Errorf("Failed to submit report: %v", err)
	}

	var submittedReport model.Report
	if err := db.First(&submittedReport, report.ID).Error; err != nil {
		t.Fatalf("Failed to retrieve submitted report: %v", err)
	}

	if submittedReport.SubmissionDate == nil {
		t.Errorf("Expected SubmissionDate to be set, but it was nil")
	}
}

func TestLoadReportsFromCSV(t *testing.T) {
	db, _, _, _, _, _, reportCtrl, dbName := Init()
	t.Cleanup(func() { cleanup(dbName) })

	filePath := "test_reports.csv"
	file, err := os.Create(filePath)
	if err != nil {
		t.Fatalf("Failed to create test CSV file: %v", err)
	}
	defer os.Remove(filePath)

	file.WriteString("1,Proposal,2025-05-01\n")
	file.WriteString("2,FinalReport,2025-06-15\n")
	file.Close()

	err = reportCtrl.LoadReportsFromCSV(filePath)
	if err != nil {
		t.Errorf("Failed to load reports from CSV: %v", err)
	}

	var reports []model.Report
	if err := db.Find(&reports).Error; err != nil {
		t.Fatalf("Failed to retrieve reports: %v", err)
	}

	if len(reports) != 2 {
		t.Errorf("Expected 2 reports, got %d", len(reports))
	}
}
