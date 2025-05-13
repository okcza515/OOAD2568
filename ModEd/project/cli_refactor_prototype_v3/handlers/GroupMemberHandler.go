package handlers

import (
	"ModEd/core"
	"ModEd/project/controller"
	"ModEd/project/model"
	"fmt"
)

type GroupMemberHandler struct {
	menuIO         *core.MenuIO
	instanceStorer *controller.InstanceStorer
}

func NewGroupMemberHandler(instanceStorer *controller.InstanceStorer) *GroupMemberHandler {
	return &GroupMemberHandler{
		menuIO:         core.NewMenuIO(),
		instanceStorer: instanceStorer,
	}
}

func (h *GroupMemberHandler) ViewAll(io *core.MenuIO) {
	io.Println("Listing All Group Members...")
	members, err := h.instanceStorer.GroupMember.List(map[string]interface{}{})
	if err != nil {
		io.Println(fmt.Sprintf("Error: %v", err))
		return
	}
	if len(members) == 0 {
		io.Println("No group members found.")
		return
	}
	io.PrintTableFromSlice(members, []string{"ID", "StudentId", "SeniorProjectId"})
}

func (h *GroupMemberHandler) Add(io *core.MenuIO) {
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

	if err := h.instanceStorer.GroupMember.Insert(member); err != nil {
		io.Println(fmt.Sprintf("Error adding group member: %v", err))
	} else {
		io.Println("Group member added successfully.")
	}
}

func (h *GroupMemberHandler) Update(io *core.MenuIO) {
	io.Print("Enter Group Member ID to update: ")
	id, err := io.ReadInputID()
	if err != nil {
		return
	}

	member, err := h.instanceStorer.GroupMember.RetrieveByID(id)
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

	if err := h.instanceStorer.GroupMember.UpdateByID(member); err != nil {
		io.Println(fmt.Sprintf("Error updating member: %v", err))
	} else {
		io.Println("Group member updated.")
	}
}

func (h *GroupMemberHandler) Delete(io *core.MenuIO) {
	io.Print("Enter Group Member ID to delete: ")
	id, err := io.ReadInputID()
	if err != nil {
		return
	}

	err = h.instanceStorer.GroupMember.DeleteByID(id)
	if err != nil {
		io.Println(fmt.Sprintf("Error deleting member: %v", err))
	} else {
		io.Println("Group member deleted.")
	}
}
