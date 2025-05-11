package handler

import (
	"ModEd/common/model"
	"ModEd/curriculum/controller"
	curriculumModel "ModEd/curriculum/model"
	"ModEd/curriculum/utils"
	"fmt"
	"strconv"
)

type createCurriculumHandler struct {
	curriculumController controller.CurriculumControllerInterface
}

func NewCreateCurriculumHandler(curriculumController controller.CurriculumControllerInterface) *createCurriculumHandler {
	return &createCurriculumHandler{
		curriculumController: curriculumController,
	}
}

func (h *createCurriculumHandler) Execute() error {
	fmt.Println("\nCreate New Curriculum:")

	name := utils.GetUserInput("Enter name: ")

	startYearStr := utils.GetUserInput("Enter start year: ")
	startYear, err := strconv.Atoi(startYearStr)
	if err != nil {
		fmt.Println("Invalid start year format")
		return err
	}

	endYearStr := utils.GetUserInput("Enter end year: ")
	endYear, err := strconv.Atoi(endYearStr)
	if err != nil {
		fmt.Println("Invalid end year format")
		return err
	}

	fmt.Println("Program Type Options:")
	for key, value := range model.ProgramTypeLabel {
		fmt.Printf("%d. %s\n", key, value)
	}

	programTypeStr := utils.GetUserInput("Select program type (enter number): ")
	programType, err := strconv.Atoi(programTypeStr)
	if err != nil {
		fmt.Println("Invalid program type")
		return err
	}

	departmentIdStr := utils.GetUserInput("Select department (enter number): ")
	departmentId, err := strconv.Atoi(departmentIdStr)
	if err != nil {
		fmt.Println("Invalid department ID")
		return err
	}

	curriculum := &curriculumModel.Curriculum{
		Name:         name,
		StartYear:    startYear,
		EndYear:      endYear,
		ProgramType:  model.ProgramType(programType),
		DepartmentId: uint(departmentId),
	}

	// Validate the curriculum
	if err := curriculum.Validate(); err != nil {
		fmt.Println("Validation error:", err)
		return err
	}

	fmt.Println("\nCurriculum to be created:")
	fmt.Printf("Name: %s\n", curriculum.Name)
	fmt.Printf("Start Year: %d\n", curriculum.StartYear)
	fmt.Printf("End Year: %d\n", curriculum.EndYear)
	fmt.Printf("Program Type: %s\n", model.ProgramTypeLabel[curriculum.ProgramType])

	confirm := utils.GetUserInput("\nConfirm creation? (y/n): ")
	if confirm != "y" {
		fmt.Println("Creation cancelled.")
		return nil
	}

	_, err = h.curriculumController.CreateCurriculum(curriculum)
	if err != nil {
		fmt.Println("Error creating curriculum:", err)
		return err
	}

	fmt.Println("\nCurriculum created successfully:")
	return nil
}
