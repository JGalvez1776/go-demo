package slice

import "fmt"

// Slices are similar to Python lists
// (Fun fact Python does not let you use arrays)
func Slice() {
	array := [7]int{1, 2, 3, 4, 5, 6, 7}
	fmt.Printf("Here we have an array of size 7: %v\n", array)
	fmt.Println("Remember arrays have a fixed size")
	slice := array[3:5]
	fmt.Printf("We can then take a \"slice\" from the array: %v", slice)
	fmt.Println("This represents a portion of the original array so if we")
	fmt.Printf("Change a value of the slice it changes the array")
	slice[0] = 99
	fmt.Println("Let's say slice[0] = 99")
	fmt.Println("Now we have:")
	fmt.Printf("Slice: %v\n", slice)
	fmt.Printf("Array: %v\n", array)
	fmt.Println("We can also modify the size of a slice")
	slice = slice[:4]
	fmt.Printf("slice[:4] = %v", slice)

}
