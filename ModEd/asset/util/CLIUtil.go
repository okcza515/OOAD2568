package util

import (
	"fmt"
	"os"
	"os/exec"
	"reflect"
	"runtime"
)

func ClearScreen() {
	if runtime.GOOS == "windows" {
		cmd := exec.Command("cmd", "/c", "cls")
		cmd.Stdout = os.Stdout
		err := cmd.Run()
		if err != nil {
			panic(err)
		}
	} else {
		cmd := exec.Command("clear")
		cmd.Stdout = os.Stdout
		err := cmd.Run()
		if err != nil {
			panic(err)
		}
	}
}

func PrintStruct(s interface{}) {
	v := reflect.ValueOf(s)
	t := reflect.TypeOf(s)

	for i := 0; i < v.NumField(); i++ {
		fmt.Printf("%s: %v\n", t.Field(i).Name, v.Field(i).Interface())
	}
	fmt.Println("___________________________________________________")
}

func PrintBanner() {
	fmt.Print(`
    __  ___          ________    __   ___                   __ 
   /  |/  /___  ____/ / ____/___/ /  /   |  _____________  / /_
  / /|_/ / __ \/ __  / __/ / __  /  / /| | / ___/ ___/ _ \/ __/
 / /  / / /_/ / /_/ / /___/ /_/ /  / ___ |(__  |__  )  __/ /_  
/_/__/_/\____/\__,_/_____/\__,_/  /_/  |_/____/____/\___/\__/  
  / ___/___  ______   __(_)_______     / ____/ /   /  _/       
  \__ \/ _ \/ ___/ | / / / ___/ _ \   / /   / /    / /         
 ___/ /  __/ /   | |/ / / /__/  __/  / /___/ /____/ /          
/____/\___/_/    |___/_/\___/\___/   \____/_____/___/          

`)
}

func PrintByeBye() {
	fmt.Println("+-------------------------------------------------+")
	fmt.Println("|            Good bye, Have a Great Day !         |")
	fmt.Println("+-------------------------------------------------+")
}

func GetCommandInput() string {
	var buffer string

	fmt.Print("Command: ")
	_, _ = fmt.Scanln(&buffer)
	fmt.Println()

	return buffer
}

func PressEnterToContinue() {
	fmt.Println()
	fmt.Print("[Press Enter to Continue]")
	var buffer string
	_, _ = fmt.Scanln(&buffer)
}
