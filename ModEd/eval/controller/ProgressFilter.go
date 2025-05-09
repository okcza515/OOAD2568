package controller

import (
	evalModel "ModEd/eval/model"
)

type ProgressFilter struct {
	AssessmentId uint
	Type         evalModel.AssessmentType
	StudentCode  string
	Status       evalModel.AssessmentStatus
}
