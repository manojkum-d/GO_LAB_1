package main

import (
	"fmt"
)

func main() {
	fmt.Println("Welcome to the Student Grade Processing System")

	for {
		var numStudents int
		fmt.Print("\nEnter the number of students: ")
		fmt.Scanln(&numStudents)

		// Input validation for number of students
		if numStudents <= 0 {
			fmt.Println("Number of students should be greater than 0. Please try again.")
			continue
		}

		var totalGrade, grade int

		// Input phase
		for i := 1; i <= numStudents; i++ {
			fmt.Printf("Enter grade for student %d: ", i)
			fmt.Scanln(&grade)

			// Data validation
			if grade < 0 || grade > 100 {
				fmt.Println("Invalid grade. Grade should be between 0 and 100. Please try again.")
				i-- // Decrement i to re-enter the grade for this student
				continue
			}

			totalGrade += grade
		}

		// Calculate average grade
		averageGrade := float64(totalGrade) / float64(numStudents)

		// Grade classification
		var letterGrade string
		switch {
		case averageGrade >= 90:
			letterGrade = "A"
		case averageGrade >= 80:
			letterGrade = "B"
		case averageGrade >= 70:
			letterGrade = "C"
		case averageGrade >= 60:
			letterGrade = "D"
		default:
			letterGrade = "F"
		}

		// Display results
		fmt.Printf("Average Grade: %.2f\n", averageGrade)
		fmt.Println("Letter Grade:", letterGrade)

		// Prompt for repeat
		var choice string
		fmt.Print("\nDo you want to process grades for another class? (yes/no): ")
		fmt.Scanln(&choice)
		if choice != "yes" {
			break
		}
	}

	fmt.Println("Thank you for using the Student Grade Processing System")
}
