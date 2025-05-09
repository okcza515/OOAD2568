package test

import (
	"fmt"
	"testing"

	"gorm.io/gorm"

	commonModel "ModEd/common/model"
	controller "ModEd/curriculum/controller"
	model "ModEd/curriculum/model"
	"ModEd/curriculum/utils"
)

func TestCreateCurriculum(t *testing.T) {
	db, err := utils.NewGormSqlite(&utils.GormConfig{
		DBPath: "../../data/curriculum.db",
		Config: &gorm.Config{},
	})
	if err != nil {
		panic(err)
	}

	curriculumController := controller.NewCurriculumController(db)

	newCurriculum := model.Curriculum{
		CurriculumId: 1001,
		Name:         "Computer Science",
		StartYear:    2023,
		EndYear:      2027,
		DepartmentId: 1,
		ProgramType:  commonModel.INTERNATIONAL,
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
	db, err := utils.NewGormSqlite(&utils.GormConfig{
		DBPath: "../../data/curriculum.db",
		Config: &gorm.Config{},
	})
	if err != nil {
		panic(err)
	}

	curriculumController := controller.NewCurriculumController(db)

	curriculum := model.Curriculum{
		CurriculumId: 2001,
		Name:         "Software Engineering",
		StartYear:    2022,
		EndYear:      2026,
		DepartmentId: 1,
		ProgramType:  commonModel.INTERNATIONAL,
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
	db, err := utils.NewGormSqlite(&utils.GormConfig{
		DBPath: "../../data/curriculum.db",
		Config: &gorm.Config{},
	})
	if err != nil {
		panic(err)
	}

	curriculumController := controller.NewCurriculumController(db)

	curriculum := model.Curriculum{
		CurriculumId: 3001,
		Name:         "Data Science",
		StartYear:    2021,
		EndYear:      2025,
		DepartmentId: 1,
		ProgramType:  commonModel.INTERNATIONAL,
	}

	_, err = curriculumController.CreateCurriculum(&curriculum)
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
	db, err := utils.NewGormSqlite(&utils.GormConfig{
		DBPath: "../../data/curriculum.db",
		Config: &gorm.Config{},
	})
	if err != nil {
		panic(err)
	}

	curriculumController := controller.NewCurriculumController(db)

	curriculum := model.Curriculum{
		CurriculumId: 4001,
		Name:         "Cybersecurity",
		StartYear:    2024,
		EndYear:      2028,
		DepartmentId: 1,
		ProgramType:  commonModel.INTERNATIONAL,
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
	db, err := utils.NewGormSqlite(&utils.GormConfig{
		DBPath: "../../data/curriculum.db",
		Config: &gorm.Config{},
	})
	if err != nil {
		panic(err)
	}

	curriculumController := controller.NewCurriculumController(db)

	curriculum := model.Curriculum{
		CurriculumId: 5001,
		Name:         "AI and Robotics",
		StartYear:    2023,
		EndYear:      2027,
		DepartmentId: 1,
		ProgramType:  commonModel.INTERNATIONAL,
	}

	id, err := curriculumController.CreateCurriculum(&curriculum)
	if err != nil {
		t.Fatalf("Failed to create curriculum: %v", err)
	}

	deleted, err := curriculumController.DeleteCurriculum(id)
	if err != nil {
		t.Fatalf("Failed to delete curriculum: %v", err)
	}
	if deleted.CurriculumId != id {
		t.Errorf("Expected deleted ID %v, got %v", id, deleted.CurriculumId)
	}

	_, err = curriculumController.GetCurriculum(id)
	if err == nil {
		t.Fatalf("Expected error after deleting curriculum, got nil")
	}
}

func TestSeedCurriculum(t *testing.T) {
	db, err := utils.NewGormSqlite(&utils.GormConfig{
		DBPath: "../../data/curriculum.db",
		Config: &gorm.Config{},
	})
	//TODO: Fix this since migration is now at core module
	// migrationController := controller.NewMigrationController(db)
	// if err := migrationController.MigrateToDB(); err != nil {
	// 	t.Fatalf("Failed to migrate to db: %v", err)
	// }

	curriculumController := controller.NewCurriculumController(db)
	curriculums, err := curriculumController.CreateSeedCurriculum("../../data/curriculum/curriculum.json")
	if err != nil {
		t.Fatalf("Failed to seed curriculum: %v", err)
	}
	if len(curriculums) == 0 {
		t.Errorf("Expected at least 1 curriculum, got 0")
	}

	for _, curriculum := range curriculums {
		fmt.Printf("Curriculum: %v\n", curriculum.CurriculumId)
		curriculumId, err := curriculumController.GetCurriculum(curriculum.CurriculumId)
		if err != nil {
			t.Fatalf("Failed to create curriculum: %v", err)
		}
		if curriculumId.CurriculumId != curriculum.CurriculumId {
			t.Errorf("Expected curriculum ID %d, got %d", curriculum.CurriculumId, curriculumId.CurriculumId)
		}
	}

}

//TODO: Fix this since migration is now at core module

// func TestMigration(t *testing.T) {
// 	db, err := utils.NewGormSqlite(&utils.GormConfig{
// 		DBPath: "../../data/curriculum.db",
// 		Config: &gorm.Config{},
// 	})
// 	if err != nil {
// 		panic(err)
// 	}
//
// 	// migrationController := controller.NewMigrationController(db)
// 	// if err := migrationController.MigrateToDB(); err != nil {
// 	// 	t.Fatalf("Failed to migrate to db: %v", err)
// 	// }
// }
