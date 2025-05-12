package menu

import (
	controller "ModEd/common/controller"
	"ModEd/common/model"
	"ModEd/common/util"
	"fmt"

	"gorm.io/gorm"
)

type RegisterFunc func()

func GenericRegister(choice int, db *gorm.DB, path string) {
	actions := map[int]RegisterFunc{
		1: func() {
			RegisterModel(db, path,
				controller.NewStudentController(db),
				[]*model.Student{},
				"student")
		},
		2: func() {
			RegisterModel(db, path,
				controller.NewInstructorController(db),
				[]*model.Instructor{},
				"instructor")
		},
		3: func() {
			RegisterModel(db, path,
				controller.NewDepartmentController(db),
				[]*model.Department{},
				"department")
		},
		4: func() {
			RegisterModel(db, path,
				controller.NewFacultyController(db),
				[]*model.Faculty{},
				"faculty")
		},
	}

	if choice == 0 {
		fmt.Println("Exit")
		return
	}

	action, ok := actions[choice]
	if !ok {
		fmt.Println("Invalid choice. Please select a number between 0 and 4.")
		return
	}

	action()
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

type RetrieveFunc func()

func GenericRetrieve(choice int, db *gorm.DB) {
	actions := map[int]RetrieveFunc{
		1: func() {
			RetrieveModel(db, controller.NewStudentController(db))
		},
		2: func() {
			RetrieveModel(db, controller.NewInstructorController(db))
		},
		3: func() {
			RetrieveModel(db, controller.NewDepartmentController(db))
		},
		4: func() {
			RetrieveModel(db, controller.NewFacultyController(db))
		},
	}

	if choice == 0 {
		fmt.Println("Exit")
		return
	}

	action, ok := actions[choice]
	if !ok {
		fmt.Println("Invalid choice. Please select a number between 0 and 4.")
		return
	}

	action()
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

type DeleteFunc func(key string)

func GenericDelete(choice int, db *gorm.DB) {
	var key string
	fmt.Print("Enter key to delete: ")
	fmt.Scan(&key)

	actions := map[int]DeleteFunc{
		1: func(key string) {
			if err := model.DeleteStudentByCode(db, key); err != nil {
				fmt.Printf("Error deleting student: %v\n", err)
				return
			}
			fmt.Printf("Student with ID %s deleted successfully.\n", key)
		},
		2: func(key string) {
			if err := model.DeleteInstructorByCode(db, key); err != nil {
				fmt.Printf("Error deleting instructor: %v\n", err)
				return
			}
			fmt.Printf("Instructor with ID %s deleted successfully.\n", key)
		},
		3: func(key string) {
			fmt.Println("Department deletion not implemented. Use CLEAR_DB to reset all departments.")
		},
		4: func(key string) {
			fmt.Println("Faculty deletion not implemented. Use CLEAR_DB to reset all faculties.")
		},
	}

	if choice == 0 {
		fmt.Println("Exit")
		return
	}

	action, ok := actions[choice]
	if !ok {
		fmt.Println("Invalid choice. Please select a number between 0 and 4.")
		return
	}

	action(key)
}

func Delete[T model.CommonDataInterface](db *gorm.DB, controller interface {
	Delete(field string, value interface{}) error },feild string, value interface{}) {

	if err := controller.Delete(feild, value); err != nil {
		fmt.Printf("Error deleteing %s data in %s: %v\n", value, feild, err)
		return
	}

	fmt.Printf("Complete task. (on deleting)\n")
}
