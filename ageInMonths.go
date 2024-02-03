package main

import (
	"fmt"
)

func main() {
	var age int
	fmt.Println("Enter age in years")
	fmt.Scan(&age)

	ageInMonths(age)

}

func ageInMonths(years int) int {
	return years * 12
}
