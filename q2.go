package main

import "fmt"

type Student struct {
	ID    int
	Name  string
	Age   int
	Grade int
}

func main() {
	var student Student

	// Print the user to details:
	fmt.Println("Enter Student ID:")
	fmt.Scan(&student.ID)

	fmt.Println("Enter Student Name:")
	fmt.Scan(&student.Name)

	fmt.Println("Enter Student Age:")
	fmt.Scan(&student.Age)

	fmt.Println("Enter Student Grade:")
	fmt.Scan(&student.Grade)

	// Print out the details of the student
	fmt.Println("Student ID:", student.ID)
	fmt.Println("Name:", student.Name)
	fmt.Println("Age:", student.Age)
	fmt.Println("Grade:", student.Grade)
}
