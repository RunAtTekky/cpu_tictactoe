package models

type Request struct {
	Row    int  `json:"row"`
	Col    int  `json:"col"`
	X_turn bool `json:"x_turn"`
}
