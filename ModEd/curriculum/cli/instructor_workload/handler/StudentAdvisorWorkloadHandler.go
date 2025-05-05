package handler

import (
	"ModEd/curriculum/controller"
	"ModEd/curriculum/model"
	"ModEd/curriculum/utils"
	"fmt"
	"time"

	commonModel "ModEd/common/model"
)

func RunStudentAdvisorWorkloadHandler(controller controller.StudentWorkloadInterface) {
	for {
		DisplayStudentAdvisorWorkloadModuleMenu()
		choice := utils.GetUserChoice()
		fmt.Println("choice: ", choice)

		// mockStudentAdvisor := &model.StudentAdvisor{
		// 	InstructorId: 1,
		// 	Instructor: commonModel.Instructor{
		// 		InstructorCode: "I001",
		// 		FirstName:      "John",
		// 		LastName:       "Doe",
		// 		Email:          "",
		// 		StartDate:      nil,
		// 		Department:     nil,
		// 	},
		// 	Students: []commonModel.Student{
		// 		{
		// 			StudentCode: "S001",
		// 			FirstName:   "Jane",
		// 			LastName:    "Smith",
		// 			Email:       "",
		// 			StartDate:   time.Time{},
		// 			BirthDate:   time.Time{},
		// 			Program:     commonModel.REGULAR,
		// 			Status:      nil,
		// 		},
		// 	},
		// }

		mockStudentRequest := &model.StudentRequest{
			Student: commonModel.Student{
				StudentCode: "S001",
				FirstName:   "Jane",
				LastName:    "Smith",
				Email:       "",
				StartDate:   time.Time{},
				BirthDate:   time.Time{},
				Program:     commonModel.REGULAR,
				Status:      nil,
			},
			InstructorId: 1,
			Instructor: commonModel.Instructor{
				InstructorCode: "I001",
				FirstName:      "John",
				LastName:       "Doe",
				Email:          "",
				StartDate:      nil,
				Department:     nil,
			},
			RequestType: "Absence",
			CreatedAt:   "2023-10-01",
			UpdatedAt:   "2023-10-01",
			Review:      "",
			Comment:     "",
		}

		switch choice {
		case "3":
			err := controller.DeleteByStudentId(1)
			if err != nil {
				fmt.Println("Error deleting student advisor:", err)
			} else {
				fmt.Println("Student advisor deleted successfully")
			}
		case "4":
			studentAdvisors, err := controller.GetStudentUnderSupervisionByInstructorId(1)
			if err != nil {
				fmt.Println("Error getting student advisors:", err)
			} else {
				fmt.Println("Student advisors under supervision:")
				for _, advisor := range studentAdvisors {
					fmt.Printf("ID: %d, Instructor: %s %s\n", advisor.ID, advisor.Instructor.FirstName, advisor.Instructor.LastName)
					for _, student := range advisor.Students {
						fmt.Printf("  Student: %s %s\n", student.FirstName, student.LastName)
					}
				}
			}
		case "5":
			err := controller.CreateStudentRequest(*mockStudentRequest)
			if err != nil {
				fmt.Println("Error creating student request:", err)
			} else {
				fmt.Println("Student request created successfully")
			}
		case "6":
			studentRequests, err := controller.GetStudentRequestsByInstructorId(1)
			if err != nil {
				fmt.Println("Error getting student requests:", err)
			} else {
				fmt.Println("Student requests:")
				for _, request := range studentRequests {
					fmt.Printf("ID: %d, Student: %s %s, Request Type: %s, Review: %s\n",
						request.ID, request.Student.FirstName, request.Student.LastName, request.RequestType, request.Review)
					if request.Review != "" {
						fmt.Printf("  Comment: %s\n", request.Comment)
					}
				}
			}
		case "7":
			var id uint
			var review, comment string
			fmt.Print("Enter Student Request ID to review: ")
			fmt.Scanln(&id)
			fmt.Print("Enter review (accept/reject): ")
			fmt.Scanln(&review)
			fmt.Print("Enter comment: ")
			fmt.Scanln(&comment)
			err := controller.ReviewStudentRequest(id, review, comment)
			if err != nil {
				fmt.Println("Error reviewing student request:", err)
			} else {
				fmt.Println("Student request reviewed successfully")
			}
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
	fmt.Println("2. Update Student Advisor")
	fmt.Println("3. Delete Student Advisor")
	fmt.Println("4. List all Student Advisors By Advisor ID")

	fmt.Println("5. Add Student Request")
	fmt.Println("6. List all Student Requests")
	fmt.Println("7. Review Student Request")

	fmt.Println("Type 'exit' to quit")
}
