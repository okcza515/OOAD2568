package main

import (
	controller "ModEd/curriculum/controller/Internship"
	"bufio"
	"flag"
	"fmt"
	"os"
	"strconv"
	"strings"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"

	model "ModEd/curriculum/model/Internship"

	"time"
)

func main() {
	var (
		database string
		path     string
	)

	flag.StringVar(&database, "database", "", "Path of SQLite Database.")
	flag.StringVar(&path, "path", "", "Path to CSV or JSON for student registration.")
	flag.Parse()

	db, err := gorm.Open(sqlite.Open(database), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}

	scanner := bufio.NewScanner(os.Stdin)

	for {
		fmt.Println("\n==== Internship Application System ====")
		fmt.Println("1. Migrate Database")
		fmt.Println("2. Create Internship Application")
		fmt.Println("3. Clear Database")
		fmt.Println("4. Evaluation Student Performance")
		fmt.Println("5. Evaluation Student Report")
		fmt.Println("6. Update Approval Status")
		fmt.Println("7. Exit")
		fmt.Print("Enter your choice: ")

		scanner.Scan()
		choice := strings.TrimSpace(scanner.Text())

		switch choice {
		case "1":
			migrationController := controller.MigrationController{Db: db}
			err = migrationController.MigrateToDB()
			if err != nil {
				fmt.Printf("Error: Migration failed: %v\n", err)
			} else {
				fmt.Println("Database migration completed successfully!")
			}

			companyDataController := controller.NewCompanyDataController(db)
			err = companyDataController.ImportCompaniesFromCSV("D:/Bangmod/(Year_3)_2.2567/CPE(326) OOAD/OOAD2568/ModEd/data/Intership/Company.csv")
			if err != nil {
				fmt.Printf("Error: Failed to import companies: %v\n", err)
			} else {
				fmt.Println("Companies imported successfully!")
			}

			internStudentController := controller.InternStudentController{Connector: db}
			err = internStudentController.RegisterInternStudentsFromFile("D:/Bangmod/(Year_3)_2.2567/CPE(326) OOAD/OOAD2568/ModEd/data/StudentList.csv")
			if err != nil {
				fmt.Printf("Error: Failed to import students: %v\n", err)
			} else {
				fmt.Println("Students imported successfully!")
			}

		case "2":
			applicationController := controller.CreateInternshipApplicationController(db)

			for {
				fmt.Println("\n==== Create Internship Application ====")
				fmt.Print("Enter StudentCode (or type 'exit' to go back): ")
				scanner.Scan()
				studentCode := strings.TrimSpace(scanner.Text())
				if strings.ToLower(studentCode) == "exit" {
					break
				}

				var student model.InternStudent
				if err := db.Where("student_code = ?", studentCode).First(&student).Error; err != nil {
					fmt.Printf("Error: Student with code '%s' not found.\n", studentCode)
					continue
				}

				report_paper := model.InternshipReport{
					ReportScore: 0,
				}
				review_paper := model.SupervisorReview{
					InstructorScore: 0,
					MentorScore:     0,
				}
				db.Create(&review_paper)
				db.Create(&report_paper)

				advisorCode := student.ID

				fmt.Print("Enter CompanyName: ")
				scanner.Scan()
				companyName := strings.TrimSpace(scanner.Text())

				var company model.Company
				if err := db.Where("company_name = ?", companyName).First(&company).Error; err != nil {
					fmt.Printf("Error: Company with name '%s' not found.\n", companyName)
					continue
				}

				application := &model.InternshipApplication{
					TurninDate:            time.Now(),
					ApprovalAdvisorStatus: model.WAIT,
					ApprovalCompanyStatus: model.WAIT,
					AdvisorCode:           advisorCode,
					InternshipReportId:    report_paper.ID,
					SupervisorReviewId:    review_paper.ID,
					CompanyId:             company.ID,
					StudentCode:           studentCode,
				}

				err = applicationController.RegisterInternshipApplications([]*model.InternshipApplication{application})
				if err != nil {
					fmt.Printf("Failed to register internship application: %v\n", err)
				} else {
					fmt.Println("Internship Application created successfully!")
				}
			}

		case "3":
			fmt.Println("Clearing the database...")
			db.Migrator().DropTable(&model.InternStudent{})
			db.Migrator().DropTable(&model.Company{})
			db.Migrator().DropTable(&model.InternshipApplication{})
			db.Migrator().DropTable(&model.InternshipReport{})
			db.Migrator().DropTable(&model.SupervisorReview{})
			fmt.Println("Database cleared successfully!")
		case "4":
			Review := controller.ReviewController{DB: db}

			fmt.Print("Enter StudentCode: ")
			scanner.Scan()
			studentCode := strings.TrimSpace(scanner.Text())

			fmt.Print("Enter Supervisor Score: ")
			scanner.Scan()
			SupervisorScoreInput := strings.TrimSpace(scanner.Text())
			SupervisorScore, err := strconv.Atoi(SupervisorScoreInput)
			if err != nil {
				fmt.Printf("Error: Invalid score input. Please enter a valid number.\n")
				break
			}

			fmt.Print("Enter Mentor Score: ")
			scanner.Scan()
			MentorScoreInput := strings.TrimSpace(scanner.Text())
			MentorScore, err := strconv.Atoi(MentorScoreInput)
			if err != nil {
				fmt.Printf("Error: Invalid score input. Please enter a valid number.\n")
				break
			}
			Review.UpdateReviewScore(studentCode, SupervisorScore, MentorScore)

		case "5":

			Report := controller.ReportController{DB: db}

			fmt.Print("Enter StudentCode: ")
			scanner.Scan()
			studentCode := strings.TrimSpace(scanner.Text())

			fmt.Print("Enter Score: ")
			scanner.Scan()
			scoreInput := strings.TrimSpace(scanner.Text())
			score, err := strconv.Atoi(scoreInput)
			if err != nil {
				fmt.Printf("Error: Invalid score input. Please enter a valid number.\n")
				break
			}
			Report.UpdateReportScore(studentCode, score)

		case "6":
			approvedController := controller.NewApprovedController(db)

			fmt.Print("Enter StudentCode: ")
			scanner.Scan()
			studentCode := strings.TrimSpace(scanner.Text())

			var application model.InternshipApplication
			if err := db.Where("student_code = ?", studentCode).First(&application).Error; err != nil {
				fmt.Printf("Error: Internship application for StudentCode '%s' not found.\n", studentCode)
				continue
			}

			fmt.Print("Enter Advisor Approval Status (APPROVED/REJECT): ")
			scanner.Scan()
			advisorStatus := strings.ToUpper(strings.TrimSpace(scanner.Text()))
			if advisorStatus != string(model.APPROVED) && advisorStatus != string(model.REJECT) {
				fmt.Println("Invalid status. Please enter 'APPROVED' or 'REJECT'.")
				continue
			}
			err = approvedController.UpdateAdvisorApprovalStatus(application.ID, model.ApprovedStatus(advisorStatus))
			if err != nil {
				fmt.Printf("Failed to update advisor approval status: %v\n", err)
				continue
			} else {
				fmt.Println("Advisor approval status updated successfully!")
			}

			fmt.Print("Enter Company Approval Status (APPROVED/REJECT): ")
			scanner.Scan()
			companyStatus := strings.ToUpper(strings.TrimSpace(scanner.Text()))
			if companyStatus != string(model.APPROVED) && companyStatus != string(model.REJECT) {
				fmt.Println("Invalid status. Please enter 'APPROVED' or 'REJECT'.")
				continue
			}
			err = approvedController.UpdateCompanyApprovalStatus(application.ID, model.ApprovedStatus(companyStatus))
			if err != nil {
				fmt.Printf("Failed to update company approval status: %v\n", err)
			} else {
				fmt.Println("Company approval status updated successfully!")
			}

			if advisorStatus == string(model.APPROVED) && companyStatus == string(model.APPROVED) {
				var student model.InternStudent
				if err := db.Where("student_code = ?", studentCode).First(&student).Error; err != nil {
					fmt.Printf("Error: Student with code '%s' not found.\n", studentCode)
					continue
				}
				student.InternStatus = model.ACTIVE
				if err := db.Save(&student).Error; err != nil {
					fmt.Printf("Error: Failed to update intern status: %v\n", err)
				} else {
					fmt.Println("Intern status updated to ACTIVE.")
				}
			}

		case "7":
			fmt.Println("Exiting the system. Goodbye!")
			return

		default:
			fmt.Println("Invalid choice. Please try again.")
		}
	}
}
