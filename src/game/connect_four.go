package game

import (
	"errors"
	"fmt"
	"strconv"
	"strings"
)

type Player string

// Syntax for a constant (Read only)
const HEIGHT = 6
const WIDTH = 7

type Game struct {
	// Note for structs we have to mention the array size as otherwise that is a slice
	board         [HEIGHT][WIDTH]Player
	players       [2]Player
	currentPlayer int
}

func (game Game) PrintBoard() {
	for i := 0; i < HEIGHT; i++ {
		fmt.Println("-" + strings.Repeat("--", WIDTH))
		for index := 0; index < WIDTH; index++ {
			fmt.Print("|" + game.board[i][index])
		}
		fmt.Println("|")
	}
	fmt.Println("-" + strings.Repeat("--", WIDTH))

}

// Syntax for a method via a pointer (This allows you to modify attributes)
func (game *Game) Place(column int) error {
	if column < 0 || column >= WIDTH {
		return errors.New("ColumnOutOfBounds")
	}

	// Alternate way of declaring variables (Can only be done in function / methods)
	rowIndex := 0

	// Go doesn't have while loops :')
	for loop := true; loop; loop = rowIndex < HEIGHT && game.board[rowIndex][column] == " " {
		rowIndex++
	}

	rowIndex--
	if game.board[rowIndex][column] != " " {
		return errors.New("ColumnFull")
	}
	fmt.Println(rowIndex)
	game.board[rowIndex][column] = game.getPlayer()
	game.currentPlayer = (game.currentPlayer + 1) % len(game.players)
	return nil
}

func (game Game) getPlayer() Player {
	return game.players[game.currentPlayer]
}

func initializeGame() (game Game) {
	var player1 Player = "X"
	var player2 Player = "O"

	// Initalize values. Note game is already initialized from the return value
	game.players = [2]Player{player1, player2}
	game.board = [HEIGHT][WIDTH]Player{}
	for y := 0; y < HEIGHT; y++ {
		for x := 0; x < WIDTH; x++ {
			game.board[y][x] = " "
		}
	}
	game.currentPlayer = 0
	return game
}

func Start() {
	var game Game = initializeGame()
	var input string
	var errorCode error
	var column int

	for {
		game.PrintBoard()
		fmt.Println(game.getPlayer() + "'s turn: ")
		fmt.Scanln(&input)
		fmt.Println(input)
		// The _ serves as a variable to save a variable and it shushs the compiler
		// as the Go compiler requires that every assignment to a variable is used
		column, _ = strconv.Atoi(input)
		fmt.Println(column)
		errorCode = game.Place(column)
		if errorCode != nil {
			fmt.Println(errorCode)
		}

		// TODO: Remember to add the game check
	}

}
