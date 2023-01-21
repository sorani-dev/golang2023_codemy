package main

import "fmt"

func main() {
	/*
		ARRAY
		Declare with explicit length:
			var varName = [size] datatype {value1, value2}
		Declare with inferred length:
			var varName = [] datatype {value1, value2}
		Get a value by index (zero-indexed):
			varName[index]

	*/
	// Declare array with explicit length
	var firstNames = [3]string{"Marlene", "Alice", "Laurence"}

	fmt.Println(firstNames[0]) // prints Marlene

	// Declare array and infer its length
	var numbers = []int{1, 2, 3, 4, 5, 6}
	fmt.Println(numbers, numbers[0]) // prints [1 2 3 4 5 6] 1

	// Shorthand for declaring array
	lastNames := [2]string{"Watson", "Holmes"}
	fmt.Println(lastNames, lastNames[0], lastNames[1]) // prints [Watson Holmes] Watson Holmes

	// Shorthand with infered length
	ourNames := []string{"John", "Sherlock"}
	fmt.Println(ourNames, ourNames[0], ourNames[1]) // prints [John Sherlock] John Sherlock

	// Change item in array
	ourNames[1] = "Mycroft"
	fmt.Println(ourNames) // prints [John Mycroft]

	// Default assignment
	var ourNumbers = [5]int{1, 2}
	fmt.Println(ourNumbers) // prints [1 2 0 0 0]

	// Assign values at certain positions
	var moreNumbers = [10]int{0: 42, 4: 99}
	fmt.Println(moreNumbers) // prints [42 0 0 0 99 0 0 0 0 0]

	// Slices
	var toppings = [5]string{"Mushrooms", "Tomato", "Magarita", "Cheese", "Supreme"}
	fmt.Println(toppings) // prints [Mushrooms Tomato Magarita Cheese Supreme]

	toppingsSlice := toppings[0:2]
	fmt.Println(toppingsSlice) // prints [Mushrooms Tomato]

	// Modify a slice
	toppingsSlice[1] = "Mozarella"
	fmt.Println(toppingsSlice) // prints [Mushrooms Mozarella]

	// Add item to slice
	toppingsSlice = append(toppingsSlice, "Apple")
	fmt.Println(toppingsSlice) // prints [Mushrooms Mozarella Apple]

	// Add two slices together
	otherSlice := toppings[3:4]
	fmt.Println(otherSlice) // prints [Cheese]

	otherSlice = append(otherSlice, toppingsSlice...)
	fmt.Println(otherSlice) // prints [Cheese Mushrooms Mozarella Apple]

}
