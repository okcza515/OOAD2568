package test

import (
	"testing"
	"time"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"

	controller "ModEd/curriculum/controller/class"
	"ModEd/curriculum/model"
)

func initDB() *gorm.DB {
	db, err := gorm.Open(sqlite.Open("../../data/curriculum.db"), &gorm.Config{})
	if err != nil {
		panic(err)
	}

	if err := db.AutoMigrate(&model.Class{}); err != nil {
		panic(err)
	}
	return db
}

func TestCreateClass(t *testing.T) {
	db := initDB()
	classController := controller.NewClassController(db)

	newClass := model.Class{
		CourseId: 1,
		ClassId:  1,
		Schedule: func() time.Time {
			t, err := time.Parse(time.RFC3339, "2021-01-01T00:00:00Z")
			if err != nil {
				panic(err)
			}
			return t
		}(),
		Section: 1,
	}
	classId, err := classController.CreateClass(&newClass)
	if err != nil {
		t.Fatalf("Failed to create class: %v", err)
	}

	retrievedClass, err := classController.GetClass(classId)
	if err != nil {
		t.Fatalf("Failed to get class: %v", err)
	}
	if retrievedClass.CourseId != newClass.CourseId {
		t.Errorf("Expected course ID %d, got %v", newClass.CourseId, retrievedClass.CourseId)
	}
}

func TestGetClass(t *testing.T) {
	db := initDB()
	classController := controller.NewClassController(db)

	newClass := model.Class{
		CourseId: 1,
		ClassId:  1,
		Schedule: func() time.Time {
			t, err := time.Parse(time.RFC3339, "2021-01-01T00:00:00Z")
			if err != nil {
				panic(err)
			}
			return t
		}(),
		Section: 1,
	}
	classId, err := classController.CreateClass(&newClass)
	if err != nil {
		t.Fatalf("Failed to create class: %v", err)
	}

	retrievedClass, err := classController.GetClass(classId)
	if err != nil {
		t.Fatalf("Failed to get class: %v", err)
	}
	if retrievedClass.CourseId != newClass.CourseId {
		t.Errorf("Expected course ID %d, got %v", newClass.CourseId, retrievedClass.CourseId)
	}
}

func TestGetClasses(t *testing.T) {
	db := initDB()
	classController := controller.NewClassController(db)

	newClass := model.Class{
		CourseId: 1,
		ClassId:  1,
		Schedule: func() time.Time {
			t, err := time.Parse(time.RFC3339, "2021-01-01T00:00:00Z")
			if err != nil {
				panic(err)
			}
			return t
		}(),
		Section: 1,
	}
	_, err := classController.CreateClass(&newClass)
	if err != nil {
		t.Fatalf("Failed to create class: %v", err)
	}

	classes, err := classController.GetClasses()
	if err != nil {
		t.Fatalf("Failed to get classes: %v", err)
	}
	if len(classes) < 1 {
		t.Errorf("Expected atleast 1 class, got %d", len(classes))
	}
}

func TestUpdateClass(t *testing.T) {
	db := initDB()
	classController := controller.NewClassController(db)

	newClass := model.Class{
		CourseId: 1,
		ClassId:  1,
		Schedule: func() time.Time {
			t, err := time.Parse(time.RFC3339, "2021-01-01T00:00:00Z")
			if err != nil {
				panic(err)
			}
			return t
		}(),
		Section: 1,
	}
	classId, err := classController.CreateClass(&newClass)
	if err != nil {
		t.Fatalf("Failed to create class: %v", err)
	}

	retrievedClass, err := classController.GetClass(classId)
	if err != nil {
		t.Fatalf("Failed to get class: %v", err)
	}

	retrievedClass.Section = 2

	updatedClass, err := classController.UpdateClass(retrievedClass)
	if err != nil {
		t.Fatalf("Failed to update class: %v", err)
	}
	if updatedClass.Section != 2 {
		t.Errorf("Expected section %v, got %v", 2, updatedClass.Section)
	}
}

func TestDeleteClass(t *testing.T) {
	db := initDB()
	classController := controller.NewClassController(db)

	newClass := model.Class{
		CourseId: 1,
		ClassId:  1,
		Schedule: func() time.Time {
			t, err := time.Parse(time.RFC3339, "2021-01-01T00:00:00Z")
			if err != nil {
				panic(err)
			}
			return t
		}(),
		Section: 1,
	}
	classId, err := classController.CreateClass(&newClass)
	if err != nil {
		t.Fatalf("Failed to create class: %v", err)
	}

	retrievedClass, err := classController.GetClass(classId)
	if err != nil {
		t.Fatalf("Failed to get class: %v", err)
	}

	deletedClass, err := classController.DeleteClass(retrievedClass.ID)
	if err != nil {
		t.Fatalf("Failed to delete class: %v", err)
	}
	if deletedClass.ID != retrievedClass.ID {
		t.Errorf("Expected class ID %v, got %v", retrievedClass.ID, deletedClass.ID)
	}

	_, err = classController.GetClass(retrievedClass.ID)
	if err == nil {
		t.Fatalf("Expected error when getting deleted class, got nil")
	}
}
