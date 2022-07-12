package tui

import (
	"unicode"

	db "github.com/Reterer/card-memo/console-app/model"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

type repeatMenuModel struct {
	currQCards []db.Card
	nextQCards []db.Card
	groupId    int

	returnModel tea.Model
}

func MakeRepeatMenuModel(returnModel tea.Model, groupId int) repeatMenuModel {
	m := repeatMenuModel{
		returnModel: returnModel,
		groupId:     groupId,
	}

	var err error

	m.currQCards, err = db.CardsOfGroup(groupId)
	if err != nil {
		panic(err)
	}
	// fmt.Println(m.currQCards, groupId)
	// var v int
	// fmt.Scan(&v)
	return m
}

func (m repeatMenuModel) Init() tea.Cmd {
	return nil
}

func (m repeatMenuModel) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	if msg, ok := msg.(tea.KeyMsg); ok && msg.String() == "ctrl+c" {
		return m, tea.Quit
	}
	if len(m.currQCards) == 0 {
		return m.returnModel, nil
	}
	return m.defaultUpdate(msg)
}

func (m repeatMenuModel) defaultUpdate(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		key := msg.String()
		if key == "esc" { // Закрыть окно
			return m.returnModel, nil
		} else if key == "y" {
			_, err := db.UpdateCardLearningVal(m.currQCards[len(m.currQCards)-1], 1.0)
			if err != nil {
				panic(err)
			}

		} else if key == "n" {
			newVal, err := db.UpdateCardLearningVal(m.currQCards[len(m.currQCards)-1], 0)
			if err != nil {
				panic(err)
			}
			m.nextQCards = append(m.nextQCards, m.currQCards[len(m.currQCards)-1])
			m.nextQCards[len(m.nextQCards)-1].LearnVal = newVal
		} else if unicode.IsDigit(rune(key[0])) {
			digit := float64(key[0] - '0')
			_, err := db.UpdateCardLearningVal(m.currQCards[len(m.currQCards)-1], 0.1*digit)
			if err != nil {
				panic(err)
			}
		}
		m.currQCards = m.currQCards[:len(m.currQCards)-1]
	}

	if len(m.currQCards) == 0 {
		m.currQCards, m.nextQCards = m.nextQCards, m.currQCards
	}
	if len(m.currQCards) == 0 {
		return m.returnModel, nil
	}

	return m, nil
}

func (m repeatMenuModel) View() string {
	currCard := dbCardToTuiCard(m.currQCards[len(m.currQCards)-1])
	item := groupRightPanelStyle.Render(currCard.View())
	return lipgloss.JoinVertical(lipgloss.Center, item)
}
