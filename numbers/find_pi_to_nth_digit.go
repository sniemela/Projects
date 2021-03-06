package main

import (
	"fmt"
	"math"
	"strconv"

	"github.com/sniemela/Projects/util"
)

func main() {
	number, err := util.ReadNumber("Enter a number: ")
	if err != nil {
		fmt.Print("Could not read a number")
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
