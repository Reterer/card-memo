package main

import (
	"fmt"
	"os"

	"github.com/Reterer/card-memo/console-app/model"
	"github.com/Reterer/card-memo/console-app/tui"
	tea "github.com/charmbracelet/bubbletea"
)

func main() {
	var err error
	err = model.Init("file:data.db?cache=shared&mode=rwc")
	if err != nil {
		panic(err)
	}
	defer model.Deinit()

	m := tui.MakeModel()
	p := tea.NewProgram(m, tea.WithAltScreen())
	if err := p.Start(); err != nil {
		fmt.Println("Error running program:", err)
		os.Exit(1)
	}
}
