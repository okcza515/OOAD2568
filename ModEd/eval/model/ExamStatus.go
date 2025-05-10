// MEP-1007
package model

type ExamStatus string

const (
    Draft   ExamStatus = "Draft"
    Publish ExamStatus = "Publish"
	Hidden  ExamStatus = "Hidden"
)