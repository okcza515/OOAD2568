package cli

import (
	commonModel "ModEd/common/model"
	curriculumModel "ModEd/curriculum/model"
	evalModel "ModEd/eval/model"
	"fmt"
	"strconv"
	"time"

	"gorm.io/gorm"
)

type AssessmentCLI struct {
	db *gorm.DB
}

func NewAssessmentCLI(db *gorm.DB) *AssessmentCLI {
	return &AssessmentCLI{db: db}
}

func (cli *AssessmentCLI) Run(args []string) {
	if len(args) < 2 {
		cli.printHelp()
		return
	}

	command := args[1]
	switch command {
	case "create":
		cli.handleCreate(args[2:])
	case "list":
		cli.handleList(args[2:])
	case "get":
		cli.handleGet(args[2:])
	case "update":
		cli.handleUpdate(args[2:])
	case "delete":
		cli.handleDelete(args[2:])
	case "status":
		cli.handleStatus(args[2:])
	case "submit":
		cli.handleSubmit(args[2:])
	default:
		fmt.Printf("Unknown command: %s\n", command)
		cli.printHelp()
	}
}

func (cli *AssessmentCLI) printHelp() {
	fmt.Println("Usage: eval assessment <command> [options]")
	fmt.Println("\nCommands:")
	fmt.Println("  create    Create a new assessment")
	fmt.Println("  list      List assessments")
	fmt.Println("  get       Get assessment details")
	fmt.Println("  update    Update an assessment")
	fmt.Println("  delete    Delete an assessment")
	fmt.Println("  status    Update assessment status")
	fmt.Println("  submit    Submit an assessment")
}

func (cli *AssessmentCLI) handleCreate(args []string) {
	if len(args) < 6 {
		fmt.Println("Usage: eval assessment create --title <title> --type <type> --start <date> --due <date> --course <id> --instructor <id>")
		return
	}

	var (
		title          string
		description    string
		assessmentType string
		startDate      string
		dueDate        string
		courseId       uint
		instructorId   uint
	)

	// Parse arguments
	for i := 0; i < len(args); i += 2 {
		if i+1 >= len(args) {
			break
		}
		switch args[i] {
		case "--title":
			title = args[i+1]
		case "--description":
			description = args[i+1]
		case "--type":
			assessmentType = args[i+1]
		case "--start":
			startDate = args[i+1]
		case "--due":
			dueDate = args[i+1]
		case "--course":
			id, _ := strconv.ParseUint(args[i+1], 10, 32)
			courseId = uint(id)
		case "--instructor":
			id, _ := strconv.ParseUint(args[i+1], 10, 32)
			instructorId = uint(id)
		}
	}

	// Validate required fields
	if title == "" || assessmentType == "" || startDate == "" || dueDate == "" || courseId == 0 || instructorId == 0 {
		fmt.Println("Error: Missing required fields")
		return
	}

	// Parse dates
	start, err := time.Parse("2006-01-02", startDate)
	if err != nil {
		fmt.Printf("Error parsing start date: %v\n", err)
		return
	}

	due, err := time.Parse("2006-01-02", dueDate)
	if err != nil {
		fmt.Printf("Error parsing due date: %v\n", err)
		return
	}

	// Create assessment using builder
	builder := evalModel.NewAssessmentBuilder(evalModel.AssessmentType(assessmentType))
	assessment := builder.
		SetTitle(title).
		SetDescription(description).
		SetDates(start, due).
		SetCourse(curriculumModel.Course{CourseId: courseId}).
		SetInstructor(commonModel.Instructor{InstructorCode: strconv.FormatUint(uint64(instructorId), 10)}).
		Build()

	if err := cli.db.Create(assessment).Error; err != nil {
		fmt.Printf("Error creating assessment: %v\n", err)
		return
	}

	fmt.Printf("Assessment created successfully with ID: %d\n", assessment.AssessmentId)
}

func (cli *AssessmentCLI) handleList(args []string) {
	var (
		courseId       uint
		instructorId   uint
		assessmentType string
	)

	// Parse filters
	for i := 0; i < len(args); i += 2 {
		if i+1 >= len(args) {
			break
		}
		switch args[i] {
		case "--course":
			id, _ := strconv.ParseUint(args[i+1], 10, 32)
			courseId = uint(id)
		case "--instructor":
			id, _ := strconv.ParseUint(args[i+1], 10, 32)
			instructorId = uint(id)
		case "--type":
			assessmentType = args[i+1]
		}
	}

	var assessments []evalModel.Assessment
	query := cli.db.Model(&evalModel.Assessment{})

	if courseId != 0 {
		query = query.Where("course_id = ?", courseId)
	}
	if instructorId != 0 {
		query = query.Where("instructor_code = ?", instructorId)
	}
	if assessmentType != "" {
		query = query.Where("type = ?", assessmentType)
	}

	if err := query.Find(&assessments).Error; err != nil {
		fmt.Printf("Error listing assessments: %v\n", err)
		return
	}

	// Print assessments in a table format
	fmt.Println("ID\tType\tTitle\tStatus\tStart Date\tDue Date")
	fmt.Println("--------------------------------------------------")
	for _, a := range assessments {
		fmt.Printf("%d\t%s\t%s\t%s\t%s\t%s\n",
			a.AssessmentId,
			a.Type,
			a.Title,
			a.Status,
			a.StartDate.Format("2006-01-02"),
			a.DueDate.Format("2006-01-02"),
		)
	}
}

func (cli *AssessmentCLI) handleGet(args []string) {
	if len(args) < 1 {
		fmt.Println("Usage: eval assessment get <id>")
		return
	}

	id := args[0]
	var assessment evalModel.Assessment
	if err := cli.db.First(&assessment, id).Error; err != nil {
		fmt.Printf("Error getting assessment: %v\n", err)
		return
	}

	// Print assessment details
	fmt.Printf("Assessment Details:\n")
	fmt.Printf("ID: %d\n", assessment.AssessmentId)
	fmt.Printf("Type: %s\n", assessment.Type)
	fmt.Printf("Title: %s\n", assessment.Title)
	fmt.Printf("Description: %s\n", assessment.Description)
	fmt.Printf("Status: %s\n", assessment.Status)
	fmt.Printf("Start Date: %s\n", assessment.StartDate.Format("2006-01-02"))
	fmt.Printf("Due Date: %s\n", assessment.DueDate.Format("2006-01-02"))
}

func (cli *AssessmentCLI) handleUpdate(args []string) {
	if len(args) < 2 {
		fmt.Println("Usage: eval assessment update <id> [--title <title>] [--description <description>] [--start <date>] [--due <date>]")
		return
	}

	id := args[0]
	var assessment evalModel.Assessment
	if err := cli.db.First(&assessment, id).Error; err != nil {
		fmt.Printf("Error finding assessment: %v\n", err)
		return
	}

	// Parse update fields
	for i := 1; i < len(args); i += 2 {
		if i+1 >= len(args) {
			break
		}
		switch args[i] {
		case "--title":
			assessment.Title = args[i+1]
		case "--description":
			assessment.Description = args[i+1]
		case "--start":
			start, err := time.Parse("2006-01-02", args[i+1])
			if err != nil {
				fmt.Printf("Error parsing start date: %v\n", err)
				return
			}
			assessment.StartDate = start
		case "--due":
			due, err := time.Parse("2006-01-02", args[i+1])
			if err != nil {
				fmt.Printf("Error parsing due date: %v\n", err)
				return
			}
			assessment.DueDate = due
		}
	}

	if err := cli.db.Save(&assessment).Error; err != nil {
		fmt.Printf("Error updating assessment: %v\n", err)
		return
	}

	fmt.Println("Assessment updated successfully")
}

func (cli *AssessmentCLI) handleDelete(args []string) {
	if len(args) < 1 {
		fmt.Println("Usage: eval assessment delete <id>")
		return
	}

	id := args[0]
	var assessment evalModel.Assessment
	if err := cli.db.First(&assessment, id).Error; err != nil {
		fmt.Printf("Error finding assessment: %v\n", err)
		return
	}

	if err := cli.db.Delete(&assessment).Error; err != nil {
		fmt.Printf("Error deleting assessment: %v\n", err)
		return
	}

	fmt.Println("Assessment deleted successfully")
}

func (cli *AssessmentCLI) handleStatus(args []string) {
	if len(args) < 2 {
		fmt.Println("Usage: eval assessment status <id> --status <status>")
		return
	}

	id := args[0]
	var status string

	// Parse status
	for i := 1; i < len(args); i += 2 {
		if i+1 >= len(args) {
			break
		}
		if args[i] == "--status" {
			status = args[i+1]
		}
	}

	if status == "" {
		fmt.Println("Error: Status is required")
		return
	}

	var assessment evalModel.Assessment
	if err := cli.db.First(&assessment, id).Error; err != nil {
		fmt.Printf("Error finding assessment: %v\n", err)
		return
	}

	assessment.SetStatus(evalModel.AssessmentStatus(status))

	if err := cli.db.Save(&assessment).Error; err != nil {
		fmt.Printf("Error updating status: %v\n", err)
		return
	}

	fmt.Println("Assessment status updated successfully")
}

func (cli *AssessmentCLI) handleSubmit(args []string) {
	if len(args) < 4 {
		fmt.Println("Usage: eval assessment submit <assessment_id> --student <id> --answers <answers>")
		return
	}

	assessmentId := args[0]
	var studentId uint
	var answers string

	// Parse submission details
	for i := 1; i < len(args); i += 2 {
		if i+1 >= len(args) {
			break
		}
		switch args[i] {
		case "--student":
			id, _ := strconv.ParseUint(args[i+1], 10, 32)
			studentId = uint(id)
		case "--answers":
			answers = args[i+1]
		}
	}

	if studentId == 0 || answers == "" {
		fmt.Println("Error: Student ID and answers are required")
		return
	}

	var assessment evalModel.Assessment
	if err := cli.db.First(&assessment, assessmentId).Error; err != nil {
		fmt.Printf("Error finding assessment: %v\n", err)
		return
	}

	submission := evalModel.AssessmentSubmission{
		StudentCode: strconv.FormatUint(uint64(studentId), 10),
		Answers:     answers,
		Submitted:   true,
		SubmittedAt: time.Now(),
	}

	// Use appropriate strategy based on assessment type
	var strategy evalModel.SubmissionStrategy
	if assessment.Type == evalModel.QuizType {
		strategy = &evalModel.QuizSubmissionStrategy{}
	} else {
		strategy = &evalModel.AssignmentSubmissionStrategy{}
	}

	if err := strategy.ValidateSubmission(&submission); err != nil {
		fmt.Printf("Error validating submission: %v\n", err)
		return
	}

	if err := cli.db.Create(&submission).Error; err != nil {
		fmt.Printf("Error submitting assessment: %v\n", err)
		return
	}

	fmt.Println("Assessment submitted successfully")
}
