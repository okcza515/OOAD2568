package model

type ExaminationCriteria struct {
	TotalScore int        `json:"total_score"` 
	Sections   []Section  `json:"sections"`
}

type Section struct {
	SectionNo     			   uint   `json:"section_no"`      
	Description   			   string `json:"description"`     
	NumQuestions  			   int    `json:"num_questions"`   
	TotalSectionScore          int    `json:"total_section_score"`           
}
