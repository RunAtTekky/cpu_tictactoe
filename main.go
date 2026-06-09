package main

import (
	"fmt"
	"os"

	"tictactoe/ui"

	tea "github.com/charmbracelet/bubbletea"
)

func main() {
	p := tea.NewProgram(ui.InitialModel())

	if _, err := p.Run(); err != nil {
		fmt.Printf("There was an error %v", err)
		os.Exit(1)
	}
}
