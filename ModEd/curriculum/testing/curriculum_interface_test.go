package test

import (
	"testing"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"

	commonModel "ModEd/common/model"
	controller "ModEd/curriculum/controller/curriculum"
	model "ModEd/curriculum/model"
)

func initCurriculumDB() *gorm.DB {
	db, err := gorm.Open(sqlite.Open("../../data/curriculum.db"), &gorm.Config{})
	if err != nil {
		panic(err)
	}

	if err := db.AutoMigrate(&model.Curriculum{}); err != nil {
		panic(err)
	}
	return db
}

func TestCreateCurriculum(t *testing.T) {
	db := initCurriculumDB()
	curriculumController := controller.NewCurriculumController(db)

	newCurriculum := model.Curriculum{
		CurriculumId:   1001,
		Name:           "Computer Science",
		StartYear:      2023,
		EndYear:        2027,
		DepartmentName: "CS Dept",
		ProgramType:    commonModel.INTERNATIONAL,
	}

	curriculumId, err := curriculumController.CreateCurriculum(&newCurriculum)
	if err != nil {
		t.Fatalf("Failed to create curriculum: %v", err)
	}

	retrieved, err := curriculumController.GetCurriculum(curriculumId)
	if err != nil {
		t.Fatalf("Failed to retrieve curriculum: %v", err)
	}
	if retrieved.Name != newCurriculum.Name {
		t.Errorf("Expected name %s, got %s", newCurriculum.Name, retrieved.Name)
	}
}

func TestGetCurriculum(t *testing.T) {
	db := initCurriculumDB()
	curriculumController := controller.NewCurriculumController(db)

	curriculum := model.Curriculum{
		CurriculumId:   2001,
		Name:           "Software Engineering",
		StartYear:      2022,
		EndYear:        2026,
		DepartmentName: "SE Dept",
		ProgramType:    commonModel.INTERNATIONAL,
	}

	id, err := curriculumController.CreateCurriculum(&curriculum)
	if err != nil {
		t.Fatalf("Failed to create curriculum: %v", err)
	}

	result, err := curriculumController.GetCurriculum(id)
	if err != nil {
		t.Fatalf("Failed to get curriculum: %v", err)
	}
	if result.CurriculumId != curriculum.CurriculumId {
		t.Errorf("Expected CurriculumId %d, got %d", curriculum.CurriculumId, result.CurriculumId)
	}
}

func TestGetCurriculums(t *testing.T) {
	db := initCurriculumDB()
	curriculumController := controller.NewCurriculumController(db)

	curriculum := model.Curriculum{
		CurriculumId:   3001,
		Name:           "Data Science",
		StartYear:      2021,
		EndYear:        2025,
		DepartmentName: "DS Dept",
		ProgramType:    commonModel.INTERNATIONAL,
	}

	_, err := curriculumController.CreateCurriculum(&curriculum)
	if err != nil {
		t.Fatalf("Failed to create curriculum: %v", err)
	}

	curriculums, err := curriculumController.GetCurriculums()
	if err != nil {
		t.Fatalf("Failed to get curriculums: %v", err)
	}
	if len(curriculums) == 0 {
		t.Errorf("Expected at least 1 curriculum, got 0")
	}
}

func TestUpdateCurriculum(t *testing.T) {
	db := initCurriculumDB()
	curriculumController := controller.NewCurriculumController(db)

	curriculum := model.Curriculum{
		CurriculumId:   4001,
		Name:           "Cybersecurity",
		StartYear:      2024,
		EndYear:        2028,
		DepartmentName: "Cyber Dept",
		ProgramType:    commonModel.INTERNATIONAL,
	}

	id, err := curriculumController.CreateCurriculum(&curriculum)
	if err != nil {
		t.Fatalf("Failed to create curriculum: %v", err)
	}

	retrieved, err := curriculumController.GetCurriculum(id)
	if err != nil {
		t.Fatalf("Failed to retrieve curriculum: %v", err)
	}

	retrieved.Name = "Cybersecurity & Digital Forensics"
	updated, err := curriculumController.UpdateCurriculum(retrieved)
	if err != nil {
		t.Fatalf("Failed to update curriculum: %v", err)
	}
	if updated.Name != "Cybersecurity & Digital Forensics" {
		t.Errorf("Expected updated name, got %s", updated.Name)
	}
}

func TestDeleteCurriculum(t *testing.T) {
	db := initCurriculumDB()
	curriculumController := controller.NewCurriculumController(db)

	curriculum := model.Curriculum{
		CurriculumId:   5001,
		Name:           "AI and Robotics",
		StartYear:      2023,
		EndYear:        2027,
		DepartmentName: "AI Dept",
		ProgramType:    commonModel.INTERNATIONAL,
	}

	id, err := curriculumController.CreateCurriculum(&curriculum)
	if err != nil {
		t.Fatalf("Failed to create curriculum: %v", err)
	}

	deleted, err := curriculumController.DeleteCurriculum(id)
	if err != nil {
		t.Fatalf("Failed to delete curriculum: %v", err)
	}
	if deleted.ID != id {
		t.Errorf("Expected deleted ID %v, got %v", id, deleted.ID)
	}

	_, err = curriculumController.GetCurriculum(id)
	if err == nil {
		t.Fatalf("Expected error after deleting curriculum, got nil")
	}
}
