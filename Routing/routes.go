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

	fmt.Fprintf(w, "Success!\n")

	board := models.Board_IN_use

	best_move := internal.Minimax(board.Board, board.X_turn, board.Depth)
	fmt.Println(best_move)

}
