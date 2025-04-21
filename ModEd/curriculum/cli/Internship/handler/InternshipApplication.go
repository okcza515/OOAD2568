//MEP-1009 Student Internship
package Internship

import (
	controller "ModEd/curriculum/controller"
	"bufio"
	"fmt"
	"os"
	"strings"
	"time"

	model "ModEd/curriculum/model"

	"gorm.io/gorm"
)

func InternshipApplication(controller *controller.InternshipApplicationController, db *gorm.DB) {
	scanner := bufio.NewScanner(os.Stdin)

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

		err := controller.RegisterInternshipApplications([]*model.InternshipApplication{application})
		if err != nil {
			fmt.Printf("Failed to register internship application: %v\n", err)
		} else {
			fmt.Println("Internship Application created successfully!")
		}
	}
}
