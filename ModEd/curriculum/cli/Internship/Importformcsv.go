package Internship

import (
	controller "ModEd/curriculum/controller/Internship"
	"fmt"

	"gorm.io/gorm"
)

func ImportCompanyData(Company *controller.CompanyDataController, db *gorm.DB, path string) {
	err := Company.ImportCompaniesFromCSV(path)
	if err != nil {
		fmt.Printf("Error: Failed to import companies: %v\n", err)
	} else {
		fmt.Println("Companies imported successfully!")
	}
}
