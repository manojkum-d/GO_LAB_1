package main

import (
	"fmt"
)

func main() {
	var location string
	var temperature float64
	var humidity float64
	var windSpeed float64

	// Taking user inputs
	fmt.Print("Enter location: ")
	fmt.Scanln(&location)

	fmt.Print("Enter temperature (in Celsius): ")
	fmt.Scanln(&temperature)

	fmt.Print("Enter humidity (in percentage): ")
	fmt.Scanln(&humidity)

	fmt.Print("Enter wind speed (in meters per second): ")
	fmt.Scanln(&windSpeed)

	// Printing the collected data
	fmt.Println("\nWeather Data for", location)
	fmt.Println("Temperature:", temperature, "°C")
	fmt.Println("Humidity:", humidity, "%")
	fmt.Println("Wind Speed:", windSpeed, "m/s")
}
