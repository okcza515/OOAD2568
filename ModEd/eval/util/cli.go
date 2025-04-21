package util

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
	"time"
)

func promptString(prompt string) string {
	reader := bufio.NewReader(os.Stdin)
	fmt.Print(prompt)
	text, _ := reader.ReadString('\n')
	return strings.TrimSpace(text)
}

func promptUint(prompt string) (uint64, error) {
	text := promptString(prompt)
	return strconv.ParseUint(text, 10, 64)
}

func promptDate(prompt string) (time.Time, error) {
	text := promptString(prompt)
	return time.Parse("2006-01-02", text)
}

func promptFloat(prompt string) (float64, error) {
	text := promptString(prompt)
	return strconv.ParseFloat(text, 64)
}