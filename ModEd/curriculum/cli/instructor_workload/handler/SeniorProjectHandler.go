// MEP-1008
package handler

type SeniorProjectWorkload struct{}

func (s SeniorProjectWorkload) Execute() {
	seniorProjectMenu := NewMenuHandler("Senior Project Workload Menu", true)
	seniorProjectMenu.Add("View Advising Project", nil)
	seniorProjectMenu.Add("View Committee Project", nil)
	seniorProjectMenu.Add("Evaluate Project", nil)
	seniorProjectMenu.SetBackHandler(Back{})
	seniorProjectMenu.SetDefaultHandler(UnknownCommand{})
	seniorProjectMenu.Execute()
}
