package main

import (
	"ModEd/core"
	"ModEd/core/migration"
	internship "ModEd/curriculum/cli/Internship"
	"ModEd/curriculum/cli/curriculum"
	instructorWorkload "ModEd/curriculum/cli/instructor_workload"
	wilproject "ModEd/curriculum/cli/wil-project"
	controller "ModEd/curriculum/controller"
	"fmt"

	"gorm.io/gorm"
)

const (
	defaultDBPath = "../w../data/ModEd.bin"
)

type Command interface {
	Execute() error
}

// CurriculumCommand
type CurriculumCommand struct {
	db                   *gorm.DB
	courseController     controller.CourseControllerInterface
	classController      controller.ClassControllerInterface
	curriculumController controller.CurriculumControllerInterface
}

func (c *CurriculumCommand) Execute() error {
	curriculum.RunCurriculumModuleCLI(&curriculum.CurriculumCLIParams{
		CourseController:     c.courseController,
		ClassController:      c.classController,
		CurriculumController: c.curriculumController,
	})
	return nil
}

// WILProjectCommand
type WILProjectCommand struct {
	db               *gorm.DB
	courseController controller.CourseControllerInterface
	classController  controller.ClassControllerInterface
}

func (w *WILProjectCommand) Execute() error {
	wilproject.RunWILModuleCLI(w.db, w.courseController, w.classController)
	return nil
}

// InstructorWorkloadCommand
type InstructorWorkloadCommand struct {
	db                   *gorm.DB
	courseController     controller.CourseControllerInterface
	classController      controller.ClassControllerInterface
	curriculumController controller.CurriculumControllerInterface
}

func (i *InstructorWorkloadCommand) Execute() error {
	instructorWorkload.RunInstructorWorkloadModuleCLI(i.db, i.courseController, i.classController, i.curriculumController)
	return nil
}

// InternshipCommand
type InternshipCommand struct {
	db                   *gorm.DB
	curriculumController controller.CurriculumControllerInterface
}

func (i *InternshipCommand) Execute() error {
	internship.RunInterShipCLI(i.db, i.curriculumController)
	return nil
}

// ResetDBCommand
type ResetDBCommand struct{}

func (r *ResetDBCommand) Execute() error {
	return resetDB()
}

type CommandExecutor struct {
	commands map[string]Command
}

func NewCommandExecutor() *CommandExecutor {
	return &CommandExecutor{commands: make(map[string]Command)}
}

func (ce *CommandExecutor) RegisterCommand(name string, command Command) {
	ce.commands[name] = command
}

func (ce *CommandExecutor) ExecuteCommand(name string) error {
	if command, exists := ce.commands[name]; exists {
		return command.Execute()
	}
	return fmt.Errorf("invalid command: %s", name)
}

func main() {
	db, err := migration.
		GetInstance().
		SetPathDB(defaultDBPath).
		MigrateModule(core.MODULE_CURRICULUM).
		MigrateModule(core.MODULE_INSTRUCTOR).
		MigrateModule(core.MODULE_INTERNSHIP).
		MigrateModule(core.MODULE_WILPROJECT).
		BuildDB()

	if err != nil {
		panic(err)
	}

	curriculumController := controller.NewCurriculumController(db)
	classController := controller.NewClassController(db)
	courseController := controller.NewCourseController(db)

	commandExecutor := NewCommandExecutor()
	commandExecutor.RegisterCommand("1", &CurriculumCommand{db, courseController, classController, curriculumController})
	commandExecutor.RegisterCommand("2", &WILProjectCommand{db, courseController, classController})
	commandExecutor.RegisterCommand("3", &InstructorWorkloadCommand{db, courseController, classController, curriculumController})
	commandExecutor.RegisterCommand("4", &InternshipCommand{db, curriculumController})
	commandExecutor.RegisterCommand("resetdb", &ResetDBCommand{})

	for {
		displayMainMenu()
		choice := getUserChoice()

		if choice == "0" {
			fmt.Println("Exiting...")
			return
		}

		if err := commandExecutor.ExecuteCommand(choice); err != nil {
			fmt.Printf("Error: %v\n", err)
		}
	}
}

func displayMainMenu() {
	fmt.Println("\nModEd CLI Menu")
	fmt.Println("1. Curriculum")
	fmt.Println("2. WIL-Project")
	fmt.Println("3. Instructor Workload")
	fmt.Println("4. Internship")
	fmt.Println("'resetdb' to re-initialize database")
}

func getUserChoice() string {
	var choice string
	fmt.Print("Enter your choice: ")
	fmt.Scanln(&choice)
	return choice
}

func resetDB() error {
	err := migration.GetInstance().DropAllTables()
	if err != nil {
		return err
	}

	_, err = migration.GetInstance().
		SetPathDB(defaultDBPath).
		MigrateModule(core.MODULE_CURRICULUM).
		MigrateModule(core.MODULE_INSTRUCTOR).
		MigrateModule(core.MODULE_INTERNSHIP).
		MigrateModule(core.MODULE_WILPROJECT).
		BuildDB()

	if err != nil {
		return err
	}

	return nil
}
