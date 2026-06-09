package ui

import "github.com/charmbracelet/lipgloss"

var (
	cellStyle       = lipgloss.NewStyle().Width(5).Height(3).Align(lipgloss.Center, lipgloss.Center).Border(lipgloss.NormalBorder())
	cursorCellStyle = cellStyle.Copy().BorderForeground(lipgloss.Color("205")).Bold(true)
)
