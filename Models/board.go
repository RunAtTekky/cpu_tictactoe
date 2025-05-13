package models

type Board struct {
	Board     [3][3]rune `json:"board"`
	X_turn    bool       `json:"x_turn"`
	Game_over bool       `json:"game_over"`
	Depth     int        `json:"depth"`
}

var Board_IN_use Board = Board{
	Board: [3][3]rune{
		{'O', '$', '$'},
		{'$', 'X', 'O'},
		{'$', 'X', '$'},
	},
	X_turn:    true,
	Game_over: false,
	Depth:     0,
}
