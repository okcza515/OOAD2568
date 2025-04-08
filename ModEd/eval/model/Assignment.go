// Assignment
// MEP-1006
package model

import (
	"time"

	"ModEd/common/model"

	"gorm.io/gorm"
)

type Assignment struct {
	// Instructor_Name  model.Instructor
	gorm.Model
	Title       string               
	Description string                
	StartDate   time.Time             
	DueDate     time.Time              
	Submission  []AssignmentSubmission 
}

type AssignmentSubmission struct { 
	gorm.Model
	StudentCode model.Student 
	Content     string       
	SubmittedAt string       
}
