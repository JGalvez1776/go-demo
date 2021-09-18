package game

import (
	"errors"
	"fmt"
	"strings"
)

type Player byte

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
	fmt.Println()

	for i := 0; i < HEIGHT; i++ {
		fmt.Println("-" + strings.Repeat("-", WIDTH))
		for index := 0; index < WIDTH; index++ {
			fmt.Print("|%v", game.board[i][index])
		}
		fmt.Println("|")
	}
	fmt.Println("-" + strings.Repeat("-", WIDTH))

}

// Syntax for a method via a pointer (This allows you to modify attributes)
func (game *Game) Place(column int) error {
	if column < 0 || column >= WIDTH {
		return errors.New("ColumnOutOfBounds")
	}

	// Alternate way of declaring variables (Can only be done in function / methods)
	rowIndex := HEIGHT - 1

	// Go doesn't have while loops :')
	for loop := true; loop; loop = rowIndex >= 0 && game.board[rowIndex][column] == ' ' {
		rowIndex++
	}

	if rowIndex == -1 {
		return errors.New("ColumnFull")
	}

	game.board[rowIndex][column] = game.getPlayer()
	game.currentPlayer = (game.currentPlayer + 1) % len(game.players)
	return nil
}

func (game Game) getPlayer() Player {
	return game.players[game.currentPlayer]
}

func initializeGame() (game Game) {
	var player1 Player = 'X'
	var player2 Player = 'O'

	// Initalize values. Note game is already initialized from the return value
	game.players = [2]Player{player1, player2}
	game.board = [HEIGHT][WIDTH]Player{}
	game.currentPlayer = 0
	return game
}

func Start() {
	var game Game = initializeGame()
	game.PrintBoard()
}
