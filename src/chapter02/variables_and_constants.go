package main

import "fmt"

// Import other packages

func main() {
	/* Variables
	var variableName datatype = value
	variableName := value
	*/
	// string, int, float32, bool
	var firstName string = "Simon"

	fmt.Println(firstName) // prints Simon

	// Shorthand variable
	lastName := "Doe"
	age := 42

	fmt.Println(lastName, age) // prints Doe 42

	// Create but don't assign
	var fullName string
	var magicNumber int
	var canVote bool
	var pi float32

	fmt.Println(fullName)    // prints empty string
	fmt.Println(magicNumber) // prints 0
	fmt.Println(canVote)     // prints false
	fmt.Println(pi)          // prints 0

	fullName = "Jene Eyre"
	magicNumber = 42
	canVote = true
	pi = 3.1459
	fmt.Println(fullName)    // prints Jene Eyre
	fmt.Println(magicNumber) // prints 42
	fmt.Println(canVote)     // prints true
	fmt.Println(pi)          // prints 3.1459

	// Multiple variable declaration
	var name1, name2 string = "Makoto", "Eileen"
	fmt.Println(name1, name2) // prints Makoto Eileen

	name1 = "Ritsuko"
	fmt.Println(name1) // prints Ritsuko

	/*
		Conntants
		const CONSTANTNAME datatype = value
		constants cannot be reassigned or changed
	*/
	const PIZZA string = "Mushroom"
	fmt.Println(PIZZA) // prints Mushroom

	// PIZZA = "Margarita" //  cannot assign to PIZZA (constant "Mushroom" of type string)

	// Multiple constants
	const (
		PIZZA1   = "Mushroom"
		PIZZA2   = "Cheese"
		MYNUMBER = 77
	)

	fmt.Println(PIZZA1, PIZZA2, MYNUMBER) // prints Mushroom Cheese 77
}
