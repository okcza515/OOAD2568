package utils

import (
	"ModEd/project/model"
	"encoding/csv"
	"fmt"
	"os"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func OpenDatabase(database string) *gorm.DB {
	db, err := gorm.Open(sqlite.Open(database), &gorm.Config{})
	if err != nil {
		fmt.Printf("Failed to connect to database: %v\n", err)
		os.Exit(1)
	}

	err = db.AutoMigrate(
		&model.SeniorProject{},
		&model.GroupMember{},
		&model.Advisor{},
		&model.Assessment{},
		&model.AssessmentCriteria{},
		&model.Assignment{},
		&model.Committee{},
		&model.Presentation{},
		&model.Progress{},
		&model.Report{},
		&model.ScoreAssessmentAdvisor{},
		&model.ScoreAssessmentCommittee{},
		&model.ScoreAssignmentAdvisor{},
		&model.ScoreAssignmentCommittee{},
		&model.ScorePresentationAdvisor{},
		&model.ScorePresentationCommittee{},
		&model.ScoreReportAdvisor{},
		&model.ScoreReportCommittee{},
	)
	if err != nil {
		fmt.Printf("AutoMigrate failed: %v\n", err)
		os.Exit(1)
	}

	return db
}

func ImportCsv(filePath string) [][]string {
	file, err := os.Open(filePath)
	if err != nil {
		fmt.Printf("Failed to open CSV %s: %v\n", filePath, err)
		return nil
	}
	defer file.Close()

	reader := csv.NewReader(file)
	records, err := reader.ReadAll()
	if err != nil {
		fmt.Printf("Failed to read CSV %s: %v\n", filePath, err)
		return nil
	}

	return records
}

func PrintTitle(title string) {
	fmt.Println("==========================================================")
	fmt.Printf("                   🎓 %s                  \n", title)
	fmt.Println("==========================================================")
	fmt.Println("Welcome to the Senior Project Management (MEP-1005) System!")
	fmt.Println("Use the menu below to navigate through the system.")
	fmt.Println()
}
