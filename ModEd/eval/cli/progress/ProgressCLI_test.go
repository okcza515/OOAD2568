package cli

import (
	"fmt"

	"testing"

	"time"

	commonModel "ModEd/common/model"

	evalModel "ModEd/eval/model"
)

type Progress struct {
	StudentCode  commonModel.Student
	Title        evalModel.Assignment
	AssignmentId evalModel.Assignment
	Status       evalModel.Assignment
	LastUpdate   time.Time
	TotalSubmit  uint
}

func TestProgressModelDirect(t *testing.T) {
	progress := Progress{
		StudentCode: commonModel.Student{
			StudentCode: "65070503469",
			FirstName:   "Sawitt",
			LastName:    "์Ngamvilaisiriwong",
		},
		Title: evalModel.Assignment{
			Title: "Assignment 1",
		},
		Status: evalModel.Assignment{
			Status: "Submitted",
		},
		LastUpdate:  time.Now(),
		TotalSubmit: 3,
	}

	fmt.Println("=== Test Progress Struct ===")
	fmt.Printf("Student: %s %s\n", progress.StudentCode.FirstName, progress.StudentCode.LastName)
	fmt.Println("Title:", progress.Title.Title)
	fmt.Println("Status:", progress.Status.Status)
	fmt.Println("Submit Count:", progress.TotalSubmit)
	fmt.Println("Last Update:", progress.LastUpdate.Format("2006-01-02 15:04:05"))
}
