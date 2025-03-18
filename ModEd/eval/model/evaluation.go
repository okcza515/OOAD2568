//Chanawat Limpanatewin 65070503445
//MEP-1006

package model

import (
	"ModEd/common/model"

	"time"

	"gorm.io/gorm"
)

type Evaluation struct {
	gorm.Model
	SID          model.Student    `gorm:"foreignKey:SID"`          //รหัสนักศึกษา
	InstructorId model.Instructor `gorm:"foreignKey:InstructorId"` //รหัสอาจารย์ที่ตรวจ
	Score        float64          `gorm:"not null"`                //คะแนนที่นักศึกษาได้
	Feedback     string           //คำแนะนำจากอาจารย์
	EvaluatedAt  time.Time        //ประเมินเมื่อไร
}
