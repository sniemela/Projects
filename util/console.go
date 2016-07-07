package util

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

// ReadNumber reads a number from stdin.
func ReadNumber(message string) (int, error) {
	reader := bufio.NewReader(os.Stdin)
	fmt.Print(message)
	text, err := reader.ReadString('\n')
	if err != nil {
		return 0, err
	}

	number, err := strconv.Atoi(strings.TrimRight(text, "\r\n"))
	if err != nil {
		fmt.Print("Could not convert text to integer. Error: ", err)
		return 0, err
	}

    return number, nil
}
