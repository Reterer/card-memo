package tui

import (
	"fmt"

	"github.com/charmbracelet/bubbles/list"
	"github.com/charmbracelet/bubbles/textinput"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"

	db "github.com/Reterer/card-memo/console-app/model"
)

type group struct {
	title     string
	shortDesc string
	fullDesc  string
	id        int

	isFocused      bool
	isTypingMode   bool
	cursor         int
	titleInput     textinput.Model
	shortDescInput textinput.Model
	fullDescInput  textinput.Model
}

func (g group) Title() string       { return g.title }
func (g group) Description() string { return g.shortDesc }
func (g group) FilterValue() string { return g.title }

func (g *group) SetFocus(focus bool) {
	g.isFocused = focus
	g.cursor = 0
	g.titleInput.SetValue(g.title)
	g.titleInput.SetCursorMode(textinput.CursorStatic)

	g.shortDescInput.SetValue(g.shortDesc)
	g.shortDescInput.SetCursorMode(textinput.CursorStatic)

	g.fullDescInput.SetValue(g.fullDesc)
	g.fullDescInput.SetCursorMode(textinput.CursorStatic)

}

func (g group) EditModeUpdate(msg tea.Msg) (group, tea.Cmd) {
	if g.isTypingMode {
		if key, ok := msg.(tea.KeyMsg); ok && key.String() == "esc" {
			g.isTypingMode = false
			g.title = g.titleInput.Value()
			g.shortDesc = g.shortDescInput.Value()
			g.fullDesc = g.fullDescInput.Value()
			if g.cursor == 0 {
				g.titleInput.Blur()
			} else if g.cursor == 1 {
				g.shortDescInput.Blur()
			} else if g.cursor == 2 {
				g.fullDescInput.Blur()
			}
			return g, nil
		}
		var cmd tea.Cmd
		if g.cursor == 0 {
			g.titleInput, cmd = g.titleInput.Update(msg)
		} else if g.cursor == 1 {
			g.shortDescInput, cmd = g.shortDescInput.Update(msg)
		} else if g.cursor == 2 {
			g.fullDescInput, cmd = g.fullDescInput.Update(msg)
		}
		return g, cmd
	}

	switch msg := msg.(type) {
	case tea.KeyMsg:
		key := msg.String()
		if key == "up" || key == "k" {
			if g.cursor > 0 {
				g.cursor--
			}
		} else if key == "down" || key == "j" {
			if g.cursor < 2 {
				g.cursor++
			}
		} else if key == "enter" {
			g.isTypingMode = true
			var cmd tea.Cmd
			if g.cursor == 0 {
				cmd = g.titleInput.Focus()
			} else if g.cursor == 1 {
				cmd = g.shortDescInput.Focus()
			} else if g.cursor == 2 {
				cmd = g.fullDescInput.Focus()
			}

			return g, cmd
		}
	}
	return g, nil
}

func (g group) View() string {
	selectedHeaderStyle := defaultItemStyles.NormalTitle.Copy().Foreground(defaultItemStyles.SelectedDesc.GetForeground())
	selectedDescHeaderStyle := groupRightPanelDescTitle.Copy().Foreground(defaultItemStyles.SelectedDesc.GetForeground())
	selectedBodyStyle := groupRightPanelDescBody.Copy().Foreground(defaultItemStyles.SelectedDesc.GetForeground())
	sections := []string{}
	{
		header := "GROUP:"
		title := g.title
		if g.isTypingMode && g.cursor == 0 {
			title = g.titleInput.View()
		}

		if g.isFocused && g.cursor == 0 {
			header = selectedHeaderStyle.Render(header)
			title = selectedHeaderStyle.Render(title)
		} else {
			header = defaultItemStyles.NormalTitle.Render(header)
			title = groupRightPanelDescBody.Render(title)
		}
		sections = append(sections, header+title+"\n")
	}
	{
		header := "SHORT DESC:"
		shortDesc := g.shortDesc
		if g.isTypingMode && g.cursor == 1 {
			shortDesc = g.shortDescInput.View()
		}

		if g.isFocused && g.cursor == 1 {
			header = selectedHeaderStyle.Render(header)
			shortDesc = selectedBodyStyle.Render(shortDesc)
		} else {
			header = defaultItemStyles.NormalTitle.Render(header)
			shortDesc = groupRightPanelDescBody.Render(shortDesc)
		}
		sections = append(sections, header+shortDesc+"\n")
	}
	{
		header := "FULL DESC"
		fullDesc := g.fullDesc
		if g.isTypingMode && g.cursor == 2 {
			fullDesc = g.fullDescInput.View()
		}

		if g.isFocused && g.cursor == 2 {
			header = selectedDescHeaderStyle.Render(header)
			fullDesc = selectedBodyStyle.Render(fullDesc)
		} else {
			header = groupRightPanelDescTitle.Render(header)
			fullDesc = groupRightPanelDescBody.Render(fullDesc)
		}
		sections = append(sections, header+fullDesc)
	}

	return lipgloss.JoinVertical(lipgloss.Left, sections...)
}

type card struct {
	title      string
	shortDesc  string
	fullDesc   string
	id         int
	groupId    int
	learnValue float64

	isFocused      bool
	isTypingMode   bool
	cursor         int
	titleInput     textinput.Model
	shortDescInput textinput.Model
	fullDescInput  textinput.Model
}

func (i card) Title() string       { return i.title }
func (i card) Description() string { return i.shortDesc }
func (i card) FilterValue() string { return i.title }

func (c *card) SetFocus(focus bool) {
	c.isFocused = focus
	c.cursor = 0
	c.titleInput.SetValue(c.title)
	c.titleInput.SetCursorMode(textinput.CursorStatic)

	c.shortDescInput.SetValue(c.shortDesc)
	c.shortDescInput.SetCursorMode(textinput.CursorStatic)

	c.fullDescInput.SetValue(c.fullDesc)
	c.fullDescInput.SetCursorMode(textinput.CursorStatic)

}

func (c card) EditModeUpdate(msg tea.Msg) (card, tea.Cmd) {
	if c.isTypingMode {
		if key, ok := msg.(tea.KeyMsg); ok && key.String() == "esc" {
			c.isTypingMode = false
			c.title = c.titleInput.Value()
			c.shortDesc = c.shortDescInput.Value()
			c.fullDesc = c.fullDescInput.Value()
			if c.cursor == 0 {
				c.titleInput.Blur()
			} else if c.cursor == 1 {
				c.shortDescInput.Blur()
			} else if c.cursor == 2 {
				c.fullDescInput.Blur()
			}
			return c, nil
		}
		var cmd tea.Cmd
		if c.cursor == 0 {
			c.titleInput, cmd = c.titleInput.Update(msg)
		} else if c.cursor == 1 {
			c.shortDescInput, cmd = c.shortDescInput.Update(msg)
		} else if c.cursor == 2 {
			c.fullDescInput, cmd = c.fullDescInput.Update(msg)
		}
		return c, cmd
	}

	switch msg := msg.(type) {
	case tea.KeyMsg:
		key := msg.String()
		if key == "up" || key == "k" {
			if c.cursor > 0 {
				c.cursor--
			}
		} else if key == "down" || key == "j" {
			if c.cursor < 2 {
				c.cursor++
			}
		} else if key == "enter" {
			c.isTypingMode = true
			var cmd tea.Cmd
			if c.cursor == 0 {
				cmd = c.titleInput.Focus()
			} else if c.cursor == 1 {
				cmd = c.shortDescInput.Focus()
			} else if c.cursor == 2 {
				cmd = c.fullDescInput.Focus()
			}

			return c, cmd
		}
	}
	return c, nil
}

func (c card) View() string {
	selectedHeaderStyle := defaultItemStyles.NormalTitle.Copy().Foreground(defaultItemStyles.SelectedDesc.GetForeground())
	selectedDescHeaderStyle := groupRightPanelDescTitle.Copy().Foreground(defaultItemStyles.SelectedDesc.GetForeground())
	selectedBodyStyle := groupRightPanelDescBody.Copy().Foreground(defaultItemStyles.SelectedDesc.GetForeground())
	sections := []string{}

	{
		header := "CARD:"
		title := c.title
		if c.isTypingMode && c.cursor == 0 {
			title = c.titleInput.View()
		}

		if c.isFocused && c.cursor == 0 {
			header = selectedHeaderStyle.Render(header)
			title = selectedHeaderStyle.Render(title)
		} else {
			header = defaultItemStyles.NormalTitle.Render(header)
			title = groupRightPanelDescBody.Render(title)
		}
		sections = append(sections, header+title+"\n")
	}
	{
		header := "SHORT DESC:"
		shortDesc := c.shortDesc
		if c.isTypingMode && c.cursor == 1 {
			shortDesc = c.shortDescInput.View()
		}

		if c.isFocused && c.cursor == 1 {
			header = selectedHeaderStyle.Render(header)
			shortDesc = selectedBodyStyle.Render(shortDesc)
		} else {
			header = defaultItemStyles.NormalTitle.Render(header)
			shortDesc = groupRightPanelDescBody.Render(shortDesc)
		}
		sections = append(sections, header+shortDesc+"\n")
	}
	{
		header := "FULL DESC"
		fullDesc := c.fullDesc
		if c.isTypingMode && c.cursor == 2 {
			fullDesc = c.fullDescInput.View()
		}

		if c.isFocused && c.cursor == 2 {
			header = selectedDescHeaderStyle.Render(header)
			fullDesc = selectedBodyStyle.Render(fullDesc)
		} else {
			header = groupRightPanelDescTitle.Render(header)
			fullDesc = groupRightPanelDescBody.Render(fullDesc)
		}
		sections = append(sections, header+fullDesc)
	}
	{
		header := defaultItemStyles.NormalTitle.Render("LEARNING VAL:")
		val := groupRightPanelDescBody.Render(fmt.Sprintf("%f", c.learnValue))
		sections = append(sections, header+val)
	}

	return lipgloss.JoinVertical(lipgloss.Left, sections...)
}

func makeGroupList(m model) list.Model {
	var items []list.Item
	groups, err := db.Groups()
	if err != nil {
		panic(err)
	}
	for _, g := range groups {
		items = append(items, group{
			id:        g.Id,
			title:     g.Title,
			shortDesc: g.ShortDesc,
			fullDesc:  g.FullDesc,
		})
	}

	groupList := list.New(items, list.NewDefaultDelegate(), groupListStyle.GetWidth(), groupListStyle.GetHeight())
	groupList.Title = "Groups"
	groupList.SetShowHelp(false)
	groupList.SetShowPagination(false)
	groupList.KeyMap.PrevPage.SetEnabled(false)
	groupList.KeyMap.NextPage.SetEnabled(false)
	groupList.DisableQuitKeybindings()
	return groupList
}

func makeCardList(m groupMenuModel) list.Model {
	var items []list.Item
	cards, err := db.CardsOfGroup(m.groupId)
	if err != nil {
		panic(err)
	}
	for _, c := range cards {
		items = append(items, dbCardToTuiCard(c))
	}
	cardList := list.New(items, list.NewDefaultDelegate(), groupListStyle.GetWidth(), groupListStyle.GetHeight())
	cardList.Title = "Groups > Cards"
	cardList.SetShowHelp(false)
	cardList.SetShowPagination(false)
	cardList.KeyMap.PrevPage.SetEnabled(false)
	cardList.KeyMap.NextPage.SetEnabled(false)
	cardList.DisableQuitKeybindings()
	return cardList
}

func dbCardToTuiCard(c db.Card) card {
	return card{
		id:         c.Id,
		groupId:    c.GroupId,
		learnValue: c.LearnVal,
		title:      c.Title,
		shortDesc:  c.ShortDesc,
		fullDesc:   c.FullDesc,
	}
}
