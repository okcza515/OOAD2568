package menu

import (
	"ModEd/core"
	"ModEd/project/controller"
	"ModEd/project/model"
	"ModEd/project/utils"
	"fmt"
)

func GroupMemberMenu(groupMemberController *controller.GroupMemberController) *utils.MenuItem {
	return &utils.MenuItem{
		Title: "Group Member Management",
		Children: []*utils.MenuItem{
			{
				Title: "View All Group Members",
				Action: func(io *core.MenuIO) {
					io.Println("Listing All Group Members...")
					members, err := groupMemberController.List(map[string]interface{}{})
					if err != nil {
						io.Println(fmt.Sprintf("Error: %v", err))
						return
					}
					if len(members) == 0 {
						io.Println("No group members found.")
						return
					}
					io.PrintTableFromSlice(members, []string{"ID", "StudentId", "SeniorProjectId"})
				},
			},
			{
				Title: "Add Group Member",
				Action: func(io *core.MenuIO) {
					io.Print("Enter Student ID (number): ")
					studentId, err := io.ReadInputID()
					if err != nil {
						return
					}

					io.Print("Enter Senior Project ID: ")
					projectId, err := io.ReadInputID()
					if err != nil {
						return
					}

					member := &model.GroupMember{
						StudentId:       uint(studentId),
						SeniorProjectId: uint(projectId),
					}

					if err := groupMemberController.Insert(member); err != nil {
						io.Println(fmt.Sprintf("Error adding group member: %v", err))
					} else {
						io.Println("Group member added successfully.")
					}
				},
			},
			{
				Title: "Update Group Member",
				Action: func(io *core.MenuIO) {
					io.Print("Enter Group Member ID to update: ")
					id, _ := io.ReadInputID()

					member, err := groupMemberController.RetrieveByID(id)
					if err != nil {
						io.Println(fmt.Sprintf("Error retrieving group member: %v", err))
						return
					}

					io.Print(fmt.Sprintf("Current Student ID (%d): ", member.StudentId))
					newStudentId, err := io.ReadInputID()
					if err != nil {
						return
					}
					member.StudentId = newStudentId

					io.Print(fmt.Sprintf("Current Project ID (%d): ", member.SeniorProjectId))
					newProjectId, err := io.ReadInputID()
					if err != nil {
						return
					}
					member.SeniorProjectId = newProjectId

					if err := groupMemberController.UpdateByID(member); err != nil {
						io.Println(fmt.Sprintf("Error updating member: %v", err))
					} else {
						io.Println("Group member updated.")
					}
				},
			},
			{
				Title: "Delete Group Member",
				Action: func(io *core.MenuIO) {
					io.Print("Enter Group Member ID to delete: ")
					id, err := io.ReadInputID()
					if err != nil {
						return
					}

					err = groupMemberController.DeleteByID(id)
					if err != nil {
						io.Println(fmt.Sprintf("Error deleting member: %v", err))
					} else {
						io.Println("Group member deleted.")
					}
				},
			},
		},
	}
}
