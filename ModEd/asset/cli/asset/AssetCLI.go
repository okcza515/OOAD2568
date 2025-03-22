package main

import (
	//controller "ModEd/asset/controller/asset"
	util "ModEd/asset/util"
	"fmt"
)

func main() {
	//facade, err := controller.CreateAssetControllerFacade()
	//if err != nil {
	//	panic("err: initialize controllers failed")
	//}

	util.PrintBanner()

	//supplies, err := facade.Supply.GetAll()
	//fmt.Println(supplies)

	for {
		inputBuffer := ""
		_, _ = fmt.Scanln(&inputBuffer)

		switch inputBuffer {
		case "1":
			fmt.Println("Menu 1: ")
			// your function goes here
		case "2":
			fmt.Println("Menu 2: ")
			// your function goes here
		case "3":
			fmt.Println("Menu 3: ")
			// your function goes here
		case "4":
			fmt.Println("Menu 4: ")
			// your function goes here
		case "exit":
			break
		}
	}

}
