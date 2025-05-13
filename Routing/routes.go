package routing

import (
	"encoding/json"
	"fmt"
	"net/http"

	internal "github.com/RunAtTekky/backend/Internal"
	models "github.com/RunAtTekky/backend/Models"
)

func Hello_handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello World\n")

}

func Place_handler(w http.ResponseWriter, r *http.Request) {
	var req models.Request
	json.NewDecoder(r.Body).Decode(&req)

	fmt.Printf(`You passed this data:
ROW: %d
COL: %d
TURN: %t
`, req.Row, req.Col, req.X_turn)

	models.Board_IN_use.Print_Board()
	if !models.Board_IN_use.Insert(req.Row, req.Col, req.X_turn) {
		fmt.Fprintf(w, "FAILED! Already present!\n")
		return
	}

	fmt.Fprintf(w, "Success!\n")
	models.Board_IN_use.Print_Board()

	board := models.Board_IN_use

	best_move := internal.Minimax(board.Board, board.X_turn, board.Depth)
	fmt.Println(best_move)

	models.Board_IN_use.Insert(best_move.Move.Row, best_move.Move.Col, board.X_turn)
	models.Board_IN_use.Print_Board()
}
