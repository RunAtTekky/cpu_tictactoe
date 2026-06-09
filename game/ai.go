package game

import (
	"math"
)

func Minimax(board [3][3]rune, is_x_turn bool, depth int) BestMove {
	if depth == 8 {
		return BestMove{
			Score: 0,
			Move:  Move{},
		}
	}
	if Check_winner(&board) {
		return BestMove{
			Score: Score(&board, depth),
			Move:  Move{},
		}
	}
	depth++
	var moves []Move
	var scores []int

	// Go through all the possible moves
	available_moves := Get_available_moves(&board)

	// fmt.Println(available_moves)

	for _, move := range available_moves {
		possible_game := Get_new_state(&board, move, is_x_turn)
		scores = append(scores, Minimax(possible_game, !is_x_turn, depth).Score)
		moves = append(moves, move)
	}

	if is_x_turn {
		best_move := BestMove{
			Score: math.MinInt,
			Move:  Move{},
		}
		for i, val := range scores {
			if val > best_move.Score {
				best_move.Score = val
				best_move.Move.Row = moves[i].Row
				best_move.Move.Col = moves[i].Col
			}
		}
		if best_move.Score == math.MinInt {
			return BestMove{
				Score: 0,
				Move:  Move{},
			}
		}

		return best_move
	} else {
		best_move := BestMove{
			Score: math.MaxInt,
			Move:  Move{},
		}
		for i, val := range scores {
			if val < best_move.Score {
				best_move.Score = val
				best_move.Move.Row = moves[i].Row
				best_move.Move.Col = moves[i].Col
			}
		}
		if best_move.Score == math.MaxInt {
			return BestMove{
				Score: 0,
				Move:  Move{},
			}
		}
		return best_move
	}
}
