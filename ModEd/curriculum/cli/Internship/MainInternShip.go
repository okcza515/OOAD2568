package Internship

import (
	controller "ModEd/curriculum/controller/Internship"
	"ModEd/curriculum/utils"
	"bufio"
	"fmt"
	"os"

	"gorm.io/gorm"
)

func RunInterShipCLI(
	db *gorm.DB,
) {

	scanner := bufio.NewScanner(os.Stdin)
	CreatemigrationController := controller.MakeMigrationController(db)
	CreateInternshipApplicationController := controller.CreateInternshipApplicationController(db)
	CreateReviewController := controller.CreateReviewController(db)
	CreateReportController := controller.CreateReportController(db)
	CreateAprovedController := controller.CreateApprovedController(db)
	CreateCompanyController := controller.CreateCompanyDataController(db)
	CreateInternStudentController := controller.CreateInternStudentController(db)

	CreatemigrationController.DropAllTables()
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

		scanner.Scan()
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
