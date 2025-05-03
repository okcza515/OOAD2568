package main

import (
	"fmt"
	"time"

	"ModEd/eval/controller"
	"ModEd/eval/model"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func main() {
	// Initialize database
	db, err := gorm.Open(sqlite.Open("moded.db"), &gorm.Config{})
	if err != nil {
		fmt.Printf("Failed to connect to database: %v\n", err)
		return
	}

	// Auto-migrate the schema
	err = db.AutoMigrate(
		&model.BaseAssessment{},
		&model.Submission{},
		&model.Question{},
		&model.Answer{},
	)
	if err != nil {
		fmt.Printf("Failed to migrate database: %v\n", err)
		return
	}

	// Create assessment controller
	assessmentController := controller.NewAssessmentController(db)

	// Create a quiz using the builder pattern
	quizBuilder := model.NewAssessmentBuilder()
	quiz, err := quizBuilder.
		SetType(model.AssessmentTypeQuiz).
		SetTitle("Midterm Quiz").
		SetDescription("Test your knowledge on the course material").
		SetTimeRange(time.Now(), time.Now().Add(1*time.Hour)).
		SetQuizOptions(30*time.Minute, 2, true, true).
		Build()
	if err != nil {
		fmt.Printf("Failed to create quiz: %v\n", err)
		return
	}

	// Save the quiz
	if err := assessmentController.Create(quiz); err != nil {
		fmt.Printf("Failed to save quiz: %v\n", err)
		return
	}
	fmt.Println("Quiz created successfully!")

	// Create an assignment using the builder pattern
	assignmentBuilder := model.NewAssessmentBuilder()
	assignment, err := assignmentBuilder.
		SetType(model.AssessmentTypeAssignment).
		SetTitle("Final Project").
		SetDescription("Create a comprehensive project demonstrating your understanding").
		SetTimeRange(time.Now(), time.Now().Add(7*24*time.Hour)).
		SetAssignmentOptions(10*1024*1024, []string{"pdf", "docx", "zip"}, 3, true).
		Build()
	if err != nil {
		fmt.Printf("Failed to create assignment: %v\n", err)
		return
	}

	// Save the assignment
	if err := assessmentController.Create(assignment); err != nil {
		fmt.Printf("Failed to save assignment: %v\n", err)
		return
	}
	fmt.Println("Assignment created successfully!")

	// List all assessments
	assessments, err := assessmentController.GetAll()
	if err != nil {
		fmt.Printf("Failed to get assessments: %v\n", err)
		return
	}

	fmt.Println("\n===== All Assessments =====")
	for _, a := range assessments {
		fmt.Printf("ID: %d | Type: %s | Title: %s | Status: %s\n",
			a.GetID(), a.GetType(), a.GetTitle(), a.GetStatus())
		fmt.Printf("  Start: %s | End: %s\n",
			a.GetStartTime().Format("2006-01-02 15:04"),
			a.GetEndTime().Format("2006-01-02 15:04"))
		fmt.Println("------------------------")
	}

	// Submit to the quiz
	quizID := quiz.GetID()
	if err := assessmentController.Submit(quizID, "STUDENT001", "Quiz answers here"); err != nil {
		fmt.Printf("Failed to submit quiz: %v\n", err)
		return
	}
	fmt.Println("Quiz submitted successfully!")

	// Submit to the assignment
	assignmentID := assignment.GetID()
	if err := assessmentController.Submit(assignmentID, "STUDENT001", "Project submission here"); err != nil {
		fmt.Printf("Failed to submit assignment: %v\n", err)
		return
	}
	fmt.Println("Assignment submitted successfully!")

	// Grade the quiz
	if err := assessmentController.Grade(quizID, "STUDENT001", 85.5, "Good work!"); err != nil {
		fmt.Printf("Failed to grade quiz: %v\n", err)
		return
	}
	fmt.Println("Quiz graded successfully!")

	// Grade the assignment
	if err := assessmentController.Grade(assignmentID, "STUDENT001", 92.0, "Excellent project!"); err != nil {
		fmt.Printf("Failed to grade assignment: %v\n", err)
		return
	}
	fmt.Println("Assignment graded successfully!")

	// Get and display a specific assessment
	quiz, err = assessmentController.GetByID(quizID)
	if err != nil {
		fmt.Printf("Failed to get quiz: %v\n", err)
		return
	}

	fmt.Println("\n===== Quiz Details =====")
	fmt.Printf("ID: %d\n", quiz.GetID())
	fmt.Printf("Title: %s\n", quiz.GetTitle())
	fmt.Printf("Description: %s\n", quiz.GetDescription())
	fmt.Printf("Type: %s\n", quiz.GetType())
	fmt.Printf("Status: %s\n", quiz.GetStatus())
	fmt.Printf("Start Time: %s\n", quiz.GetStartTime().Format("2006-01-02 15:04"))
	fmt.Printf("End Time: %s\n", quiz.GetEndTime().Format("2006-01-02 15:04"))

	// Display quiz-specific properties
	if q, ok := quiz.(*model.Quiz); ok {
		fmt.Printf("Time Limit: %v\n", q.GetTimeLimit())
		fmt.Printf("Max Attempts: %d\n", q.GetMaxAttempts())
		fmt.Printf("Show Answers: %v\n", q.GetShowAnswers())
		fmt.Printf("Randomize Questions: %v\n", q.GetRandomize())
	}

	// Get and display assignment details
	assignment, err = assessmentController.GetByID(assignmentID)
	if err != nil {
		fmt.Printf("Failed to get assignment: %v\n", err)
		return
	}

	fmt.Println("\n===== Assignment Details =====")
	fmt.Printf("ID: %d\n", assignment.GetID())
	fmt.Printf("Title: %s\n", assignment.GetTitle())
	fmt.Printf("Description: %s\n", assignment.GetDescription())
	fmt.Printf("Type: %s\n", assignment.GetType())
	fmt.Printf("Status: %s\n", assignment.GetStatus())
	fmt.Printf("Start Time: %s\n", assignment.GetStartTime().Format("2006-01-02 15:04"))
	fmt.Printf("End Time: %s\n", assignment.GetEndTime().Format("2006-01-02 15:04"))

	// Display assignment-specific properties
	if a, ok := assignment.(*model.Assignment); ok {
		fmt.Printf("Max File Size: %d bytes\n", a.GetMaxFileSize())
		fmt.Printf("Allowed Types: %v\n", a.GetAllowedTypes())
		fmt.Printf("Group Size: %d\n", a.GetGroupSize())
		fmt.Printf("Is Group Assignment: %v\n", a.GetIsGroup())
	}
} 