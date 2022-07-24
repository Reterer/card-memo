package tui

import (
	db "github.com/Reterer/card-memo/console-app/model"
	"github.com/charmbracelet/bubbles/list"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

type groupMenuModel struct {
	cardList list.Model
	help     helpModel

	isEditMode  bool
	focusedCard card

	returnModel tea.Model
	groupId     int
}

func makeGroupMenu(returnModel tea.Model, groupId int) groupMenuModel {
	m := groupMenuModel{
		returnModel: returnModel,
		groupId:     groupId,
	}
	m.cardList = makeCardList(m)
	m.help = makeHelp()
	return m
}

func (m groupMenuModel) Init() tea.Cmd {
	return nil
}

func (m groupMenuModel) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	if msg, ok := msg.(tea.KeyMsg); ok && msg.String() == "ctrl+c" {
		return m, tea.Quit
	}
	if m.isEditMode {
		return m.editModeUpdate(msg)
	}
	return m.defaultUpdate(msg)
}

func (m groupMenuModel) editModeUpdate(msg tea.Msg) (tea.Model, tea.Cmd) {
	if !m.focusedCard.isTypingMode {
		items := m.cardList.Items()
		switch msg := msg.(type) {
		case tea.KeyMsg:
			key := msg.String()
			if key == "left" || key == "h" || key == "esc" {
				if len(items) != 0 {
					m.focusedCard.SetFocus(false)
					dbcard := db.Card{
						Title:     m.focusedCard.title,
						ShortDesc: m.focusedCard.shortDesc,
						FullDesc:  m.focusedCard.fullDesc,
						Id:        m.focusedCard.id,
						GroupId:   m.groupId,
						LearnVal:  m.focusedCard.learnValue,
					}
					if m.focusedCard.id == -1 {
						err := db.AddCard(dbcard)
						if err != nil {
							panic(err)
						}
					} else {
						err := db.UpdateCard(dbcard)
						if err != nil {
							panic(err)
						}
					}
					cmd := m.cardList.SetItem(m.cardList.Index(), m.focusedCard)
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
	m.focusedCard, cmd = m.focusedCard.EditModeUpdate(msg)
	return m, cmd
}

func (m groupMenuModel) defaultUpdate(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		key := msg.String()
		if key == "esc" { // Закрыть окно группы
			return m.returnModel, nil
		} else if key == "right" || key == "l" { // Редактировать выбранную группу
			items := m.cardList.Items()
			if len(items) != 0 {
				m.focusedCard = items[m.cardList.Index()].(card)
				m.focusedCard.SetFocus(true)
				m.isEditMode = true
			}
			return m, nil
		} else if key == "d" { // Удалить выбранную группу
			if item := m.cardList.SelectedItem(); item != nil {
				m.cardList.RemoveItem(m.cardList.Index())
				m.cardList.ResetSelected()
				err := db.RemoveCardById(item.(card).id)
				if err != nil {
					panic(err)
				}
			}
			return m, nil
		} else if key == "a" { // Создать группу
			newCard := card{
				title:     "new card",
				isFocused: true,
				id:        -1,
				groupId:   m.groupId,
			}
			newCard.SetFocus(true)
			m.cardList.InsertItem(len(m.cardList.Items()), newCard)
			m.cardList.Select(len(m.cardList.Items()) - 1)
			m.focusedCard = newCard
			m.isEditMode = true

			return m, nil
		}
	}
	var cmd tea.Cmd
	m.cardList, cmd = m.cardList.Update(msg)
	return m, cmd
}

func (m groupMenuModel) View() string {
	list := groupListStyle.Render(m.cardList.View())

	selectedItem := ""
	if item, ok := m.cardList.SelectedItem().(card); ok {
		selectedItem = item.View()
	}

	if m.isEditMode {
		selectedItem = groupRightPanelStyle.Copy().BorderForeground(defaultItemStyles.SelectedDesc.GetForeground()).Render(m.focusedCard.View())
	} else {
		selectedItem = groupRightPanelStyle.Render(selectedItem)
	}
	helpPanel := helpRightPanelStyle.Render(m.help.View())
	rightPanel := lipgloss.JoinVertical(lipgloss.Left, selectedItem, helpPanel)
	return lipgloss.JoinHorizontal(lipgloss.Top, list, rightPanel)
}
