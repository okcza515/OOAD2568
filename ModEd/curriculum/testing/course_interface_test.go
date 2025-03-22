package test

import (
	"testing"
	"time"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"

	controller "ModEd/curriculum/controller/course"
	"ModEd/curriculum/model"
)

func setupDB() *gorm.DB {
	db, err := gorm.Open(sqlite.Open("../../data/curriculum.db"), &gorm.Config{})
	if err != nil {
		panic(err)
	}

	if err := db.AutoMigrate(&model.Course{}); err != nil {
		panic(err)
	}
	return db
}

func TestCreateCourse(t *testing.T) {
	db := setupDB()
	courseController := controller.NewCourseController(db)

	newCourse := &model.Course{
		Name:         "Test Course",
		Description:  "Test Description",
		Optional:     false,
		CourseStatus: model.ACTIVE,
		CreatedAt:    time.Now(),
		UpdatedAt:    time.Now(),
	}

	err := courseController.CreateCourse(newCourse)
	if err != nil {
		t.Fatalf("Failed to create course: %v", err)
	}

	retrievedCourse, err := courseController.GetCourseByID(newCourse.ID)
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

func TestGetCourseByID(t *testing.T) {
	db := setupDB()
	courseController := controller.NewCourseController(db)

	newCourse := &model.Course{
		Name:         "Test Course",
		Description:  "Test Description",
		Optional:     false,
		CourseStatus: model.ACTIVE,
		CreatedAt:    time.Now(),
		UpdatedAt:    time.Now(),
	}

	err := courseController.CreateCourse(newCourse)
	if err != nil {
		t.Fatalf("Failed to create course: %v", err)
	}

	retrievedCourse, err := courseController.GetCourseByID(newCourse.ID)
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
	db := setupDB()
	courseController := controller.NewCourseController(db)

	newCourse := &model.Course{
		Name:         "Test Course",
		Description:  "Test Description",
		Optional:     false,
		CourseStatus: model.ACTIVE,
		CreatedAt:    time.Now(),
		UpdatedAt:    time.Now(),
	}

	err := courseController.CreateCourse(newCourse)
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
	db := setupDB()
	courseController := controller.NewCourseController(db)

	newCourse := &model.Course{
		Name:         "Test Course",
		Description:  "Test Description",
		Optional:     false,
		CourseStatus: model.ACTIVE,
		CreatedAt:    time.Now(),
		UpdatedAt:    time.Now(),
	}

	err := courseController.CreateCourse(newCourse)
	if err != nil {
		t.Fatalf("Failed to create course: %v", err)
	}

	updatedCourse := &model.Course{
		Name:         "Updated Course",
		Description:  "Updated Description",
		Optional:     true,
		CourseStatus: model.INACTIVE,
	}

	err = courseController.UpdateCourse(newCourse.ID, updatedCourse)
	if err != nil {
		t.Fatalf("Failed to update course: %v", err)
	}

	retrievedCourse, err := courseController.GetCourseByID(newCourse.ID)
	if err != nil {
		t.Fatalf("Failed to get updated course: %v", err)
	}
	if retrievedCourse.Name != "Updated Course" {
		t.Errorf("Expected updated name 'Updated Course', got %v", retrievedCourse.Name)
	}
	if retrievedCourse.Description != "Updated Description" {
		t.Errorf("Expected updated description 'Updated Description', got %v", retrievedCourse.Description)
	}
	if !retrievedCourse.Optional {
		t.Errorf("Expected Optional to be true, got %v", retrievedCourse.Optional)
	}
	if retrievedCourse.CourseStatus != model.INACTIVE {
		t.Errorf("Expected CourseStatus to be INACTIVE, got %v", retrievedCourse.CourseStatus)
	}
}

func TestUpdateCourseStatus(t *testing.T) {
	db := setupDB()
	courseController := controller.NewCourseController(db)

	newCourse := &model.Course{
		Name:         "Status Test Course",
		Description:  "Testing Status Transitions",
		Optional:     false,
		CourseStatus: model.INACTIVE, // Starting with Draft status
		CreatedAt:    time.Now(),
		UpdatedAt:    time.Now(),
	}

	err := courseController.CreateCourse(newCourse)
	if err != nil {
		t.Fatalf("Failed to create course: %v", err)
	}

	// Transition from INACTIVE to ACTIVE
	updatedCourse := &model.Course{
		Name:         newCourse.Name,
		Description:  newCourse.Description,
		Optional:     newCourse.Optional,
		CourseStatus: model.ACTIVE,
	}

	err = courseController.UpdateCourse(newCourse.ID, updatedCourse)
	if err != nil {
		t.Fatalf("Failed to update course status: %v", err)
	}

	retrievedCourse, err := courseController.GetCourseByID(newCourse.ID)
	if err != nil {
		t.Fatalf("Failed to get updated course: %v", err)
	}
	if retrievedCourse.CourseStatus != model.ACTIVE {
		t.Errorf("Expected CourseStatus to be ACTIVE, got %v", retrievedCourse.CourseStatus)
	}

	// Transition from ACTIVE to INACTIVE
	updatedCourse.CourseStatus = model.INACTIVE
	err = courseController.UpdateCourse(newCourse.ID, updatedCourse)
	if err != nil {
		t.Fatalf("Failed to update course status: %v", err)
	}

	retrievedCourse, err = courseController.GetCourseByID(newCourse.ID)
	if err != nil {
		t.Fatalf("Failed to get updated course: %v", err)
	}
	if retrievedCourse.CourseStatus != model.INACTIVE {
		t.Errorf("Expected CourseStatus to be Archived, got %v", retrievedCourse.CourseStatus)
	}
}

func TestDeleteCourse(t *testing.T) {
	db := setupDB()
	courseController := controller.NewCourseController(db)

	newCourse := &model.Course{
		Name:         "Test Course",
		Description:  "Test Description",
		Optional:     false,
		CourseStatus: model.ACTIVE,
		CreatedAt:    time.Now(),
		UpdatedAt:    time.Now(),
	}

	err := courseController.CreateCourse(newCourse)
	if err != nil {
		t.Fatalf("Failed to create course: %v", err)
	}

	err = courseController.DeleteCourse(newCourse.ID)
	if err != nil {
		t.Fatalf("Failed to delete course: %v", err)
	}

	_, err = courseController.GetCourseByID(newCourse.ID)
	if err == nil {
		t.Fatalf("Expected error when getting deleted course, got nil")
	}
}
