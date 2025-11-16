// Author: William Provost
// Version: 1.0.0
// Date: 2025-11-15
// Fileoverview: This program shows comparing strings with different cases.

package main

import (
	"fmt"
	"strings"
)

func main() {
	// variable declaration
	var string1 string = "Hello"
	var string2 string = "HELLO"
	var string3 string = "I LOVE CS!"
	var string4 string = "I LOVE CS!"

	// check if string are the same
	// NOTE: using regular "equal" operator
	if string1 == string2 {
		fmt.Println("\"" + string1 + "\" is the same as \"" + string2 + "\".")
	} else if strings.ToLower(string1) == strings.ToLower(string2) {
		fmt.Println("\"" + string1 + "\" is the same as \"" + string2 + "\" (lower case insensitive).")
	} else {
		fmt.Println("\"" + string1 + "\" is NOT the same as \"" + string2 + "\".")
	}

	// check if string are the same
	// NOTE: there is no strict (also checking type) "equal" operator in Go
	if string3 == string4 {
		fmt.Println("\"" + string3 + "\" is the same as \"" + string4 + "\" (upper case insensitive).")
	} else if strings.ToUpper(string3) == strings.ToUpper(string4) {
		fmt.Println("\"" + string3 + "\" is the same as \"" + string4 + "\" (upper case insensitive).")
	} else {
		fmt.Println("\"" + string3 + "\" is NOT the same as \"" + string4 + "\".")
	}

	fmt.Println("\nDone.")
}
