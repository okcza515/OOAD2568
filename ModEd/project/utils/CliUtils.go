package utils

import (
	"ModEd/project/model"
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

func PrintTitle(title string) {
	fmt.Println("==========================================================")
	fmt.Printf("                   🎓 %s                  \n", title)
	fmt.Println("==========================================================")
	fmt.Println("Welcome to the Senior Project Management (MEP-1005) System!")
	fmt.Println("Use the menu below to navigate through the system.")
	fmt.Println()
}
