package routing

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	models "github.com/RunAtTekky/backend/Models"
	"github.com/RunAtTekky/backend/game"
)

func Hello_handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello World\n")

}

func Place_handler(w http.ResponseWriter, r *http.Request) {
	var req models.Request
	json.NewDecoder(r.Body).Decode(&req)
	print_request(req)

	inserted := models.Board_IN_use.Insert(req.Row, req.Col, req.X_turn)
	if !inserted {
		fmt.Fprintf(w, "FAILED! Already present!\n")
		return
	}

	fmt.Fprintf(w, "Success!\n")

	// Print board after insertion
	models.Board_IN_use.Print_Board()

	board := models.Board_IN_use
	best_move := game.Minimax(board.Board, board.X_turn, board.Depth)
	log.Printf("Best move %v\n", best_move)

	// Make the move from AI side
	models.Board_IN_use.Insert(best_move.Move.Row, best_move.Move.Col, board.X_turn)
	models.Board_IN_use.Print_Board()
}

func print_request(req models.Request) {
	log.Printf(`
You passed this data:
	ROW: %d
	COL: %d
	TURN: %t
`, req.Row, req.Col, req.X_turn)
}

func Has_won_handler(w http.ResponseWriter, r *http.Request) {
	if game.Check_winner(&models.Board_IN_use.Board) {
		fmt.Fprint(w, true)
	} else {
		fmt.Fprint(w, false)
	}
}

func Game_over_handler(w http.ResponseWriter, r *http.Request) {
	if game.Is_game_over(&models.Board_IN_use.Board) {
		fmt.Fprint(w, true)
	} else {
		fmt.Fprint(w, false)
	}
}
