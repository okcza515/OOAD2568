// MEP-1008
package handler

import (
	instructorWorkloadController "ModEd/curriculum/controller"
	"ModEd/curriculum/model"
	projectModel "ModEd/project/model"
	"bufio"
	"fmt"
	"os"
	"strings"

	"gorm.io/gorm"
)

type SeniorProjectWorkloadHandler struct {
	db *gorm.DB
}

func NewSeniorProjectWorkloadHandler(db *gorm.DB) SeniorProjectWorkloadHandler {
	return SeniorProjectWorkloadHandler{db: db}
}

func (s SeniorProjectWorkloadHandler) Execute() {
	seniorProjectMenu := NewMenuHandler("Senior Project Workload Menu", true)
	seniorProjectMenu.Add(string(MENU_SENIOR_PROJECT_VIEW_ADVISOR_PROJECT), ViewAdvisingProject{db: s.db})
	seniorProjectMenu.Add(string(MENU_SENIOR_PROJECT_VIEW_COMMITTEE_PROJECT), ViewCommitteeProject{db: s.db})
	seniorProjectMenu.Add(string(MENU_SENIOR_PROJECT_EVALUATE_PROJECT), EvaluateProject{db: s.db})
	seniorProjectMenu.SetBackHandler(Back{})
	seniorProjectMenu.SetDefaultHandler(UnknownCommand{})
	seniorProjectMenu.Execute()
}

type ViewAdvisingProject struct {
	db *gorm.DB
}

func (v ViewAdvisingProject) Execute() {
	// controller := projectController.NewAdvisorController(v.db)
	// controller.ListAdvisorsByInstructor(1)
}

type ViewCommitteeProject struct {
	db *gorm.DB
}

func (v ViewCommitteeProject) Execute() {
	// controller := projectController.NewCommitteeController(v.db)
	// controller.ListCommitteesByInstructor(1)
}

type EvaluateProject struct {
	db *gorm.DB
}

func (e EvaluateProject) Execute() {
	controller := instructorWorkloadController.NewProjectController(e.db)

	reader := bufio.NewReader(os.Stdin)

	fmt.Println("\nEnter evaluation type (Assignment, Proposal, Report, Presentation):")
	fmt.Print("Enter: ")

	input, _ := reader.ReadString('\n')
	input = strings.TrimSpace(input)

	if !model.IsValidAssignmentType(input) {
		fmt.Println("Invalid type. Defaulting to 'Assignment'.")
		input = string(model.ASSIGNMENT)
	}

	taskType := model.ProjectEvaluationTypeEnum(input)
	mockEvaluation := &model.ProjectEvaluation{
		GroupId:        1,
		AssignmentId:   1,
		AssignmentType: string(taskType),
		Score:          0.0,
		Comment:        "",
	}
	mockCriteria := []projectModel.AssessmentCriteria{
		{
			CriteriaName: "Criteria A",
		},
		{
			CriteriaName: "Criteria B",
		},
		{
			CriteriaName: "Criteria C",
		},
	}
	controller.CreateEvaluation(mockEvaluation, string(taskType), mockCriteria)
}
