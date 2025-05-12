package handlers

import (
	"ModEd/core"
	"ModEd/project/controller"
	"fmt"
)

type ReportHandler struct {
	menuIO         *core.MenuIO
	instanceStorer *controller.InstanceStorer
}

func NewReportHandler(instanceStorer *controller.InstanceStorer) *ReportHandler {
	return &ReportHandler{
		menuIO:         core.NewMenuIO(),
		instanceStorer: instanceStorer,
	}
}

func (h *ReportHandler) ViewAll(io *core.MenuIO) {
	io.Println("Viewing Report...")

	reports, err := h.instanceStorer.Report.GetFormattedReportList()
	if err != nil {
		io.Println(fmt.Sprintf("Error retrieving reports: %v", err))
		return
	}

	if len(reports) == 0 {
		io.Println("No reports found.")
		return
	}

	io.Println("Report (Based on Due Dates):")
	for _, r := range reports {
		io.Println(r)
	}
}

func (h *ReportHandler) Add(io *core.MenuIO) {
	io.Println("Adding New Report...")

	io.Print("Enter Senior Project ID: ")
	projectID, err := io.ReadInputID()
	if err != nil {
		return
	}

	io.Print("Enter Report Type (Idea, Proposal, Progress, Midterm, Final): ")
	reportType, err := io.ReadInput()
	if err != nil {
		io.Println(fmt.Sprintf("Error reading input: %v", err))
		return
	}

	io.Print("Enter Due Date (YYYY-MM-DD): ")
	dueDate, err := io.ReadInputTime()
	if err != nil {
		io.Println(fmt.Sprintf("Invalid Due Date format: %v", err))
		return
	}

	err = h.instanceStorer.Report.AddNewReport(projectID, reportType, dueDate)
	if err != nil {
		io.Println(fmt.Sprintf("Error adding new report: %v", err))
	} else {
		io.Println("Report added successfully!")
	}
}

func (h *ReportHandler) ViewByID(io *core.MenuIO) {
	io.Println("Viewing Report by ID...")

	io.Print("Enter Report ID: ")
	reportID, err := io.ReadInputID()
	if err != nil {
		return
	}

	report, err := h.instanceStorer.Report.RetrieveByID(reportID)
	if err != nil {
		io.Println(fmt.Sprintf("Error retrieving report: %v", err))
		return
	}

	io.Println(report.ToString())
}

func (h *ReportHandler) Update(io *core.MenuIO) {
	io.Println("Updating Report...")

	io.Print("Enter Report ID to update: ")
	reportID, err := io.ReadInputID()
	if err != nil {
		return
	}

	io.Print("Enter new Due Date (YYYY-MM-DD): ")
	newDueDate, err := io.ReadInputTime()
	if err != nil {
		return
	}

	err = h.instanceStorer.Report.UpdateReport(reportID, newDueDate)
	if err != nil {
		io.Println(fmt.Sprintf("Error updating report: %v", err))
	} else {
		io.Println("Report updated successfully!")
	}
}

func (h *ReportHandler) Delete(io *core.MenuIO) {
	io.Println("Deleting Report...")

	io.Print("Enter Report ID to delete: ")
	reportID, err := io.ReadInputID()
	if err != nil {
		return
	}

	err = h.instanceStorer.Report.DeleteByID(reportID)
	if err != nil {
		io.Println(fmt.Sprintf("Error deleting report: %v", err))
	} else {
		io.Println("Report deleted successfully!")
	}
}

func (h *ReportHandler) Submit(io *core.MenuIO) {
	io.Println("Submitting Report...")

	io.Print("Enter Report ID to submit: ")
	reportID, err := io.ReadInputID()
	if err != nil {
		return
	}

	err = h.instanceStorer.Report.SubmitReport(reportID)
	if err != nil {
		io.Println(fmt.Sprintf("Error submitting report: %v", err))
	} else {
		io.Println("Report submitted successfully!")
	}
}
