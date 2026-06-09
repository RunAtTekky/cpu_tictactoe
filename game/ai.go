package game

import (
	"math"
	"tictactoe/game"
)

func Minimax(board [3][3]rune, is_x_turn bool, depth int) game.BestMove {
	if depth == 8 {
		return game.BestMove{
			Score: 0,
			Move:  game.Move{},
		}
	}
	if game.Check_winner(&board) {
		return game.BestMove{
			Score: game.Score(&board, depth),
			Move:  game.Move{},
		}
	}
	depth++
	var moves []game.Move
	var scores []int

	// Go through all the possible moves
	available_moves := game.Get_available_moves(&board)

	// fmt.Println(available_moves)

	for _, move := range available_moves {
		possible_game := game.Get_new_state(&board, move, is_x_turn)
		scores = append(scores, Minimax(possible_game, !is_x_turn, depth).Score)
		moves = append(moves, move)
	}

	if is_x_turn {
		best_move := game.BestMove{
			Score: math.MinInt,
			Move:  game.Move{},
		}
		for i, val := range scores {
			if val > best_move.Score {
				best_move.Score = val
				best_move.Move.Row = moves[i].Row
				best_move.Move.Col = moves[i].Col
			}
		}
		if best_move.Score == math.MinInt {
			return game.BestMove{
				Score: 0,
				Move:  game.Move{},
			}
		}

		return best_move
	} else {
		best_move := game.BestMove{
			Score: math.MaxInt,
			Move:  game.Move{},
		}
		for i, val := range scores {
			if val < best_move.Score {
				best_move.Score = val
				best_move.Move.Row = moves[i].Row
				best_move.Move.Col = moves[i].Col
			}
		}
		if best_move.Score == math.MaxInt {
			return game.BestMove{
				Score: 0,
				Move:  game.Move{},
			}
		}
		return best_move
	}
}
