package game

import (
	"errors"
	"fmt"
	"strconv"
	"strings"
)

type player string

// Syntax for a constant (Read only)
const HEIGHT = 6
const WIDTH int = 7
const DRAW_MESSAGE string = "GAME OVER: It is a draw"
const X_MESSAGE = "GAME OVER: X Wins!"
const O_MESSAGE = "GAME OVER: O Wins!"

type Game struct {
	// Note for structs we have to mention the array size as otherwise that is a slice
	board         [HEIGHT][WIDTH]player
	players       [2]player
	currentPlayer int
}

type winner struct {
	message string
	isOver  bool
}

// Syntax for a method of a struct
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

func (game Game) get(x int, y int) player {
	return game.board[y][x]
}

func (game Game) isOver() (result winner) {
	empty_space := false

	// Could do some fancy recursive backtracking but instead brute force

	// Checks if the board is full denoting a potential draw
	for y := 0; y < HEIGHT; y++ {
		for x := 0; x < WIDTH; x++ {
			if game.get(x, y) == " " {
				empty_space = true
				result.message = DRAW_MESSAGE
			}
		}
	}

	// Checks for horizontal wins
	for y := 0; y < HEIGHT; y++ {
		for x := 0; x < WIDTH-3; x++ {
			if game.get(x, y) != " " &&
				game.get(x, y) == game.get(x+1, y) &&
				game.get(x+1, y) == game.get(x+2, y) &&
				game.get(x+2, y) == game.get(x+3, y) {
				result.isOver = true
				if game.get(x, y) == "X" {
					result.message = X_MESSAGE
				} else {
					result.message = O_MESSAGE
				}
				return result
			}
		}
	}

	// Checks for vertical wins
	for y := 0; y < HEIGHT-3; y++ {
		for x := 0; x < WIDTH; x++ {
			if game.get(x, y) != " " &&
				game.get(x, y) == game.get(x, y+1) &&
				game.get(x, y+1) == game.get(x, y+2) &&
				game.get(x, y+2) == game.get(x, y+3) {
				result.isOver = true
				if game.get(x, y) == "X" {
					result.message = X_MESSAGE
				} else {
					result.message = O_MESSAGE
				}
				return result
			}
		}
	}

	// Checks for / diagonal wins
	for y := 3; y < HEIGHT; y++ {
		for x := 0; x < WIDTH-3; x++ {
			if game.get(x, y) != " " &&
				game.get(x, y) == game.get(x+1, y-1) &&
				game.get(x+1, y-1) == game.get(x+2, y-2) &&
				game.get(x+2, y-2) == game.get(x+3, y-3) {
				result.isOver = true
				if game.get(x, y) == "X" {
					result.message = X_MESSAGE
				} else {
					result.message = O_MESSAGE
				}
				return result
			}
		}
	}

	// Checks for \ diagonal wins
	for y := 0; y < HEIGHT-3; y++ {
		for x := 0; x < WIDTH-3; x++ {
			if game.get(x, y) != " " &&
				game.get(x, y) == game.get(x+1, y+1) &&
				game.get(x+1, y+1) == game.get(x+2, y+2) &&
				game.get(x+2, y+2) == game.get(x+3, y+3) {
				result.isOver = true
				if game.get(x, y) == "X" {
					result.message = X_MESSAGE
				} else {
					result.message = O_MESSAGE
				}
				return result
			}
		}
	}

	if !empty_space {
		result.isOver = true
		result.message = DRAW_MESSAGE
	} else {
		result.isOver = false
	}
	return result
}

func (game Game) getPlayer() player {
	return game.players[game.currentPlayer]
}

func initializeGame() (game Game) {
	var player1 player = "X"
	var player2 player = "O"

	// Initalize values. Note game is already initialized from the return value
	game.players = [2]player{player1, player2}
	game.board = [HEIGHT][WIDTH]player{}
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
	var status winner

	for {
		game.PrintBoard()
		fmt.Println(game.getPlayer() + "'s turn: ")
		// Sends a pointer so the value of input can change
		fmt.Scanln(&input)

		// The _ serves as a variable to save a variable and it shushs the compiler
		// as the Go compiler requires that every assignment to a variable is used
		column, _ = strconv.Atoi(input)
		errorCode = game.Place(column)
		if errorCode != nil {
			if errorCode.Error() == "ColumnOutOfBounds" {
				fmt.Printf("Please insert and integer between 0 and %v.\n", WIDTH-1)
			} else if errorCode.Error() == "ColumnFull" {
				fmt.Println("Please select a column that is not full.")
			} else {
				fmt.Println("Unknown Error")
			}
			fmt.Println("Try Again:")
		}
		status = game.isOver()
		if status.isOver {
			game.PrintBoard()
			fmt.Println(status.message)
			fmt.Println("Press ENTER to exit")
			fmt.Scanln(&input)
			return
		}
	}

}
