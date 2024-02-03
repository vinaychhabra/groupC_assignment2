package main

import "fmt"

//making function to concate 2 strings
func concatenateStrings(str1, str2 string) string {
	result := str1 + str2
	return result
}

func main() {
	// returning the concated strings
	string1 := "This is 1st string, "
	string2 := "This is 2nd string"

	concatenatedString := concatenateStrings(string1, string2)

	fmt.Println(concatenatedString)
}
