// Created by: Maryam Nona
// Created on: May 2021
//
// This program does basic math

package main

import (
	"fmt"
)

func main() {
	// This function does addition
	var fahrenheit int

	// input
	fmt.Println("This program converts Fahrenheit to Celsius ")
	fmt.Println()
	fmt.Print("Enter fahrenheit (F°): ")
	fmt.Scanln(&fahrenheit)

	// output
	fmt.Println("C°= ", (fahrenheit-32)*5/9)
	fmt.Println("Done.")
}
