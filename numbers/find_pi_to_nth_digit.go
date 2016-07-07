package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
)

const decimalLimit = 1000000

func main() {
	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Enter a number: ")
	text, err := reader.ReadString('\n')
	if err != nil {
		fmt.Print("Could not read input.")
		return
	}

	number, err := strconv.Atoi(strings.TrimRight(text, "\r\n"))
	if err != nil {
		fmt.Print("Could not convert text to integer. Error: ", err)
		return
	}

	pi := calcPi()

	textToPrint := "Calculated PI = %." + strconv.Itoa(number) + "f"
	fmt.Printf(textToPrint, pi)
}

// Find PI to the Nth digit. Based on John Machin's formula.
func calcPi() float64 {
	return 16*math.Atan(1/5.0) - 4*math.Atan(1/239.0)
}
