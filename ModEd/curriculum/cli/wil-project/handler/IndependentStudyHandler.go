// MEP-1010 Work Integrated Learning (WIL)
package handler

import (
	"ModEd/curriculum/controller"
	"ModEd/curriculum/model"
	"ModEd/curriculum/utils"
	"ModEd/utils/deserializer"
	"fmt"
)

func RunIndependentStudyHandler(controller *controller.WILModuleFacade) {
	for {
		printIndependentMenu()
		choice := utils.GetUserChoice()
		switch choice {
		case "1":
			fmt.Println("Migrate Independent Study")
			path := ""
			fmt.Println("Please enter the path of the independent study(s) file (csv or json): ")
			_, _ = fmt.Scanln(&path)
			fd, err := deserializer.NewFileDeserializer(path)
			if err != nil {
				fmt.Println(err)
				continue
			}
			isModel := new([]model.IndependentStudy)
			if err := fd.Deserialize(isModel); err != nil {
				fmt.Println(err)
				continue
			}
			for _, is := range *isModel {
				controller.IndependentStudyController.Insert(&is)
			}
			fmt.Println("Migration Success!")
		case "0":
			fmt.Println("Exiting...")
			return
		default:
			fmt.Println("Invalid option")
		}
	}
}

func printIndependentMenu() {
	fmt.Println("\nIndependent Study Menu:")
	fmt.Println("1.Migrate independent study")
	fmt.Println("0.Exit Independent Study module")
}
