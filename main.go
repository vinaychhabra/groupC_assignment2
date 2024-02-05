package main

import (
	"fmt"
	"math"
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

// Function to check if the number is a perfect square
func isPerfectSquare(num int) bool { //ByAshbir 500228410
	if num < 0 {
		return false
	}

	// Calculate the square root of the number
	sqrt := int(math.Sqrt(float64(num)))

	// Check if the square of the square root is equal to the original number
	return sqrt*sqrt == num
}

//Function isLeapYear()

func isLeapYear(year int) bool {

	if year%4 == 0 {

		if year%100 != 0 || (year%100 == 0 && year%400 == 0) {

			return true
		}

	}
	return false
}

// this function will take int as integer and return factorial of that number
func factorial(number int) int {
	result := 1

	for i := 2; i <= number; i++ {
		result = result * i
	}
	return result
}

// this function is to check number is positive, negative or zero
func checkNumber(num int) string {
	switch {
	case num > 0:
		return "Positive"
	case num < 0:
		return "Negative"
	default:
		return "Zero"
	}
}

// this function prints if a number is prime or not
func isPrime(num int) {
	if num < 2 {
		fmt.Printf("%d is not a prime number.\n", num)
		return
	}
	maxDivisor := int(math.Sqrt(float64(num)))
	for i := 2; i <= maxDivisor; i++ {
		if num%i == 0 {
			fmt.Printf("%d is not a prime number.\n", num)
			return
		}
	}
	fmt.Printf("%d is a prime number.\n", num)
}

// Function to check if a number is even
func isEven(num int) bool {
	return num%2 == 0
}

func temperatureConverter(temperatureInput string) {

	// Convert input to float
	temperatureCelsius, err := strconv.ParseFloat(temperatureInput[:len(temperatureInput)-1], 64)
	//This line converts the input string to a floating-point number using strconv.ParseFloat
	// It removes the newline character at the end using [:len(temperatureInput)-1]
	if err != nil {
		fmt.Println("Invalid input. Please enter a valid temperature.")
		return
	}

	// Convert Celsius to Fahrenheit using the formula
	temperatureFahrenheit := (9 * temperatureCelsius / 5) + 32

	// Print the result with 2 decimal places
	fmt.Printf("Temperature in Fahrenheit: %.2f\n", temperatureFahrenheit)
}

// ++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++//

// Main Function
func main() {

	// createfibonacci() function call -500218849
	fmt.Println("Kindly enter the input value to create fibonacci series")
	var fibno string
	fmt.Scanln(&fibno)

	number, err := strconv.Atoi(fibno) // Convert the input to an integer

	if err != nil || number <= 0 {
		fmt.Println("Invalid input. Please enter a positive integer.")
		return
	}

	fibSeries := createfibonacci(number)
	fmt.Printf("Fibonacci Series up to %d terms: %v\n", number, fibSeries) // fmt.Print(fibno)

	fmt.Println("------------------------------------------------------------------------------------------------")

	// isPerfectSquare() function call - 500228410
	var perfectnumber int
	fmt.Println("Enter a number to check a perfect square: ")

	fmt.Scan(&perfectnumber)                 // Get input from user
	result := isPerfectSquare(perfectnumber) // Check if the input is a perfect square

	fmt.Printf("%d is a perfect square: %t\n", perfectnumber, result) // Print the result

	fmt.Println("------------------------------------------------------------------------------------------------")

	// isArmstrongNumber() function call
	var isArmstrongNumber int
	fmt.Println("Kindly enter the input value to check if number is armstrong number")
	fmt.Scan(&isArmstrongNumber)
	armstrongNumber(isArmstrongNumber)

	fmt.Println("------------------------------------------------------------------------------------------------")

	// ageInMonths() function call
	var age int
	fmt.Println("Enter age in years to convert it to months")
	fmt.Scan(&age)

	fmt.Println(ageInMonths(age))
	fmt.Println("------------------------------------------------------------------------------------------------")

	// isLeapYear() function call
	fmt.Println("Kindly enter the input value to check year is leap year") //user prompt
	var isLeapYearNumber string
	fmt.Scanln(&isLeapYearNumber) // user input

	isLeapYearNumberInt, _ := strconv.Atoi(isLeapYearNumber) // convert string to int
	if isLeapYear(isLeapYearNumberInt) {
		fmt.Println("It is a LEAP year")
	} else {
		fmt.Println("NOT a leap year")
	}

	fmt.Println("------------------------------------------------------------------------------------------------")

	// factorial() function call
	fmt.Println(("Enter a number to find factorial: "))
	var facto_number int
	fmt.Scan(&facto_number)
	fmt.Println(factorial(facto_number))

	fmt.Println("------------------------------------------------------------------------------------------------")

	//checkNumber() function call
	fmt.Println(("Enter a number to check if it's positive or negative: "))
	var number_check int
	fmt.Scan(&number_check)
	fmt.Println(checkNumber(number_check))

	fmt.Println("------------------------------------------------------------------------------------------------")

	// checkPrimeNumber() function call
	fmt.Println("Enter a number to check if number is prime: ")
	var checkPrimeNumber int
	fmt.Scan(&checkPrimeNumber)
	isPrime(checkPrimeNumber)

	fmt.Println("------------------------------------------------------------------------------------------------")

	//isEven()  function call
	fmt.Print("Enter a number: ")
	var num int
	fmt.Scan(&num)

	if isEven(num) {
		fmt.Printf("%d is an even number.\n", num)
	} else {
		fmt.Printf("%d is an odd number.\n", num)
	}

	fmt.Println("------------------------------------------------------------------------------------------------")

	// temperatureConverter() function call
	fmt.Println("Enter temperature in Celsius: ")
	var temperature string
	fmt.Scanln(&temperature)
	temperatureConverter(temperature)
}
