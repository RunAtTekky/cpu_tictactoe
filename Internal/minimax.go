package internal

import (
	"math"
)

var EMPTY rune = '$'
var game_over bool = false
var board [3][3]rune
var x_turn bool = true

type Move struct {
	Row int
	Col int
}

type BestMove struct {
	Move  Move
	Score int
}

func Minimax(board [3][3]rune, is_x_turn bool, depth int) BestMove {
	if depth == 8 {
		return BestMove{
			Score: 0,
			Move:  Move{},
		}
	}
	if check_winner(&board) {
		return BestMove{
			Score: Score(&board, depth),
			Move:  Move{},
		}
	}
	depth++
	var Moves []Move
	var Scores []int

	// Go through all the possible Moves
	available_Moves := get_available_Moves(&board)

	// fmt.Println(available_Moves)

	for _, Move := range available_Moves {
		possible_game := get_new_state(&board, Move, is_x_turn)
		Scores = append(Scores, Minimax(possible_game, !is_x_turn, depth).Score)
		Moves = append(Moves, Move)
	}

	if is_x_turn {
		best_Move := BestMove{
			Score: math.MinInt,
			Move:  Move{},
		}
		for i, val := range Scores {
			if val > best_Move.Score {
				best_Move.Score = val
				best_Move.Move.Row = Moves[i].Row
				best_Move.Move.Col = Moves[i].Col
			}
		}
		if best_Move.Score == math.MinInt {
			return BestMove{
				Score: 0,
				Move:  Move{},
			}
		}

		return best_Move
	} else {
		best_Move := BestMove{
			Score: math.MaxInt,
			Move:  Move{},
		}
		for i, val := range Scores {
			if val < best_Move.Score {
				best_Move.Score = val
				best_Move.Move.Row = Moves[i].Row
				best_Move.Move.Col = Moves[i].Col
			}
		}
		if best_Move.Score == math.MaxInt {
			return BestMove{
				Score: 0,
				Move:  Move{},
			}
		}
		return best_Move
	}
}

func get_new_state(board *[3][3]rune, Move Move, is_x_turn bool) [3][3]rune {
	var new_state [3][3]rune

	for i, val := range board {
		new_state[i] = val
	}

	if is_x_turn {
		new_state[Move.Row][Move.Col] = 'X'
	} else {
		new_state[Move.Row][Move.Col] = 'O'
	}

	return new_state
}

func get_available_Moves(board *[3][3]rune) []Move {
	var Moves []Move

	for i := 0; i < 3; i++ {
		for j := 0; j < 3; j++ {
			if board[i][j] == EMPTY {
				var new_Move Move = Move{
					Row: i,
					Col: j,
				}
				Moves = append(Moves, new_Move)
			}
		}
	}

	return Moves
}

func Score(board *[3][3]rune, depth int) int {
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

func check_valid(Row int, Col int) bool {
	if Row < 1 || Row > 3 {
		return false
	}
	if Col < 1 || Col > 3 {
		return false
	}
	if board[Row-1][Col-1] != EMPTY {
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
