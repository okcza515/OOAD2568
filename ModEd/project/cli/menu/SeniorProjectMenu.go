package menu

import (
	"ModEd/project/controller"
	"ModEd/project/model"
	"ModEd/project/utils"
)

func BuildSeniorProjectMenu(seniorProjectController *controller.SeniorProjectController) *utils.MenuItem {
	return &utils.MenuItem{
		Title: "Manage Senior Project",
		Children: []*utils.MenuItem{
			{
				Title: "Create Senior Project",
				Action: func(io *utils.MenuIO) {
					io.Print("Enter the group name (-1 to cancel): ")
					groupNameStr, err := io.ReadInput()
					if err != nil || groupNameStr == "-1" {
						io.Println("Cancelled.")
						return
					}

					if err := seniorProjectController.Insert(&model.SeniorProject{
						GroupName: groupNameStr,
					}); err != nil {
						io.Println(err.Error())
						return
					}
				},
			},
			{
				Title: "List Senior Projects",
				Action: func(io *utils.MenuIO) {
					records, err := seniorProjectController.List(map[string]interface{}{})
					if err != nil {
						io.Println(err.Error())
						return
					}

					io.PrintTableFromSlice(records, []string{"ID", "GroupName", "CreatedAt"})
				},
			},
		},
	}
}
