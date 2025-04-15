// MEP-1013
package handler

import("fmt"
"ModEd/asset/util")

func printAssetManagementOption(){
	fmt.Println("========== Asset Management ==========")
	fmt.Println("Choose your asset management type")
	fmt.Println("1. Instrument Management")
	fmt.Println("2. Supply Management")
}

func AssetManagementHandler(){
	inputBuffer := ""
	util.ClearScreen()
	util.PrintBanner()
	printAssetManagementOption()
	inputBuffer = util.GetCommandInput()
	fmt.Println(inputBuffer)
}