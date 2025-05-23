package util

import (
	"fmt"
	"os"
	"os/exec"
	"reflect"
	"runtime"
	"strconv"
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

func PrintSpaceManagementBanner() {
	fmt.Print(`
░██████╗██████╗░░█████╗░░█████╗░███████╗     ███╗░░░███╗░█████╗░███╗░░██╗░█████╗░░██████╗░███████╗███╗░░░███╗███████╗███╗░░██╗████████╗
██╔════╝██╔══██╗██╔══██╗██╔══██╗██╔════╝     ████╗░████║██╔══██╗████╗░██║██╔══██╗██╔════╝░██╔════╝████╗░████║██╔════╝████╗░██║╚══██╔══╝
╚█████╗░██████╔╝███████║██║░░╚═╝█████╗░░     ██╔████╔██║███████║██╔██╗██║███████║██║░░██╗░█████╗░░██╔████╔██║█████╗░░██╔██╗██║░░░██║░░░
░╚═══██╗██╔═══╝░██╔══██║██║░░██╗██╔══╝░░     ██║╚██╔╝██║██╔══██║██║╚████║██╔══██║██║░░╚██╗██╔══╝░░██║╚██╔╝██║██╔══╝░░██║╚████║░░░██║░░░
██████╔╝██║░░░░░██║░░██║╚█████╔╝███████╗     ██║░╚═╝░██║██║░░██║██║░╚███║██║░░██║╚██████╔╝███████╗██║░╚═╝░██║███████╗██║░╚███║░░░██║░░░
╚═════╝░╚═╝░░░░░╚═╝░░╚═╝░╚════╝░╚══════╝     ╚═╝░░░░░╚═╝╚═╝░░╚═╝╚═╝░░╚══╝╚═╝░░╚═╝░╚═════╝░╚══════╝╚═╝░░░░░╚═╝╚══════╝╚═╝░░╚══╝░░░╚═╝░░░

░██████╗███████╗██████╗░██╗░░░██╗██╗░█████╗░███████╗
██╔════╝██╔════╝██╔══██╗██║░░░██║██║██╔══██╗██╔════╝
╚█████╗░█████╗░░██████╔╝╚██╗░██╔╝██║██║░░╚═╝█████╗░░
░╚═══██╗██╔══╝░░██╔══██╗░╚████╔╝░██║██║░░██╗██╔══╝░░
██████╔╝███████╗██║░░██║░░╚██╔╝░░██║╚█████╔╝███████╗
╚═════╝░╚══════╝╚═╝░░╚═╝░░░╚═╝░░░╚═╝░╚════╝░╚══════╝`)
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

func GetUintInput(prompt string) uint {
	var input uint
	fmt.Print(prompt)
	_, err := fmt.Scanln(&input)
	if err != nil {
		fmt.Println("Invalid input. Please enter a positive number.")
		return GetUintInput(prompt)
	}
	return input
}

func GetUintPointerInput(prompt string) *uint {
	fmt.Print(prompt)
	var input string
	_, err := fmt.Scanln(&input)
	if err != nil || input == "" {
		return nil
	}

	var value uint
	_, err = fmt.Sscanf(input, "%d", &value)
	if err != nil {
		return nil
	}

	return &value
}

func GetFloatInput(prompt string) float64 {
	var input float64
	fmt.Print(prompt)
	_, err := fmt.Scanln(&input)
	if err != nil {
		fmt.Println("Invalid input. Please enter a number.")
		return GetFloatInput(prompt)
	}
	return input
}

func GetStringInput(prompt string) string {
	var input string
	fmt.Print(prompt)
	fmt.Scanln(&input)
	return input
}

func GetOptionalStringInput(prompt, current string) string {
	input := GetStringInput(fmt.Sprintf("%s (Current: %s): ", prompt, current))
	if input == "" {
		return current
	}
	return input
}

func GetOptionalUintInput(prompt string, current uint) uint {
	input := GetStringInput(fmt.Sprintf("%s (Current: %d): ", prompt, current))
	if input == "" {
		return current
	}
	value, err := strconv.ParseUint(input, 10, 32)
	if err != nil {
		fmt.Println("Invalid number. Keeping the current value.")
		return current
	}
	return uint(value)
}

func GetOptionalFloatInput(prompt string, current float64) float64 {
	input := GetStringInput(fmt.Sprintf("%s (Current: %.2f): ", prompt, current))
	if input == "" {
		return current
	}
	value, err := strconv.ParseFloat(input, 64)
	if err != nil {
		fmt.Println("Invalid number. Keeping the current value.")
		return current
	}
	return value
}

func DereferenceString(s *string) string {
	if s == nil {
		return ""
	}
	return *s
}
