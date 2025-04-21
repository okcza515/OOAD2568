//MEP-1009 Student Internship
package Internship

import (
	handler "ModEd/curriculum/cli/Internship/handler"
	controller "ModEd/curriculum/controller"
	"ModEd/curriculum/utils"
	"fmt"

	"gorm.io/gorm"
)

func RunInterShipCLI(
	db *gorm.DB,
) {

	const (
		defaultStudentPath = "../../data/StudentList.csv"
	)

	const (
		defaultCompaniesPath = "../../data/Internship/Company.csv"
	)

	CreateInternshipApplicationController := controller.CreateInternshipApplicationController(db)
	CreateReviewController := controller.CreateReviewController(db)
	CreateReportController := controller.CreateReportController(db)
	CreateAprovedController := controller.CreateApprovedController(db)
	CreateInternStudentController := controller.CreateInternStudentController(db)
	CreateGenericImportController := controller.CreateGenericImportController(db)

	CreateGenericImportController.ImportCompaniesFromCSV(defaultCompaniesPath)
	CreateInternStudentController.RegisterInternStudentsFromFile(defaultStudentPath)

	for {
		fmt.Println("\n==== Internship Application System ====")
		fmt.Println("1. Create Internship Application")
		fmt.Println("2. Evaluation Student Performance")
		fmt.Println("3. Evaluation Student Report")
		fmt.Println("4. Update Approval Status")
		fmt.Println("5. Exit")
		fmt.Print("Enter your choice: ")

		choice := utils.GetUserChoice()
		switch choice {
		case "1":
			handler.InternshipApplication(CreateInternshipApplicationController, db)
		case "2":
			handler.InternshipReview(CreateReviewController)
		case "3":
			handler.InternshipReport(CreateReportController)
		case "4":
			handler.InternShipApproved(CreateAprovedController)
		case "5":
			fmt.Println("Exiting the system. Goodbye!")
			return
		default:
			fmt.Println("Invalid choice. Please try again.")
		}
	}
}
