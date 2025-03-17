// 65070503409 Chayaphon Chaisangkha
// MEP-1006
package model

import (
	"time"

	"gorm.io/gorm"

	"ModEd/common/model"
)

// Model to store quiz detail
type Quiz struct {
	gorm.Model
	ID          uint                   //Quiz ID
	Title       string                 //Quiz Title
	Description string                 //Quiz Description
	Released    bool                   //to check if quiz is released
	QuizStart   time.Time              //Time that quiz start
	QuizEnd     time.Time              //Time that quiz finish
	Status      string                 //Quiz status
	Submission  []AssignmentSubmission //Submission detail
	TotalSubmit uint                   //Toltal number of submited quiz
}

// Model to store quiz submission detail
type QuizSubmission struct {
	gorm.Model
	SID        model.Student //Student ID
	FirstName  model.Student //Student FirstName
	LastName   model.Student //Student LastName
	Email      model.Student //Student Email
	Answers    string        //Student answer
	Submitted  bool          //to check if student is submited
	SubmitTime time.Time     //Student submit time
}
