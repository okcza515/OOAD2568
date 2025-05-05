package cli

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func AdminDeleteInterviewCLI(interviewService AdminInterviewService) {
	reader := bufio.NewReader(os.Stdin)

	fmt.Print("Enter Interview ID to delete: ")
	input, _ := reader.ReadString('\n')
	input = strings.TrimSpace(input)

	interviewID, err := strconv.Atoi(input)
	if err != nil {
		fmt.Println("Invalid Interview ID.")
		return
	}

	err = interviewService.DeleteInterview(uint(interviewID))
	if err != nil {
		fmt.Println("Failed to delete interview:", err)
	} else {
		fmt.Println("Interview deleted successfully.")
	}
}
