// MEP-1009 Student Internship
package internship

import (
    "ModEd/core/cli"
    handler "ModEd/curriculum/cli/Internship/handler"
    controller "ModEd/curriculum/controller"
    "ModEd/curriculum/utils"

    "gorm.io/gorm"
)

func RunInterShipCLI(
    db *gorm.DB,
    curriculumController controller.CurriculumControllerInterface,
) {

    menuManager := cli.NewCLIMenuManager()
    wrapper := controller.NewInternshipModuleWrapper(
        db,
        curriculumController,
    )
    instructorWorkloadModuleState := handler.NewInternShipModuleMenuStateHandler(menuManager, wrapper)
    menuManager.SetState(instructorWorkloadModuleState)

    for {
        menuManager.Render()
        menuManager.UserInput = utils.GetUserChoice()
        err := menuManager.HandleUserInput()
        if err != nil {
            panic(err)
        }
    }
}