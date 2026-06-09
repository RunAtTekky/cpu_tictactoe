package main

import (
	"fmt"
	"math"
	"os"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"

	"tictactoe/game"
)

var game_over bool = false
var board [3][3]rune
var x_turn bool = true

type model struct {
	board       [3][3]rune
	x_turn_m    bool
	game_over_m bool
	cursorX     int
	cursorY     int
	depth       int
}

func InitialModel() tea.Model {
	return model{
		board: [3][3]rune{
			{'$', '$', '$'},
			{'$', '$', '$'},
			{'$', '$', '$'},
		},
		x_turn_m:    true,
		game_over_m: false,
		cursorX:     1,
		cursorY:     1,
	}
}

var (
	cellStyle       = lipgloss.NewStyle().Width(5).Height(3).Align(lipgloss.Center, lipgloss.Center).Border(lipgloss.NormalBorder())
	cursorCellStyle = cellStyle.Copy().BorderForeground(lipgloss.Color("205")).Bold(true)
)

func (m model) Init() tea.Cmd {
	return tea.ClearScreen
}

func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.String() {
		case "ctrl+c", "q":
			return m, tea.Quit
		case "c":
			return m, tea.ClearScreen
		case "r":
			restart_game(&m)
			return m, tea.ClearScreen
		case "x":
			make_move(&m)
			return m, tea.ClearScreen

		case "up", "k":
			if m.cursorY > 0 {
				m.cursorY--
			}
		case "down", "j":
			if m.cursorY < 2 {
				m.cursorY++
			}
		case "left", "h":
			if m.cursorX > 0 {
				m.cursorX--
			}
		case "right", "l":
			if m.cursorX < 2 {
				m.cursorX++
			}
		}
	}
	return m, nil

}

func (m model) View() string {
	if m.game_over_m {
		if !game.Check_winner(&m.board) {
			return "DRAW"
		}

		s := "GAME OVER\n"
		if m.x_turn_m {
			s += "X WON!"
		} else {
			s += "O WON!"
		}
		return s
	}

	rows := make([]string, 3)

	for y := 0; y < 3; y++ {
		cells := make([]string, 3)
		for x := 0; x < 3; x++ {
			cell := ' '
			if m.board[y][x] != 0 {
				cell = m.board[y][x]
			}
			style := cellStyle
			if cell == 'O' {
				style = style.Foreground(lipgloss.Color("34"))
			}
			if x == m.cursorX && y == m.cursorY {
				style = cursorCellStyle
			}
			cells[x] = style.Render(string(cell))
		}
		// Join all cells in a row horizontally
		rows[y] = lipgloss.JoinHorizontal(lipgloss.Top, cells...)
	}

	// Join all rows vertically
	grid := lipgloss.JoinVertical(lipgloss.Left, rows...)

	return grid + "\n\nUse VIM Bindings to move. Press q to quit. Press r to restart"
}

func restart_game(m *model) {
	m.board = [3][3]rune{
		{'$', '$', '$'},
		{'$', '$', '$'},
		{'$', '$', '$'},
	}
	m.x_turn_m = !m.x_turn_m
	if !m.x_turn_m {
		best_move := minimax(m.board, m.x_turn_m, m.depth)
		if m.x_turn_m {
			m.board[best_move.Move.Row][best_move.Move.Col] = 'X'
		} else {
			m.board[best_move.Move.Row][best_move.Move.Col] = 'O'
		}
		m.x_turn_m = !m.x_turn_m
	}
	m.game_over_m = false
	m.cursorX = 1
	m.cursorY = 1
}

func make_move(m *model) {
	if m.board[m.cursorY][m.cursorX] != game.EMPTY {
		return
	}
	if m.x_turn_m {
		m.board[m.cursorY][m.cursorX] = 'X'
	} else {
		m.board[m.cursorY][m.cursorX] = 'O'
	}

	if game.Is_game_over(&m.board) {
		m.game_over_m = true
		return
	}

	m.x_turn_m = !m.x_turn_m

	best_move := minimax(m.board, m.x_turn_m, m.depth)
	if m.x_turn_m {
		m.board[best_move.Move.Row][best_move.Move.Col] = 'X'
	} else {
		m.board[best_move.Move.Row][best_move.Move.Col] = 'O'
	}
	if game.Is_game_over(&m.board) {
		m.game_over_m = true
		return
	}
	m.x_turn_m = !m.x_turn_m
}

func minimax(board [3][3]rune, is_x_turn bool, depth int) game.BestMove {
	if depth == 8 {
		return game.BestMove{
			Score: 0,
			Move:  game.Move{},
		}
	}
	if game.Check_winner(&board) {
		return game.BestMove{
			Score: game.Score(&board, depth),
			Move:  game.Move{},
		}
	}
	depth++
	var moves []game.Move
	var scores []int

	// Go through all the possible moves
	available_moves := game.Get_available_moves(&board)

	// fmt.Println(available_moves)

	for _, move := range available_moves {
		possible_game := game.Get_new_state(&board, move, is_x_turn)
		scores = append(scores, minimax(possible_game, !is_x_turn, depth).Score)
		moves = append(moves, move)
	}

	if is_x_turn {
		best_move := game.BestMove{
			Score: math.MinInt,
			Move:  game.Move{},
		}
		for i, val := range scores {
			if val > best_move.Score {
				best_move.Score = val
				best_move.Move.Row = moves[i].Row
				best_move.Move.Col = moves[i].Col
			}
		}
		if best_move.Score == math.MinInt {
			return game.BestMove{
				Score: 0,
				Move:  game.Move{},
			}
		}

		return best_move
	} else {
		best_move := game.BestMove{
			Score: math.MaxInt,
			Move:  game.Move{},
		}
		for i, val := range scores {
			if val < best_move.Score {
				best_move.Score = val
				best_move.Move.Row = moves[i].Row
				best_move.Move.Col = moves[i].Col
			}
		}
		if best_move.Score == math.MaxInt {
			return game.BestMove{
				Score: 0,
				Move:  game.Move{},
			}
		}
		return best_move
	}
}

func main() {

	p := tea.NewProgram(InitialModel())

	if _, err := p.Run(); err != nil {
		fmt.Printf("There was an error %v", err)
		os.Exit(1)
	}

}
