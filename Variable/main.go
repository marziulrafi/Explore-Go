package main

import "fmt"

func main() {
	var firstName string = "Marziul"
	lastName := "Rafi" // Explicitly declared

	fmt.Println(firstName)
	fmt.Println(lastName)

	// Grouped Variable
	var (
		country   string = "Bangladesh"
		city      string = "Dhaka"
		continent string = "Asia"
	)
	fmt.Println(country)
	fmt.Println(city)
	fmt.Println(continent)

	// Multiple Variable
	var a, b, c int = 1, 2, 3
	fmt.Println(a, b, c)

	// Constant
	const pi = 3.141
	// pi = 3.14 // Can't be reassigned
	fmt.Println(pi)
}
