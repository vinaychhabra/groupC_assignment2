package main

import (
	"fmt"
)

func greet(name string, age int) {
	fmt.Printf("Hello, %s! You are %d years old.\n", name, age)
}

func getName() string {
	var name string
	fmt.Print("Enter your name: ")
	fmt.Scanln(&name)
	return name
}

func getAge(age int) int {
	return age
}

func main() {
	name := getName()

	var ageInput int
	fmt.Print("Enter your age: ")
	fmt.Scanln(&ageInput)

	age := getAge(ageInput)
	greet(name, age)
}
