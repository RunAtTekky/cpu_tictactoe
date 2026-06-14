package routing

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	internal "github.com/RunAtTekky/backend/Internal"
	models "github.com/RunAtTekky/backend/Models"
	"github.com/RunAtTekky/backend/game"
)

func Hello_handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello World\n")

}

func Place_handler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("content-type", "application/json")
	var req models.Request
	json.NewDecoder(r.Body).Decode(&req)
	internal.Print_request(req)

	inserted := models.Board_IN_use.Insert(req.Row, req.Col, req.X_turn)
	if !inserted {
		res := &models.GameResponse{
			Success:  false,
			ErrorMsg: "FAILED! Already present!",
		}

		json.NewEncoder(w).Encode(res)
		return
	}

	// Print board after insertion
	models.Board_IN_use.Print_Board()

	board := models.Board_IN_use
	best_move := game.Minimax(board.Board, board.X_turn, board.Depth)
	log.Printf("Best move %v\n", best_move)

	// Make the move from AI side
	models.Board_IN_use.Insert(best_move.Move.Row, best_move.Move.Col, board.X_turn)
	models.Board_IN_use.Print_Board()

	res := &models.GameResponse{
		Row:      best_move.Move.Row,
		Col:      best_move.Move.Col,
		GameOver: game.Is_game_over(&models.Board_IN_use.Board),
		Success:  true,
		ErrorMsg: "",
	}

	json.NewEncoder(w).Encode(res)
}

func Has_won_handler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("content-type", "application/json")
	if game.Check_winner(&models.Board_IN_use.Board) {
		res := &models.GameResponse{
			Success: true,
			HasWon:  true,
		}

		json.NewEncoder(w).Encode(res)
	} else {
		res := &models.GameResponse{
			Success: true,
			HasWon:  false,
		}

		json.NewEncoder(w).Encode(res)
	}
}

func Game_over_handler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("content-type", "application/json")
	if game.Is_game_over(&models.Board_IN_use.Board) {
		res := &models.GameResponse{
			Success:  true,
			GameOver: true,
		}
		json.NewEncoder(w).Encode(res)
	} else {
		res := &models.GameResponse{
			Success:  true,
			GameOver: false,
		}
		json.NewEncoder(w).Encode(res)
	}
}
