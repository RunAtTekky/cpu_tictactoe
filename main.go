package main

import (
	"fmt"
	"math"
)

var EMPTY rune = '$'
var game_over bool = false
var board [9]rune
var x_turn bool = true

type Move struct {
	row int
	col int
}

func minimax(board [9]rune, is_x_turn bool, depth int) int {
	if check_winner(&board) {
		return score(&board, depth)
	}
	depth++
	var moves []Move
	var scores []int

	// Go through all the possible moves
	available_moves := get_available_moves(&board)

	fmt.Println(available_moves)

	for _, move := range available_moves {
		possible_game := get_new_state(&board, move, is_x_turn)
		fmt.Println("MOVE ", move)
		print_board(&possible_game)
		scores = append(scores, minimax(possible_game, !is_x_turn, depth))
		fmt.Println(scores[len(scores)-1])
		// minimax(possible_game)
		moves = append(moves, move)
	}

	if is_x_turn {
		maxi := math.MinInt
		for _, val := range scores {
			maxi = max(maxi, val)
		}

		return maxi
	} else {
		mini := math.MaxInt
		for _, val := range scores {
			mini = min(mini, val)
		}

		return mini
	}
}

func get_new_state(board *[9]rune, move Move, is_x_turn bool) [9]rune {
	var new_state [9]rune

	for i, val := range board {
		new_state[i] = val
	}

	row := move.row
	col := move.col
	location := row*3 + col

	if is_x_turn {
		new_state[location] = 'X'
	} else {
		new_state[location] = 'O'
	}

	return new_state
}

func get_available_moves(board *[9]rune) []Move {
	var moves []Move

	for i, char := range board {
		if char == EMPTY {
			var new_move Move
			new_move.row = i / 3
			new_move.col = i % 3

			moves = append(moves, new_move)
		}
	}

	return moves
}

func score(board *[9]rune, depth int) int {
	// Horizontal
	for i := 0; i+2 < 9; i += 3 {
		if board[i] != EMPTY && board[i] == board[i+1] && board[i] == board[i+2] {
			if board[i] == 'X' {
				return 10 - depth
			} else {
				return depth - 10
			}
		}
	}

	// Vertical
	for i := 0; i < 3; i++ {
		if board[i] != EMPTY && board[i] == board[i+3] && board[i] == board[i+6] {
			if board[i] == 'X' {
				return 10 - depth
			} else {
				return depth - 10
			}
		}
	}

	// Diagonal
	if board[4] != EMPTY && board[0] == board[4] && board[4] == board[8] {
		if board[4] == 'X' {
			return 10 - depth
		} else {
			return depth - 10
		}
	}
	if board[4] != EMPTY && board[2] == board[4] && board[4] == board[6] {
		if board[4] == 'X' {
			return 10 - depth
		} else {
			return depth - 10
		}
	}

	// It must be draw
	return 0
}

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
		board[i] = EMPTY
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

		ans := minimax(board, x_turn, 0)

		fmt.Println("Chances of winning: ", ans)

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

// func main() {
// 	var board [9]rune = [9]rune{
// 		// 'O', 'O', '$',
// 		// 'X', 'X', '$',
// 		// 'O', 'X', '$',
// 		// 'O', 'O', '$',
// 		// 'X', 'O', 'X',
// 		// '$', 'X', '$',
// 		'O', 'O', '$',
// 		'X', 'O', 'X',
// 		'$', 'X', '$',
// 	}
//
// 	ans := minimax(board, false, 0)
//
// 	fmt.Println(ans)
//
// }
