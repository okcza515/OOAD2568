package cli

import (
	"fmt"
	"time"

	"ModEd/eval/controller"
	"ModEd/eval/model"
)

func AddAssignment(controller *controller.AssignmentController) {
	var assignment model.Assignment
	var title, description string
	var dueDateStr string
	var statusChoice int

	fmt.Print("Enter assignment title: ")
	fmt.Scan(&title)
	fmt.Print("Enter assignment description: ")
	fmt.Scan(&description)

	var status string
	for {
		fmt.Println("\nSelect assignment status:")
		fmt.Println("1. Active")
		fmt.Println("2. Inactive")
		fmt.Print("Enter your choice (1-2): ")
		fmt.Scan(&statusChoice)

		switch statusChoice {
		case 1:
			status = "Active"
			break
		case 2:
			status = "Inactive"
			break
		default:
			fmt.Println("Invalid choice. Please select 1 or 2.")
			continue
		}
		break
	}

	fmt.Print("Enter due date (YYYY-MM-DD): ")
	fmt.Scan(&dueDateStr)

	dueDate, err := time.Parse("2006-01-02", dueDateStr)
	if err != nil {
		fmt.Println("Invalid date format. Please use YYYY-MM-DD")
		return
	}

	assignment = model.Assignment{
		Title:       title,
		Description: description,
		Status:      status,
		StartDate:   time.Now(),
		DueDate:     dueDate,
		Released:    true,
	}

	err = controller.AddAssignment(assignment)
	if err != nil {
		fmt.Println("Error adding assignment:", err)
		return
	}
	fmt.Println("Assignment added successfully!")
}

func DeleteAssignment(controller *controller.AssignmentController) {
	var id uint
	fmt.Print("Enter assignment ID to delete: ")
	fmt.Scan(&id)

	err := controller.DeleteAssignment(id)
	if err != nil {
		fmt.Println("Error deleting assignment:", err)
		return
	}
	fmt.Println("Assignment deleted successfully!")
}

func UpdateAssignment(controller *controller.AssignmentController) {
	var id uint
	fmt.Print("Enter assignment ID to update: ")
	fmt.Scan(&id)

	// First get the existing assignment
	assignment, err := controller.GetAssignmentByID(id)
	if err != nil {
		fmt.Println("Error finding assignment:", err)
		return
	}

	var title, description, status string
	var dueDateStr string

	fmt.Printf("Current title (%s): ", assignment.Title)
	fmt.Scan(&title)
	if title != "" {
		assignment.Title = title
	}

	fmt.Printf("Current description (%s): ", assignment.Description)
	fmt.Scan(&description)
	if description != "" {
		assignment.Description = description
	}

	fmt.Printf("Current status (%s): ", assignment.Status)
	fmt.Scan(&status)
	if status != "" {
		assignment.Status = status
	}

	fmt.Print("New due date (YYYY-MM-DD) or press enter to keep current: ")
	fmt.Scan(&dueDateStr)
	if dueDateStr != "" {
		dueDate, err := time.Parse("2006-01-02", dueDateStr)
		if err != nil {
			fmt.Println("Invalid date format. Please use YYYY-MM-DD")
			return
		}
		assignment.DueDate = dueDate
	}

	err = controller.UpdateAssignment(*assignment)
	if err != nil {
		fmt.Println("Error updating assignment:", err)
		return
	}
	fmt.Println("Assignment updated successfully!")
}

func DisplayAllAssignments(controller *controller.AssignmentController) {
	assignments := controller.GetAllAssignments()

	if len(assignments) == 0 {
		fmt.Println("No assignments found.")
		return
	}

	fmt.Println("\nAll Assignments:")
	fmt.Println("===============")
	for _, a := range assignments {
		fmt.Printf("ID: %d\n", a.ID)
		fmt.Printf("Title: %s\n", a.Title)
		fmt.Printf("Description: %s\n", a.Description)
		fmt.Printf("Status: %s\n", a.Status)
		fmt.Printf("Start Date: %s\n", a.StartDate.Format("2006-01-02"))
		fmt.Printf("Due Date: %s\n", a.DueDate.Format("2006-01-02"))
		fmt.Println("---------------")
	}
}

func QueryAssignmentByID(controller *controller.AssignmentController) {
	var id uint
	fmt.Print("Enter assignment ID: ")
	fmt.Scan(&id)

	assignment, err := controller.GetAssignmentByID(id)
	if err != nil {
		fmt.Println("Error retrieving assignment:", err)
		return
	}

	fmt.Printf("\nAssignment Details:\n")
	fmt.Printf("ID: %d\n", assignment.ID)
	fmt.Printf("Title: %s\n", assignment.Title)
	fmt.Printf("Description: %s\n", assignment.Description)
	fmt.Printf("Status: %s\n", assignment.Status)
	fmt.Printf("Start Date: %s\n", assignment.StartDate.Format("2006-01-02"))
	fmt.Printf("Due Date: %s\n", assignment.DueDate.Format("2006-01-02"))
}

func QueryAssignmentsByStatus(controller *controller.AssignmentController) {
	var status string
	fmt.Print("Enter status to search for: ")
	fmt.Scan(&status)

	assignments := controller.GetAssignmentsByStatus(status)

	if len(assignments) == 0 {
		fmt.Println("No assignments found with status:", status)
		return
	}

	fmt.Printf("\nAssignments with status '%s':\n", status)
	fmt.Println("=========================")
	for _, a := range assignments {
		fmt.Printf("ID: %d\n", a.ID)
		fmt.Printf("Title: %s\n", a.Title)
		fmt.Printf("Description: %s\n", a.Description)
		fmt.Printf("Start Date: %s\n", a.StartDate.Format("2006-01-02"))
		fmt.Printf("Due Date: %s\n", a.DueDate.Format("2006-01-02"))
		fmt.Println("---------------")
	}
}

func QueryAssignmentsByDateRange(controller *controller.AssignmentController) {
	var startDateStr, endDateStr string
	fmt.Print("Enter start date (YYYY-MM-DD): ")
	fmt.Scan(&startDateStr)
	fmt.Print("Enter end date (YYYY-MM-DD): ")
	fmt.Scan(&endDateStr)

	startDate, err := time.Parse("2006-01-02", startDateStr)
	if err != nil {
		fmt.Println("Invalid start date format. Please use YYYY-MM-DD")
		return
	}

	endDate, err := time.Parse("2006-01-02", endDateStr)
	if err != nil {
		fmt.Println("Invalid end date format. Please use YYYY-MM-DD")
		return
	}

	assignments := controller.GetAssignmentsByDateRange(startDate, endDate)

	if len(assignments) == 0 {
		fmt.Println("No assignments found in the specified date range.")
		return
	}

	fmt.Printf("\nAssignments between %s and %s:\n", startDateStr, endDateStr)
	fmt.Println("=========================")
	for _, a := range assignments {
		fmt.Printf("ID: %d\n", a.ID)
		fmt.Printf("Title: %s\n", a.Title)
		fmt.Printf("Description: %s\n", a.Description)
		fmt.Printf("Status: %s\n", a.Status)
		fmt.Printf("Start Date: %s\n", a.StartDate.Format("2006-01-02"))
		fmt.Printf("Due Date: %s\n", a.DueDate.Format("2006-01-02"))
		fmt.Println("---------------")
	}
}
