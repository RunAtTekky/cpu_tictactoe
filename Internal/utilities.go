package internal

import (
	"log"

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

func Print_request(req models.Request) {
	log.Printf(`
You passed this data:
	ROW: %d
	COL: %d
	TURN: %t
`, req.Row, req.Col, req.X_turn)
}
