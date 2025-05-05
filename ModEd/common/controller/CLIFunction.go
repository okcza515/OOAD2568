package controller

import (
	"ModEd/common/model"
	"ModEd/common/util"
	"ModEd/utils/deserializer"
	"fmt"

	"gorm.io/gorm"
)

func GenericRegister(choice int, db *gorm.DB, path string) {
	deserializer := util.CommonDeserializer(path)

	switch choice {
	case 0:
		fmt.Println("Exit")
		return
	case 1:
		registerModel(db, &deserializer,
			CreateStudentController(db),
			[]*model.Student{},
			"student")
	case 2:
		registerModel(db, &deserializer,
			CreateInstructorController(db),
			[]*model.Instructor{},
			"instructor")
	case 3:
		registerModel(db, &deserializer,
			CreateDepartmentController(db),
			[]*model.Department{},
			"department")
	case 4:
		registerModel(db, &deserializer,
			CreateFacultyController(db),
			[]*model.Faculty{},
			"faculty")
	default:
		fmt.Println("Invalid choice. Please select a number between 0 and 4.")
	}
}

func registerModel[T model.CommonDataInterface](db *gorm.DB, deserializer *deserializer.FileDeserializer,
	controller interface{ Register(models []*T) error },
	models []*T, modelName string) {

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
	controller := CreateStudentController(db)
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

func Delete[T model.CommonDataInterface](db *gorm.DB, controller interface{ Delete(field string, value interface{}) error },
	feild string, value interface{}) {

	if err := controller.Delete(feild, value); err != nil {
		fmt.Printf("Error deleteing %s data in %s: %v\n", value, feild, err)
		return
	}

	fmt.Printf("Complete task. (on deleting)\n")
}