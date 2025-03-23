package testing

import (
	"ModEd/project/controller"
	"ModEd/project/model"
	"os"
	"testing"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func Init() (*gorm.DB, controller.IAssessmentController, controller.IAssignmentController, *controller.PresentationController, controller.IReportController, string) {
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

	assessmentController := controller.NewAssessmentController(db)
	assignmentController := controller.NewAssignmentController(db)
	presentationController := controller.NewPresentationController(db)
	reportController := controller.NewReportController(db)

	return db, assessmentController, assignmentController, presentationController, reportController, dbName
}

func cleanup(dbName string) {
	os.Remove(dbName)
}

func TestListAllAssessments(t *testing.T) {
	db, assessmentCtrl, _, _, _, dbName := Init()
	t.Cleanup(func() { cleanup(dbName) })

	assessment := model.Assessment{}
	if err := db.Create(&assessment).Error; err != nil {
		t.Error(err)
		return
	}

	assessments, err := assessmentCtrl.ListAllAssessments()
	if err != nil || len(assessments) == 0 {
		t.Errorf("Expected assessments, got error: %v", err)
	}
}

func TestRetrieveAssessment(t *testing.T) {
	db, assessmentCtrl, _, _, _, dbName := Init()
	t.Cleanup(func() { cleanup(dbName) })

	assessment := model.Assessment{}
	if err := db.Create(&assessment).Error; err != nil {
		t.Fatalf("Failed to create assessment: %v", err)
	}

	res, err := assessmentCtrl.RetrieveAssessment(assessment.ID)
	if err != nil || res == nil {
		t.Errorf("Failed to retrieve assessment: %v", err)
	}
}

func TestInsertAssessment(t *testing.T) {
	_, assessmentCtrl, _, _, _, dbName := Init()
	t.Cleanup(func() { cleanup(dbName) })

	assessment := model.Assessment{}
	err := assessmentCtrl.InsertAssessment(&assessment)
	if err != nil {
		t.Errorf("Failed to insert assessment: %v", err)
	}
}

func TestDeleteAssessment(t *testing.T) {
	db, assessmentCtrl, _, _, _, dbName := Init()
	t.Cleanup(func() { cleanup(dbName) })

	assessment := model.Assessment{Model: gorm.Model{ID: 1}}
	db.Create(&assessment)

	err := assessmentCtrl.DeleteAssessment(1)
	if err != nil {
		t.Errorf("Failed to delete assessment: %v", err)
	}
}

func TestListAllAssignments(t *testing.T) {
	db, _, assignmentCtrl, _, _, dbName := Init()
	t.Cleanup(func() { cleanup(dbName) })

	assignment := model.Assignment{}
	if err := db.Create(&assignment).Error; err != nil {
		t.Error(err)
		return
	}

	assignments, err := assignmentCtrl.ListAllAssignments()
	if err != nil || len(assignments) == 0 {
		t.Errorf("Expected assignments, got error: %v", err)
	}
}

func TestRetrieveAssignment(t *testing.T) {
	db, _, assignmentCtrl, _, _, dbName := Init()
	t.Cleanup(func() { cleanup(dbName) })

	assignment := model.Assignment{}
	if err := db.Create(&assignment).Error; err != nil {
		t.Fatalf("Failed to create assignment: %v", err)
	}

	res, err := assignmentCtrl.RetrieveAssignment(assignment.ID)
	if err != nil || res == nil {
		t.Errorf("Failed to retrieve assignment: %v", err)
	}
}

func TestInsertAssignment(t *testing.T) {
	_, _, assignmentCtrl, _, _, dbName := Init()
	t.Cleanup(func() { cleanup(dbName) })

	assignment := model.Assignment{}
	err := assignmentCtrl.InsertAssignment(&assignment)
	if err != nil {
		t.Errorf("Failed to insert assignment: %v", err)
	}
}

func TestDeleteAssignment(t *testing.T) {
	db, _, assignmentCtrl, _, _, dbName := Init()
	t.Cleanup(func() { cleanup(dbName) })

	assignment := model.Assignment{Model: gorm.Model{ID: 1}}
	db.Create(&assignment)

	err := assignmentCtrl.DeleteAssignment(1)
	if err != nil {
		t.Errorf("Failed to delete assignment: %v", err)
	}
}

func TestListAllPresentations(t *testing.T) {
	db, _, _, presentationCtrl, _, dbName := Init()
	t.Cleanup(func() { cleanup(dbName) })

	presentation := model.Presentation{}
	if err := db.Create(&presentation).Error; err != nil {
		t.Error(err)
		return
	}

	presentations, err := presentationCtrl.ListAllPresentations()
	if err != nil || len(presentations) == 0 {
		t.Errorf("Expected presentations, got error: %v", err)
	}
}

func TestRetrievePresentation(t *testing.T) {
	db, _, _, presentationCtrl, _, dbName := Init()
	t.Cleanup(func() { cleanup(dbName) })

	presentation := model.Presentation{}
	if err := db.Create(&presentation).Error; err != nil {
		t.Fatalf("Failed to create presentation: %v", err)
	}

	res, err := presentationCtrl.RetrievePresentation(presentation.ID)
	if err != nil || res == nil {
		t.Errorf("Failed to retrieve presentation: %v", err)
	}
}

func TestInsertPresentation(t *testing.T) {
	_, _, _, presentationCtrl, _, dbName := Init()
	t.Cleanup(func() { cleanup(dbName) })

	presentation := model.Presentation{}
	err := presentationCtrl.InsertPresentation(presentation)
	if err != nil {
		t.Errorf("Failed to insert presentation: %v", err)
	}
}

func TestDeletePresentation(t *testing.T) {
	db, _, _, presentationCtrl, _, dbName := Init()
	t.Cleanup(func() { cleanup(dbName) })

	presentation := model.Presentation{Model: gorm.Model{ID: 1}}
	db.Create(&presentation)

	err := presentationCtrl.DeletePresentation(1)
	if err != nil {
		t.Errorf("Failed to delete presentation: %v", err)
	}
}

func TestListAllReports(t *testing.T) {
	db, _, _, _, reportCtrl, dbName := Init()
	t.Cleanup(func() { cleanup(dbName) })

	report := model.Report{}
	if err := db.Create(&report).Error; err != nil {
		t.Error(err)
		return
	}

	reports, err := reportCtrl.ListAllReports()
	if err != nil || len(reports) == 0 {
		t.Errorf("Expected reports, got error: %v", err)
	}
}

func TestRetrieveReport(t *testing.T) {
	db, _, _, _, reportCtrl, dbName := Init()
	t.Cleanup(func() { cleanup(dbName) })

	report := model.Report{}
	if err := db.Create(&report).Error; err != nil {
		t.Fatalf("Failed to create report: %v", err)
	}

	res, err := reportCtrl.RetrieveReport(report.ID)
	if err != nil || res == nil {
		t.Errorf("Failed to retrieve report: %v", err)
	}
}

func TestInsertReport(t *testing.T) {
	_, _, _, _, reportCtrl, dbName := Init()
	t.Cleanup(func() { cleanup(dbName) })

	report := model.Report{}
	err := reportCtrl.InsertReport(&report)
	if err != nil {
		t.Errorf("Failed to insert report: %v", err)
	}
}

func TestDeleteReport(t *testing.T) {
	db, _, _, _, reportCtrl, dbName := Init()
	t.Cleanup(func() { cleanup(dbName) })

	report := model.Report{Model: gorm.Model{ID: 1}}
	db.Create(&report)

	err := reportCtrl.DeleteReport(1)
	if err != nil {
		t.Errorf("Failed to delete report: %v", err)
	}
}
