package main

import (
	"fmt"

	exortexample "github.com/JGalvez1776/go-demo/src/syntax/export-example"
	math "github.com/JGalvez1776/go-demo/src/utils"
)

func main() {
	fmt.Println("This is working")
	fmt.Println(exortexample.Public)
	fmt.Println(math.Factorial((4)))
}
