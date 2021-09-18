package main

import (
	"fmt"

	game "github.com/JGalvez1776/go-demo/src/game"
	exportexample "github.com/JGalvez1776/go-demo/src/syntax/export-example"
	math "github.com/JGalvez1776/go-demo/src/utils"
)

func main() {
	// Note semicolons are not part of the source code
	// they are added during compilation to help with grammar
	fmt.Println("This is working")
	fmt.Println(exportexample.Public)
	fmt.Println(math.Factorial((4)))
	exportexample.Export()
	game.Start()
}
