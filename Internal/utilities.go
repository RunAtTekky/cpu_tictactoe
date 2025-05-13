package internal

import models "github.com/RunAtTekky/backend/Models"

func Insert(row, col int, x_turn bool) bool {
	board := models.Board_IN_use
	if board.Board[row][col] != EMPTY {
		return false
	}

	if x_turn {
		board.Board[row][col] = 'X'
	} else {
		board.Board[row][col] = 'O'
	}
	return true
}
