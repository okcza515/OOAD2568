// MEP-1003 Student Recruitment
package cli

import (
	"ModEd/recruit/util"
	"bufio"
	"fmt"
	"os"
)

func ApplicantRegistrationCLI(
	ApplicantRegistrationService ApplicantRegistrationService,
) {
	scanner := bufio.NewScanner(os.Stdin)
	util.ClearScreen()

	fmt.Println("==== Applicant Registration ====")
	fmt.Println("1. Register manually")
	fmt.Println("2. Register using CSV/JSON file")
	fmt.Print("Select option: ")
	scanner.Scan()
	choice := scanner.Text()

	switch choice {
	case "1":
		ApplicantRegistrationService.RegisterManually(scanner)
	case "2":
		ApplicantRegistrationService.RegisterFromFile(scanner)
	default:
		fmt.Println("Invalid option.")
	}

}
