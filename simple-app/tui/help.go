package tui

import (
	"github.com/charmbracelet/bubbles/help"
	"github.com/charmbracelet/bubbles/key"
)

type keyMap struct{}

func (k keyMap) ShortHelp() []key.Binding {
	return []key.Binding{}
}
func (k keyMap) FullHelp() [][]key.Binding {
	return [][]key.Binding{
		{
			key.NewBinding(
				key.WithKeys("up", "k"),
				key.WithHelp("↑/k", "up"),
			),
			key.NewBinding(
				key.WithKeys("down", "j"),
				key.WithHelp("↓/j", "down"),
			),
			key.NewBinding(
				key.WithKeys("right", "l"),
				key.WithHelp("→/l", "open edit mode"),
			),
			key.NewBinding(
				key.WithKeys("left", "h"),
				key.WithHelp("←/h", "quit edit mode"),
			),
		},
		{
			key.NewBinding(
				key.WithKeys("ctrl+c"),
				key.WithHelp("ctrl+c", "exit"),
			),
		},
	}
}

type helpModel struct {
	help.Model
	k help.KeyMap
}

func (m helpModel) View() string {
	return m.FullHelpView(m.k.FullHelp())
}

func makeHelp() helpModel {
	km := keyMap{}
	help := helpModel{
		Model: help.New(),
		k:     km,
	}
	help.Width = helpRightPanelStyle.GetWidth()
	return help
}
