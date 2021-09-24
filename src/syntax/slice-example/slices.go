package slice

import "fmt"

func Slice() {
	array := [7]int{1, 2, 3, 4, 5, 6, 7}
	fmt.Printf("Here we have an array of size 7: %v\n", array)
	fmt.Println("Remember arrays have a fixed size")
	slice := array[3:5]
	fmt.Printf("We can then take a \"slice\" from the array: %v", slice)
}
