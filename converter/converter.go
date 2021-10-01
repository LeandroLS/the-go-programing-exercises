package main

import (
	"fmt"
	"os"
	"strconv"
	"the-go-programing-exercises/section-2-6/tempconv"
)

func main() {
	args := os.Args[1:]

	if len(args) <= 0 {
		fmt.Println("Arguments not found")
		os.Exit(1)
	}
	arg := args[0]
	int, err := strconv.Atoi(arg)
	if err != nil {
		fmt.Println("Please insert an integer, not string")
		os.Exit(1)
	}
	intToCelcius := tempconv.IToC(int)
	celciusToFahrenheit := tempconv.CToF(intToCelcius)

	intToFahrenheit := tempconv.IToF(int)
	fahrenheitToCelcius := tempconv.FToC(intToFahrenheit)
	fmt.Printf("%v is %v\n", intToCelcius, celciusToFahrenheit)
	fmt.Printf("%v is %v\n", intToFahrenheit, fahrenheitToCelcius)
}
