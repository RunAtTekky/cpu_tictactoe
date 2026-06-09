package models


import (
	"log"

	"github.com/RunAtTekky/backend/game"
)

type Board struct {
	Board     game.Board `json:"board"`
	X_turn    bool       `json:"x_turn"`
	Game_over bool       `json:"game_over"`
	Depth     int        `json:"depth"`
}

var Board_IN_use Board = Board{
	Board: game.Board{
		{'$', '$', '$'},
		{'$', '$', '$'},
		{'$', '$', '$'},
	},
	X_turn:    true,
	Game_over: false,
	Depth:     0,
}

var EMPTY = '$'

func (board *Board) Insert(row, col int, x_turn bool) bool {
	if board.Board[row][col] != EMPTY {
		return false
	}

	if x_turn {
		board.Board[row][col] = 'X'
	} else {
		board.Board[row][col] = 'O'
	}

	board.X_turn = !board.X_turn

	return true
}

func (board *Board) Print_Board() {
	log.Printf("Board:\n%s\n%s\n%s\n",
		string(board.Board[0][:]),
		string(board.Board[1][:]),
		string(board.Board[2][:]),
	)
}
