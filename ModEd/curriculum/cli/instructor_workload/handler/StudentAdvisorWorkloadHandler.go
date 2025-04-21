package handler

import (
	"ModEd/curriculum/controller"
	"ModEd/curriculum/model"
	"ModEd/curriculum/utils"
	"fmt"
	"time"

	commonModel "ModEd/common/model"
)

func RunStudentAdvisorWorkloadHandler(controller controller.StudentWorkloadService) {
	for {
		DisplayStudentAdvisorWorkloadModuleMenu()
		choice := utils.GetUserChoice()
		fmt.Println("choice: ", choice)

		switch choice {
		case "1":
			mockStudentAdvisor := &model.StudentAdvisor{
				InstructorId: 1,
				Instructor: commonModel.Instructor{
					InstructorCode: "I001",
					FirstName:      "John",
					LastName:       "Doe",
					Email:          "",
					StartDate:      nil,
					Department:     nil,
				},
				Students: []commonModel.Student{
					{
						StudentCode: "S001",
						FirstName:   "Jane",
						LastName:    "Smith",
						Email:       "",
						StartDate:   time.Time{},
						BirthDate:   time.Time{},
						Program:     commonModel.REGULAR,
						Status:      nil,
					},
				},
			}
			err := controller.CreateStudentAdvisor(*mockStudentAdvisor)
			if err != nil {
				fmt.Println("Error creating student advisor:", err)
			} else {
				fmt.Println("Student advisor created successfully")
			}
		case "2":
			fmt.Println("Edit Student Advisor Not implemented yet...")
		case "3":
			fmt.Println("Delete Student Advisor Not implemented yet...")
		case "4":
			fmt.Println("List all Student Advisors By Advisor ID Not implemented yet...")
		case "5":
			fmt.Println("Get Student Advisor By ID Not implemented yet...")
		case "6":
			fmt.Println("List all Student Requests Not implemented yet...")
		case "7":
			fmt.Println("Review Student Request Not implemented yet...")
		case "8":
			fmt.Println("Create Meeting Not implemented yet...")
		case "9":
			fmt.Println("Edit Meeting Not implemented yet...")
		case "10":
			fmt.Println("Delete Meeting Not implemented yet...")
		case "11":
			fmt.Println("List all Meetings by attendee ID Not implemented yet...")
		case "12":
			fmt.Println("Get Meeting By ID Not implemented yet...")
		case "exit":
			fmt.Println("Exiting...")
			return
		default:
			fmt.Println("Invalid option")
		}
	}
}

func DisplayStudentAdvisorWorkloadModuleMenu() {
	fmt.Println("\nStudentAdvisor Workload Module Menu:")
	fmt.Println("1. Add Student Advisor")
	fmt.Println("2. Edit Student Advisor")
	fmt.Println("3. Delete Student Advisor")
	fmt.Println("4. List all Student Advisors By Advisor ID")

	fmt.Println("6. List all Student Requests")
	fmt.Println("7. Review Student Request")

	fmt.Println("8. Create Meeting")
	fmt.Println("9. Edit Meeting")
	fmt.Println("10. Delete Meeting")
	fmt.Println("11. List all Meetings by attendee ID")
	fmt.Println("12. Get Meeting By ID")

	fmt.Println("Type 'exit' to quit")
}
