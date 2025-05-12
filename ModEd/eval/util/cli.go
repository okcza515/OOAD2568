package util

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"time"
)

func GetDateTimeInput(prompt string) (time.Time, error) {
	reader := bufio.NewReader(os.Stdin)
	fmt.Print(prompt)
	text, _ := reader.ReadString('\n')
	return time.Parse("2006-01-02 15:04:05", strings.TrimSpace(text))
}