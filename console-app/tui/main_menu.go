package tui

import (
	"github.com/charmbracelet/bubbles/list"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

type model struct {
	groupList list.Model
	help      helpModel

	isEditMode   bool
	focusedGroup group
}

func MakeModel() model {
	m := model{}
	m.groupList = makeGroupList(m)
	m.help = makeHelp(&m)
	return m
}

func (m model) Init() tea.Cmd {
	return nil
}

func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	if msg, ok := msg.(tea.KeyMsg); ok && msg.String() == "ctrl+c" {
		return m, tea.Quit
	}
	if m.isEditMode {
		return m.editModeUpdate(msg)
	}

	return m.defaultUpdate(msg)

}

func (m model) defaultUpdate(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		key := msg.String()
		if key == "enter" { // Открыть группу
			return m.openGroup()
		} else if key == "right" || key == "l" { // Редактировать выбранную группу
			items := m.groupList.Items()
			if len(items) != 0 {
				m.focusedGroup = items[m.groupList.Index()].(group)
				m.focusedGroup.SetFocus(true)
				m.isEditMode = true
			}
			return m, nil
		} else if key == "d" { // Удалить выбранную группу
			if item := m.groupList.SelectedItem(); item != nil {
				// todo delete from db
				m.groupList.RemoveItem(m.groupList.Index())
			}
			return m, nil
		} else if key == "a" { // Создать группу
			newGroup := group{
				title:     "new group",
				isFocused: true,
				id:        -1,
			}
			newGroup.SetFocus(true)
			m.groupList.InsertItem(len(m.groupList.Items()), newGroup)
			m.groupList.Select(len(m.groupList.Items()) - 1)
			m.focusedGroup = newGroup
			m.isEditMode = true

			return m, nil
		}
	case tea.WindowSizeMsg:
		h, v := groupListStyle.GetFrameSize()
		m.groupList.SetSize(msg.Width-h, msg.Height-v)
	}
	var cmd tea.Cmd
	m.groupList, cmd = m.groupList.Update(msg)
	return m, cmd
}

func (m model) editModeUpdate(msg tea.Msg) (tea.Model, tea.Cmd) {
	if !m.focusedGroup.isTypingMode {
		items := m.groupList.Items()
		switch msg := msg.(type) {
		case tea.KeyMsg:
			key := msg.String()
			if key == "left" || key == "h" {
				if len(items) != 0 {
					m.focusedGroup.SetFocus(false)
					cmd := m.groupList.SetItem(m.groupList.Index(), m.focusedGroup)
					m.isEditMode = false
					return m, cmd
				}
			}
		}
		if len(items) == 0 {
			m.isEditMode = false
			return m, nil
		}
	}

	var cmd tea.Cmd
	m.focusedGroup, cmd = m.focusedGroup.EditModeUpdate(msg)
	return m, cmd
}

func (m model) View() string {
	list := groupListStyle.Render(m.groupList.View())

	selectedItem := ""
	if item, ok := m.groupList.SelectedItem().(group); ok {
		selectedItem = item.View()
	}

	if m.isEditMode {
		selectedItem = groupRightPanelStyle.Copy().BorderForeground(defaultItemStyles.SelectedDesc.GetForeground()).Render(m.focusedGroup.View())
	} else {
		selectedItem = groupRightPanelStyle.Render(selectedItem)
	}
	helpPanel := helpRightPanelStyle.Render(m.help.View())
	rightPanel := lipgloss.JoinVertical(lipgloss.Left, selectedItem, helpPanel)
	return lipgloss.JoinHorizontal(lipgloss.Top, list, rightPanel)
}

func (m model) openGroup() (tea.Model, tea.Cmd) {
	if item := m.groupList.SelectedItem(); item != nil {
		return makeGroupMenu(m, item.(group).id), nil
	}
	return m, nil
}
