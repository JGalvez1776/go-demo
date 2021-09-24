package main

// Imports are switched between modules already implemented in go like fmt
// and those that are written, which have to provide where Go has to find time
// Note the names game, exportexample, etc are the names of the packages.
import (
	"fmt"

	game "github.com/JGalvez1776/go-demo/src/game"
	exportexample "github.com/JGalvez1776/go-demo/src/syntax/export-example"
	slice "github.com/JGalvez1776/go-demo/src/syntax/slice-example"
)

func main() {
	// Note semicolons are not part of the source code
	// they are added during compilation to help with grammar
	var input string
	for {
		// Takes in user input
		fmt.Print("> ")
		fmt.Scanln(&input)
		fmt.Println()

		if input == "game" {
			game.Start()
		} else if input == "export" {
			message := exportexample.Export()
			fmt.Printf("We know have \"%v\" saved as a variable in main\n", message)
			fmt.Printf("We can also access the public value of \"%v\"", exportexample.Public)
		} else if input == "slice" {
			slice.Slice()
		} else if input == "help" {
			help()
		} else {
			fmt.Println("Please enter a valid command. Enter \"help\" for more information")
		}

	}
}

func help() {
	fmt.Println("Valid commands:")
	fmt.Println("game = A basic connect-four game")
	fmt.Println("export = A demo to go alongside src/syntax/export-example/exports.go that shows how exports work in Go.")
	fmt.Println("slice = A demo to go alongside src/syntax/slice-example/slices.go that shows of the slice type.")

}
