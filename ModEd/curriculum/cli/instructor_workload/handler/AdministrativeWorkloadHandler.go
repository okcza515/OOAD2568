// MEP-1008
package handler

type AdminstrativeWorkload struct{}

func (a AdminstrativeWorkload) Execute() {
	adminstrativeMenu := NewMenuHandler("Adminstrative Workload Menu", true)
	adminstrativeMenu.Add("Meeting", nil)
	adminstrativeMenu.Add("Student Request", nil)
	adminstrativeMenu.SetBackHandler(Back{})
	adminstrativeMenu.SetDefaultHandler(UnknownCommand{})
	adminstrativeMenu.Execute()
}
