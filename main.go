package main

import (
	"fmt"

	game "github.com/JGalvez1776/go-demo/src/game"
	exportexample "github.com/JGalvez1776/go-demo/src/syntax/export-example"
)

func main() {
	// Note semicolons are not part of the source code
	// they are added during compilation to help with grammar
	fmt.Println("Start")
	exportexample.Export()
	game.Start()

}
