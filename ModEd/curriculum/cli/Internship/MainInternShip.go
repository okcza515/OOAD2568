//MEP-1009 Student Internship
package Internship

import (
	controller "ModEd/curriculum/controller"
	"ModEd/curriculum/utils"
	"fmt"

	"gorm.io/gorm"
)

func RunInterShipCLI(
	db *gorm.DB,
) {

	CreateInternshipApplicationController := controller.CreateInternshipApplicationController(db)
	CreateReviewController := controller.CreateReviewController(db)
	CreateReportController := controller.CreateReportController(db)
	CreateAprovedController := controller.CreateApprovedController(db)
	CreateCompanyController := controller.CreateCompanyDataController(db)
	CreateInternStudentController := controller.CreateInternStudentController(db)

	CreateCompanyController.ImportCompaniesFromCSV("")
	CreateInternStudentController.RegisterInternStudentsFromFile("")

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
			InternshipApplication(CreateInternshipApplicationController, db)
		case "2":
			InternshipReview(CreateReviewController)
		case "3":
			InternshipReport(CreateReportController)
		case "4":
			InternShipApproved(CreateAprovedController)
		case "5":
			fmt.Println("Exiting the system. Goodbye!")
			return
		default:
			fmt.Println("Invalid choice. Please try again.")
		}
	}
}
