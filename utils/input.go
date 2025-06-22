package utils

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

var scanner = bufio.NewScanner(os.Stdin)

func GetInput(prompt string) string {
	fmt.Print(prompt)
	scanner.Scan()
	return strings.TrimSpace(scanner.Text())
}

func GetInputLower(prompt string) string {
	return strings.ToLower(GetInput(prompt))
}

func WaitForEnter(message string) {
	fmt.Print(message)
	scanner.Scan()
}

func ClearScreen() {
	fmt.Print("\033[H\033[2J")
}
