package handler

import (
	"fmt"
)

type AcademicWorkload struct{}
type Back struct{}

func (b Back) Execute() {
	return
}

type UnknownCommand struct{}

func (u UnknownCommand) Execute() {
	fmt.Println("Unknown command, please try again.")
}

func (c AcademicWorkload) Execute() {
	academicMenu := NewMenuHandler("Academic Workload Menu", true)
	academicMenu.Add("Curriculum", nil)
	academicMenu.Add("Course", nil)
	academicMenu.Add("Class", nil)
	academicMenu.Add("Class Material", nil)
	academicMenu.Add("Course Plan", nil)
	academicMenu.SetBackHandler(Back{})
	academicMenu.SetDefaultHandler(UnknownCommand{})
	academicMenu.Execute()
}
