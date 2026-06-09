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
