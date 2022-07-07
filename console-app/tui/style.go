package tui

import (
	"github.com/charmbracelet/bubbles/list"
	"github.com/charmbracelet/lipgloss"
)

var (
	defaultBorder = lipgloss.Border{
		Top:         "─",
		Bottom:      "─",
		Left:        "│",
		Right:       "│",
		TopLeft:     "╭",
		TopRight:    "╮",
		BottomLeft:  "╰",
		BottomRight: "╯",
	}
	defaultItemStyles = list.NewDefaultItemStyles()

	groupListStyle = lipgloss.NewStyle().
			Width(50).
			Height(40).
			Margin(1, 2)

	groupRightPanelStyle = lipgloss.NewStyle().
				Width(70).
				Height(30).
				Margin(0, 2).
				Padding(0, 1).
				Border(defaultBorder).
				BorderForeground(defaultItemStyles.NormalDesc.GetForeground())
	groupRightPanelDescTitle = defaultItemStyles.NormalTitle.Copy().
					Width(64).Align(lipgloss.Center)
	groupRightPanelDescBody = defaultItemStyles.NormalDesc.Copy().
				Width(64).Align(lipgloss.Left)

	helpRightPanelStyle = lipgloss.NewStyle().
				Width(70).
				Height(10).
				Margin(0, 2).
				Padding(0, 1).
				PaddingLeft(2).
				Border(defaultBorder).
				BorderForeground(defaultItemStyles.NormalDesc.GetForeground())
)
