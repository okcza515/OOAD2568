package util

import (
	"bufio"
	"fmt"
	"os"
)

func WaitForEnter() {
    fmt.Println("\nPress Enter to continue...")
    bufio.NewReader(os.Stdin).ReadBytes('\n')
}
