package controller

import (
	"sync"

	"gorm.io/gorm"
)

var (
	facultyInstance    *FacultyController
	studentInstance    *StudentController
	instructorInstance *InstructorController
	departmentInstance *DepartmentController
	facultyLock        sync.Once
	studentLock        sync.Once
	instructorLock     sync.Once
	departmentLock     sync.Once
)

// GetFacultyController returns the singleton instance of FacultyController
func GetFacultyController(db *gorm.DB) *FacultyController {
	facultyLock.Do(func() {
		facultyInstance = newFacultyController(db)
	})
	return facultyInstance
}

// GetStudentController returns the singleton instance of StudentController
func GetStudentController(db *gorm.DB) *StudentController {
	studentLock.Do(func() {
		studentInstance = newStudentController(db)
	})
	return studentInstance
}

// GetInstructorController returns the singleton instance of InstructorController
func GetInstructorController(db *gorm.DB) *InstructorController {
	instructorLock.Do(func() {
		instructorInstance = newInstructorController(db)
	})
	return instructorInstance
}

// GetDepartmentController returns the singleton instance of DepartmentController
func GetDepartmentController(db *gorm.DB) *DepartmentController {
	departmentLock.Do(func() {
		departmentInstance = newDepartmentController(db)
	})
	return departmentInstance
}
