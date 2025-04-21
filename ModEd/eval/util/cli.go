package util

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
	"time"
)

func PromptString(prompt string) string {
	reader := bufio.NewReader(os.Stdin)
	fmt.Print(prompt)
	text, _ := reader.ReadString('\n')
	return strings.TrimSpace(text)
}

func PromptUint(prompt string) (uint64, error) {
	text := PromptString(prompt)
	return strconv.ParseUint(text, 10, 64)
}

func PromptDate(prompt string) (time.Time, error) {
	text := PromptString(prompt)
	return time.Parse("2006-01-02", text)
}

func PromptFloat(prompt string) (float64, error) {
	text := PromptString(prompt)
	return strconv.ParseFloat(text, 64)
}