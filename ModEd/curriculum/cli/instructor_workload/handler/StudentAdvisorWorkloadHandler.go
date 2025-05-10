// MEP-1008
package handler
import (
	"gorm.io/gorm"
)

type StudentAdvisorWorkloadHandler struct{
	db *gorm.DB
}
func NewStudentAdvisorWorkloadHandler(db *gorm.DB) StudentAdvisorWorkloadHandler {
	return StudentAdvisorWorkloadHandler{db: db}
}

func (s StudentAdvisorWorkloadHandler) Execute() {
	StudentAdvisorMenu := NewMenuHandler("Senior Project Workload Menu", true)
	StudentAdvisorMenu.Add("View Advising Project", ViewAdvisingProject{db: s.db})
	StudentAdvisorMenu.Add("View Committee Project", ViewCommitteeProject{db: s.db})
	StudentAdvisorMenu.Add("Evaluate Project", EvaluateProject{db: s.db})
	StudentAdvisorMenu.SetBackHandler(Back{})
	StudentAdvisorMenu.SetDefaultHandler(UnknownCommand{})
	StudentAdvisorMenu.Execute()
}
