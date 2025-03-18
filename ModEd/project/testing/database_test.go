package testing

import (
	"ModEd/project/controller"
	"ModEd/project/model"
	"os"
	"testing"
	"time"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func Init() (*gorm.DB, controller.IAssessmentController, string) {
	dbName := "test.db"
	db, _ := gorm.Open(sqlite.Open(dbName), &gorm.Config{})
	db.Exec("PRAGMA foreign_keys = ON;")

	if err := db.AutoMigrate(
		&model.Advisor{},
		&model.Assessment{},
		&model.AssessmentCriteria{},
		&model.Assignment{},
		&model.Committee{},
		&model.GroupMember{},
		&model.Presentation{},
		&model.Progress{},
		&model.Report{},
		&model.ScoreAssessmentAdvisor{},
		&model.ScoreAssessmentCommittee{},
		&model.ScoreAssignmentAdvisor{},
		&model.ScoreAssignmentCommittee{},
		&model.ScorePresentationAdvisor{},
		&model.ScorePresentationCommittee{},
		&model.ScoreReportAdvisor{},
		&model.ScoreReportCommittee{},
		&model.SeniorProject{},
	); err != nil {
		panic(err)
	}

	controller := controller.NewAssessmentController(db)

	return db, controller, dbName
}

func cleanup(dbName string) {
	os.Remove(dbName)
}

func TestListAllAssessments(t *testing.T) {
	db, ctrl, dbName := Init()
	t.Cleanup(func() { cleanup(dbName) })

	assessment := model.Assessment{}
	if err := db.Create(&assessment).Error; err != nil {
		t.Error(err)
		return
	}

	assessments, err := ctrl.ListAllAssessments()
	if err != nil || len(assessments) == 0 {
		t.Errorf("Expected assessments, got error: %v", err)
	}
}

func TestRetrieveAssessment(t *testing.T) {
	db, ctrl, dbName := Init()
	t.Cleanup(func() { cleanup(dbName) })

	assessment := model.Assessment{}
	if err := db.Create(&assessment).Error; err != nil {
		t.Fatalf("Failed to create assessment: %v", err)
	}

	res, err := ctrl.RetrieveAssessment(assessment.ID)
	if err != nil || res == nil {
		t.Errorf("Failed to retrieve assessment: %v", err)
	}
}

func TestInsertAssessment(t *testing.T) {
	_, ctrl, dbName := Init()
	t.Cleanup(func() { cleanup(dbName) })

	assessment := model.Assessment{}
	err := ctrl.InsertAssessment(&assessment)
	if err != nil {
		t.Errorf("Failed to insert assessment: %v", err)
	}
}

func TestDeleteAssessment(t *testing.T) {
	db, ctrl, dbName := Init()
	t.Cleanup(func() { cleanup(dbName) })

	assessment := model.Assessment{Model: gorm.Model{ID: 1}}
	db.Create(&assessment)

	err := ctrl.DeleteAssessment(1)
	if err != nil {
		t.Errorf("Failed to delete assessment: %v", err)
	}
}

func TestListAllAssignments(t *testing.T) {
	db, _, dbName := Init()
	ctrl := controller.NewAssignmentController(db)
	t.Cleanup(func() { cleanup(dbName) })

	assignment := model.Assignment{}
	if err := db.Create(&assignment).Error; err != nil {
		t.Error(err)
		return
	}

	assignments, err := ctrl.ListAllAssignments()
	if err != nil || len(assignments) == 0 {
		t.Errorf("Expected assignments, got error: %v", err)
	}
}

func TestRetrieveAssignment(t *testing.T) {
	db, _, dbName := Init()
	ctrl := controller.NewAssignmentController(db)
	t.Cleanup(func() { cleanup(dbName) })

	assignment := model.Assignment{}
	if err := db.Create(&assignment).Error; err != nil {
		t.Fatalf("Failed to create assignment: %v", err)
	}

	res, err := ctrl.RetrieveAssignment(assignment.ID)
	if err != nil || res == nil {
		t.Errorf("Failed to retrieve assignment: %v", err)
	}
}

func TestInsertAssignment(t *testing.T) {
	db, _, dbName := Init()
	ctrl := controller.NewAssignmentController(db)
	t.Cleanup(func() { cleanup(dbName) })

	assignment := model.Assignment{}
	err := ctrl.InsertAssignment(&assignment)
	if err != nil {
		t.Errorf("Failed to insert assignment: %v", err)
	}
}

func TestUpdateAssignment(t *testing.T) {
	db, _, dbName := Init()
	ctrl := controller.NewAssignmentController(db)
	t.Cleanup(func() { cleanup(dbName) })

	assignment := model.Assignment{}
	db.Create(&assignment)

	assignment.Name = "Updated Title"
	err := ctrl.UpdateAssignment(&assignment)
	if err != nil {
		t.Errorf("Failed to update assignment: %v", err)
	}

	var updated model.Assignment
	db.First(&updated, assignment.ID)
	if updated.Name != "Updated Title" {
		t.Errorf("Expected title 'Updated Title', got %s", updated.Name)
	}
}

func TestDeleteAssignment(t *testing.T) {
	db, _, dbName := Init()
	ctrl := controller.NewAssignmentController(db)
	t.Cleanup(func() { cleanup(dbName) })

	assignment := model.Assignment{Model: gorm.Model{ID: 1}}
	db.Create(&assignment)

	err := ctrl.DeleteAssignment(1)
	if err != nil {
		t.Errorf("Failed to delete assignment: %v", err)
	}
}

func TestListAllPresentations(t *testing.T) {
	db, _, dbName := Init()
	ctrl := controller.NewPresentationController(db)
	t.Cleanup(func() { cleanup(dbName) })

	presentation := model.Presentation{
		SeniorProjectId:  1,
		PresentationType: model.PresentationType("Final"),
		Date:             time.Now(),
	}

	if err := db.Create(&presentation).Error; err != nil {
		t.Error(err)
		return
	}

	presentations, err := ctrl.ListAllPresentations()
	if err != nil || len(presentations) == 0 {
		t.Errorf("Expected presentations, got error: %v", err)
	}
}

func TestRetrievePresentation(t *testing.T) {
	db, _, dbName := Init()
	ctrl := controller.NewPresentationController(db)
	t.Cleanup(func() { cleanup(dbName) })

	presentation := model.Presentation{
		SeniorProjectId:  1,
		PresentationType: model.PresentationType("Final"),
		Date:             time.Now(),
	}

	if err := db.Create(&presentation).Error; err != nil {
		t.Fatalf("Failed to create presentation: %v", err)
	}

	res, err := ctrl.RetrievePresentation(presentation.ID)
	if err != nil || res == nil {
		t.Errorf("Failed to retrieve presentation: %v", err)
	}
}

func TestInsertPresentation(t *testing.T) {
	db, _, dbName := Init()
	ctrl := controller.NewPresentationController(db)
	t.Cleanup(func() { cleanup(dbName) })

	presentation := model.Presentation{
		SeniorProjectId:  1,
		PresentationType: model.PresentationType("Final"),
		Date:             time.Now(),
	}

	err := ctrl.InsertPresentation(&presentation)
	if err != nil {
		t.Errorf("Failed to insert presentation: %v", err)
	}
}

func TestUpdatePresentation(t *testing.T) {
	db, _, dbName := Init()
	ctrl := controller.NewPresentationController(db)
	t.Cleanup(func() { cleanup(dbName) })

	presentation := model.Presentation{
		SeniorProjectId:  1,
		PresentationType: model.PresentationType("Initial"),
		Date:             time.Now(),
	}
	db.Create(&presentation)

	presentation.PresentationType = model.PresentationType("Final")

	err := ctrl.UpdatePresentation(&presentation)
	if err != nil {
		t.Errorf("Failed to update presentation: %v", err)
	}

	var updated model.Presentation
	db.First(&updated, presentation.ID)
	if updated.PresentationType != "Final" {
		t.Errorf("Expected presentation type 'Final', got %s", updated.PresentationType)
	}
}

func TestDeletePresentation(t *testing.T) {
	db, _, dbName := Init()
	ctrl := controller.NewPresentationController(db)
	t.Cleanup(func() { cleanup(dbName) })

	presentation := model.Presentation{
		Model:            gorm.Model{ID: 1},
		SeniorProjectId:  1,
		PresentationType: model.PresentationType("Final"),
		Date:             time.Now(),
	}
	db.Create(&presentation)

	err := ctrl.DeletePresentation(1)
	if err != nil {
		t.Errorf("Failed to delete presentation: %v", err)
	}
}

func TestListAllReports(t *testing.T) {
	db, _, dbName := Init()
	ctrl := controller.NewReportController(db)
	t.Cleanup(func() { cleanup(dbName) })

	report := model.Report{
		SeniorProjectId: 1,
		ReportType:      model.ReportType("Final"),
		DueDate:         time.Now().AddDate(0, 1, 0), // One month from now
	}

	if err := db.Create(&report).Error; err != nil {
		t.Error(err)
		return
	}

	reports, err := ctrl.ListAllReports()
	if err != nil || len(reports) == 0 {
		t.Errorf("Expected reports, got error: %v", err)
	}
}

func TestRetrieveReport(t *testing.T) {
	db, _, dbName := Init()
	ctrl := controller.NewReportController(db)
	t.Cleanup(func() { cleanup(dbName) })

	report := model.Report{
		SeniorProjectId: 1,
		ReportType:      model.ReportType("Initial"),
		DueDate:         time.Now().AddDate(0, 1, 0),
	}

	if err := db.Create(&report).Error; err != nil {
		t.Fatalf("Failed to create report: %v", err)
	}

	res, err := ctrl.RetrieveReport(report.ID)
	if err != nil || res == nil {
		t.Errorf("Failed to retrieve report: %v", err)
	}
}

func TestInsertReport(t *testing.T) {
	db, _, dbName := Init()
	ctrl := controller.NewReportController(db)
	t.Cleanup(func() { cleanup(dbName) })

	report := model.Report{
		SeniorProjectId: 1,
		ReportType:      model.ReportType("Final"),
		DueDate:         time.Now().AddDate(0, 1, 0),
	}

	err := ctrl.InsertReport(&report)
	if err != nil {
		t.Errorf("Failed to insert report: %v", err)
	}
}

func TestUpdateReport(t *testing.T) {
	db, _, dbName := Init()
	ctrl := controller.NewReportController(db)
	t.Cleanup(func() { cleanup(dbName) })

	report := model.Report{
		SeniorProjectId: 1,
		ReportType:      model.ReportType("Initial"),
		DueDate:         time.Now().AddDate(0, 1, 0),
	}
	db.Create(&report)

	// Update values
	report.ReportType = model.ReportType("Final")

	err := ctrl.UpdateReport(&report)
	if err != nil {
		t.Errorf("Failed to update report: %v", err)
	}

	var updated model.Report
	db.First(&updated, report.ID)
	if updated.ReportType != "Final" {
		t.Errorf("Expected report type 'Final', got %s", updated.ReportType)
	}
}

func TestDeleteReport(t *testing.T) {
	db, _, dbName := Init()
	ctrl := controller.NewReportController(db)
	t.Cleanup(func() { cleanup(dbName) })

	report := model.Report{
		Model:           gorm.Model{ID: 1},
		SeniorProjectId: 1,
		ReportType:      model.ReportType("Final"),
		DueDate:         time.Now().AddDate(0, 1, 0),
	}
	db.Create(&report)

	err := ctrl.DeleteReport(1)
	if err != nil {
		t.Errorf("Failed to delete report: %v", err)
	}
}
