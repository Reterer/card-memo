package main

import (
	"fmt"
	"os"

	"github.com/Reterer/card-memo/console-app/tui"
	tea "github.com/charmbracelet/bubbletea"
)

func main() {
	m := tui.MakeModel()
	p := tea.NewProgram(m, tea.WithAltScreen())
	if err := p.Start(); err != nil {
		fmt.Println("Error running program:", err)
		os.Exit(1)
	}
}
