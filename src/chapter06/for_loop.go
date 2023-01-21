package main

import "fmt"

func main() {
	// For loop
	for i := 0; i < 5; i++ {
		fmt.Println(i)
	}
	/*
		prints:
			0
			1
			2
			3
			4

	*/

	// loop through an array
	var names = [4]string{"Jane", "Will", "Jean", "Julian"}

	for index, name := range names {
		fmt.Println(index, name)
	}
	/*
		prints
			0 Jane
			1 Will
			2 Jean
			3 Julian
	*/

	for _, name := range names {
		fmt.Println(name)
	}
	/*
		prints
			Jane
			Will
			Jean
			Julian
	*/

	for index, _ := range names {
		fmt.Println(index)
	}
	/*
		prints
			0
			1
			2
			3
	*/
}
