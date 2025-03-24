package main

import (
	controller "ModEd/curriculum/controller/Internship"
	"bufio"
	"fmt"
	"strings"

	"errors"
	"flag"
	"os"

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

	if _, err := os.Stat(path); errors.Is(err, os.ErrNotExist) {
		panic("*** Error: " + path + " does not exist.\n")
	}

	migrationController := controller.MigrationController{Db: db}
	err = migrationController.MigrateToDB()
	if err != nil {
		panic("err: migration failed")
	}

	companyDataController := controller.NewCompanyDataController(db)
	err = companyDataController.ImportCompaniesFromCSV("")
	if err != nil {
		panic("err: failed to import companies")
	}

	internStudentController := controller.InternStudentController{Connector: db}
	err = internStudentController.RegisterInternStudentsFromFile("")
	if err != nil {
		panic("err: failed to import students")
	}

	applicationController := controller.CreateInternshipApplicationController(db)

	scanner := bufio.NewScanner(os.Stdin)

	for {
		fmt.Println("\n==== Create Internship Application ====")
		fmt.Print("Enter StudentCode (or type 'exit' to quit): ")
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
			CompanyId:             company.ID,
			StudentCode:           studentCode,
		}

		err = applicationController.RegisterInternshipApplications([]*model.InternshipApplication{application})
		if err != nil {
			fmt.Printf("Failed to register internship application: %v\n", err)
		} else {
			fmt.Println("InternshipApplication created successfully!")
		}
	}
}
