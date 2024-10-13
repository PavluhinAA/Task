package main

import (
	"TaskCalcV2/calculator"
	"fmt"
)

func main() {
	fmt.Println("Input example: XX*X\nIf the result is not an integer, it will be rounded")
	for {
		fmt.Println("Enter a request")
		err := calculator.Input()
		if err != nil {
			fmt.Println(err)
		}
	}
}
