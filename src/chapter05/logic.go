package main

import "fmt"

func main() {

	// If ... statement
	var firstName = "Marina"

	if firstName == "Marina" {
		fmt.Println("Hello", firstName)
	}
	/*
		prints Hello Marina
	*/

	firstName = "Ranma"

	if firstName == "Marina" {
		fmt.Println("Hello", firstName)
	}
	/*
		prints nothing
	*/

	// If ... else ... statement
	firstName = "Michel"

	if firstName == "Marina" {
		fmt.Println("Hello", firstName)
	} else {
		fmt.Println("Sorry,", firstName, "I don't know you!")
	}
	/*
		prints Sorry, Michel I don't know you!
	*/

	// If ... else if ... else ... statement
	firstName = "Michel"

	if firstName == "Marina" {
		fmt.Println("Hello", firstName)
	} else if firstName == "Michel" {
		fmt.Println("What's up", firstName, "!")
	} else {
		fmt.Println("Sorry,", firstName, "I don't know you!")
	}
	/*
		prints What's up Michel !
	*/

	myNum := 42
	if myNum > 30 {
		fmt.Println("42 is greater than 30")
	} else {
		fmt.Println("42 is less than 30")
	}
	/*
		prints 42 is greater than 30
	*/

	myNum = 3
	if myNum > 30 {
		fmt.Println("3 is greater than 30")
	} else {
		fmt.Println("3 is less than 30")
	}
	/*
		prints 3 is less than 30
	*/
}
