package main

import (
	"fmt"
)

var EMPTY rune = '$'
var game_over bool = false
var board [9]rune
var x_turn bool = true

func check_winner(board *[9]rune) bool {
	// Horizontal
	for i := 0; i+2 < 9; i += 3 {
		if board[i] != EMPTY && board[i] == board[i+1] && board[i] == board[i+2] {
			return true
		}
	}

	// Vertical
	for i := 0; i < 3; i++ {
		if board[i] != EMPTY && board[i] == board[i+3] && board[i] == board[i+6] {
			return true
		}
	}

	// Diagonal
	if board[4] != EMPTY && board[0] == board[4] && board[4] == board[8] {
		return true
	}
	if board[4] != EMPTY && board[2] == board[4] && board[4] == board[6] {
		return true
	}

	return false

}

func check_valid(row int, col int) bool {
	if row < 1 || row > 3 {
		return false
	}
	if col < 1 || col > 3 {
		return false
	}
	location := (row-1)*3 - 1 + col

	if board[location] != EMPTY {
		return false
	}

	return true
}

func print_board(board *[9]rune) {
	for i := 0; i+2 < 9; i += 3 {
		fmt.Printf("%c %c %c\n", board[i], board[i+1], board[i+2])
	}
}

func main() {

	for i := range board {
		board[i] = '$'
	}

	turns := 0

	for !game_over {

		var row, col int
		print_board(&board)

		fmt.Println("Enter row: ")
		fmt.Scan(&row)

		fmt.Println("Enter col: ")
		fmt.Scan(&col)

		if !check_valid(row, col) {
			fmt.Println("Enter in empty cell")
			continue
		}

		location := (row-1)*3 - 1 + col
		if x_turn {
			board[location] = 'X'
		} else {
			board[location] = 'O'
		}

		x_turn = !x_turn

		game_over = check_winner(&board)
		turns++

		if turns == 9 {
			break
		}
	}

	if !check_winner(&board) {
		fmt.Println("GAME ENDED IN DRAW")
	} else if x_turn {
		fmt.Println("Game Over!!!! O WON")
	} else {
		fmt.Println("Game Over!!!! X WON")
	}
}
