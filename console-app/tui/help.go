package tui

import (
	"github.com/charmbracelet/bubbles/help"
	"github.com/charmbracelet/bubbles/key"
)

type keyMap struct {
	m *model
}

func (k keyMap) ShortHelp() []key.Binding {
	return []key.Binding{}
}
func (k keyMap) FullHelp() [][]key.Binding {
	return [][]key.Binding{
		{
			k.m.groupList.KeyMap.CursorUp,
			k.m.groupList.KeyMap.CursorDown,
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

func makeHelp(m *model) helpModel {
	km := keyMap{
		m: m,
	}
	help := helpModel{
		Model: help.New(),
		k:     km,
	}
	help.Width = helpRightPanelStyle.GetWidth()
	return help
}
