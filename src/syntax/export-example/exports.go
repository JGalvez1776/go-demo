package exportexample

import "fmt"

// This variable is exported
var Public = "Hi this string is exported!"

// This variable is not exported as it begins with a lower case letter
var private = "Hi this string is not exported!"

// We can export this function and have it reference a variable that is not exported
func Export() string {
	fmt.Println("Here you are accessing a public method!")
	fmt.Println("Since this method is from the exportexample package")
	fmt.Printf("We can print out the private (Package-viewable):\"%v\"\n", private)
	fmt.Println("This method is also returning this private variable so it")
	fmt.Println("now available in the main file.")
	return private
}

// Also note, it is a compilation error to declare a variable and never use it
