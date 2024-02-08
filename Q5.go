package main

import (
	"fmt"
)

func main() {
	var numSets int

	fmt.Println("Enter the number of sets of data:")
	fmt.Scanln(&numSets)

	principalAmounts := make([]float64, numSets)
	annualInterestRates := make([]float64, numSets)
	years := make([]float64, numSets)

	// Input data for each set
	for i := 0; i < numSets; i++ {
		fmt.Printf("\nSet %d:\n", i+1)

		fmt.Print("Enter the principal amount: $")
		fmt.Scanln(&principalAmounts[i])

		fmt.Print("Enter the annual interest rate (%): ")
		fmt.Scanln(&annualInterestRates[i])

		fmt.Print("Enter the number of years: ")
		fmt.Scanln(&years[i])
	}

	// Calculate and display simple interest for each set
	fmt.Println("\nResults:")
	for i := 0; i < numSets; i++ {
		simpleInterest := (principalAmounts[i] * annualInterestRates[i] * years[i]) / 100
		fmt.Printf("Set %d - Simple Interest: $%.2f\n", i+1, simpleInterest)
	}
}
