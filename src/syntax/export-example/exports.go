package exportexample

// This variable is exported
var Public = "Hi this string is exported!"

// This variable is not exported as it begins with a lower case letter
var private = "Hi this string is not exported!"

// We can export this function and have it reference a variable that is not exported
func Export() {
	print(private)
}

// Also note, it is a compilation error to declare a variable and never use it
