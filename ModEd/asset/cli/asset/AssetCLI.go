package main

import (
	controller "ModEd/asset/controller/asset"
	util "ModEd/asset/util"
	"fmt"
)

func main() {
	facade, err := controller.CreateAssetControllerFacade()
	if err != nil {
		panic("err: initialize controllers failed")
	}

	util.PrintBanner()

	supplies, err := facade.Supply.GetAll()
	fmt.Println(supplies)
}
