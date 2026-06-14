package models

type GameResponse struct {
	Row      int    `json:"row"`
	Col      int    `json:"col"`
	GameOver bool   `json:"game_over"`
	HasWon   bool   `json:"has_won"`
	Success  bool   `json:"success"`
	ErrorMsg string `json:"errMsg"`
}
