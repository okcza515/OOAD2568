package test

import (
	"testing"
	"time"

	"gorm.io/gorm"

	controller "ModEd/curriculum/controller/course"
	"ModEd/curriculum/model"
	"ModEd/curriculum/utils"
)

func TestCreateCourse(t *testing.T) {
	db, err := utils.NewGormSqlite(&utils.GormConfig{
		DBPath: "../../data/curriculum.db",
		Config: &gorm.Config{},
	})
	if err != nil {
		panic(err)
	}

	courseController := controller.NewCourseController(db)

	newCourse := model.Course{
		Name:         "Test Course",
		Description:  "Test Description",
		Optional:     false,
		CourseStatus: model.ACTIVE,
		CreatedAt:    time.Now(),
		UpdatedAt:    time.Now(),
	}

	courseId, err := courseController.CreateCourse(&newCourse)
	if err != nil {
		t.Fatalf("Failed to create course: %v", err)
	}

	retrievedCourse, err := courseController.GetCourseByID(courseId)
	if err != nil {
		t.Fatalf("Failed to get course: %v", err)
	}
	if retrievedCourse.Name != newCourse.Name {
		t.Errorf("Expected course name %s, got %v", newCourse.Name, retrievedCourse.Name)
	}
}

func TestGetCourseByID(t *testing.T) {
	db, err := utils.NewGormSqlite(&utils.GormConfig{
		DBPath: "../../data/curriculum.db",
		Config: &gorm.Config{},
	})
	if err != nil {
		panic(err)
	}

	courseController := controller.NewCourseController(db)

	newCourse := model.Course{
		Name:         "Test Course",
		Description:  "Test Description",
		Optional:     false,
		CourseStatus: model.ACTIVE,
		CreatedAt:    time.Now(),
		UpdatedAt:    time.Now(),
	}

	courseId, err := courseController.CreateCourse(&newCourse)
	if err != nil {
		t.Fatalf("Failed to create course: %v", err)
	}

	retrievedCourse, err := courseController.GetCourseByID(courseId)
	if err != nil {
		t.Fatalf("Failed to get course: %v", err)
	}
	if retrievedCourse.Name != newCourse.Name {
		t.Errorf("Expected course name %s, got %v", newCourse.Name, retrievedCourse.Name)
	}
	if retrievedCourse.CourseStatus != model.ACTIVE {
		t.Errorf("Expected course status %v, got %v", model.ACTIVE, retrievedCourse.CourseStatus)
	}
}

func TestListCourses(t *testing.T) {
	db, err := utils.NewGormSqlite(&utils.GormConfig{
		DBPath: "../../data/curriculum.db",
		Config: &gorm.Config{},
	})
	if err != nil {
		panic(err)
	}

	courseController := controller.NewCourseController(db)

	newCourse := model.Course{
		Name:         "Test Course",
		Description:  "Test Description",
		Optional:     false,
		CourseStatus: model.ACTIVE,
		CreatedAt:    time.Now(),
		UpdatedAt:    time.Now(),
	}

	_, err = courseController.CreateCourse(&newCourse)
	if err != nil {
		t.Fatalf("Failed to create course: %v", err)
	}

	courses, err := courseController.ListCourses()
	if err != nil {
		t.Fatalf("Failed to list courses: %v", err)
	}
	if len(courses) < 1 {
		t.Errorf("Expected at least 1 course, got %d", len(courses))
	}
}

func TestUpdateCourse(t *testing.T) {
	db, err := utils.NewGormSqlite(&utils.GormConfig{
		DBPath: "../../data/curriculum.db",
		Config: &gorm.Config{},
	})
	if err != nil {
		panic(err)
	}

	courseController := controller.NewCourseController(db)

	newCourse := model.Course{
		Name:         "Test Course",
		Description:  "Test Description",
		Optional:     false,
		CourseStatus: model.ACTIVE,
		CreatedAt:    time.Now(),
		UpdatedAt:    time.Now(),
	}

	courseId, err := courseController.CreateCourse(&newCourse)
	if err != nil {
		t.Fatalf("Failed to create course: %v", err)
	}

	// First retrieve the course to get its ID
	retrievedCourse, err := courseController.GetCourseByID(courseId)
	if err != nil {
		t.Fatalf("Failed to get course: %v", err)
	}

	// Update the retrieved course
	retrievedCourse.Name = "Updated Course"
	retrievedCourse.Description = "Updated Description"
	retrievedCourse.Optional = true
	retrievedCourse.CourseStatus = model.INACTIVE
	retrievedCourse.UpdatedAt = time.Now()

	result, err := courseController.UpdateCourse(retrievedCourse)
	if err != nil {
		t.Fatalf("Failed to update course: %v", err)
	}

	if result.Name != "Updated Course" {
		t.Errorf("Expected updated name 'Updated Course', got %v", result.Name)
	}
	if result.Description != "Updated Description" {
		t.Errorf("Expected updated description 'Updated Description', got %v", result.Description)
	}
	if !result.Optional {
		t.Errorf("Expected Optional to be true, got %v", result.Optional)
	}
	if result.CourseStatus != model.INACTIVE {
		t.Errorf("Expected CourseStatus to be INACTIVE, got %v", result.CourseStatus)
	}
}

func TestDeleteCourse(t *testing.T) {
	db, err := utils.NewGormSqlite(&utils.GormConfig{
		DBPath: "../../data/curriculum.db",
		Config: &gorm.Config{},
	})
	if err != nil {
		panic(err)
	}

	courseController := controller.NewCourseController(db)

	newCourse := model.Course{
		Name:         "Test Course",
		Description:  "Test Description",
		Optional:     false,
		CourseStatus: model.ACTIVE,
		CreatedAt:    time.Now(),
		UpdatedAt:    time.Now(),
	}

	courseId, err := courseController.CreateCourse(&newCourse)
	if err != nil {
		t.Fatalf("Failed to create course: %v", err)
	}

	deletedCourse, err := courseController.DeleteCourse(courseId)
	if err != nil {
		t.Fatalf("Failed to delete course: %v", err)
	}
	if deletedCourse.ID != courseId {
		t.Errorf("Expected course ID %v, got %v", courseId, deletedCourse.ID)
	}

	_, err = courseController.GetCourseByID(courseId)
	if err == nil {
		t.Fatalf("Expected error when getting deleted course, got nil")
	}
}
