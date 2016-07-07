package main

import (
    "fmt"
    "github.com/sniemela/Projects/util"
)

func main() {
    number, err := util.ReadNumber("Enter a number: ")
    if err != nil {
        fmt.Print("Could not read a number.")
        return
    }

    fmt.Printf("Fibonacci %d = %d", number, fib(number))
}

func fib(n int) int {
    if n < 2 {
        return n
    }
    return fib(n - 1) + fib(n - 2)
}