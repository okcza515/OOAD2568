package model

import (
	"time"
	"github.com/google/uuid"
)

type Instructor struct {
	Name 			string
	InstructorID 	string
	FirstName 		string
	LastName  		string  
	Email     		string  
	StartDate 		time.Time
	Faculty   		Faculty
	Department 		Department
	Courses   		[]uuid.UUID		//UUID to avoid circular dependency
}