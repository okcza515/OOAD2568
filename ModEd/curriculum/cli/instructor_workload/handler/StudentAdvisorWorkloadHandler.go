// MEP-1008
package handler

type StudentAdvisorWorkload struct{}

func (s StudentAdvisorWorkload) Execute() {
	StudentAdvisorMenu := NewMenuHandler("Senior Project Workload Menu", true)
	StudentAdvisorMenu.Add("View Advising Project", nil)
	StudentAdvisorMenu.Add("View Committee Project", nil)
	StudentAdvisorMenu.Add("Evaluate Project", nil)
	StudentAdvisorMenu.SetBackHandler(Back{})
	StudentAdvisorMenu.SetDefaultHandler(UnknownCommand{})
	StudentAdvisorMenu.Execute()
}
