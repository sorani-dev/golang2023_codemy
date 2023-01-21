package main

import (
	"fmt"
	"math"
)

func main() {
	/*
		OPERATORS:
		Math
		Assignment = += -= *= /=
		Comparaison == != > < >= <=
	*/

	// Mathematical operators
	fmt.Println("Math Operators")
	fmt.Println("10 + 5 =", 10+5) // 10 + 5 = 15
	fmt.Println("10 - 5 =", 10-5) // 10 - 5 = 5
	fmt.Println("10 * 5 =", 10*5) // 10 * 5 = 50
	fmt.Println("10 / 5 =", 10/5) // 10 / 5 = 2
	fmt.Println("10 % 5 =", 10%5) // 10 % 5 = 0

	var base, exponent float64
	base = 2
	exponent = 3
	fmt.Println("2 to the third power =", math.Pow(base, exponent)) // 2 to the third power = 8

	// Increment/Decrement operators (plus one/minus one) ++, --
	num1 := 10
	num1++
	fmt.Println("10 + 1 =", num1) // 10 + 1 = 11

	num2 := 10
	num2--
	fmt.Println("10 - 1 =", num2) // 10 - 1 = 9

	// Assignment operators =, +=, -=, *=, /=
	myNum := 10
	myNum += 5
	fmt.Println("10 + 5 =", myNum) // 10 + 5 = 15

	// Comparaison Operators ==, !=, >, <, >=, <=
	fmt.Println("5 == 5", 5 == 5) // 5 == 5 true
	fmt.Println("5 != 5", 5 != 5) // 5 != 5 false
	fmt.Println("5 != 6", 5 != 6) // 5 != 5 false
	fmt.Println("5 > 6", 5 > 6)   // 5 > 6 false
	fmt.Println("5 < 6", 5 <= 6)  // 5 < 6 true
	fmt.Println("5 <= 6", 5 <= 6) // 5 <= 6 true
	fmt.Println("5 <= 5", 5 <= 5) // 5 <= 5 true
	fmt.Println("5 >= 6", 5 >= 6) // 5 >= 6 false
	fmt.Println("5 >= 5", 5 >= 5) // 5 >= 5 true
}
