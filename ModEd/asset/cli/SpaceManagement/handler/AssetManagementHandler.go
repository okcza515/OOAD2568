// MEP-1013
package handler

import("fmt"
controller "ModEd/asset/controller/spacemanagement"
model "ModEd/asset/model/spacemanagement"
"ModEd/asset/util")

func printAssetManagementOption(){
	fmt.Println("========== Asset Management ==========")
	fmt.Println("Choose your asset management type")
	fmt.Println("1. Instrument Management")
	fmt.Println("2. Supply Management")
}

func AssetManagementHandler(facade *controller.SpaceManagementControllerFacade){
	inputBuffer := ""
	util.ClearScreen()
	util.PrintBanner()
	printAssetManagementOption()
	inputBuffer = util.GetCommandInput()
	fmt.Println(inputBuffer)
}