package main

import (
	"fmt"
)

func main() {

	var mathScore, scienceScore, englishScore float64
	fmt.Println("Please enter your scores for Math, Science, and English:")

	fmt.Print("Math: ")
	fmt.Scanln(&mathScore)

	fmt.Print("Science: ")
	fmt.Scanln(&scienceScore)

	fmt.Print("English: ")
	fmt.Scanln(&englishScore)

	// Calculating average grade
	averageScore := (mathScore + scienceScore + englishScore) / 3

	// Determining letter grade based on average score
	var letterGrade string
	switch {
	case averageScore >= 90:
		letterGrade = "A"
	case averageScore >= 80:
		letterGrade = "B"
	case averageScore >= 70:
		letterGrade = "C"
	case averageScore >= 60:
		letterGrade = "D"
	default:
		letterGrade = "F"
	}

	// Displaying the result
	fmt.Printf("Your average grade is %.2f, which corresponds to the letter grade %s.\n", averageScore, letterGrade)
}
