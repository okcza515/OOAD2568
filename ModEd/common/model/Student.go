package model

import (
	"time"

	"gorm.io/gorm"
)

type Student struct {
	gorm.Model
	StudentCode string        `gorm:"not null;unique" csv:"student_code"`
	FirstName 	string        `csv:"first_name"`
	LastName  	string        `csv:"last_name"`
	Email     	string        `csv:"email"`
	StartDate 	time.Time     `csv:"start_date"`
	BirthDate 	time.Time     `csv:"birth_date"`
	Program   	ProgramType   `csv:"program"`
	Status    	StudentStatus `csv:"status"`
}
