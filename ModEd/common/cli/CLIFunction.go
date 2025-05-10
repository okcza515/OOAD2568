package main

import (
	controller "ModEd/common/controller"
	"ModEd/common/model"
	"ModEd/common/util"
	"fmt"

	"gorm.io/gorm"
)

func GenericRegister(choice int, db *gorm.DB, path string) {
	switch choice {
	case 0:
		fmt.Println("Exit")
		return
	case 1:
		RegisterModel(db, path,
			controller.NewStudentController(db),
			[]*model.Student{},
			"student")
	case 2:
		RegisterModel(db, path,
			controller.NewInstructorController(db),
			[]*model.Instructor{},
			"instructor")
	case 3:
		RegisterModel(db, path,
			controller.NewDepartmentController(db),
			[]*model.Department{},
			"department")
	case 4:
		RegisterModel(db, path,
			controller.NewFacultyController(db),
			[]*model.Faculty{},
			"faculty")
	default:
		fmt.Println("Invalid choice. Please select a number between 0 and 4.")
	}
}

func RegisterModel[T model.CommonDataInterface](db *gorm.DB, path string,
	controller interface{ Register(models []*T) error },
	models []*T, modelName string) {

	deserializer := util.CommonDeserializer(path)

	if err := deserializer.Deserialize(&models); err != nil {
		fmt.Printf("Error deserializing %s data: %v\n", modelName, err)
		return
	}

	if err := controller.Register(models); err != nil {
		fmt.Printf("Error registering %s data: %v\n", modelName, err)
		return
	}

	fmt.Printf("Complete task. (on %s)\n", modelName)
}

func GenericRetrieve(choice int, db *gorm.DB) {
	switch choice {
	case 0:
		fmt.Println("Exit")
		return
	case 1:
		RetrieveModel(db, controller.NewStudentController(db))
	case 2:
		RetrieveModel(db, controller.NewInstructorController(db))
	case 3:
		RetrieveModel(db, controller.NewDepartmentController(db))
	case 4:
		RetrieveModel(db, controller.NewFacultyController(db))
	default:
		fmt.Println("Invalid choice. Please select a number between 0 and 4.")
	}
}

func RetrieveModel[T model.CommonDataInterface](db *gorm.DB,
	controller interface {
		GetAll() (model []*T, err error)
	}) {

	data, err := controller.GetAll()
	if err != nil {
		fmt.Printf("Error retrieving data: %v\n", err)
		return
	}
	for _, student := range data {
		fmt.Printf("%+v\n", student)
	}
}

func GenericDelete(choice int, db *gorm.DB) {

	var field string
	fmt.Print("Enter model's field: ")
	fmt.Scan(&field)

	var key string
	fmt.Print("Which record is to be deleted: ")
	fmt.Scan(&key)

	switch choice {
	case 0:
		return
	case 1: // Student
		if err := model.DeleteStudentByCode(db, key); err != nil {
			fmt.Printf("Error deleting student: %v\n", err)
			return
		}
		fmt.Printf("Student with ID %s deleted successfully.\n", key)
	case 2: // Instructor
		if err := model.DeleteInstructorByCode(db, key); err != nil {
			fmt.Printf("Error deleting instructor: %v\n", err)
			return
		}
		fmt.Printf("Instructor with ID %s deleted successfully.\n", key)
	case 3: // Department
		fmt.Println("Department deletion not implemented. Use CLEAR_DB to reset all departments.")
	case 4: // Faculty
		fmt.Println("Faculty deletion not implemented. Use CLEAR_DB to reset all faculties.")
	default:
		fmt.Println("Invalid choice. Please select a number between 0 and 4.")
		return
	}
}

func Delete[T model.CommonDataInterface](db *gorm.DB, controller interface {
	Delete(field string, value interface{}) error },feild string, value interface{}) {

	if err := controller.Delete(feild, value); err != nil {
		fmt.Printf("Error deleteing %s data in %s: %v\n", value, feild, err)
		return
	}

	fmt.Printf("Complete task. (on deleting)\n")
}
