package cli

import (
	"ModEd/eval/controller"
	"ModEd/eval/model"
	"fmt"
	"log"
	"strconv"
	"strings"
	"time"

	"gorm.io/gorm"
)

// AssessmentCLI provides command-line interface for managing assessments
type AssessmentCLI struct {
	controller controller.AssessmentController
}

// NewAssessmentCLI creates a new assessment CLI
func NewAssessmentCLI(controller controller.AssessmentController) *AssessmentCLI {
	return &AssessmentCLI{controller: controller}
}

// Run starts the assessment CLI
func (cli *AssessmentCLI) Run() {
	for {
		fmt.Println("\n===== Assessment Management =====")
		fmt.Println("1. List All Assessments")
		fmt.Println("2. Create Assessment")
		fmt.Println("3. Update Assessment")
		fmt.Println("4. Delete Assessment")
		fmt.Println("5. Submit Assessment")
		fmt.Println("6. Grade Assessment")
		fmt.Println("7. Exit")
		fmt.Print("Choose an option: ")

		var option int
		fmt.Scan(&option)

		switch option {
		case 1:
			cli.listAssessments()
		case 2:
			cli.createAssessment()
		case 3:
			cli.updateAssessment()
		case 4:
			cli.deleteAssessment()
		case 5:
			cli.submitAssessment()
		case 6:
			cli.gradeAssessment()
		case 7:
			fmt.Println("Exiting...")
			return
		default:
			fmt.Println("Invalid option")
		}
	}
}

func (cli *AssessmentCLI) listAssessments() {
	assessments, err := cli.controller.GetAll()
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	fmt.Println("\n===== Assessments =====")
	for _, a := range assessments {
		fmt.Printf("ID: %d | Type: %s | Title: %s | Status: %s\n",
			a.GetID(), a.GetType(), a.GetTitle(), a.GetStatus())
		fmt.Printf("  Start: %s | End: %s\n",
			a.GetStartTime().Format("2006-01-02 15:04"),
			a.GetEndTime().Format("2006-01-02 15:04"))
		fmt.Println("------------------------")
	}
}

func (cli *AssessmentCLI) createAssessment() {
	builder := model.NewAssessmentBuilder()

	// Get common assessment details
	var title, description string
	fmt.Print("Enter title: ")
	fmt.Scanln(&title)
	fmt.Print("Enter description: ")
	fmt.Scanln(&description)

	// Get assessment type
	var assessmentType string
	fmt.Print("Enter assessment type (quiz/assignment): ")
	fmt.Scanln(&assessmentType)

	// Set common fields
	builder.SetTitle(title).
		SetDescription(description).
		SetType(model.AssessmentType(assessmentType))

	// Get time range
	startTime := cli.getTimeInput("start")
	endTime := cli.getTimeInput("end")
	builder.SetTimeRange(startTime, endTime)

	// Get type-specific options
	switch model.AssessmentType(assessmentType) {
	case model.AssessmentTypeQuiz:
		cli.getQuizOptions(builder)
	case model.AssessmentTypeAssignment:
		cli.getAssignmentOptions(builder)
	default:
		fmt.Println("Invalid assessment type")
		return
	}

	// Build and create assessment
	assessment, err := builder.Build()
	if err != nil {
		fmt.Println("Error creating assessment:", err)
		return
	}

	if err := cli.controller.Create(assessment); err != nil {
		fmt.Println("Error saving assessment:", err)
		return
	}

	fmt.Println("Assessment created successfully")
}

func (cli *AssessmentCLI) getTimeInput(prompt string) time.Time {
	var year, month, day, hour, minute int
	fmt.Printf("Enter %s time (YYYY MM DD HH mm): ", prompt)
	fmt.Scanf("%d %d %d %d %d", &year, &month, &day, &hour, &minute)
	return time.Date(year, time.Month(month), day, hour, minute, 0, 0, time.Local)
}

func (cli *AssessmentCLI) getQuizOptions(builder *model.AssessmentBuilder) {
	var timeLimit int
	var maxAttempts int
	var showAnswers, randomize string

	fmt.Print("Enter time limit in minutes: ")
	fmt.Scanln(&timeLimit)
	fmt.Print("Enter maximum attempts: ")
	fmt.Scanln(&maxAttempts)
	fmt.Print("Show answers after submission? (y/n): ")
	fmt.Scanln(&showAnswers)
	fmt.Print("Randomize questions? (y/n): ")
	fmt.Scanln(&randomize)

	builder.SetQuizOptions(
		time.Duration(timeLimit)*time.Minute,
		maxAttempts,
		strings.ToLower(showAnswers) == "y",
		strings.ToLower(randomize) == "y",
	)
}

func (cli *AssessmentCLI) getAssignmentOptions(builder *model.AssessmentBuilder) {
	var maxFileSize int64
	var groupSize int
	var isGroup string
	var allowedTypes string

	fmt.Print("Enter maximum file size in MB: ")
	fmt.Scanln(&maxFileSize)
	fmt.Print("Enter allowed file types (comma-separated): ")
	fmt.Scanln(&allowedTypes)
	fmt.Print("Enter group size: ")
	fmt.Scanln(&groupSize)
	fmt.Print("Is this a group assignment? (y/n): ")
	fmt.Scanln(&isGroup)

	builder.SetAssignmentOptions(
		maxFileSize*1024*1024,
		strings.Split(allowedTypes, ","),
		groupSize,
		strings.ToLower(isGroup) == "y",
	)
}

func (cli *AssessmentCLI) updateAssessment() {
	var id uint
	fmt.Print("Enter assessment ID: ")
	fmt.Scanln(&id)

	assessment, err := cli.controller.GetByID(id)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	// Get updated details
	var title, description string
	fmt.Print("Enter new title (press Enter to keep current): ")
	fmt.Scanln(&title)
	fmt.Print("Enter new description (press Enter to keep current): ")
	fmt.Scanln(&description)

	if title != "" {
		assessment.(*model.BaseAssessment).Title = title
	}
	if description != "" {
		assessment.(*model.BaseAssessment).Description = description
	}

	if err := cli.controller.Update(assessment); err != nil {
		fmt.Println("Error updating assessment:", err)
		return
	}

	fmt.Println("Assessment updated successfully")
}

func (cli *AssessmentCLI) deleteAssessment() {
	var id uint
	fmt.Print("Enter assessment ID: ")
	fmt.Scanln(&id)

	if err := cli.controller.Delete(id); err != nil {
		fmt.Println("Error:", err)
		return
	}

	fmt.Println("Assessment deleted successfully")
}

func (cli *AssessmentCLI) submitAssessment() {
	var id uint
	var studentID, content string

	fmt.Print("Enter assessment ID: ")
	fmt.Scanln(&id)
	fmt.Print("Enter student ID: ")
	fmt.Scanln(&studentID)
	fmt.Print("Enter submission content: ")
	fmt.Scanln(&content)

	if err := cli.controller.Submit(id, studentID, content); err != nil {
		fmt.Println("Error:", err)
		return
	}

	fmt.Println("Submission successful")
}

func (cli *AssessmentCLI) gradeAssessment() {
	var id uint
	var studentID string
	var score float64
	var feedback string

	fmt.Print("Enter assessment ID: ")
	fmt.Scanln(&id)
	fmt.Print("Enter student ID: ")
	fmt.Scanln(&studentID)
	fmt.Print("Enter score: ")
	fmt.Scanln(&score)
	fmt.Print("Enter feedback: ")
	fmt.Scanln(&feedback)

	if err := cli.controller.Grade(id, studentID, score, feedback); err != nil {
		fmt.Println("Error:", err)
		return
	}

	fmt.Println("Grading successful")
} 