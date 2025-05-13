package internal

import (
	"math"
)

var EMPTY rune = '$'
var game_over bool = false
var board [3][3]rune
var x_turn bool = true

type Move struct {
	row int
	col int
}

type BestMove struct {
	move  Move
	score int
}

func Minimax(board [3][3]rune, is_x_turn bool, depth int) BestMove {
	if depth == 8 {
		return BestMove{
			score: 0,
			move:  Move{},
		}
	}
	if check_winner(&board) {
		return BestMove{
			score: score(&board, depth),
			move:  Move{},
		}
	}
	depth++
	var moves []Move
	var scores []int

	// Go through all the possible moves
	available_moves := get_available_moves(&board)

	// fmt.Println(available_moves)

	for _, move := range available_moves {
		possible_game := get_new_state(&board, move, is_x_turn)
		scores = append(scores, Minimax(possible_game, !is_x_turn, depth).score)
		moves = append(moves, move)
	}

	if is_x_turn {
		best_move := BestMove{
			score: math.MinInt,
			move:  Move{},
		}
		for i, val := range scores {
			if val > best_move.score {
				best_move.score = val
				best_move.move.row = moves[i].row
				best_move.move.col = moves[i].col
			}
		}
		if best_move.score == math.MinInt {
			return BestMove{
				score: 0,
				move:  Move{},
			}
		}

		return best_move
	} else {
		best_move := BestMove{
			score: math.MaxInt,
			move:  Move{},
		}
		for i, val := range scores {
			if val < best_move.score {
				best_move.score = val
				best_move.move.row = moves[i].row
				best_move.move.col = moves[i].col
			}
		}
		if best_move.score == math.MaxInt {
			return BestMove{
				score: 0,
				move:  Move{},
			}
		}
		return best_move
	}
}

func get_new_state(board *[3][3]rune, move Move, is_x_turn bool) [3][3]rune {
	var new_state [3][3]rune

	for i, val := range board {
		new_state[i] = val
	}

	if is_x_turn {
		new_state[move.row][move.col] = 'X'
	} else {
		new_state[move.row][move.col] = 'O'
	}

	return new_state
}

func get_available_moves(board *[3][3]rune) []Move {
	var moves []Move

	for i := 0; i < 3; i++ {
		for j := 0; j < 3; j++ {
			if board[i][j] == EMPTY {
				var new_move Move = Move{
					row: i,
					col: j,
				}
				moves = append(moves, new_move)
			}
		}
	}

	return moves
}

func score(board *[3][3]rune, depth int) int {
	// Horizontal
	for i := 0; i < 3; i++ {
		if board[i][0] != EMPTY && board[i][0] == board[i][1] && board[i][0] == board[i][2] {
			if board[i][0] == 'X' {
				return 10 - depth
			} else {
				return depth - 10
			}
		}
	}

	// Vertical
	for i := 0; i < 3; i++ {
		if board[0][i] != EMPTY && board[0][i] == board[1][i] && board[0][i] == board[2][i] {
			if board[0][i] == 'X' {
				return 10 - depth
			} else {
				return depth - 10
			}
		}
	}

	// Diagonal
	if board[1][1] != EMPTY && board[1][1] == board[0][0] && board[1][1] == board[2][2] {
		if board[1][1] == 'X' {
			return 10 - depth
		} else {
			return depth - 10
		}
	}
	if board[1][1] != EMPTY && board[1][1] == board[0][2] && board[1][1] == board[2][0] {
		if board[1][1] == 'X' {
			return 10 - depth
		} else {
			return depth - 10
		}
	}

	// It must be draw
	return 0
}

func check_winner(board *[3][3]rune) bool {
	// Horizontal
	for i := 0; i < 3; i++ {
		if board[i][0] != EMPTY && board[i][0] == board[i][1] && board[i][0] == board[i][2] {
			return true
		}
	}

	// Vertical
	for i := 0; i < 3; i++ {
		if board[0][i] != EMPTY && board[0][i] == board[1][i] && board[0][i] == board[2][i] {
			return true
		}
	}

	// Diagonal
	if board[1][1] != EMPTY && board[1][1] == board[0][0] && board[1][1] == board[2][2] {
		return true
	}
	if board[1][1] != EMPTY && board[1][1] == board[0][2] && board[1][1] == board[2][0] {
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
	if board[row-1][col-1] != EMPTY {
		return false
	}

	return true
}

func is_game_over(board *[3][3]rune) bool {
	if check_winner(board) {
		return true
	}

	for i := 0; i < 3; i++ {
		for j := 0; j < 3; j++ {
			if board[i][j] == EMPTY {
				return false
			}
		}
	}

	return true
}
