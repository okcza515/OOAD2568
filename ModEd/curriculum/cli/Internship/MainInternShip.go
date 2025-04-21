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
		fmt.Println("1. Migrate Database")
		fmt.Println("2. Create Internship Application")
		fmt.Println("3. Evaluation Student Performance")
		fmt.Println("4. Evaluation Student Report")
		fmt.Println("5. Update Approval Status")
		fmt.Println("6. Exit")
		fmt.Print("Enter your choice: ")

		scanner.Scan()
		choice := utils.GetUserChoice()
		switch choice {
		case "1":
			InternshipMigration(CreatemigrationController)
		case "2":
			InternshipApplication(CreateInternshipApplicationController, db)
		case "3":
			InternshipReview(CreateReviewController)
		case "4":
			InternshipReport(CreateReportController)
		case "5":
			InternShipApproved(CreateAprovedController)
		case "6":
			fmt.Println("Exiting the system. Goodbye!")
			return
		default:
			fmt.Println("Invalid choice. Please try again.")
		}
	}
}
