package models

import (
	"strings"
	"terminal-illness/utils"

	"github.com/charmbracelet/bubbles/textinput"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

type LoadSetupModel struct {
	inputs   []textinput.Model
	selected int
	focused  bool
}

func InitialLoadSetupModel(selected int) LoadSetupModel {
	m := LoadSetupModel{
		inputs:   make([]textinput.Model, 3),
		selected: selected,
	}

	var t textinput.Model
	for i := range m.inputs {
		t = textinput.New()
		t.Cursor.Style = utils.CursorStyle
		t.CharLimit = 256
		t.Width = 40

		switch i {
		case 0:
			t.Placeholder = "https://api.thecatapi.com/v1/images/search"
			t.Prompt = "Target: "
			t.TextStyle = utils.CursorModeHelpStyle.Italic(true)
			t.CompletionStyle = utils.CursorModeHelpStyle
		case 1:
			t.Prompt = "Method: "
			t.Placeholder = "Get"
			t.CharLimit = 64
			t.Width = 40
		case 2:
			t.Placeholder = "test"
			t.Prompt = "Load Test Field: "
			t.CharLimit = 50
			t.Width = 20
		}

		m.inputs[i] = t
	}

	return m
}

func (m LoadSetupModel) View() string {
	var b strings.Builder

	for i := range m.inputs {
		b.WriteString(m.inputs[i].View())
		if i < len(m.inputs)-1 {
			b.WriteRune('\n')
		}
	}

	var menu = b.String()
	return lipgloss.NewStyle().
		Width(52).
		Height(8).
		Border(lipgloss.RoundedBorder()).
		Padding(1).
		Render(menu)
}

func (m LoadSetupModel) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.String() {
		case "ctrl+c", "esc":
			return m, tea.Quit

		case "tab", "shift+tab", "enter", "up", "down":
			s := msg.String()

			if s == "enter" && m.selected == 2 {
				return m, nil
			}

			// Cycle indexes
			if s == "up" || s == "shift+tab" {
				m.selected--
			} else {
				m.selected++
			}

			if m.selected > len(m.inputs) {
				m.selected = 0
			} else if m.selected < 0 {
				m.selected = len(m.inputs)
			}

			var cmds []tea.Cmd
			// Only create blink cmd for the focused input
			for i := range m.inputs {
				if i == m.selected {
					// Set focused state
					cmds = append(cmds, m.inputs[i].Focus())
					m.inputs[i].PromptStyle = utils.FocusedStyle
					m.inputs[i].TextStyle = utils.FocusedStyle
					// Add the blink command
					cmds = append(cmds, textinput.Blink)
				} else {
					// Remove focused state
					m.inputs[i].Blur()
					m.inputs[i].PromptStyle = utils.NoStyle
					m.inputs[i].TextStyle = utils.NoStyle
				}
			}

			return m, tea.Batch(cmds...)
		}
	}

	// Handle character input and blinking
	cmd := m.updateInputs(msg)
	return m, cmd
}

func (m *LoadSetupModel) updateInputs(msg tea.Msg) tea.Cmd {
	cmds := make([]tea.Cmd, 0)

	if m.selected >= 0 && m.selected < len(m.inputs) {
		// Update the focused input and collect its command
		newInput, cmd := m.inputs[m.selected].Update(msg)
		m.inputs[m.selected] = newInput

		if cmd != nil {
			cmds = append(cmds, cmd)
			// Add blink command for the active input
			cmds = append(cmds, textinput.Blink)
		}
	}

	return tea.Batch(cmds...)
}

func (m LoadSetupModel) Init() tea.Cmd {
	return nil
}

func (m *LoadSetupModel) Focus() {
	m.focused = true
	m.selected = 0
	if len(m.inputs) > 0 {
		m.inputs[0].Focus()
		m.inputs[0].PromptStyle = utils.FocusedStyle
		m.inputs[0].TextStyle = utils.FocusedStyle
	}
}

func (m *LoadSetupModel) Blur() {
	m.focused = false
	m.selected = -1
	for i := range m.inputs {
		currentValue := m.inputs[i].Value()
		m.inputs[i].Blur()
		m.inputs[i].PromptStyle = utils.NoStyle
		m.inputs[i].TextStyle = utils.NoStyle
		m.inputs[i].SetValue(currentValue)
	}
}
