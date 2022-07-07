package tui

import (
	"github.com/charmbracelet/bubbles/list"
	"github.com/charmbracelet/bubbles/textinput"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
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
	desc       string
	id         int
	groupId    int
	learnValue float64
}

func (i card) Title() string       { return i.title }
func (i card) Description() string { return i.desc }
func (i card) FilterValue() string { return i.title }

func makeGroupList(m model) list.Model {
	items := []list.Item{
		group{title: "Raspberry Pi’s", shortDesc: "I have ’em all over my house"},
		group{title: "Nutella", shortDesc: "It's good on toast"},
		group{title: "Bitter melon", shortDesc: "It cools you down"},
		group{title: "Nice socks", shortDesc: "And by that I mean socks without holes"},
		group{title: "Eight hours of sleep", shortDesc: "I had this once"},
		group{title: "Cats", shortDesc: "Usually"},
		group{title: "Plantasia, the album", shortDesc: "My plants love it too"},
		group{title: "Pour over coffee", shortDesc: "It takes forever to make though"},
		group{title: "VR", shortDesc: "Virtual reality...what is there to say?"},
		group{title: "Noguchi Lamps", shortDesc: "Such pleasing organic forms"},
		group{title: "Linux", shortDesc: "Pretty much the best OS", fullDesc: `
Quit is a convenience function for quitting Bubble Tea programs.Use it when you need to shut down a Bubble Tea program from the 
outside. If you wish to quit from within a Bubble 
Tea program use \ the Quit command. 
If the program is not running this will be 
a no-op, so it's safe to 
call if the program is unstarted or has 
already exited.
		`},
		group{title: "Business school", shortDesc: "Just kidding"},
		group{title: "Pottery", shortDesc: "Wet clay is a great feeling"},
		group{title: "Shampoo", shortDesc: "Nothing like clean hair"},
		group{title: "Table tennis", shortDesc: "It’s surprisingly exhausting"},
	}

	groupList := list.New(items, list.NewDefaultDelegate(), 0, 0)
	groupList.Title = "Groups"
	groupList.SetShowHelp(false)
	groupList.SetShowPagination(false)
	groupList.KeyMap.PrevPage.SetEnabled(false)
	groupList.KeyMap.NextPage.SetEnabled(false)
	groupList.DisableQuitKeybindings()
	return groupList
}
