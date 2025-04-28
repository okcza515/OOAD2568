package test

import (
	"fmt"
	"testing"
	"time"

	"gorm.io/gorm"

	"ModEd/curriculum/controller"
	"ModEd/curriculum/model"
	"ModEd/curriculum/utils"
)

func TestCreateClass(t *testing.T) {
	db, err := utils.NewGormSqlite(&utils.GormConfig{
		DBPath: "../../data/curriculum.db",
		Config: &gorm.Config{},
	})
	if err != nil {
		panic(err)
	}

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
	db, err := utils.NewGormSqlite(&utils.GormConfig{
		DBPath: "../../data/curriculum.db",
		Config: &gorm.Config{},
	})
	if err != nil {
		panic(err)
	}

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
	db, err := utils.NewGormSqlite(&utils.GormConfig{
		DBPath: "../../data/curriculum.db",
		Config: &gorm.Config{},
	})
	if err != nil {
		panic(err)
	}

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
	_, err = classController.CreateClass(&newClass)
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
	db, err := utils.NewGormSqlite(&utils.GormConfig{
		DBPath: "../../data/curriculum.db",
		Config: &gorm.Config{},
	})
	if err != nil {
		panic(err)
	}

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
	db, err := utils.NewGormSqlite(&utils.GormConfig{
		DBPath: "../../data/curriculum.db",
		Config: &gorm.Config{},
	})
	if err != nil {
		panic(err)
	}

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

	deletedClass, err := classController.DeleteClass(retrievedClass.ClassId)
	if err != nil {
		t.Fatalf("Failed to delete class: %v", err)
	}
	if deletedClass.ClassId != retrievedClass.ClassId {
		t.Errorf("Expected class ID %v, got %v", retrievedClass.ClassId, deletedClass.ClassId)
	}

	_, err = classController.GetClass(retrievedClass.ClassId)
	if err == nil {
		t.Fatalf("Expected error when getting deleted class, got nil")
	}
}

func TestSeedClass(t *testing.T) {
	db, err := utils.NewGormSqlite(&utils.GormConfig{
		DBPath: "../../data/curriculum.db",
		Config: &gorm.Config{},
	})
	//TODO: Fix this since migration is now at core module
	// migrationController := controller.NewMigrationController(db)
	// if err := migrationController.MigrateToDB(); err != nil {
	// 	t.Fatalf("Failed to migrate to db: %v", err)
	// }

	classController := controller.NewClassController(db)
	classes, err := classController.CreateSeedClass("../../data/curriculum/class.json")
	if err != nil {
		t.Fatalf("Failed to seed class: %v", err)
	}
	if len(classes) == 0 {
		t.Errorf("Expected at least 1 class, got 0")
	}

	for _, class := range classes {
		fmt.Printf("Class: %v\n", class.ClassId)
		rclass, err := classController.GetClass(class.ClassId)
		if err != nil {
			t.Fatalf("Failed to create class: %v", err)
		}
		if rclass.ClassId != class.ClassId {
			t.Errorf("Expected class ID %d, got %d", class.ClassId, rclass.ClassId)
		}
	}

}
