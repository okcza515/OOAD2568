//Chanawat Limpanatewin 65070503445

package model

import (
	"time"
	//"ModEd/common/model"
	//"gorm.io/gorm"
)

type Evaluation struct {
	StudentID    int
	InstructorID int
	Score        float64
	Feedback     string
	EvaluatedAt  time.Time
}

////อีกแบบนึง
// type Evaluation struct {
// 	gorm.Model
// 	StudentID		model.Student
// 	InstructorID	model.Instructor
// 	Score			float64
// 	Feedback		string
// 	EvaluatedAt		time.Time
// }
