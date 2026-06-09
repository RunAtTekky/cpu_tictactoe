package internal

import (
	models "github.com/RunAtTekky/backend/Models"
	"github.com/RunAtTekky/backend/game"
)

func Insert(row, col int, x_turn bool) bool {
	board := models.Board_IN_use
	if board.Board[row][col] != game.EMPTY {
		return false
	}

	if x_turn {
		board.Board[row][col] = 'X'
	} else {
		board.Board[row][col] = 'O'
	}
	return true
}
