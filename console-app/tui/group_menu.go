package tui

import (
	"fmt"

	tea "github.com/charmbracelet/bubbletea"
)

type groupMenuModel struct {
	isEditMode bool

	returnModel tea.Model
	groupId     int
}

func makeGroupMenu(returnModel tea.Model, groupId int) groupMenuModel {
	m := groupMenuModel{
		returnModel: returnModel,
		groupId:     groupId,
	}
	return m
}

func (m groupMenuModel) Init() tea.Cmd {
	return nil
}

func (m groupMenuModel) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	if msg, ok := msg.(tea.KeyMsg); ok && msg.String() == "ctrl+c" {
		return m, tea.Quit
	}
	return m.defaultUpdate(msg)

}

func (m groupMenuModel) defaultUpdate(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		key := msg.String()
		if key == "esc" { // Закрыть окно группы
			return m.returnModel, nil
		}
	}
	var cmd tea.Cmd
	return m, cmd
}

func (m groupMenuModel) View() string {
	return fmt.Sprintf("list of cards. Group id %d", m.groupId)
}
