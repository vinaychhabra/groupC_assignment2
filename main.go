package main

import (
	"fmt"
	"strconv"
)

// Armstrong Number Function
func armstrongNumber(value int) bool {

	//Declaring variables
	var sum int = 0
	var save int
	var rem int
	var cube int

	//Save the inputNumber
	save = value

	//Multiplying each digit of the inputInt 3 times
	//And Adding each digit cube of inputInt to variable sum
	for i := value; i > 0; i = i / 10 {
		rem = i % 10
		cube = rem * rem * rem
		sum = sum + cube
	}

	// //Check the condition if the sum is equal to the save number
	if sum == save {
		fmt.Printf("%d,%s", save, "Is a armstrong number")
		return true

	}
	fmt.Println("NOT an armstrong number")
	return false

}

// ageinmonths function
func ageInMonths(years int) int {
	return years * 12
}

// Function to generate Fibonacci series up to n terms
func createfibonacci(n int) []int { // By Jevica

	if n <= 0 {
		fmt.Println("Inavlid Input")
		main()
	}
	fibSeries := make([]int, n)
	fibSeries[0], fibSeries[1] = 0, 1

	for i := 2; i < n; i++ {
		fibSeries[i] = fibSeries[i-1] + fibSeries[i-2]
	}

	return fibSeries

}

// Main Function
func main() {

	// Fibonacci function call -500218849
	fmt.Println("Kindly enter the input value to create fibonacci series")
	var fibno string
	fmt.Scanln(&fibno)
	// Convert the input to an integer
	number, err := strconv.Atoi(fibno)

	fmt.Print(fibno)
	if err != nil || number <= 0 {
		fmt.Println("Invalid input. Please enter a positive integer.")
		return
	}

	fibSeries := createfibonacci(number)
	fmt.Printf("Fibonacci Series up to %d terms: %v\n", number, fibSeries)
	//Armstrong Main Function
	fmt.Println("Kindly enter the input value to check if number is armstrong number")
	var isArmstrongNumber string
	fmt.Scanln(&isArmstrongNumber)
	inputInt, _ := strconv.Atoi(isArmstrongNumber)
	armstrongNumber(inputInt)
	//age in months
	var age int
	fmt.Println("Enter age in years")
	fmt.Scan(&age)

	fmt.Println(ageInMonths(age))

}
