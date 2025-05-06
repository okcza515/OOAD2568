package main

import (
	"ModEd/project/controller"
	"ModEd/project/model"
	"ModEd/project/utils"
	"fmt"
	"strconv"
)

func GroupMemberMenu(groupMemberController *controller.GroupMemberController) *utils.MenuItem {
	return &utils.MenuItem{
		Title: "Group Member Management",
		Children: []*utils.MenuItem{
			{
				Title: "View All Group Members",
				Action: func(io *utils.MenuIO) {
					io.Println("Listing All Group Members...")
					members, err := groupMemberController.ListAllGroupMembers()
					if err != nil {
						io.Println(fmt.Sprintf("Error: %v", err))
						return
					}
					if len(members) == 0 {
						io.Println("No group members found.")
						return
					}
					for _, m := range members {
						io.Println(fmt.Sprintf("ID: %d, Student ID: %d, Project ID: %d", m.ID, m.StudentId, m.SeniorProjectId))
					}
				},
			},
			{
				Title: "Add Group Member",
				Action: func(io *utils.MenuIO) {
					io.Print("Enter Student ID (number): ")
					studentIdStr, _ := io.ReadInput()
					studentId, err := strconv.ParseUint(studentIdStr, 10, 32)
					if err != nil {
						io.Println("Invalid Student ID.")
						return
					}

					io.Print("Enter Senior Project ID: ")
					projectIdStr, _ := io.ReadInput()
					projectId, err := strconv.ParseUint(projectIdStr, 10, 32)
					if err != nil {
						io.Println("Invalid Project ID.")
						return
					}

					member := &model.GroupMember{
						StudentId:       uint(studentId),
						SeniorProjectId: uint(projectId),
					}

					if err := groupMemberController.InsertGroupMember(member); err != nil {
						io.Println(fmt.Sprintf("Error adding group member: %v", err))
					} else {
						io.Println("Group member added successfully.")
					}
				},
			},
			{
				Title: "Update Group Member",
				Action: func(io *utils.MenuIO) {
					io.Print("Enter Group Member ID to update: ")
					idStr, _ := io.ReadInput()
					id, err := strconv.ParseUint(idStr, 10, 32)
					if err != nil {
						io.Println("Invalid ID.")
						return
					}

					member, err := groupMemberController.RetrieveGroupMember(uint(id))
					if err != nil {
						io.Println(fmt.Sprintf("Error retrieving group member: %v", err))
						return
					}

					io.Print(fmt.Sprintf("Current Student ID (%d): ", member.StudentId))
					newStudentIdStr, _ := io.ReadInput()
					if newStudentIdStr != "" {
						newStudentId, err := strconv.ParseUint(newStudentIdStr, 10, 32)
						if err != nil {
							io.Println("Invalid Student ID.")
							return
						}
						member.StudentId = uint(newStudentId)
					}

					io.Print(fmt.Sprintf("Current Project ID (%d): ", member.SeniorProjectId))
					newProjectIdStr, _ := io.ReadInput()
					if newProjectIdStr != "" {
						newProjectId, err := strconv.ParseUint(newProjectIdStr, 10, 32)
						if err != nil {
							io.Println("Invalid Project ID.")
							return
						}
						member.SeniorProjectId = uint(newProjectId)
					}

					if err := groupMemberController.UpdateGroupMember(member); err != nil {
						io.Println(fmt.Sprintf("Error updating member: %v", err))
					} else {
						io.Println("Group member updated.")
					}
				},
			},
			{
				Title: "Delete Group Member",
				Action: func(io *utils.MenuIO) {
					io.Print("Enter Group Member ID to delete: ")
					idStr, _ := io.ReadInput()
					id, err := strconv.ParseUint(idStr, 10, 32)
					if err != nil {
						io.Println("Invalid ID.")
						return
					}

					err = groupMemberController.DeleteGroupMember(uint(id))
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
