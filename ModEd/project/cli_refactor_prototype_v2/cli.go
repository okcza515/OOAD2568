package main

import (
	"ModEd/project/cli_refactor_prototype_v2/handlers"
	"ModEd/project/controller"
	"ModEd/project/utils"
)

func main() {
	db := utils.OpenDatabase("project.db")
	db.Exec("PRAGMA foreign_keys = ON;")

	// Create instance storer
	instance := controller.CreateInstance(db)

	// Handlers
	seniorProjectHandler := handlers.NewSeniorProjectHandler(instance)
	advisorHandler := handlers.NewAdvisorHandler(instance)
	committeeHandler := handlers.NewCommitteeHandler(instance)
	reportHandler := handlers.NewReportHandler(instance)
	progressHandler := handlers.NewProgressHandler(instance)
	presentationHandler := handlers.NewPresentationHandler(instance)
	groupMemberHandler := handlers.NewGroupMemberHandler(instance)
	assignmentHandler := handlers.NewAssignmentHandler(instance)

	assessmentScoreHandler := handlers.NewAssessmentScoreHandler(instance)
	assignmentScoreHandler := handlers.NewAssignmentScoreHandler(instance)
	presentationScoreHandler := handlers.NewPresentationScoreHandler(instance)
	reportScoreHandler := handlers.NewReportScoreHandler(instance)

	utils.MenuTitle("Senior Project CLI")
	builder := utils.NewMenuBuilder(&utils.MenuItem{
		Title: "Main Menu",
		Children: []*utils.MenuItem{
			{
				Title: "Senior Project Setup",
				Children: []*utils.MenuItem{
					{
						Title: "Manage Senior Project",
						Children: []*utils.MenuItem{
							{
								Title:  "Create Senior Project",
								Action: seniorProjectHandler.Create,
							},
							{
								Title:  "List Senior Projects",
								Action: seniorProjectHandler.List,
							},
						},
					},
					{
						Title: "Group Member Management",
						Children: []*utils.MenuItem{
							{Title: "View All Group Members", Action: groupMemberHandler.ViewAll},
							{Title: "Add Group Member", Action: groupMemberHandler.Add},
							{Title: "Update Group Member", Action: groupMemberHandler.Update},
							{Title: "Delete Group Member", Action: groupMemberHandler.Delete},
						},
					},
					{
						Title: "Assign Groups, Advisors, and Committees",
						Children: []*utils.MenuItem{
							{
								Title: "Advisor Manager",
								Children: []*utils.MenuItem{
									{Title: "Assign Advisor to Project", Action: advisorHandler.AssignAdvisor},
									{Title: "Update Advisor Role", Action: advisorHandler.UpdateAdvisorRole},
									{Title: "Remove Advisor", Action: advisorHandler.RemoveAdvisor},
									{Title: "List Advisors by Project", Action: advisorHandler.ListAdvisorsByProject},
									{Title: "List Projects by Instructor", Action: advisorHandler.ListProjectsByInstructor},
								},
							},
							{
								Title: "Committee Manager",
								Children: []*utils.MenuItem{
									{Title: "Add Committee Member", Action: committeeHandler.Add},
									{Title: "List Committee Members by Project", Action: committeeHandler.ListByProject},
									{Title: "List Projects by Committee Member", Action: committeeHandler.ListProjectsByInstructor},
									{Title: "Remove Committee Member", Action: committeeHandler.Remove},
								},
							},
						},
					},
				},
			},
			{
				Title: "Project Execution and Monitoring",
				Children: []*utils.MenuItem{
					{
						Title: "Assignments Management",
						Children: []*utils.MenuItem{
							{Title: "View All Assignments", Action: assignmentHandler.ViewAll},
							{Title: "Add New Assignment", Action: assignmentHandler.Add},
							{Title: "Delete Assignment", Action: assignmentHandler.Delete},
						},
					},
					{
						Title: "Presentations Management",
						Children: []*utils.MenuItem{
							{Title: "View All Presentations", Action: presentationHandler.ViewAll},
							{Title: "Add New Presentation", Action: presentationHandler.Add},
							{Title: "View Presentation by ID", Action: presentationHandler.ViewByID},
							{Title: "Update Presentation", Action: presentationHandler.Update},
							{Title: "Delete Presentation", Action: presentationHandler.Delete},
						},
					},
					{
						Title: "Track Progress",
						Children: []*utils.MenuItem{
							{Title: "View All Progress", Action: progressHandler.ViewAll},
							{Title: "Add New Progress", Action: progressHandler.Add},
							{Title: "View Progress by ID", Action: progressHandler.ViewByID},
							{Title: "Update Progress Name", Action: progressHandler.UpdateName},
							{Title: "Delete Progress", Action: progressHandler.Delete},
							{Title: "Mark Progress as Completed", Action: progressHandler.MarkCompleted},
							{Title: "Mark Progress as Incomplete", Action: progressHandler.MarkIncomplete},
						},
					},
					{
						Title: "Reports Management",
						Children: []*utils.MenuItem{
							{Title: "View All Reports", Action: reportHandler.ViewAll},
							{Title: "Add New Report", Action: reportHandler.Add},
							{Title: "View Report by ID", Action: reportHandler.ViewByID},
							{Title: "Update Report", Action: reportHandler.Update},
							{Title: "Delete Report", Action: reportHandler.Delete},
							{Title: "Submit Report", Action: reportHandler.Submit},
						},
					},
				},
			},
			{
				Title: "Evaluation & Assessment",
				Children: []*utils.MenuItem{
					{
						Title: "Evaluate Assignment",
						Children: []*utils.MenuItem{
							{Title: "Insert Score For Advisor", Action: assignmentScoreHandler.InsertAdvisorScore},
							{Title: "Insert Score For Committee", Action: assignmentScoreHandler.InsertCommitteeScore},
							{Title: "Check Score", Action: assignmentScoreHandler.CheckScore},
						},
					},
					{
						Title: "Evaluate Report",
						Children: []*utils.MenuItem{
							{Title: "For Advisor", Action: reportScoreHandler.InsertAdvisorScore},
							{Title: "For Committee", Action: reportScoreHandler.InsertCommitteeScore},
							{Title: "Check Score", Action: reportScoreHandler.CheckScore},
						},
					},
					{
						Title: "Evaluate Presentation",
						Children: []*utils.MenuItem{
							{Title: "For Advisor", Action: presentationScoreHandler.InsertAdvisorScore},
							{Title: "For Committee", Action: presentationScoreHandler.InsertCommitteeScore},
							{Title: "Check Score", Action: presentationScoreHandler.CheckScore},
						},
					},
					{
						Title: "Assessment Score Manager",
						Children: []*utils.MenuItem{
							{Title: "View Project Scores", Action: assessmentScoreHandler.ViewProjectScores},
							{Title: "Submit Advisor Score", Action: assessmentScoreHandler.SubmitAdvisorScore},
							{Title: "Submit Committee Score", Action: assessmentScoreHandler.SubmitCommitteeScore},
						},
					},
				},
			},
		},
	}, nil, nil)

	builder.Show()
}
