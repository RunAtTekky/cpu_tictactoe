package models

type MoveResponse struct {
	Row      int    `json:"row"`
	Col      int    `json:"col"`
	Success  bool   `json:"success"`
	ErrorMsg string `json:"errMsg"`
}

type GameStateResponse struct {
	GameOver bool `json:"game_over"`
	HasWon   bool `json:"has_won"`
	Success  bool `json:"success"`
}
