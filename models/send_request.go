package models

import (
	"terminal-illness/utils"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

// SendRequestModel represents the request execution button in the UI
type SendRequestModel struct {
	focused bool
}

func (m *SendRequestModel) Focus() {
	m.focused = true
}

func (m *SendRequestModel) Blur() {
	m.focused = false
}

func (m SendRequestModel) Init() tea.Cmd {
	return nil
}

func (m SendRequestModel) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.String() {
		case "ctrl+c", "q":
			return m, tea.Quit
		}
	}
	return m, nil
}

func (m SendRequestModel) View() string {
	style := lipgloss.NewStyle().
		Width(20).
		Height(8).
		Border(lipgloss.RoundedBorder()).
		Padding(1)

	buttonStyle := utils.AuthorStyle
	if m.focused {
		buttonStyle = buttonStyle.Copy().Foreground(lipgloss.Color("205"))
		style = style.BorderForeground(lipgloss.Color("205"))
	}

	return style.Render(buttonStyle.Render(utils.SendButton))
}
