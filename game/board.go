package game

type Board [3][3]rune

var EMPTY rune = '$'
var MINIMAX_CONSTANT int = 10

type Move struct {
	Row int
	Col int
}

type BestMove struct {
	Move  Move
	Score int
}

func Get_available_moves(board *Board) []Move {
	var moves []Move

	for i := range 3 {
		for j := range 3 {
			if board[i][j] == EMPTY {
				var new_move Move = Move{
					Row: i,
					Col: j,
				}
				moves = append(moves, new_move)
			}
		}
	}

	return moves
}

func Score(board *Board, depth int) int {
	// Horizontal
	for i := range 3 {
		if board[i][0] != EMPTY && board[i][0] == board[i][1] && board[i][0] == board[i][2] {
			if board[i][0] == 'X' {
				return MINIMAX_CONSTANT - depth
			} else {
				return depth - MINIMAX_CONSTANT
			}
		}
	}

	// Vertical
	for i := range 3 {
		if board[0][i] != EMPTY && board[0][i] == board[1][i] && board[0][i] == board[2][i] {
			if board[0][i] == 'X' {
				return MINIMAX_CONSTANT - depth
			} else {
				return depth - MINIMAX_CONSTANT
			}
		}
	}

	// Diagonal
	if board[1][1] != EMPTY && board[1][1] == board[0][0] && board[1][1] == board[2][2] {
		if board[1][1] == 'X' {
			return MINIMAX_CONSTANT - depth
		} else {
			return depth - MINIMAX_CONSTANT
		}
	}
	if board[1][1] != EMPTY && board[1][1] == board[0][2] && board[1][1] == board[2][0] {
		if board[1][1] == 'X' {
			return MINIMAX_CONSTANT - depth
		} else {
			return depth - MINIMAX_CONSTANT
		}
	}

	// It must be draw
	return 0
}

func Check_winner(board *Board) bool {
	// Horizontal
	for i := range 3 {
		if board[i][0] != EMPTY && board[i][0] == board[i][1] && board[i][0] == board[i][2] {
			return true
		}
	}

	// Vertical
	for i := range 3 {
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

func Check_valid(row int, col int, board Board) bool {
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

func Is_game_over(board *Board) bool {
	if Check_winner(board) {
		return true
	}

	for i := range 3 {
		for j := range 3 {
			if board[i][j] == EMPTY {
				return false
			}
		}
	}

	return true
}

func Get_new_state(board *Board, move Move, is_x_turn bool) Board {
	var new_state [3][3]rune

	new_state = *board

	if is_x_turn {
		new_state[move.Row][move.Col] = 'X'
	} else {
		new_state[move.Row][move.Col] = 'O'
	}

	return new_state
}
