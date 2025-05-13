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

func GetFacultyController(db *gorm.DB) *FacultyController {
	facultyLock.Do(func() {
		facultyInstance = newFacultyController(db)
	})
	return facultyInstance
}

func GetStudentController(db *gorm.DB) *StudentController {
	studentLock.Do(func() {
		studentInstance = newStudentController(db)
	})
	return studentInstance
}

func GetInstructorController(db *gorm.DB) *InstructorController {
	instructorLock.Do(func() {
		instructorInstance = newInstructorController(db)
	})
	return instructorInstance
}

func GetDepartmentController(db *gorm.DB) *DepartmentController {
	departmentLock.Do(func() {
		departmentInstance = newDepartmentController(db)
	})
	return departmentInstance
}
