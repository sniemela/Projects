package main

import (
	"fmt"
	"strconv"

	"github.com/sniemela/Projects/util"
	"math"
)

func main() {
	number, err := util.ReadNumber("Enter a number: ")
	if err != nil {
		fmt.Print("Could not read a number.")
		return
	}

	textToPrint := "Calculated PI = %." + strconv.Itoa(number) + "f"
	fmt.Printf(textToPrint, math.E)
}
