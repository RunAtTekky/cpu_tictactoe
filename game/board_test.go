package game

import "testing"

func TestGameOver(t *testing.T) {
	t.Run("All empty", func(t *testing.T) {
		board := Board{
			{'$', '$', '$'},
			{'$', '$', '$'},
			{'$', '$', '$'},
		}

		got := Is_game_over(&board)
		want := false

		if got != want {
			t.Errorf("Expected %t but got %t", want, got)
		}
	})

	t.Run("Vertical ALL X, game over", func(t *testing.T) {
		board := Board{
			{'$', 'X', '$'},
			{'$', 'X', '$'},
			{'$', 'X', '$'},
		}

		got := Is_game_over(&board)
		want := true

		if got != want {
			t.Errorf("Expected %t but got %t", want, got)
		}
	})

	t.Run("Horizontal ALL O, game over", func(t *testing.T) {
		board := Board{
			{'$', '$', '$'},
			{'O', 'O', 'O'},
			{'$', '$', '$'},
		}

		got := Is_game_over(&board)
		want := true

		if got != want {
			t.Errorf("Expected %t but got %t", want, got)
		}
	})

	t.Run("Diagonal ALL O, game over", func(t *testing.T) {
		board := Board{
			{'$', '$', 'O'},
			{'$', 'O', '$'},
			{'O', '$', '$'},
		}

		got := Is_game_over(&board)
		want := true

		if got != want {
			t.Errorf("Expected %t but got %t", want, got)
		}
	})
}

func TestValid(t *testing.T) {
	board := Board{
		{'$', '$', 'O'},
		{'$', 'O', '$'},
		{'O', '$', '$'},
	}
	t.Run("Row 1 Col 1 is valid", func(t *testing.T) {
		got := Check_valid(1, 1, board)
		want := true

		if got != want {
			t.Errorf("Wanted %t but got %t", want, got)
		}
	})

	t.Run("Row is negative", func(t *testing.T) {
		got := Check_valid(-2, 1, board)
		want := false

		if got != want {
			t.Errorf("Wanted %t but got %t", want, got)
		}
	})
}
